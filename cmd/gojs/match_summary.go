package main

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/Feresey/luxpanel/internal/splitter"
)

type matchSummaryTeam struct {
	TeamID int     `json:"team_id"`
	Name   string  `json:"name"`
	Role   string  `json:"role"`
	Damage float64 `json:"damage"`
	Heal   float64 `json:"heal"`
	Kills  int     `json:"kills"`
}

type matchSummaryResult struct {
	MapName      string  `json:"map_name,omitempty"`
	GameMode     string  `json:"game_mode,omitempty"`
	StartTime    string  `json:"start_time,omitempty"`
	DurationSec  float64 `json:"duration_sec"`
	TeamLeftID   int     `json:"team_left_id,omitempty"`
	TeamRightID  int     `json:"team_right_id,omitempty"`
	TeamLeftName string  `json:"team_left_name,omitempty"`
	TeamRightName string `json:"team_right_name,omitempty"`
	DamageLeft   float64 `json:"damage_left"`
	DamageRight  float64 `json:"damage_right"`
	HealLeft     float64 `json:"heal_left"`
	HealRight    float64 `json:"heal_right"`
	KillsLeft    int     `json:"kills_left"`
	KillsRight   int     `json:"kills_right"`
	Teams        []matchSummaryTeam `json:"teams,omitempty"`
}

func (r *Runtime) marshalMatchSummaryJSON(ctx context.Context, level *splitter.Level) (string, error) {
	if level == nil {
		b, err := json.Marshal(matchSummaryResult{})
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	out := matchSummaryResult{
		DurationSec: levelSpanSeconds(level),
	}
	if !level.StartLevelTime.IsZero() {
		out.StartTime = level.StartLevelTime.Format("2006-01-02 15:04:05")
	}
	if cl := level.CombatLog; cl != nil {
		out.MapName = cl.Start.MapName
		out.GameMode = cl.Start.GameMode
	}

	leftID, rightID, leftPlayers, rightPlayers, ok := chartTeamSides(level)
	if !ok {
		b, err := json.Marshal(out)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	out.TeamLeftID = leftID
	out.TeamRightID = rightID
	out.TeamLeftName = "Team " + strconv.Itoa(leftID)
	out.TeamRightName = "Team " + strconv.Itoa(rightID)

	nameToSide := make(map[string]int, len(leftPlayers)+len(rightPlayers))
	for _, p := range leftPlayers {
		nameToSide[p.Name] = 0
	}
	for _, p := range rightPlayers {
		nameToSide[p.Name] = 1
	}
	teamIDs := sortedNonEmptyTeamIDs(level)
	allyID, _ := allyTeamIDForTimeline(level)
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
			if !timeInRangeFromStart(ts, t0, 0, out.DurationSec) {
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
		role := "Враги"
		if tid == allyID {
			role = "Союзники"
		} else if len(teamIDs) > 2 {
			role = "Враги (Team " + strconv.Itoa(tid) + ")"
		}
		side := ""
		if tid == 1 {
			side = " · левый спавн"
		} else if tid == 2 {
			side = " · правый спавн"
		}
		out.Teams = append(out.Teams, matchSummaryTeam{
			TeamID: tid,
			Name:   "Team " + strconv.Itoa(tid),
			Role:   role + side,
			Damage: a.damage,
			Heal:   a.heal,
			Kills:  a.kills,
		})
	}

	b, err := json.Marshal(out)
	if err != nil {
		return "", err
	}
	r.lg.For(ctx).Debugw("marshalMatchSummaryJSON", "left", leftID, "right", rightID, "duration", out.DurationSec, "out_len", len(b))
	return string(b), nil
}

