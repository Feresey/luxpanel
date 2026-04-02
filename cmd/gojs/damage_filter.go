package main

import (
	"context"
	"encoding/json"
	"sort"
	"strings"

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
	Initiator  string          `json:"initiator"`
	Recipient  string          `json:"recipient"`
	Weapon     string          `json:"weapon"`
	Modifiers  map[string]bool `json:"modifiers,omitempty"`
	DamageType string          `json:"damage_type"`
}

type damageTableSummary struct {
	Hits      int     `json:"hits"`
	Damage    float64 `json:"damage"`
	Hull      float64 `json:"hull"`
	Shield    float64 `json:"shield"`
	AvgDamage float64 `json:"avg_damage"`
}

type damageTableRow struct {
	Source    string             `json:"source"`
	Targets   []string           `json:"targets"`
	Modifiers map[string]bool    `json:"modifiers"`
	Summary   damageTableSummary `json:"summary"`
}

type damageTableResult struct {
	Row   damageTableRow `json:"row"`
	Error string         `json:"error,omitempty"`
}

func (r *Runtime) marshalDamageFilterMetaJSON(level *splitter.Level, initiator string) (string, error) {
	meta := r.collectDamageFilterMeta(level, initiator)
	b, err := json.Marshal(meta)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (r *Runtime) collectDamageFilterMeta(level *splitter.Level, initiator string) damageFilterMeta {
	playersNames := r.getPlayersNames(level)

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

	for _, dmg := range level.CombatLog.Damage {
		if dmg == nil {
			continue
		}
		if dmg.Initiator.Name != "" {
			if _, ok := playersNames[dmg.Initiator.Name]; !ok {
				// Bot (NPC) detected from game log: PlayerID == 0.
				continue
			}
			players[dmg.Initiator.Name] = struct{}{}
		}
		if dmg.Recipient.Name != "" {
			if _, ok := playersNames[dmg.Recipient.Name]; !ok {
				continue
			}
			players[dmg.Recipient.Name] = struct{}{}
		}
		if strings.TrimSpace(initiator) != "" && dmg.Initiator.Name == initiator {
			if s := strings.TrimSpace(dmg.Source); s != "" {
				weapons[s] = struct{}{}
			}
			for _, m := range dmg.DamageModifiers {
				mods[string(m)] = struct{}{}
			}
			if dmg.Recipient.Name != "" {
				recipients[dmg.Recipient.Name] = struct{}{}
			}
		}
	}

	out.Players = keysSorted(players)
	out.Recipients = keysSorted(recipients)
	out.Weapons = keysSorted(weapons)
	out.Modifiers = keysSorted(mods)

	r.lg.For(context.Background()).Infow("filters meta", "data", out)
	return out
}

func (r *Runtime) getPlayersNames(level *splitter.Level) map[string]struct{} {
	out := make(map[string]struct{})
	if level == nil || level.GameLog == nil {
		return out
	}
	for _, p := range level.GameLog.AddPlayer {
		if p == nil {
			continue
		}
		// User request: bots have ID == 0 in game log.
		if p.PlayerID != 0 && strings.TrimSpace(p.Name) != "" {
			out[p.Name] = struct{}{}
		}
	}
	r.lg.For(context.Background()).Infow("players", "players", out)
	return out
}

func keysSorted(m map[string]struct{}) []string {
	s := make([]string, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	sort.Strings(s)
	return s
}

func (r *Runtime) marshalDamageTableJSON(level *splitter.Level, filterJSON string) (string, error) {
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
			return "", err
		}
		return string(b), nil
	}

	var req wasmDamageFilterRequest
	if err := json.Unmarshal([]byte(filterJSON), &req); err != nil {
		res := damageTableResult{Error: "invalid filter json"}
		b, _ := json.Marshal(res)
		return string(b), nil
	}
	if strings.TrimSpace(req.Initiator) == "" {
		b, err := json.Marshal(damageTableResult{
			Row:   damageTableRow{Source: "", Targets: []string{}, Modifiers: map[string]bool{}, Summary: damageTableSummary{}},
			Error: "initiator required",
		})
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	players := r.getPlayersNames(level)
	if _, ok := players[req.Initiator]; !ok {
		b, err := json.Marshal(damageTableResult{
			Row:   damageTableRow{Source: "", Targets: []string{}, Modifiers: map[string]bool{}, Summary: damageTableSummary{}},
			Error: "initiator must be a player",
		})
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	row := r.filterDamageRow(level, &req, players)
	res := damageTableResult{Row: row}
	b, err := json.Marshal(res)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (r *Runtime) filterDamageRow(level *splitter.Level, req *wasmDamageFilterRequest, players map[string]struct{}) damageTableRow {
	var hits int
	var damageSum float64
	var hullSum float64
	var shieldSum float64

	targetSet := make(map[string]struct{})
	needTargets := strings.TrimSpace(req.Recipient) == ""
	requestedRecipient := strings.TrimSpace(req.Recipient)

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
		if !r.matchDamageModifiers(dmg.DamageModifiers, req.Modifiers) {
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
				if _, ok := players[dmg.Recipient.Name]; !ok {
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

	targets := []string{}
	if requestedRecipient != "" {
		// UI should prevent bots, but keep the check for safety.
		if _, ok := players[requestedRecipient]; ok {
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
			Hits:      hits,
			Damage:    damageSum,
			Hull:      hullSum,
			Shield:    shieldSum,
			AvgDamage: avg,
		},
	}
}

func (r *Runtime) filterDamageSummary(level *splitter.Level, req *wasmDamageFilterRequest) damageTableSummary {
	t0 := level.StartLevelTime
	_ = t0 // summary doesn't depend on time, but we keep t0 for consistency

	var hits int
	var damageSum float64
	var hullSum float64
	var shieldSum float64

	for _, dmg := range level.CombatLog.Damage {
		if dmg == nil || dmg.IsEmpty() {
			continue
		}
		if dmg.Initiator.Name != req.Initiator {
			continue
		}
		if r := strings.TrimSpace(req.Recipient); r != "" && dmg.Recipient.Name != r {
			continue
		}
		if w := strings.TrimSpace(req.Weapon); w != "" && dmg.Source != w {
			continue
		}
		if !r.matchDamageModifiers(dmg.DamageModifiers, req.Modifiers) {
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
	}

	avg := 0.0
	if hits > 0 {
		avg = damageSum / float64(hits)
	}

	return damageTableSummary{
		Hits:      hits,
		Damage:    damageSum,
		Hull:      hullSum,
		Shield:    shieldSum,
		AvgDamage: avg,
	}
}

func (r *Runtime) matchDamageModifiers(have []combat.DamageModifier, want map[string]bool) bool {
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
