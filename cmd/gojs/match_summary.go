package main

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/Feresey/luxpanel/internal/splitter"
)

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
		}
	}

	b, err := json.Marshal(out)
	if err != nil {
		return "", err
	}
	r.lg.For(ctx).Debugw("marshalMatchSummaryJSON", "left", leftID, "right", rightID, "duration", out.DurationSec, "out_len", len(b))
	return string(b), nil
}

