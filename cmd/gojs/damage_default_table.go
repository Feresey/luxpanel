//go:build js

package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Feresey/luxpanel/internal/splitter"
)

type damageDefaultTableBaseRequest struct {
	Initiator  string `json:"initiator"`
	Recipient  string `json:"recipient"`
	Weapon     string `json:"weapon"`
	DamageType string `json:"damage_type"`
}

type damageDefaultTableRow struct {
	Source    string             `json:"source"`
	Targets   []string           `json:"targets"`
	Modifiers map[string]bool    `json:"modifiers"`
	Summary   damageTableSummary `json:"summary"`
}

type damageDefaultTableResult struct {
	Rows  []damageDefaultTableRow `json:"rows"`
	Error string                  `json:"error,omitempty"`
}

func defaultDamageModifierFilters() []map[string]bool {
	// Keep order aligned with internal/service/service.go:makeDamageFilters.
	return []map[string]bool{
		{}, // no modifiers
		{"CRIT": true},
		{"EXPLOSION": true},
		{"EMP": true},
		{"KINETIC": true},
		{"THERMAL": true},
		{"PRIMARY_WEAPON": true},
		{
			"PRIMARY_WEAPON":   false,
			"SECONDARY_WEAPON": false,
			"COLLISION":        false,
			"CRIT":             false,
			"IGNORE_SHIELD":    false,
		},
		{"IGNORE_SHIELD": true},
		{"COLLISION": true},
		{"MODULE": true},
	}
}

func (r *Runtime) marshalDamageDefaultFiltersTableJSON(level *splitter.Level, filterJSON string) (string, error) {
	var req damageDefaultTableBaseRequest
	if err := json.Unmarshal([]byte(filterJSON), &req); err != nil {
		return jsonStringify(damageDefaultTableResult{Error: "invalid filter json"})
	}
	if strings.TrimSpace(req.Initiator) == "" {
		return jsonStringify(damageDefaultTableResult{Error: "initiator required"})
	}

	if level == nil || level.CombatLog == nil {
		emptyRows := make([]damageDefaultTableRow, 0)
		return jsonStringify(damageDefaultTableResult{Rows: emptyRows})
	}

	players := r.getPlayersNames(level)
	if _, ok := players[req.Initiator]; !ok {
		return jsonStringify(damageDefaultTableResult{Error: "initiator must be a player"})
	}

	modFilters := defaultDamageModifierFilters()
	rows := make([]damageDefaultTableRow, 0, len(modFilters))
	for _, mods := range modFilters {
		summary, targets := r.filterDamageRowWithModifiers(level, req, mods, players)
		rows = append(rows, damageDefaultTableRow{
			Source:    req.Initiator,
			Targets:   targets,
			Modifiers: mods,
			Summary:   summary,
		})
	}

	return jsonStringify(damageDefaultTableResult{Rows: rows})
}

func (r *Runtime) filterDamageRowWithModifiers(
	level *splitter.Level,
	req damageDefaultTableBaseRequest,
	modifiers map[string]bool,
	bots map[string]struct{},
) (damageTableSummary, []string) {
	var hits int
	var damageSum float64
	var hullSum float64
	var shieldSum float64
	targetSet := make(map[string]struct{})

	requestedRecipient := strings.TrimSpace(req.Recipient)
	needTargets := requestedRecipient == ""

	for _, dmg := range level.CombatLog.Damage {
		if dmg == nil || dmg.IsEmpty() {
			continue
		}
		if dmg.Initiator.Name != req.Initiator {
			continue
		}
		if requestedRecipient != "" && dmg.Recipient.Name != requestedRecipient {
			continue
		}
		if w := strings.TrimSpace(req.Weapon); w != "" && dmg.Source != w {
			continue
		}
		if !r.matchDamageModifiers(dmg.DamageModifiers, modifiers) {
			continue
		}

		hull := float64(dmg.DamageHull)
		shield := float64(dmg.DamageShield)
		full := float64(dmg.DamageFull)

		selected := full
		switch strings.ToLower(strings.TrimSpace(req.DamageType)) {
		case "hull":
			selected = hull
		case "shield":
			selected = shield
		default:
			selected = full
		}

		hits++
		damageSum += selected
		hullSum += hull
		shieldSum += shield

		if needTargets {
			if dmg.Recipient.Name != "" {
				if _, ok := bots[dmg.Recipient.Name]; ok {
					continue
				}
				targetSet[dmg.Recipient.Name] = struct{}{}
			}
		}
	}

	avg := 0.0
	if hits > 0 {
		avg = damageSum / float64(hits)
	}

	summary := damageTableSummary{
		Hits:      hits,
		Damage:    damageSum,
		Hull:      hullSum,
		Shield:    shieldSum,
		AvgDamage: avg,
	}

	targets := []string{}
	if requestedRecipient != "" {
		if _, ok := bots[requestedRecipient]; !ok {
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
