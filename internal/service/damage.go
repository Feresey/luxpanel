package service

import (
	"slices"
	"strconv"
	"strings"

	"github.com/Feresey/luxpanel/internal/parser"
	"golang.org/x/exp/maps"
)

//go:generate stringer -type=DamageType
type DamageType int

const (
	DamageTypeTotal DamageType = iota
	DamageTypeHull
	DamageTypeShield
)

type DamageModifiersMap map[parser.DamageModifier]bool

type PlayerDamageFilterConfig struct {
	InitiatorName    string
	RecipientName    string
	RecipientObject  string
	DamageType       DamageType
	DamageModifiers  DamageModifiersMap
	FriendlyFireOnly bool
}

func (f *PlayerDamageFilterConfig) String() string {
	var sb strings.Builder
	var needSpace bool
	space := func() {
		if needSpace {
			sb.WriteRune(' ')
		} else {
			needSpace = true
		}
	}

	if f.InitiatorName != "" {
		space()
		sb.WriteString("initiator: " + f.InitiatorName)
	}
	if f.FriendlyFireOnly {
		space()
		sb.WriteString("friendly_fire")
	}
	if f.RecipientName != "" {
		space()
		sb.WriteString("recipient: " + f.RecipientName)
	}
	if f.RecipientObject != "" {
		space()
		sb.WriteString("recipient_object: " + f.RecipientObject)
	}

	space()
	sb.WriteString("damage_type: " + f.DamageType.String())

	if f.DamageModifiers != nil {
		space()
		var needComma bool
		sb.WriteRune('(')

		keys := maps.Keys(f.DamageModifiers)
		slices.Sort(keys)
		for _, modifier := range keys {
			exists := f.DamageModifiers[modifier]
			if needComma {
				sb.WriteRune(',')
			}
			needComma = true
			sb.WriteString(string(modifier) + ":" + strconv.FormatBool(exists))
		}
		sb.WriteRune(')')
	}

	return sb.String()
}

type DamageResult struct {
	BySource map[string]float32
}

func SummDamageBySource(res *DamageResult, elem *DamageWithSource) *DamageResult {
	if res == nil {
		res = new(DamageResult)
		res.BySource = make(map[string]float32)
	}

	res.BySource[elem.Source] += elem.Value
	return res
}

type DamageWithSource struct {
	Source string
	Value  float32
}

// FilterPlayerDamage суммирует урон игрока по указанным модификаторам
func FilterPlayerDamage(filter *PlayerDamageFilterConfig) Filter[*parser.CombatLogLineDamage, *DamageWithSource] {
	return func(line *parser.CombatLogLineDamage) (res *DamageWithSource, ok bool) {
		if filter.InitiatorName != "" && line.Players.Initiator != filter.InitiatorName {
			return res, false
		}
		if filter.RecipientName != "" && line.Players.Recipient != filter.RecipientName {
			return res, false
		}
		if filter.RecipientObject != "" && line.Players.RecipientObject != filter.RecipientObject {
			return res, false
		}

		if filter.FriendlyFireOnly && !line.IsFriendlyFire {
			return res, false
		}

		if filter.DamageModifiers != nil {
			for wantModifier, exists := range filter.DamageModifiers {
				if _, ok := line.DamageModifiers[wantModifier]; ok != exists {
					return res, false
				}
			}
		}

		switch filter.DamageType {
		case DamageTypeTotal:
			return &DamageWithSource{Value: line.DamageTotal, Source: line.DamageSource}, true
		case DamageTypeHull:
			return &DamageWithSource{Value: line.DamageHull, Source: line.DamageSource}, true
		case DamageTypeShield:
			return &DamageWithSource{Value: line.DamageShield, Source: line.DamageSourc