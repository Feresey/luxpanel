package main

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"time"

	"golang.org/x/exp/maps"

	"github.com/Feresey/luxpanel/internal/prettyfmt"
	"github.com/Feresey/luxpanel/internal/splitter"
)

// sortedNonEmptyTeamIDs: ненулевые команды с непустым ростером, по возрастанию id.
func sortedNonEmptyTeamIDs(level *splitter.Level) []int {
	if level == nil {
		return nil
	}
	ids := maps.Keys(level.Teams)
	slices.Sort(ids)
	var out []int
	for _, id := range ids {
		if id == 0 {
			continue
		}
		if pl := level.Teams[id]; len(pl) > 0 {
			out = append(out, id)
		}
	}
	return out
}

// chartTeamSides: левый график всегда союзная команда (localClientTeamID из combat Start),
// правый — первая непустая не-нулевая команда, отличная от союзной.
func chartTeamSides(level *splitter.Level) (leftID, rightID int, pla, plb []splitter.Player, ok bool) {
	if level == nil || len(level.Teams) == 0 {
		return 0, 0, nil, nil, false
	}
	allyID, okAlly := allyTeamIDForTimeline(level)
	if !okAlly || allyID == 0 {
		return 0, 0, nil, nil, false
	}
	pla = level.Teams[allyID]
	if len(pla) == 0 {
		return 0, 0, nil, nil, false
	}
	for _, id := range sortedNonEmptyTeamIDs(level) {
		if id == 0 || id == allyID {
			continue
		}
		plb = level.Teams[id]
		if len(plb) > 0 {
			return allyID, id, pla, plb, true
		}
	}
	return 0, 0, nil, nil, false
}

func chartTeams(level *splitter.Level) (pla, plb []splitter.Player, ok bool) {
	_, _, a, b, ok := chartTeamSides(level)
	return a, b, ok
}

// chartCurve matches the Vue/eta-chart structure (name, data, step).
type chartCurve struct {
	Name string    `json:"name"`
	Data []float64 `json:"data"`
	Step []float64 `json:"step"`
}

func filterRosterPlayers(players []splitter.Player, includeBots bool) []splitter.Player {
	if includeBots {
		return players
	}
	out := make([]splitter.Player, 0, len(players))
	for _, p := range players {
		if p.PlayerID != 0 {
			out = append(out, p)
		}
	}
	return out
}

func twoTeamRosters(level *splitter.Level, includeBots bool) (
	pla, plb []splitter.Player,
	nameToTeam map[string]int,
	nameToIdx map[string]int,
	ok bool,
) {
	var okTeams bool
	pla, plb, okTeams = chartTeams(level)
	if !okTeams {
		return nil, nil, nil, nil, false
	}
	pla = filterRosterPlayers(pla, includeBots)
	plb = filterRosterPlayers(plb, includeBots)
	if len(pla) == 0 || len(plb) == 0 {
		return nil, nil, nil, nil, false
	}
	nameToTeam = make(map[string]int)
	nameToIdx = make(map[string]int)
	for i, p := range pla {
		nameToTeam[p.Name] = 0
		nameToIdx[p.Name] = i
	}
	for i, p := range plb {
		nameToTeam[p.Name] = 1
		nameToIdx[p.Name] = i
	}
	return pla, plb, nameToTeam, nameToIdx, true
}

func buildDamageCharts(level *splitter.Level, lo, hi float64) [][]chartCurve {
	if level == nil || level.CombatLog == nil {
		return emptyCharts(level)
	}

	pla, plb, nameToTeam, nameToIdx, ok := twoTeamRosters(level, false)
	if !ok {
		return emptyCharts(level)
	}

	curves0 := newTeamCurves("Damage", pla)
	curves1 := newTeamCurves("Damage", plb)
	capHint := len(level.CombatLog.Damage)/2 + 16
	preGrowCurveSlices(curves0, capHint)
	preGrowCurveSlices(curves1, capHint)

	var sum0, sum1 float64
	pTot0 := make([]float64, len(pla))
	pTot1 := make([]float64, len(plb))

	t0 := level.StartLevelTime

	for _, dmg := range level.CombatLog.Damage {
		if dmg == nil || dmg.IsEmpty() {
			continue
		}
		srcName := dmg.Initiator.Name
		tgtName := dmg.Recipient.Name
		st, okS := nameToTeam[srcName]
		tt, okT := nameToTeam[tgtName]
		if !okS || !okT || st == tt {
			continue
		}
		amount := float64(dmg.DamageFull)
		t := dmg.GetTime(t0)
		if t.Before(t0) {
			t = t0
		}
		if !timeInRangeFromStart(t, t0, lo, hi) {
			continue
		}
		step := float64(t.Sub(t0)) / float64(time.Second)

		if st == 0 {
			sum0 += amount / float64(len(pla))
			curves0[0].Step = append(curves0[0].Step, step)
			curves0[0].Data = append(curves0[0].Data, sum0)
			idx := nameToIdx[srcName]
			pTot0[idx] += amount
			curves0[idx+1].Step = append(curves0[idx+1].Step, step)
			curves0[idx+1].Data = append(curves0[idx+1].Data, pTot0[idx])
		} else {
			sum1 += amount / float64(len(plb))
			curves1[0].Step = append(curves1[0].Step, step)
			curves1[0].Data = append(curves1[0].Data, sum1)
			idx := nameToIdx[srcName]
			pTot1[idx] += amount
			curves1[idx+1].Step = append(curves1[idx+1].Step, step)
			curves1[idx+1].Data = append(curves1[idx+1].Data, pTot1[idx])
		}
	}

	return [][]chartCurve{curves0, curves1}
}

func finalizeStaticPies(curves0, curves1 []chartCurve, pla, plb []splitter.Player, pTot0, pTot1 []float64) {
	finalizeOneTeam(curves0, pla, pTot0)
	finalizeOneTeam(curves1, plb, pTot1)
}

func finalizeOneTeam(curves []chartCurve, players []splitter.Player, totals []float64) {
	var sum float64
	for _, v := range totals {
		sum += v
	}
	n := len(players)
	avg := 0.0
	if n > 0 {
		avg = sum / float64(n)
	}
	curves[0].Data = []float64{avg}
	curves[0].Step = []float64{0}
	for i := range players {
		curves[1+i].Data = []float64{totals[i]}
		curves[1+i].Step = []float64{0}
	}
}

func buildHealCharts(level *splitter.Level, lo, hi float64) [][]chartCurve {
	if level == nil || level.CombatLog == nil {
		return emptyChartsMetric(level, "heal")
	}
	pla, plb, nameToTeam, nameToIdx, ok := twoTeamRosters(level, false)
	if !ok {
		return emptyChartsMetric(level, "heal")
	}
	curves0 := newTeamCurves("Heal", pla)
	curves1 := newTeamCurves("Heal", plb)
	pTot0 := make([]float64, len(pla))
	pTot1 := make([]float64, len(plb))

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
		srcName := h.Initiator.Name
		tgtName := h.Recipient.Name
		st, okS := nameToTeam[srcName]
		tt, okT := nameToTeam[tgtName]
		// same-team heals only (matches JS store: heal.side === heal.other)
		if !okS || !okT || st != tt {
			continue
		}
		amount := float64(h.Heal)
		if st == 0 {
			idx := nameToIdx[srcName]
			pTot0[idx] += amount
		} else {
			idx := nameToIdx[srcName]
			pTot1[idx] += amount
		}
	}
	finalizeStaticPies(curves0, curves1, pla, plb, pTot0, pTot1)
	return [][]chartCurve{curves0, curves1}
}

func buildKillCharts(level *splitter.Level, lo, hi float64) [][]chartCurve {
	if level == nil || level.CombatLog == nil {
		return emptyChartsMetric(level, "kill")
	}
	pla, plb, nameToTeam, nameToIdx, ok := twoTeamRosters(level, false)
	if !ok {
		return emptyChartsMetric(level, "kill")
	}
	curves0 := newTeamCurves("Kill", pla)
	curves1 := newTeamCurves("Kill", plb)
	pTot0 := make([]float64, len(pla))
	pTot1 := make([]float64, len(plb))

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
		kName := k.Killer.Name
		vName := k.Killed.Name
		st, okK := nameToTeam[kName]
		vt, okV := nameToTeam[vName]
		if !okK || !okV || st == vt {
			continue
		}
		if st == 0 {
			idx := nameToIdx[kName]
			pTot0[idx]++
		} else {
			idx := nameToIdx[kName]
			pTot1[idx]++
		}
	}
	finalizeStaticPies(curves0, curves1, pla, plb, pTot0, pTot1)
	return [][]chartCurve{curves0, curves1}
}

func emptyCharts(level *splitter.Level) [][]chartCurve {
	return emptyChartsMetric(level, "damage")
}

func emptyChartsMetric(level *splitter.Level, metric string) [][]chartCurve {
	if level == nil {
		return [][]chartCurve{{}, {}}
	}
	prefix := "Damage"
	switch metric {
	case "heal":
		prefix = "Heal"
	case "kill":
		prefix = "Kill"
	}
	pla, plb, ok := chartTeams(level)
	if !ok {
		return [][]chartCurve{{}, {}}
	}
	return [][]chartCurve{newTeamCurves(prefix, pla), newTeamCurves(prefix, plb)}
}

func newTeamCurves(prefix string, players []splitter.Player) []chartCurve {
	out := make([]chartCurve, 1+len(players))
	out[0] = chartCurve{Name: "Average " + prefix, Data: []float64{}, Step: []float64{}}
	for i, p := range players {
		out[1+i] = chartCurve{Name: p.Name, Data: []float64{}, Step: []float64{}}
	}
	return out
}

// preGrowCurveSlices reduces reallocations while appending damage time-series points.
func preGrowCurveSlices(curves []chartCurve, capHint int) {
	if capHint < 8 {
		capHint = 8
	}
	if capHint > 16384 {
		capHint = 16384
	}
	for i := range curves {
		curves[i].Data = make([]float64, 0, capHint)
		curves[i].Step = make([]float64, 0, capHint)
	}
}

func (r *Runtime) marshalChartsJSON(ctx context.Context, level *splitter.Level, mode string, timeRangeJSON string) (s string, err error) {
	t0 := time.Now()
	heap0 := heapAlloc()
	defer func() {
		ob := 0
		if err == nil {
			ob = len(s)
		}
		logPerfWasm(ctx, r.lg, "marshal_charts_json", t0, heap0,
			"mode", mode,
			"out_size", prettyfmt.FormatBytes(uint64(ob)),
			"has_err", err != nil,
		)
	}()
	lo, hi := clampTimeRange(level, timeRangeJSON)
	var ch [][]chartCurve
	switch mode {
	case "heal":
		ch = buildHealCharts(level, lo, hi)
	case "kill":
		ch = buildKillCharts(level, lo, hi)
	default:
		ch = buildDamageCharts(level, lo, hi)
	}
	var b []byte
	b, err = json.Marshal(ch)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalChartsJSON", "err", err, "mode", mode)
		return "", fmt.Errorf("marshal charts: %w", err)
	}
	s = string(b)
	r.lg.For(ctx).Debugw("marshalChartsJSON", "mode", mode, "time_lo", lo, "time_hi", hi, "out_len", len(s))
	return s, nil
}
