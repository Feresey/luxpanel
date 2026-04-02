package main

import (
	"encoding/json"
	"fmt"
	"slices"

	"golang.org/x/exp/maps"

	"github.com/Feresey/luxpanel/internal/splitter"
)

// chartCurve matches the Vue/eta-chart structure (name, data, step).
type chartCurve struct {
	Name string    `json:"name"`
	Data []float64 `json:"data"`
	Step []float64 `json:"step"`
}

func twoTeamRosters(level *splitter.Level) (
	pla, plb []splitter.Player,
	nameToTeam map[string]int,
	nameToIdx map[string]int,
	ok bool,
) {
	if level == nil || len(level.Teams) == 0 {
		return nil, nil, nil, nil, false
	}
	teamIDs := maps.Keys(level.Teams)
	slices.Sort(teamIDs)
	teamIDs = filterNonZero(teamIDs)
	if len(teamIDs) < 2 {
		return nil, nil, nil, nil, false
	}
	pla = level.Teams[teamIDs[0]]
	plb = level.Teams[teamIDs[1]]
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

func buildDamageCharts(level *splitter.Level) [][]chartCurve {
	if level == nil || level.CombatLog == nil {
		return emptyCharts(level)
	}

	pla, plb, nameToTeam, nameToIdx, ok := twoTeamRosters(level)
	if !ok {
		return emptyCharts(level)
	}

	curves0 := newTeamCurves("Damage", pla)
	curves1 := newTeamCurves("Damage", plb)

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
		step := t.Sub(t0).Seconds()

		if st == 0 {
			sum0 += amount / float64(len(pla))
			curves0[0].Step = append(curves0[0].Step, step)
			curves0[0].Data = append(curves0[0].Data, sum0)
			idx := nameToIdx[srcName]
			pTot0[idx] += amount
			c := curves0[idx+1]
			c.Step = append(c.Step, step)
			c.Data = append(c.Data, pTot0[idx])
			curves0[idx+1] = c
		} else {
			sum1 += amount / float64(len(plb))
			curves1[0].Step = append(curves1[0].Step, step)
			curves1[0].Data = append(curves1[0].Data, sum1)
			idx := nameToIdx[srcName]
			pTot1[idx] += amount
			c := curves1[idx+1]
			c.Step = append(c.Step, step)
			c.Data = append(c.Data, pTot1[idx])
			curves1[idx+1] = c
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

func buildHealCharts(level *splitter.Level) [][]chartCurve {
	if level == nil || level.CombatLog == nil {
		return emptyChartsMetric(level, "heal")
	}
	pla, plb, nameToTeam, nameToIdx, ok := twoTeamRosters(level)
	if !ok {
		return emptyChartsMetric(level, "heal")
	}
	curves0 := newTeamCurves("Heal", pla)
	curves1 := newTeamCurves("Heal", plb)
	pTot0 := make([]float64, len(pla))
	pTot1 := make([]float64, len(plb))

	for _, h := range level.CombatLog.Heal {
		if h == nil || h.IsEmpty() {
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

func buildKillCharts(level *splitter.Level) [][]chartCurve {
	if level == nil || level.CombatLog == nil {
		return emptyChartsMetric(level, "kill")
	}
	pla, plb, nameToTeam, nameToIdx, ok := twoTeamRosters(level)
	if !ok {
		return emptyChartsMetric(level, "kill")
	}
	curves0 := newTeamCurves("Kill", pla)
	curves1 := newTeamCurves("Kill", plb)
	pTot0 := make([]float64, len(pla))
	pTot1 := make([]float64, len(plb))

	for _, k := range level.CombatLog.Kill {
		if k == nil || k.IsEmpty() {
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
	teamIDs := maps.Keys(level.Teams)
	slices.Sort(teamIDs)
	teamIDs = filterNonZero(teamIDs)
	if len(teamIDs) < 2 {
		return [][]chartCurve{{}, {}}
	}
	pla := level.Teams[teamIDs[0]]
	plb := level.Teams[teamIDs[1]]
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

func filterNonZero(ids []int) []int {
	var out []int
	for _, id := range ids {
		if id != 0 {
			out = append(out, id)
		}
	}
	return out
}

func marshalDamageCharts(level *splitter.Level) (string, error) {
	return marshalChartsJSON(level, "damage")
}

func marshalChartsJSON(level *splitter.Level, mode string) (string, error) {
	var ch [][]chartCurve
	switch mode {
	case "heal":
		ch = buildHealCharts(level)
	case "kill":
		ch = buildKillCharts(level)
	default:
		ch = buildDamageCharts(level)
	}
	b, err := json.Marshal(ch)
	if err != nil {
		return "", fmt.Errorf("marshal charts: %w", err)
	}
	return string(b), nil
}
