package front

import (
	"encoding/json"
	"strings"

	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/Feresey/luxpanel/internal/splitter"
)

// TeamPairPanel: columns = col_players, rows = row_players.
type TeamPairPanel struct {
	ColPlayers       []string       `json:"col_players"`
	RowPlayers       []string       `json:"row_players"`
	NewbieBonusMax   map[string]int `json:"newbie_bonus_max,omitempty"`
	Matrix           [][]float64    `json:"matrix"`
	Total            float64        `json:"total"`
}

// ChartMatricesResult is JSON for chart matrix widgets.
type ChartMatricesResult struct {
	Metric string           `json:"metric"`
	Panels []TeamPairPanel  `json:"panels"`
}

type ChartQueryOptions struct {
	DamageType  string `json:"damage_type,omitempty"`
	IncludeBots bool   `json:"include_bots,omitempty"`
}

func ParseChartQueryOptions(raw string) ChartQueryOptions {
	opts := ChartQueryOptions{DamageType: "all", IncludeBots: false}
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

// MaxNewbieBonusByPlayer maps canonical player name → max Normalizer from Spell events.
func MaxNewbieBonusByPlayer(level *splitter.Level, players []splitter.Player) map[string]int {
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

// BuildChartMatrices returns matrix payload for mode (damage|heal|kill).
func BuildChartMatrices(level *splitter.Level, mode string, lo, hi float64, opts ChartQueryOptions) ChartMatricesResult {
	if level == nil {
		return ChartMatricesResult{}
	}
	switch mode {
	case "heal":
		return buildHealMatrices(level, lo, hi, opts)
	case "kill":
		return buildKillMatrices(level, lo, hi, opts)
	default:
		return buildDamageMatrices(level, lo, hi, opts)
	}
}

func buildDamageMatrices(level *splitter.Level, lo, hi float64, opts ChartQueryOptions) ChartMatricesResult {
	if level.CombatLog == nil {
		return ChartMatricesResult{Metric: "damage"}
	}
	pla, plb, nameToTeam, nameToIdx, ok := TwoTeamRosters(level, opts.IncludeBots)
	if !ok {
		return ChartMatricesResult{Metric: "damage"}
	}
	na, nb := len(pla), len(plb)
	allPlayers := append(append([]splitter.Player{}, pla...), plb...)
	newbieMax := MaxNewbieBonusByPlayer(level, allPlayers)
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
		if !TimeInRangeFromStart(t, t0, lo, hi) {
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
	p0 := TeamPairPanel{
		ColPlayers:     playerNameSlice(pla),
		RowPlayers:     playerNameSlice(plb),
		NewbieBonusMax: newbieMax,
		Matrix:         m0,
		Total:          sumMatrix(m0),
	}
	p1 := TeamPairPanel{
		ColPlayers:     playerNameSlice(plb),
		RowPlayers:     playerNameSlice(pla),
		NewbieBonusMax: newbieMax,
		Matrix:         m1,
		Total:          sumMatrix(m1),
	}
	return ChartMatricesResult{Metric: "damage", Panels: []TeamPairPanel{p0, p1}}
}

func buildKillMatrices(level *splitter.Level, lo, hi float64, opts ChartQueryOptions) ChartMatricesResult {
	if level.CombatLog == nil {
		return ChartMatricesResult{Metric: "kill"}
	}
	pla, plb, nameToTeam, nameToIdx, ok := TwoTeamRosters(level, opts.IncludeBots)
	if !ok {
		return ChartMatricesResult{Metric: "kill"}
	}
	na, nb := len(pla), len(plb)
	allPlayers := append(append([]splitter.Player{}, pla...), plb...)
	newbieMax := MaxNewbieBonusByPlayer(level, allPlayers)
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
		if !TimeInRangeFromStart(t, t0, lo, hi) {
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
	p0 := TeamPairPanel{
		ColPlayers:     playerNameSlice(pla),
		RowPlayers:     playerNameSlice(plb),
		NewbieBonusMax: newbieMax,
		Matrix:         m0,
		Total:          sumMatrix(m0),
	}
	p1 := TeamPairPanel{
		ColPlayers:     playerNameSlice(plb),
		RowPlayers:     playerNameSlice(pla),
		NewbieBonusMax: newbieMax,
		Matrix:         m1,
		Total:          sumMatrix(m1),
	}
	return ChartMatricesResult{Metric: "kill", Panels: []TeamPairPanel{p0, p1}}
}

func buildHealMatrices(level *splitter.Level, lo, hi float64, opts ChartQueryOptions) ChartMatricesResult {
	if level.CombatLog == nil {
		return ChartMatricesResult{Metric: "heal"}
	}
	pla, plb, nameToTeam, nameToIdx, ok := TwoTeamRosters(level, opts.IncludeBots)
	if !ok {
		return ChartMatricesResult{Metric: "heal"}
	}
	na, nb := len(pla), len(plb)
	allPlayers := append(append([]splitter.Player{}, pla...), plb...)
	newbieMax := MaxNewbieBonusByPlayer(level, allPlayers)
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
		if !TimeInRangeFromStart(t, t0, lo, hi) {
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
	p0 := TeamPairPanel{
		ColPlayers:     playerNameSlice(pla),
		RowPlayers:     playerNameSlice(pla),
		NewbieBonusMax: newbieMax,
		Matrix:         m0,
		Total:          sumMatrix(m0),
	}
	p1 := TeamPairPanel{
		ColPlayers:     playerNameSlice(plb),
		RowPlayers:     playerNameSlice(plb),
		NewbieBonusMax: newbieMax,
		Matrix:         m1,
		Total:          sumMatrix(m1),
	}
	return ChartMatricesResult{Metric: "heal", Panels: []TeamPairPanel{p0, p1}}
}
