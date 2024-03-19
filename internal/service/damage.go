package service

import (
	"strconv"
	"strings"

	"github.com/Feresey/luxpanel/internal/parser"
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
	InitiatorName   string
	RecipientName   string
	RecipientObject string
	DamageType      DamageType
	DamageModifiers DamageModifiersMap
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
		for modifier, exists := range f.DamageModifiers {
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

// FilterPlayerDamage суммирует урон игрока по указанным модификаторам
func FilterPlayerDamage(filter *PlayerDamageFilterConfig) Filter[*parser.CombatLogLineDamage, float32] {
	return func(line *parser.CombatLogLineDamage) (res float32, ok bool) {
		if filter.InitiatorName != "" && line.Players.Initiator != filter.InitiatorName {
			return 0, false
		}
		if filter.RecipientName != "" && line.Players.Recipient != filter.RecipientName {
			return 0, false
		}
		if filter.RecipientObject != "" && line.Players.RecipientObject != filter.RecipientObject {
			return 0, false
		}

		if filter.DamageModifiers != nil {
			for wantModifier, exists := range filter.DamageModifiers {
				if _, ok := line.DamageModifiers[wantModifier]; ok != exists {
					return 0, false
				}
			}
		}

		switch filter.DamageType {
		case DamageTypeTotal:
			return line.DamageTotal, true
		case DamageTypeHull:
			return line.DamageHull, true
		case DamageTypeShield:
			return line.DamageShield, true
		default:
			return 0, false
		}
	}
}
