package front

import (
	"encoding/json"
	"sort"
	"strings"
	"time"

	"github.com/Feresey/luxpanel/internal/damagefilters"
	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/Feresey/luxpanel/internal/splitter"
)

// DamageFilterMeta lists unique values for building filter UI.
type DamageFilterMeta struct {
	Players    []string `json:"players"`
	Recipients []string `json:"recipients"`
	Weapons    []string `json:"weapons"`
	Modifiers  []string `json:"modifiers"`
}

// WasmDamageFilterRequest is the JSON body from the browser damage table.
type WasmDamageFilterRequest struct {
	Initiator   string          `json:"initiator"`
	Recipient   string          `json:"recipient"`
	Perspective string          `json:"perspective,omitempty"`
	Weapon      string          `json:"weapon"`
	Modifiers   map[string]bool `json:"modifiers,omitempty"`
	Aggregate   *bool           `json:"aggregate,omitempty"`
	TimeFromSec *float64        `json:"time_from_sec,omitempty"`
	TimeToSec   *float64        `json:"time_to_sec,omitempty"`
}

type DamageTableSummary struct {
	Hits   int     `json:"hits"`
	Damage float64 `json:"damage"`
}

// DamageTableRowDraft is a damage table row before optional incoming suffix on Source (cmd/gojs).
type DamageTableRowDraft struct {
	SourceBase          string
	IncomingPerspective bool
	Targets             []string
	Modifiers           map[string]bool
	Summary             DamageTableSummary
}

type DamageTableResult struct {
	Row    DamageTableRowOut   `json:"row"`
	Events []DamageEventRow    `json:"events,omitempty"`
	Error  string              `json:"error,omitempty"`
}

type DamageTableRowOut struct {
	Source                string             `json:"source"`
	Targets               []string           `json:"targets"`
	Modifiers             map[string]bool    `json:"modifiers"`
	Summary               DamageTableSummary `json:"summary"`
	IncomingPerspective   bool               `json:"-"` // cmd/gojs appends localized suffix to Source
}

type DamageEventRow struct {
	TimeSec   float64         `json:"time_sec"`
	Initiator string          `json:"initiator"`
	Recipient string          `json:"recipient"`
	Weapon    string          `json:"weapon"`
	Modifiers map[string]bool `json:"modifiers"`
	Amount    float64         `json:"amount"`
}

func IsRecipientPerspective(v string) bool {
	return strings.EqualFold(strings.TrimSpace(v), "recipient")
}

// HumanAndBotPlayerNames classifies roster names from game.log AddPlayer.
func HumanAndBotPlayerNames(level *splitter.Level) (humans, bots map[string]struct{}) {
	humans = make(map[string]struct{})
	bots = make(map[string]struct{})
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
		if p.PlayerID != 0 {
			humans[name] = struct{}{}
		} else {
			bots[name] = struct{}{}
		}
	}
	return humans, bots
}

func HumanPlayerNames(level *splitter.Level) map[string]struct{} {
	humans, _ := HumanAndBotPlayerNames(level)
	return humans
}

func keysSortedStringSet(m map[string]struct{}) []string {
	s := make([]string, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	sort.Strings(s)
	return s
}

// CollectDamageFilterMeta gathers distinct filter dimensions for an initiator slice.
func CollectDamageFilterMeta(level *splitter.Level, initiator string, lo, hi float64) DamageFilterMeta {
	humans := HumanPlayerNames(level)
	out := DamageFilterMeta{
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
	for name := range humans {
		players[name] = struct{}{}
	}
	t0 := level.StartLevelTime
	for _, dmg := range level.CombatLog.Damage {
		if dmg == nil {
			continue
		}
		t := dmg.GetTime(t0)
		if t.Before(t0) {
			t = t0
		}
		if !TimeInRangeFromStart(t, t0, lo, hi) {
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
	out.Players = keysSortedStringSet(players)
	out.Recipients = keysSortedStringSet(recipients)
	out.Weapons = keysSortedStringSet(weapons)
	out.Modifiers = keysSortedStringSet(mods)
	return out
}

func ToDamageFilter(req *WasmDamageFilterRequest) damagefilters.PlayerDamageFilterConfig {
	mods := damagefilters.DamageModifiersMap{}
	for k, v := range req.Modifiers {
		mods[combat.DamageModifier(k)] = v
	}
	isIncoming := IsRecipientPerspective(req.Perspective)
	initiatorName := strings.TrimSpace(req.Initiator)
	recipientName := strings.TrimSpace(req.Recipient)
	if isIncoming {
		recipientName = initiatorName
		initiatorName = strings.TrimSpace(req.Recipient)
	}
	return damagefilters.PlayerDamageFilterConfig{
		InitiatorName:   initiatorName,
		RecipientName:   recipientName,
		DamageType:      damagefilters.DamageTypeTotal,
		DamageModifiers: mods,
		Weapon:          strings.TrimSpace(req.Weapon),
	}
}

// CollectDamageEvents returns individual damage lines matching the filter.
func CollectDamageEvents(level *splitter.Level, req *WasmDamageFilterRequest, humans map[string]struct{}, lo, hi float64) []DamageEventRow {
	if level == nil || level.CombatLog == nil {
		return []DamageEventRow{}
	}
	t0 := level.StartLevelTime
	filter := ToDamageFilter(req)
	out := make([]DamageEventRow, 0)
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
		if !TimeInRangeFromStart(t, t0, lo, hi) {
			continue
		}
		sec := float64(t.Sub(t0)) / float64(time.Second)
		amount := float64(detailed.Damage)
		mods := make(map[string]bool, len(dmg.DamageModifiers))
		for _, m := range dmg.DamageModifiers {
			mods[string(m)] = true
		}
		out = append(out, DamageEventRow{
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

// FilterDamageRowDraft aggregates one damage table row (Source localized in cmd/gojs).
func FilterDamageRowDraft(level *splitter.Level, req *WasmDamageFilterRequest, humans map[string]struct{}, lo, hi float64) DamageTableRowDraft {
	var hits int
	var damageSum float64
	targetSet := make(map[string]struct{})
	needTargets := strings.TrimSpace(req.Recipient) == ""
	requestedRecipient := strings.TrimSpace(req.Recipient)
	isIncoming := IsRecipientPerspective(req.Perspective)
	filter := ToDamageFilter(req)
	t0 := level.StartLevelTime
	for _, dmg := range level.CombatLog.Damage {
		if dmg == nil || dmg.IsEmpty() {
			continue
		}
		t := dmg.GetTime(t0)
		if t.Before(t0) {
			t = t0
		}
		if !TimeInRangeFromStart(t, t0, lo, hi) {
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
			if !isIncoming {
				if dmg.Recipient.Name != "" {
					if _, ok := humans[dmg.Recipient.Name]; !ok {
						continue
					}
					targetSet[dmg.Recipient.Name] = struct{}{}
				}
			} else if dmg.Initiator.Name != "" {
				if _, ok := humans[dmg.Initiator.Name]; !ok {
					continue
				}
				targetSet[dmg.Initiator.Name] = struct{}{}
			}
		}
	}
	targets := []string{}
	if requestedRecipient != "" {
		if _, ok := humans[requestedRecipient]; ok {
			targets = []string{requestedRecipient}
		}
	} else {
		targets = keysSortedStringSet(targetSet)
	}
	mods := req.Modifiers
	if mods == nil {
		mods = map[string]bool{}
	}
	return DamageTableRowDraft{
		SourceBase:          strings.TrimSpace(req.Initiator),
		IncomingPerspective: isIncoming,
		Targets:             targets,
		Modifiers:           mods,
		Summary: DamageTableSummary{
			Hits:   hits,
			Damage: damageSum,
		},
	}
}

// DamageDefaultTableBaseRequest is JSON for default-filters table.
type DamageDefaultTableBaseRequest struct {
	Initiator   string   `json:"initiator"`
	Recipient   string   `json:"recipient"`
	Perspective string   `json:"perspective,omitempty"`
	Weapon      string   `json:"weapon"`
	TimeFromSec *float64 `json:"time_from_sec,omitempty"`
	TimeToSec   *float64 `json:"time_to_sec,omitempty"`
}

type DamageDefaultTableRow struct {
	Source    string              `json:"source"`
	Targets   []string            `json:"targets"`
	Modifiers map[string]bool     `json:"modifiers"`
	Summary   DamageTableSummary  `json:"summary"`
	Events    []DamageEventRow    `json:"events,omitempty"`
}

type DamageDefaultTableResult struct {
	Rows  []DamageDefaultTableRow `json:"rows"`
	Error string                  `json:"error,omitempty"`
}

func DefaultDamageModifierFilters() []map[string]bool {
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

func filterDamageRowWithModifiers(level *splitter.Level, req DamageDefaultTableBaseRequest, modifiers map[string]bool, humans map[string]struct{}, lo, hi float64) (DamageTableSummary, []string) {
	isIncoming := IsRecipientPerspective(req.Perspective)
	initiatorName := strings.TrimSpace(req.Initiator)
	recipientName := strings.TrimSpace(req.Recipient)
	if isIncoming {
		recipientName = initiatorName
		initiatorName = strings.TrimSpace(req.Recipient)
	}
	cfg := damagefilters.PlayerDamageFilterConfig{
		InitiatorName: initiatorName,
		RecipientName: recipientName,
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
		if !TimeInRangeFromStart(t, t0, lo, hi) {
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
			if !isIncoming {
				if dmg.Recipient.Name != "" {
					if _, ok := humans[dmg.Recipient.Name]; !ok {
						continue
					}
					targetSet[dmg.Recipient.Name] = struct{}{}
				}
			} else if dmg.Initiator.Name != "" {
				if _, ok := humans[dmg.Initiator.Name]; !ok {
					continue
				}
				targetSet[dmg.Initiator.Name] = struct{}{}
			}
		}
	}
	summary := DamageTableSummary{Hits: hits, Damage: damageSum}
	targets := []string{}
	if requestedRecipient != "" {
		if _, ok := humans[requestedRecipient]; ok {
			targets = []string{requestedRecipient}
		}
	} else {
		targets = keysSortedStringSet(targetSet)
	}
	return summary, targets
}

// BuildDamageDefaultTableResult validates input and builds rows for default modifier presets.
func BuildDamageDefaultTableResult(level *splitter.Level, filterJSON string) DamageDefaultTableResult {
	var req DamageDefaultTableBaseRequest
	if err := json.Unmarshal([]byte(filterJSON), &req); err != nil {
		return DamageDefaultTableResult{Error: "invalid filter json"}
	}
	if strings.TrimSpace(req.Initiator) == "" {
		return DamageDefaultTableResult{Error: "initiator required"}
	}
	if level == nil || level.CombatLog == nil {
		return DamageDefaultTableResult{Rows: []DamageDefaultTableRow{}}
	}
	humans := HumanPlayerNames(level)
	if _, ok := humans[req.Initiator]; !ok {
		return DamageDefaultTableResult{Error: "initiator must be a player"}
	}
	lo, hi := ApplyTimeBoundsToReq(level, req.TimeFromSec, req.TimeToSec)
	modFilters := DefaultDamageModifierFilters()
	rows := make([]DamageDefaultTableRow, 0, len(modFilters))
	for _, mods := range modFilters {
		summary, targets := filterDamageRowWithModifiers(level, req, mods, humans, lo, hi)
		wasmReq := WasmDamageFilterRequest{
			Initiator:   req.Initiator,
			Recipient:   req.Recipient,
			Perspective: req.Perspective,
			Weapon:      req.Weapon,
			Modifiers:   mods,
			TimeFromSec: req.TimeFromSec,
			TimeToSec:   req.TimeToSec,
		}
		rows = append(rows, DamageDefaultTableRow{
			Source:    req.Initiator,
			Targets:   targets,
			Modifiers: mods,
			Summary:   summary,
			Events:    CollectDamageEvents(level, &wasmReq, humans, lo, hi),
		})
	}
	return DamageDefaultTableResult{Rows: rows}
}

// BuildDamageTableResult validates and builds a single damage table response.
func BuildDamageTableResult(level *splitter.Level, filterJSON string) DamageTableResult {
	if level == nil || level.CombatLog == nil {
		return DamageTableResult{
			Row: DamageTableRowOut{
				Source:    "",
				Targets:   []string{},
				Modifiers: map[string]bool{},
				Summary:   DamageTableSummary{},
			},
		}
	}
	var req WasmDamageFilterRequest
	if err := json.Unmarshal([]byte(filterJSON), &req); err != nil {
		return DamageTableResult{Error: "invalid filter json"}
	}
	if strings.TrimSpace(req.Initiator) == "" {
		return DamageTableResult{
			Row:   DamageTableRowOut{Source: "", Targets: []string{}, Modifiers: map[string]bool{}, Summary: DamageTableSummary{}},
			Error: "initiator required",
		}
	}
	humans := HumanPlayerNames(level)
	if _, ok := humans[req.Initiator]; !ok {
		return DamageTableResult{
			Row:   DamageTableRowOut{Source: "", Targets: []string{}, Modifiers: map[string]bool{}, Summary: DamageTableSummary{}},
			Error: "initiator must be a player",
		}
	}
	lo, hi := ApplyTimeBoundsToReq(level, req.TimeFromSec, req.TimeToSec)
	draft := FilterDamageRowDraft(level, &req, humans, lo, hi)
	row := DamageTableRowOut{
		Source:                draft.SourceBase,
		Targets:               draft.Targets,
		Modifiers:             draft.Modifiers,
		Summary:               draft.Summary,
		IncomingPerspective:   draft.IncomingPerspective,
	}
	events := CollectDamageEvents(level, &req, humans, lo, hi)
	return DamageTableResult{Row: row, Events: events}
}
