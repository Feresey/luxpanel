//go:build js

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Feresey/luxpanel/internal/damagefilters"
	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/Feresey/luxpanel/internal/splitter"
)

type damageDefaultTableBaseRequest struct {
	Initiator   string   `json:"initiator"`
	Recipient   string   `json:"recipient"`
	Weapon      string   `json:"weapon"`
	TimeFromSec *float64 `json:"time_from_sec,omitempty"`
	TimeToSec   *float64 `json:"time_to_sec,omitempty"`
}

type damageDefaultTableRow struct {
	Source    string              `json:"source"`
	Targets   []string            `json:"targets"`
	Modifiers map[string]bool     `json:"modifiers"`
	Summary   damageTableSummary  `json:"summary"`
	Events    []damageEventRow    `json:"events,omitempty"`
}

type damageDefaultTableResult struct {
	Rows  []damageDefaultTableRow `json:"rows"`
	Error string                  `json:"error,omitempty"`
}

func defaultDamageModifierFilters() []map[string]bool {
	src := damagefilters.DefaultModifierFilters()
	out := make([]map[string]bool, 0, len(src))
	for _, m := range src {
		dst := make(map[string]bool, len(m))
		for k, v := range m {
			dst[string(k)] = v
		}
		out = append(out, dst)
	}
	return out
}

func (r *Runtime) marshalDamageDefaultFiltersTableJSON(ctx context.Context, level *splitter.Level, filterJSON string) (string, error) {
	var req damageDefaultTableBaseRequest
	if err := json.Unmarshal([]byte(filterJSON), &req); err != nil {
		s, err := jsonStringify(damageDefaultTableResult{Error: "invalid filter json"})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalDamageDefaultFiltersTableJSON", "err", err)
			return "", err
		}
		r.lg.For(ctx).Debugw("marshalDamageDefaultFiltersTableJSON", "client_error", "invalid filter json", "out_len", len(s))
		return s, nil
	}
	if strings.TrimSpace(req.Initiator) == "" {
		s, err := jsonStringify(damageDefaultTableResult{Error: "initiator required"})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalDamageDefaultFiltersTableJSON", "err", err)
			return "", err
		}
		r.lg.For(ctx).Debugw("marshalDamageDefaultFiltersTableJSON", "client_error", "initiator required", "out_len", len(s))
		return s, nil
	}

	if level == nil || level.CombatLog == nil {
		emptyRows := make([]damageDefaultTableRow, 0)
		s, err := jsonStringify(damageDefaultTableResult{Rows: emptyRows})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalDamageDefaultFiltersTableJSON", "err", err)
			return "", err
		}
		r.lg.For(ctx).Debugw("marshalDamageDefaultFiltersTableJSON", "empty_level", true, "out_len", len(s))
		return s, nil
	}

	humans := humanPlayerNames(level)
	if _, ok := humans[req.Initiator]; !ok {
		s, err := jsonStringify(damageDefaultTableResult{Error: "initiator must be a player"})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalDamageDefaultFiltersTableJSON", "err", err)
			return "", err
		}
		r.lg.For(ctx).Debugw("marshalDamageDefaultFiltersTableJSON", "client_error", "initiator must be a player", "out_len", len(s))
		return s, nil
	}

	lo, hi := applyTimeBoundsToReq(level, req.TimeFromSec, req.TimeToSec)

	modFilters := defaultDamageModifierFilters()
	rows := make([]damageDefaultTableRow, 0, len(modFilters))
	for _, mods := range modFilters {
		summary, targets := filterDamageRowWithModifiers(level, req, mods, humans, lo, hi)
		wasmReq := wasmDamageFilterRequest{
			Initiator:   req.Initiator,
			Recipient:   req.Recipient,
			Weapon:      req.Weapon,
			Modifiers:   mods,
			TimeFromSec: req.TimeFromSec,
			TimeToSec:   req.TimeToSec,
		}
		rows = append(rows, damageDefaultTableRow{
			Source:    req.Initiator,
			Targets:   targets,
			Modifiers: mods,
			Summary:   summary,
			Events:    collectDamageEvents(level, &wasmReq, humans, lo, hi),
		})
	}

	s, err := jsonStringify(damageDefaultTableResult{Rows: rows})
	if err != nil {
		r.lg.For(ctx).Errorw("marshalDamageDefaultFiltersTableJSON", "err", err, "initiator", req.Initiator, "rows", len(rows))
		return "", err
	}
	r.lg.For(ctx).Debugw("marshalDamageDefaultFiltersTableJSON", "initiator", req.Initiator, "rows", len(rows), "time_lo", lo, "time_hi", hi, "out_len", len(s))
	return s, nil
}

func filterDamageRowWithModifiers(level *splitter.Level, req damageDefaultTableBaseRequest, modifiers map[string]bool, humans map[string]struct{}, lo, hi float64) (damageTableSummary, []string) {
	cfg := damagefilters.PlayerDamageFilterConfig{
		InitiatorName: strings.TrimSpace(req.Initiator),
		RecipientName: strings.TrimSpace(req.Recipient),
		DamageType:    damagefilters.DamageTypeTotal,
		Weapon:        strings.TrimSpace(req.Weapon),
	}
	if len(modifiers) > 0 {
		cfg.DamageModifiers = make(damagefilters.DamageModifiersMap, len(modifiers))
		for k, v := range modifiers {
			cfg.DamageModifiers[combat.DamageModifier(k)] = v
		}
	}

	var hits int
	var damageSum float64
	targetSet := make(map[string]struct{})

	requestedRecipient := strings.TrimSpace(req.Recipient)
	needTargets := requestedRecipient == ""
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
		if dmg.Initiator.Name != req.Initiator {
			continue
		}
		detailed, ok := cfg.Filter(dmg)
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

	summary := damageTableSummary{
		Hits:   hits,
		Damage: damageSum,
	}

	targets := []string{}
	if requestedRecipient != "" {
		if _, ok := humans[requestedRecipient]; ok {
			targets = []string{requestedRecipient}
		}
	} else {
		targets = keysSorted(targetSet)
	}

	return summary, targets
}

func jsonStringify(v any) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("json: %w", err)
	}
	return string(b), nil
}

// no-op

