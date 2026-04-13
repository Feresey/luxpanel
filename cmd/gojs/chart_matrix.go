package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/Feresey/luxpanel/internal/prettyfmt"
	"github.com/Feresey/luxpanel/internal/splitter"
)

// teamPairPanel: columns = col_players (allies for this panel), rows = row_players (enemies).
// For heal, both dimensions are the same team (intra-team healer → recipient).
type teamPairPanel struct {
	ColPlayers []string    `json:"col_players"`
	RowPlayers []string    `json:"row_players"`
	NewbieBonusMax map[string]int `json:"newbie_bonus_max,omitempty"`
	Matrix     [][]float64 `json:"matrix"`
	Total      float64     `json:"total"`
}

type chartMatricesResult struct {
	Metric string          `json:"metric"`
	Panels []teamPairPanel `json:"panels"`
}

type chartQueryOptions struct {
	DamageType  string `json:"damage_type,omitempty"`  // all | thermal | kinetic | emp
	IncludeBots bool   `json:"include_bots,omitempty"` // include bot players and bot damage
}

func parseChartQueryOptions(raw string) chartQueryOptions {
	opts := chartQueryOptions{DamageType: "all", IncludeBots: false}
	if strings.TrimSpace(raw) == "" {
		return opts
	}
	_ = json.Unmarshal([]byte(raw), &opts)
	opts.DamageType = strings.ToLower(strings.TrimSpace(opts.DamageType))
	switch opts.DamageType {
	case "thermal", "kinetic", "emp":
	default:
		opts.DamageType = "all"
	}
	return opts
}

func damageMatchesType(dmg *combat.Damage, damageType string) bool {
	if dmg == nil || damageType == "" || damageType == "all" {
		return true
	}
	want := ""
	switch damageType {
	case "thermal":
		want = "THERMAL"
	case "kinetic":
		want = "KINETIC"
	case "emp":
		want = "EMP"
	default:
		return true
	}
	for _, m := range dmg.DamageModifiers {
		if string(m) == want {
			return true
		}
	}
	return false
}

func (r *Runtime) marshalChartMatricesJSON(ctx context.Context, level *splitter.Level, mode string, timeRangeJSON, optsJSON string) (result string, err error) {
	t0 := time.Now()
	heap0 := heapAlloc()
	defer func() {
		logPerfWasm(ctx, r.lg, "marshal_chart_matrices", t0, heap0,
			"mode", mode,
			"out_size", prettyfmt.FormatBytes(uint64(len(result))),
			"has_err", err != nil,
			"opts_size", prettyfmt.FormatBytes(uint64(len(optsJSON))),
		)
	}()
	if level == nil {
		var b []byte
		b, err = json.Marshal(chartMatricesResult{})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalChartMatricesJSON", "err", err, "nil_level", true)
			return "", err
		}
		result = string(b)
		r.lg.For(ctx).Debugw("marshalChartMatricesJSON", "nil_level", true, "out_len", len(result))
		return result, nil
	}
	lo, hi := clampTimeRange(level, timeRangeJSON)
	opts := parseChartQueryOptions(optsJSON)
	switch mode {
	case "heal":
		result, err = r.marshalHealMatrices(ctx, level, lo, hi, opts)
	case "kill":
		result, err = r.marshalKillMatrices(ctx, level, lo, hi, opts)
	default:
		result, err = r.marshalDamageMatrices(ctx, level, lo, hi, opts)
	}
	return result, err
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

func maxNewbieBonusByPlayer(level *splitter.Level, players []splitter.Player) map[string]int {
	if level == nil || level.CombatLog == nil || len(level.CombatLog.Spell) == 0 || len(players) == 0 {
		return nil
	}
	byLower := make(map[string]string, len(players))
	for _, p := range players {
		n := strings.TrimSpace(p.Name)
		if n == "" {
			continue
		}
		byLower[strings.ToLower(n)] = n
	}
	out := make(map[string]int)
	for _, sp := range level.CombatLog.Spell {
		if sp == nil || sp.Normalizer <= 0 {
			continue
		}
		for _, targetRaw := range sp.Targets {
			target := strings.TrimSpace(targetRaw)
			if target == "" {
				continue
			}
			canon, ok := byLower[strings.ToLower(target)]
			if !ok || canon == "" {
				continue
			}
			if sp.Normalizer > out[canon] {
				out[canon] = sp.Normalizer
			}
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func (r *Runtime) marshalDamageMatrices(ctx context.Context, level *splitter.Level, lo, hi float64, opts chartQueryOptions) (string, error) {
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
	pla, plb, nameToTeam, nameToIdx, ok := twoTeamRosters(level, opts.IncludeBots)
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
	allPlayers := append(append([]splitter.Player{}, pla...), plb...)
	newbieMax := maxNewbieBonusByPlayer(level, allPlayers)
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
		if !damageMatchesType(dmg, opts.DamageType) {
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
		NewbieBonusMax: newbieMax,
		Matrix:     m0,
		Total:      sumMatrix(m0),
	}
	p1 := teamPairPanel{
		ColPlayers: playerNameSlice(plb),
		RowPlayers: playerNameSlice(pla),
		NewbieBonusMax: newbieMax,
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

func (r *Runtime) marshalKillMatrices(ctx context.Context, level *splitter.Level, lo, hi float64, opts chartQueryOptions) (string, error) {
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
	pla, plb, nameToTeam, nameToIdx, ok := twoTeamRosters(level, opts.IncludeBots)
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
	allPlayers := append(append([]splitter.Player{}, pla...), plb...)
	newbieMax := maxNewbieBonusByPlayer(level, allPlayers)
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
		NewbieBonusMax: newbieMax,
		Matrix:     m0,
		Total:      sumMatrix(m0),
	}
	p1 := teamPairPanel{
		ColPlayers: playerNameSlice(plb),
		RowPlayers: playerNameSlice(pla),
		NewbieBonusMax: newbieMax,
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

func (r *Runtime) marshalHealMatrices(ctx context.Context, level *splitter.Level, lo, hi float64, opts chartQueryOptions) (string, error) {
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
	pla, plb, nameToTeam, nameToIdx, ok := twoTeamRosters(level, opts.IncludeBots)
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
	allPlayers := append(append([]splitter.Player{}, pla...), plb...)
	newbieMax := maxNewbieBonusByPlayer(level, allPlayers)
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
		NewbieBonusMax: newbieMax,
		Matrix:     m0,
		Total:      sumMatrix(m0),
	}
	p1 := teamPairPanel{
		ColPlayers: playerNameSlice(plb),
		RowPlayers: playerNameSlice(plb),
		NewbieBonusMax: newbieMax,
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
