package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Feresey/luxpanel/internal/splitter"
)

// teamPairPanel: columns = col_players (allies for this panel), rows = row_players (enemies).
// For heal, both dimensions are the same team (intra-team healer → recipient).
type teamPairPanel struct {
	ColPlayers []string    `json:"col_players"`
	RowPlayers []string    `json:"row_players"`
	Matrix     [][]float64 `json:"matrix"`
	Total      float64     `json:"total"`
}

type chartMatricesResult struct {
	Metric string          `json:"metric"`
	Panels []teamPairPanel `json:"panels"`
}

func (r *Runtime) marshalChartMatricesJSON(ctx context.Context, level *splitter.Level, mode string, timeRangeJSON string) (string, error) {
	if level == nil {
		b, err := json.Marshal(chartMatricesResult{})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalChartMatricesJSON", "err", err, "nil_level", true)
			return "", err
		}
		s := string(b)
		r.lg.For(ctx).Debugw("marshalChartMatricesJSON", "nil_level", true, "out_len", len(s))
		return s, nil
	}
	lo, hi := clampTimeRange(level, timeRangeJSON)
	switch mode {
	case "heal":
		return r.marshalHealMatrices(ctx, level, lo, hi)
	case "kill":
		return r.marshalKillMatrices(ctx, level, lo, hi)
	default:
		return r.marshalDamageMatrices(ctx, level, lo, hi)
	}
}

func playerNameSlice(pl []splitter.Player) []string {
	out := make([]string, len(pl))
	for i, p := range pl {
		out[i] = p.Name
	}
	return out
}

func sumMatrix(m [][]float64) float64 {
	var t float64
	for _, row := range m {
		for _, v := range row {
			t += v
		}
	}
	return t
}

func (r *Runtime) marshalDamageMatrices(ctx context.Context, level *splitter.Level, lo, hi float64) (string, error) {
	if level.CombatLog == nil {
		b, err := json.Marshal(chartMatricesResult{Metric: "damage"})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalDamageMatrices", "err", err)
			return "", err
		}
		s := string(b)
		r.lg.For(ctx).Debugw("marshalDamageMatrices", "no_combat_log", true, "out_len", len(s))
		return s, nil
	}
	pla, plb, nameToTeam, nameToIdx, ok := twoTeamRosters(level)
	if !ok {
		b, err := json.Marshal(chartMatricesResult{Metric: "damage"})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalDamageMatrices", "err", err)
			return "", err
		}
		s := string(b)
		r.lg.For(ctx).Debugw("marshalDamageMatrices", "no_two_teams", true, "out_len", len(s))
		return s, nil
	}
	na, nb := len(pla), len(plb)
	m0 := make([][]float64, nb)
	for i := range m0 {
		m0[i] = make([]float64, na)
	}
	m1 := make([][]float64, na)
	for i := range m1 {
		m1[i] = make([]float64, nb)
	}
	t0 := level.StartLevelTime
	for _, dmg := range level.CombatLog.Damage {
		if dmg == nil || dmg.IsEmpty() {
			continue
		}
		st, okS := nameToTeam[dmg.Initiator.Name]
		tt, okT := nameToTeam[dmg.Recipient.Name]
		if !okS || !okT || st == tt {
			continue
		}
		t := dmg.GetTime(t0)
		if t.Before(t0) {
			t = t0
		}
		if !timeInRangeFromStart(t, t0, lo, hi) {
			continue
		}
		amount := float64(dmg.DamageFull)
		if st == 0 && tt == 1 {
			ri := nameToIdx[dmg.Recipient.Name]
			ci := nameToIdx[dmg.Initiator.Name]
			if ri >= 0 && ri < nb && ci >= 0 && ci < na {
				m0[ri][ci] += amount
			}
		} else if st == 1 && tt == 0 {
			ri := nameToIdx[dmg.Recipient.Name]
			ci := nameToIdx[dmg.Initiator.Name]
			if ri >= 0 && ri < na && ci >= 0 && ci < nb {
				m1[ri][ci] += amount
			}
		}
	}
	p0 := teamPairPanel{
		ColPlayers: playerNameSlice(pla),
		RowPlayers: playerNameSlice(plb),
		Matrix:     m0,
		Total:      sumMatrix(m0),
	}
	p1 := teamPairPanel{
		ColPlayers: playerNameSlice(plb),
		RowPlayers: playerNameSlice(pla),
		Matrix:     m1,
		Total:      sumMatrix(m1),
	}
	res := chartMatricesResult{Metric: "damage", Panels: []teamPairPanel{p0, p1}}
	b, err := json.Marshal(res)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalDamageMatrices", "err", err)
		return "", fmt.Errorf("marshal damage matrices: %w", err)
	}
	s := string(b)
	r.lg.For(ctx).Debugw("marshalDamageMatrices", "time_lo", lo, "time_hi", hi, "out_len", len(s))
	return s, nil
}

func (r *Runtime) marshalKillMatrices(ctx context.Context, level *splitter.Level, lo, hi float64) (string, error) {
	if level.CombatLog == nil {
		b, err := json.Marshal(chartMatricesResult{Metric: "kill"})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalKillMatrices", "err", err)
			return "", err
		}
		s := string(b)
		r.lg.For(ctx).Debugw("marshalKillMatrices", "no_combat_log", true, "out_len", len(s))
		return s, nil
	}
	pla, plb, nameToTeam, nameToIdx, ok := twoTeamRosters(level)
	if !ok {
		b, err := json.Marshal(chartMatricesResult{Metric: "kill"})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalKillMatrices", "err", err)
			return "", err
		}
		s := string(b)
		r.lg.For(ctx).Debugw("marshalKillMatrices", "no_two_teams", true, "out_len", len(s))
		return s, nil
	}
	na, nb := len(pla), len(plb)
	m0 := make([][]float64, nb)
	for i := range m0 {
		m0[i] = make([]float64, na)
	}
	m1 := make([][]float64, na)
	for i := range m1 {
		m1[i] = make([]float64, nb)
	}
	t0 := level.StartLevelTime
	for _, k := range level.CombatLog.Kill {
		if k == nil || k.IsEmpty() {
			continue
		}
		t := k.GetTime(t0)
		if t.Before(t0) {
			t = t0
		}
		if !timeInRangeFromStart(t, t0, lo, hi) {
			continue
		}
		st, okK := nameToTeam[k.Killer.Name]
		vt, okV := nameToTeam[k.Killed.Name]
		if !okK || !okV || st == vt {
			continue
		}
		if st == 0 && vt == 1 {
			ri := nameToIdx[k.Killed.Name]
			ci := nameToIdx[k.Killer.Name]
			if ri >= 0 && ri < nb && ci >= 0 && ci < na {
				m0[ri][ci]++
			}
		} else if st == 1 && vt == 0 {
			ri := nameToIdx[k.Killed.Name]
			ci := nameToIdx[k.Killer.Name]
			if ri >= 0 && ri < na && ci >= 0 && ci < nb {
				m1[ri][ci]++
			}
		}
	}
	p0 := teamPairPanel{
		ColPlayers: playerNameSlice(pla),
		RowPlayers: playerNameSlice(plb),
		Matrix:     m0,
		Total:      sumMatrix(m0),
	}
	p1 := teamPairPanel{
		ColPlayers: playerNameSlice(plb),
		RowPlayers: playerNameSlice(pla),
		Matrix:     m1,
		Total:      sumMatrix(m1),
	}
	res := chartMatricesResult{Metric: "kill", Panels: []teamPairPanel{p0, p1}}
	b, err := json.Marshal(res)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalKillMatrices", "err", err)
		return "", fmt.Errorf("marshal kill matrices: %w", err)
	}
	s := string(b)
	r.lg.For(ctx).Debugw("marshalKillMatrices", "time_lo", lo, "time_hi", hi, "out_len", len(s))
	return s, nil
}

func (r *Runtime) marshalHealMatrices(ctx context.Context, level *splitter.Level, lo, hi float64) (string, error) {
	if level.CombatLog == nil {
		b, err := json.Marshal(chartMatricesResult{Metric: "heal"})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalHealMatrices", "err", err)
			return "", err
		}
		s := string(b)
		r.lg.For(ctx).Debugw("marshalHealMatrices", "no_combat_log", true, "out_len", len(s))
		return s, nil
	}
	pla, plb, nameToTeam, nameToIdx, ok := twoTeamRosters(level)
	if !ok {
		b, err := json.Marshal(chartMatricesResult{Metric: "heal"})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalHealMatrices", "err", err)
			return "", err
		}
		s := string(b)
		r.lg.For(ctx).Debugw("marshalHealMatrices", "no_two_teams", true, "out_len", len(s))
		return s, nil
	}
	na, nb := len(pla), len(plb)
	m0 := make([][]float64, na)
	for i := range m0 {
		m0[i] = make([]float64, na)
	}
	m1 := make([][]float64, nb)
	for i := range m1 {
		m1[i] = make([]float64, nb)
	}
	t0 := level.StartLevelTime
	for _, h := range level.CombatLog.Heal {
		if h == nil || h.IsEmpty() {
			continue
		}
		t := h.GetTime(t0)
		if t.Before(t0) {
			t = t0
		}
		if !timeInRangeFromStart(t, t0, lo, hi) {
			continue
		}
		st, okS := nameToTeam[h.Initiator.Name]
		tt, okT := nameToTeam[h.Recipient.Name]
		if !okS || !okT || st != tt {
			continue
		}
		amount := float64(h.Heal)
		if st == 0 {
			ri := nameToIdx[h.Recipient.Name]
			ci := nameToIdx[h.Initiator.Name]
			if ri >= 0 && ri < na && ci >= 0 && ci < na {
				m0[ri][ci] += amount
			}
		} else {
			ri := nameToIdx[h.Recipient.Name]
			ci := nameToIdx[h.Initiator.Name]
			if ri >= 0 && ri < nb && ci >= 0 && ci < nb {
				m1[ri][ci] += amount
			}
		}
	}
	p0 := teamPairPanel{
		ColPlayers: playerNameSlice(pla),
		RowPlayers: playerNameSlice(pla),
		Matrix:     m0,
		Total:      sumMatrix(m0),
	}
	p1 := teamPairPanel{
		ColPlayers: playerNameSlice(plb),
		RowPlayers: playerNameSlice(plb),
		Matrix:     m1,
		Total:      sumMatrix(m1),
	}
	res := chartMatricesResult{Metric: "heal", Panels: []teamPairPanel{p0, p1}}
	b, err := json.Marshal(res)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalHealMatrices", "err", err)
		return "", fmt.Errorf("marshal heal matrices: %w", err)
	}
	s := string(b)
	r.lg.For(ctx).Debugw("marshalHealMatrices", "time_lo", lo, "time_hi", hi, "out_len", len(s))
	return s, nil
}
