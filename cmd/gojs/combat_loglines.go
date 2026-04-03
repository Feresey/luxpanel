//go:build js

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/Feresey/luxpanel/internal/splitter"
)

type combatLogLineRow struct {
	Time    string `json:"time"`
	Kind    string `json:"kind"`
	Summary string `json:"summary"`
}

func (r *Runtime) marshalCombatLogLinesJSON(ctx context.Context, level *splitter.Level, timeRangeJSON string) (string, error) {
	if level == nil || level.CombatLog == nil {
		r.lg.For(ctx).Debugw("marshalCombatLogLinesJSON", "empty", true)
		return "[]", nil
	}
	cl := level.CombatLog
	lo, hi := clampTimeRange(level, timeRangeJSON)
	t0 := level.StartLevelTime
	rows := make([]combatLogLineRow, 0)
	for _, ln := range cl.LogLines {
		t := ln.GetTime(t0)
		if t.Before(t0) {
			t = t0
		}

		res := !timeInRangeFromStart(t, t0, lo, hi)
		if res {
			continue
		}
		kind, summary := combatLogLineKindSummary(ln)
		rows = append(rows, combatLogLineRow{
			Time:    t.Format("15:04:05.000"),
			Kind:    kind,
			Summary: summary,
		})
	}
	b, err := json.Marshal(rows)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalCombatLogLinesJSON", "err", err, "lines", len(cl.LogLines))
		return "", err
	}
	s := string(b)
	r.lg.For(ctx).Debugw("marshalCombatLogLinesJSON", "lines_in", len(cl.LogLines), "lines_out", len(rows),
		"time_lo", lo, "time_hi", hi, "out_len", len(s))
	return s, nil
}

func joinDamageMods(mods combat.DamageModifiers) string {
	if len(mods) == 0 {
		return ""
	}
	parts := make([]string, len(mods))
	for i, m := range mods {
		parts[i] = string(m)
	}
	return strings.Join(parts, ", ")
}

func joinParticipationMods(mods combat.ParticipationModifiers) string {
	if len(mods) == 0 {
		return ""
	}
	parts := make([]string, len(mods))
	for i, m := range mods {
		parts[i] = string(m)
	}
	return strings.Join(parts, ", ")
}

func combatLogLineKindSummary(ln combat.LogLine) (kind string, summary string) {
	switch x := ln.(type) {
	case *combat.ConnectToGameSession:
		return "ConnectToGameSession", fmt.Sprintf("SessionID: %d", x.SessionID)
	case *combat.Start:
		return "Start", fmt.Sprintf("What: %q · GameMode: %q · MapName: %q · LocalClientTeamID: %d",
			x.What, x.GameMode, x.MapName, x.LocalClientTeamID)
	case *combat.Finished:
		return "Finished", fmt.Sprintf("WinnerTeamID: %d · WinReason: %q · FinishReason: %q · GameTime: %.2f",
			x.WinnerTeamID, x.WinReason, x.FinishReason, x.GameTime)
	case *combat.Reward:
		return "Reward", fmt.Sprintf("Recipient: %q · Ship: %q · Reward: %d · RewardType: %q · Reason: %q",
			x.Recipient, x.Ship, x.Reward, x.RewardType, x.Reason)
	case *combat.Damage:
		mods := joinDamageMods(x.DamageModifiers)
		s := summarizeObject("Initiator", x.Initiator) + " · " + summarizeObject("Recipient", x.Recipient) +
			fmt.Sprintf(" · Source: %q · DamageFull: %.4f · DamageHull: %.4f · DamageShield: %.4f",
				x.Source, x.DamageFull, x.DamageHull, x.DamageShield)
		if mods != "" {
			s += " · DamageModifiers: [" + mods + "]"
		}
		s += fmt.Sprintf(" · FriendlyFire: %t · Rocket: %d", x.FriendlyFire, x.Rocket)
		return "Damage", s
	case *combat.Heal:
		return "Heal", summarizeObject("Initiator", x.Initiator) + " · " + summarizeObject("Recipient", x.Recipient) +
			fmt.Sprintf(" · Source: %q · Heal: %.4f", x.Source, x.Heal)
	case *combat.Kill:
		s := summarizeObject("Killer", x.Killer) + " · " + summarizeObject("Killed", x.Killed) +
			fmt.Sprintf(" · Source: %q · FriendlyFire: %t", x.Source, x.FriendlyFire)
		return "Kill", s
	case *combat.Participant:
		mods := joinParticipationMods(x.Modifiers)
		s := fmt.Sprintf("Name: %q · Ship: %q · Damage: %.4f · MostDamageWith: %q · FriendlyFire: %t",
			x.Name, x.Ship, x.Damage, x.MostDamageWith, x.FriendlyFire)
		if mods != "" {
			s += " · Modifiers: [" + mods + "]"
		}
		return "Participant", s
	case *combat.Spawn:
		return "Spawn", fmt.Sprintf("Name: %q · ID: %d · Ship: %q", x.Name, x.ID, x.Ship)
	default:
		return fmt.Sprintf("%T", ln), ""
	}
}

func summarizeObject(prefix string, o combat.Object) string {
	return fmt.Sprintf("%s.Name: %q · %s.ObjectName: %q · %s.ObjectOwner: %q · %s.ObjectID: %d",
		prefix, o.Name, prefix, o.ObjectName, prefix, o.ObjectOwner, prefix, o.ObjectID)
}
