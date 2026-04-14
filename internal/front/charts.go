package front

import (
	"time"

	"github.com/Feresey/luxpanel/internal/splitter"
)

type ChartCurve struct {
	Name      string    `json:"name"`
	Data      []float64 `json:"data"`
	Step      []float64 `json:"step"`
	CurveKind string    `json:"-"`
	Metric    string    `json:"-"`
}

func NewTeamCurves(metric string, players []splitter.Player) []ChartCurve {
	out := make([]ChartCurve, 1+len(players))
	out[0] = ChartCurve{Name: "", CurveKind: "avg", Metric: metric, Data: []float64{}, Step: []float64{}}
	for i, p := range players {
		out[1+i] = ChartCurve{Name: p.Name, CurveKind: "player", Metric: metric, Data: []float64{}, Step: []float64{}}
	}
	return out
}

func preGrowCurveSlices(curves []ChartCurve, capHint int) {
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

func finalizeStaticPies(curves0, curves1 []ChartCurve, pla, plb []splitter.Player, pTot0, pTot1 []float64) {
	finalizeOneTeam(curves0, pla, pTot0)
	finalizeOneTeam(curves1, plb, pTot1)
}

func finalizeOneTeam(curves []ChartCurve, players []splitter.Player, totals []float64) {
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

func BuildDamageCharts(level *splitter.Level, lo, hi float64) [][]ChartCurve {
	if level == nil || level.CombatLog == nil {
		return emptyChartsMetric(level, "damage")
	}
	pla, plb, nameToTeam, nameToIdx, ok := TwoTeamRosters(level, false)
	if !ok {
		return emptyChartsMetric(level, "damage")
	}
	curves0 := NewTeamCurves("damage", pla)
	curves1 := NewTeamCurves("damage", plb)
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
		if !TimeInRangeFromStart(t, t0, lo, hi) {
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
	return [][]ChartCurve{curves0, curves1}
}

func BuildHealCharts(level *splitter.Level, lo, hi float64) [][]ChartCurve {
	if level == nil || level.CombatLog == nil {
		return emptyChartsMetric(level, "heal")
	}
	pla, plb, nameToTeam, nameToIdx, ok := TwoTeamRosters(level, false)
	if !ok {
		return emptyChartsMetric(level, "heal")
	}
	curves0 := NewTeamCurves("heal", pla)
	curves1 := NewTeamCurves("heal", plb)
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
		if !TimeInRangeFromStart(t, t0, lo, hi) {
			continue
		}
		srcName := h.Initiator.Name
		tgtName := h.Recipient.Name
		st, okS := nameToTeam[srcName]
		tt, okT := nameToTeam[tgtName]
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
	return [][]ChartCurve{curves0, curves1}
}

func BuildKillCharts(level *splitter.Level, lo, hi float64) [][]ChartCurve {
	if level == nil || level.CombatLog == nil {
		return emptyChartsMetric(level, "kill")
	}
	pla, plb, nameToTeam, nameToIdx, ok := TwoTeamRosters(level, false)
	if !ok {
		return emptyChartsMetric(level, "kill")
	}
	curves0 := NewTeamCurves("kill", pla)
	curves1 := NewTeamCurves("kill", plb)
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
		if !TimeInRangeFromStart(t, t0, lo, hi) {
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
	return [][]ChartCurve{curves0, curves1}
}

func emptyChartsMetric(level *splitter.Level, metric string) [][]ChartCurve {
	if level == nil {
		return [][]ChartCurve{{}, {}}
	}
	pla, plb, ok := ChartTeams(level)
	if !ok {
		return [][]ChartCurve{{}, {}}
	}
	return [][]ChartCurve{NewTeamCurves(metric, pla), NewTeamCurves(metric, plb)}
}

func BuildChartsForMode(level *splitter.Level, mode string, lo, hi float64) [][]ChartCurve {
	switch mode {
	case "heal":
		return BuildHealCharts(level, lo, hi)
	case "kill":
		return BuildKillCharts(level, lo, hi)
	default:
		return BuildDamageCharts(level, lo, hi)
	}
}
