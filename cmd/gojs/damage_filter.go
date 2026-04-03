package main

import (
	"context"
	"encoding/json"
	"sort"
	"strings"
	"time"

	"github.com/Feresey/luxpanel/internal/damagefilters"
	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/Feresey/luxpanel/internal/splitter"
)

// damageFilterMeta lists unique values for building filter UI.
type damageFilterMeta struct {
	Players    []string `json:"players"`
	Recipients []string `json:"recipients"`
	Weapons    []string `json:"weapons"`
	Modifiers  []string `json:"modifiers"`
}

type wasmDamageFilterRequest struct {
	Initiator   string          `json:"initiator"`
	Recipient   string          `json:"recipient"`
	Weapon      string          `json:"weapon"`
	Modifiers   map[string]bool `json:"modifiers,omitempty"`
	Aggregate   *bool           `json:"aggregate,omitempty"`
	TimeFromSec *float64        `json:"time_from_sec,omitempty"`
	TimeToSec   *float64        `json:"time_to_sec,omitempty"`
}

type damageTableSummary struct {
	Hits   int     `json:"hits"`
	Damage float64 `json:"damage"`
}

type damageTableRow struct {
	Source    string             `json:"source"`
	Targets   []string           `json:"targets"`
	Modifiers map[string]bool    `json:"modifiers"`
	Summary   damageTableSummary `json:"summary"`
}

type damageTableResult struct {
	Row    damageTableRow   `json:"row"`
	Events []damageEventRow `json:"events,omitempty"`
	Error  string           `json:"error,omitempty"`
}

type damageEventRow struct {
	TimeSec   float64         `json:"time_sec"`
	Initiator string          `json:"initiator"`
	Recipient string          `json:"recipient"`
	Weapon    string          `json:"weapon"`
	Modifiers map[string]bool `json:"modifiers"`
	Amount    float64         `json:"amount"`
}

func (r *Runtime) marshalDamageFilterMetaJSON(ctx context.Context, level *splitter.Level, initiator string, timeRangeJSON string) (string, error) {
	lo, hi := clampTimeRange(level, timeRangeJSON)
	meta := collectDamageFilterMeta(level, initiator, lo, hi)
	b, err := json.Marshal(meta)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalDamageFilterMetaJSON", "err", err, "initiator", initiator)
		return "", err
	}
	s := string(b)
	r.lg.For(ctx).Debugw("marshalDamageFilterMetaJSON", "initiator", initiator, "time_lo", lo, "time_hi", hi,
		"players", len(meta.Players), "recipients", len(meta.Recipients), "weapons", len(meta.Weapons), "out_len", len(s))
	return s, nil
}

func collectDamageFilterMeta(level *splitter.Level, initiator string, lo, hi float64) damageFilterMeta {
	humans := humanPlayerNames(level)
	out := damageFilterMeta{
		Players:    []string{},
		Recipients: []string{},
		Weapons:    []string{},
		Modifiers:  []string{},
	}
	if level == nil || level.CombatLog == nil {
		return out
	}
	players := make(map[string]struct{})
	weapons := make(map[string]struct{})
	mods := make(map[string]struct{})
	recipients := make(map[string]struct{})
	t0 := level.StartLevelTime

	for name := range humans {
		players[name] = struct{}{}
	}

	for _, dmg := range level.CombatLog.Damage {
		if dmg == nil {
			continue
		}
		t := dmg.GetTime(t0)
		if t.Before(t0) {
			t = t0
		}
		if !timeInRangeFromStart(t, t0, lo, hi) {
			continue
		}
		if strings.TrimSpace(initiator) != "" && dmg.Initiator.Name == initiator {
			if s := strings.TrimSpace(dmg.Source); s != "" {
				weapons[s] = struct{}{}
			}
			for _, m := range dmg.DamageModifiers {
				mods[string(m)] = struct{}{}
			}
			if dmg.Recipient.Name != "" {
				if _, ok := humans[dmg.Recipient.Name]; !ok {
					continue
				}
				recipients[dmg.Recipient.Name] = struct{}{}
			}
		}
	}

	out.Players = keysSorted(players)
	out.Recipients = keysSorted(recipients)
	out.Weapons = keysSorted(weapons)
	out.Modifiers = keysSorted(mods)
	return out
}

func humanAndBotPlayerNames(level *splitter.Level) (map[string]struct{}, map[string]struct{}) {
	humans := make(map[string]struct{})
	bots := make(map[string]struct{})
	if level == nil || level.GameLog == nil {
		return humans, bots
	}
	for _, p := range level.GameLog.AddPlayer {
		if p == nil {
			continue
		}
		name := strings.TrimSpace(p.Name)
		if name == "" {
			continue
		}
		// User request: real players have PlayerID != 0 in game log.
		if p.PlayerID != 0 {
			humans[name] = struct{}{}
		} else {
			bots[name] = struct{}{}
		}
	}
	return humans, bots
}

func humanPlayerNames(level *splitter.Level) map[string]struct{} {
	humans, _ := humanAndBotPlayerNames(level)
	return humans
}

func keysSorted(m map[string]struct{}) []string {
	s := make([]string, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	sort.Strings(s)
	return s
}

func (r *Runtime) marshalDamageTableJSON(ctx context.Context, level *splitter.Level, filterJSON string) (string, error) {
	if level == nil || level.CombatLog == nil {
		b, err := json.Marshal(damageTableResult{
			Row: damageTableRow{
				Source:    "",
				Targets:   []string{},
				Modifiers: map[string]bool{},
				Summary:   damageTableSummary{},
			},
		})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalDamageTableJSON", "err", err, "reason", "empty_level")
			return "", err
		}
		s := string(b)
		r.lg.For(ctx).Debugw("marshalDamageTableJSON", "empty_level", true, "out_len", len(s))
		return s, nil
	}

	var req wasmDamageFilterRequest
	if err := json.Unmarshal([]byte(filterJSON), &req); err != nil {
		res := damageTableResult{Error: "invalid filter json"}
		b, mErr := json.Marshal(res)
		if mErr != nil {
			r.lg.For(ctx).Errorw("marshalDamageTableJSON", "err", mErr, "reason", "invalid_filter_marshal")
			return "", mErr
		}
		r.lg.For(ctx).Debugw("marshalDamageTableJSON", "client_error", "invalid filter json", "out_len", len(b))
		return string(b), nil
	}
	if strings.TrimSpace(req.Initiator) == "" {
		b, err := json.Marshal(damageTableResult{
			Row:   damageTableRow{Source: "", Targets: []string{}, Modifiers: map[string]bool{}, Summary: damageTableSummary{}},
			Error: "initiator required",
		})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalDamageTableJSON", "err", err)
			return "", err
		}
		r.lg.For(ctx).Debugw("marshalDamageTableJSON", "client_error", "initiator required", "out_len", len(b))
		return string(b), nil
	}

	humans := humanPlayerNames(level)
	if _, ok := humans[req.Initiator]; !ok {
		b, err := json.Marshal(damageTableResult{
			Row:   damageTableRow{Source: "", Targets: []string{}, Modifiers: map[string]bool{}, Summary: damageTableSummary{}},
			Error: "initiator must be a player",
		})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalDamageTableJSON", "err", err)
			return "", err
		}
		r.lg.For(ctx).Debugw("marshalDamageTableJSON", "client_error", "initiator must be a player", "out_len", len(b))
		return string(b), nil
	}

	lo, hi := applyTimeBoundsToReq(level, req.TimeFromSec, req.TimeToSec)
	row := filterDamageRow(level, &req, humans, lo, hi)
	events := collectDamageEvents(level, &req, humans, lo, hi)
	res := damageTableResult{
		Row:    row,
		Events: events,
	}
	b, err := json.Marshal(res)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalDamageTableJSON", "err", err, "initiator", req.Initiator)
		return "", err
	}
	s := string(b)
	r.lg.For(ctx).Debugw("marshalDamageTableJSON", "initiator", req.Initiator, "hits", row.Summary.Hits, "events", len(events), "out_len", len(s))
	return s, nil
}

func collectDamageEvents(level *splitter.Level, req *wasmDamageFilterRequest, humans map[string]struct{}, lo, hi float64) []damageEventRow {
	if level == nil || level.CombatLog == nil {
		return []damageEventRow{}
	}
	t0 := level.StartLevelTime
	filter := toDamageFilter(req)

	out := make([]damageEventRow, 0)
	for _, dmg := range level.CombatLog.Damage {
		if dmg == nil || dmg.IsEmpty() {
			continue
		}
		detailed, ok := filter.Filter(dmg)
		if !ok {
			continue
		}
		if _, ok := humans[dmg.Recipient.Name]; !ok {
			continue
		}

		t := dmg.GetTime(t0)
		if t.Before(t0) {
			t = t0
		}
		if !timeInRangeFromStart(t, t0, lo, hi) {
			continue
		}
		sec := float64(t.Sub(t0)) / float64(time.Second)

		amount := float64(detailed.Damage)

		mods := make(map[string]bool, len(dmg.DamageModifiers))
		for _, m := range dmg.DamageModifiers {
			mods[string(m)] = true
		}
		out = append(out, damageEventRow{
			TimeSec:   sec,
			Initiator: dmg.Initiator.Name,
			Recipient: dmg.Recipient.Name,
			Weapon:    dmg.Source,
			Modifiers: mods,
			Amount:    amount,
		})
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i].TimeSec != out[j].TimeSec {
			return out[i].TimeSec < out[j].TimeSec
		}
		return out[i].Recipient < out[j].Recipient
	})
	return out
}

func filterDamageRow(level *splitter.Level, req *wasmDamageFilterRequest, humans map[string]struct{}, lo, hi float64) damageTableRow {
	var hits int
	var damageSum float64

	targetSet := make(map[string]struct{})
	needTargets := strings.TrimSpace(req.Recipient) == ""
	requestedRecipient := strings.TrimSpace(req.Recipient)
	filter := toDamageFilter(req)
	t0 := level.StartLevelTime

	for _, dmg := range level.CombatLog.Damage {
		if dmg == nil || dmg.IsEmpty() {
			continue
		}
		t := dmg.GetTime(t0)
		if t.Before(t0) {
			t = t0
		}
		if !timeInRangeFromStart(t, t0, lo, hi) {
			continue
		}
		detailed, ok := filter.Filter(dmg)
		if !ok {
			continue
		}

		selected := float64(detailed.Damage)

		hits++
		damageSum += selected

		if needTargets {
			if dmg.Recipient.Name != "" {
				if _, ok := humans[dmg.Recipient.Name]; !ok {
					continue
				}
				targetSet[dmg.Recipient.Name] = struct{}{}
			}
		}
	}

	targets := []string{}
	if requestedRecipient != "" {
		// UI should prevent bots, but keep the check for safety.
		if _, ok := humans[requestedRecipient]; ok {
			targets = []string{requestedRecipient}
		}
	} else {
		targets = keysSorted(targetSet)
	}

	mods := req.Modifiers
	if mods == nil {
		mods = map[string]bool{}
	}

	return damageTableRow{
		Source:    req.Initiator,
		Targets:   targets,
		Modifiers: mods,
		Summary: damageTableSummary{
			Hits:   hits,
			Damage: damageSum,
		},
	}
}

func toDamageFilter(req *wasmDamageFilterRequest) damagefilters.PlayerDamageFilterConfig {
	mods := damagefilters.DamageModifiersMap{}
	for k, v := range req.Modifiers {
		mods[combat.DamageModifier(k)] = v
	}
	return damagefilters.PlayerDamageFilterConfig{
		InitiatorName:   strings.TrimSpace(req.Initiator),
		RecipientName:   strings.TrimSpace(req.Recipient),
		DamageType:      damagefilters.DamageTypeTotal,
		DamageModifiers: mods,
		Weapon:          strings.TrimSpace(req.Weapon),
	}
}

func matchDamageModifiers(have []combat.DamageModifier, want map[string]bool) bool {
	if len(want) == 0 {
		return true
	}

	present := make(map[string]struct{}, len(have))
	for _, h := range have {
		present[string(h)] = struct{}{}
	}

	for mod, shouldExist := range want {
		_, ok := present[mod]
		if shouldExist && !ok {
			return false
		}
		if !shouldExist && ok {
			return false
		}
	}
	return true
}
