package front

import (
	"strings"

	"github.com/Feresey/luxpanel/internal/splitter"
)

// MatchSummaryTeamAgg is per-team combat totals for summary cards.
type MatchSummaryTeamAgg struct {
	TeamID  int
	IsAlly  bool
	SpawnID int // 1=left spawn label, 2=right, 0=none
	Damage  float64
	Heal    float64
	Kills   int
}

// MatchSummaryComputed is numeric/geo summary; localized strings applied in cmd/gojs.
type MatchSummaryComputed struct {
	MapName       string
	GameMode      string
	StartTime     string
	DurationSec   float64
	TeamLeftID    int
	TeamRightID   int
	DamageLeft    float64
	DamageRight   float64
	HealLeft      float64
	HealRight     float64
	KillsLeft     int
	KillsRight    int
	TeamAggs      []MatchSummaryTeamAgg
	MultiTeamMode bool // len(sorted team ids) > 2
}

// BuildMatchSummary aggregates match-wide stats for chart-ordered teams.
func BuildMatchSummary(level *splitter.Level) MatchSummaryComputed {
	out := MatchSummaryComputed{}
	if level == nil {
		return out
	}
	out.DurationSec = LevelSpanSeconds(level)
	if !level.StartLevelTime.IsZero() {
		out.StartTime = level.StartLevelTime.Format("2006-01-02 15:04:05")
	}
	if cl := level.CombatLog; cl != nil {
		out.MapName = cl.Start.MapName
		out.GameMode = cl.Start.GameMode
	}
	leftID, rightID, leftPlayers, rightPlayers, ok := ChartTeamSides(level)
	if !ok {
		return out
	}
	out.TeamLeftID = leftID
	out.TeamRightID = rightID
	nameToSide := make(map[string]int, len(leftPlayers)+len(rightPlayers))
	for _, p := range leftPlayers {
		nameToSide[p.Name] = 0
	}
	for _, p := range rightPlayers {
		nameToSide[p.Name] = 1
	}
	teamIDs := SortedNonEmptyTeamIDs(level)
	allyID, _ := AllyTeamIDForTimeline(level)
	out.MultiTeamMode = len(teamIDs) > 2
	type teamAgg struct {
		damage float64
		heal   float64
		kills  int
	}
	agg := make(map[int]*teamAgg, len(teamIDs))
	for _, tid := range teamIDs {
		agg[tid] = &teamAgg{}
	}
	nameToTeamID := make(map[string]int)
	for _, tid := range teamIDs {
		for _, p := range level.Teams[tid] {
			n := strings.TrimSpace(p.Name)
			if n != "" {
				nameToTeamID[n] = tid
			}
		}
	}
	if level.CombatLog != nil {
		t0 := level.StartLevelTime
		for _, dmg := range level.CombatLog.Damage {
			if dmg == nil || dmg.IsEmpty() {
				continue
			}
			s, okS := nameToSide[dmg.Initiator.Name]
			t, okT := nameToSide[dmg.Recipient.Name]
			if !okS || !okT || s == t {
				continue
			}
			ts := dmg.GetTime(t0)
			if ts.Before(t0) {
				ts = t0
			}
			if !TimeInRangeFromStart(ts, t0, 0, out.DurationSec) {
				continue
			}
			if s == 0 {
				out.DamageLeft += float64(dmg.DamageFull)
			} else {
				out.DamageRight += float64(dmg.DamageFull)
			}
			if tid, ok := nameToTeamID[dmg.Initiator.Name]; ok {
				if a := agg[tid]; a != nil {
					a.damage += float64(dmg.DamageFull)
				}
			}
		}
		for _, h := range level.CombatLog.Heal {
			if h == nil || h.IsEmpty() {
				continue
			}
			s, okS := nameToSide[h.Initiator.Name]
			t, okT := nameToSide[h.Recipient.Name]
			if !okS || !okT || s != t {
				continue
			}
			if s == 0 {
				out.HealLeft += float64(h.Heal)
			} else {
				out.HealRight += float64(h.Heal)
			}
			if tid, ok := nameToTeamID[h.Initiator.Name]; ok {
				if a := agg[tid]; a != nil {
					a.heal += float64(h.Heal)
				}
			}
		}
		for _, k := range level.CombatLog.Kill {
			if k == nil || k.IsEmpty() {
				continue
			}
			s, okS := nameToSide[k.Killer.Name]
			t, okT := nameToSide[k.Killed.Name]
			if !okS || !okT || s == t {
				continue
			}
			if s == 0 {
				out.KillsLeft++
			} else {
				out.KillsRight++
			}
			if tid, ok := nameToTeamID[k.Killer.Name]; ok {
				if a := agg[tid]; a != nil {
					a.kills++
				}
			}
		}
	}
	for _, tid := range teamIDs {
		a := agg[tid]
		if a == nil {
			continue
		}
		spawnID := 0
		if tid == 1 {
			spawnID = 1
		} else if tid == 2 {
			spawnID = 2
		}
		out.TeamAggs = append(out.TeamAggs, MatchSummaryTeamAgg{
			TeamID:  tid,
			IsAlly:  tid == allyID,
			SpawnID: spawnID,
			Damage:  a.damage,
			Heal:    a.heal,
			Kills:   a.kills,
		})
	}
	return out
}
