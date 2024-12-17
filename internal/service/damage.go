package service

import (
	"slices"
	"strconv"
	"strings"

	"github.com/Feresey/luxpanel/internal/parser/combat"
	"golang.org/x/exp/maps"
)

//go:generate go run github.com/alvaroloes/enumer -type DamageType -trimprefix DamageType -text -transform snake
type DamageType int

const (
	DamageTypeTotal DamageType = iota
	DamageTypeHull
	DamageTypeShield
)

type DamageModifiersMap map[combat.DamageModifier]bool

type PlayerDamageFilterConfig struct {
	InitiatorName   string             `json:"initiator_name,omitempty"`
	RecipientName   string             `json:"recipient_name,omitempty"`
	DamageToObject  bool               `json:"damage_to_object,omitempty"`
	DamageType      DamageType         `json:"damage_type,omitempty"`
	DamageModifiers DamageModifiersMap `json:"damage_modifiers,omitempty"`
	FriendlyFire    bool               `json:"friendly_fire,omitempty"`
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
	if f.FriendlyFire {
		space()
		sb.WriteString("friendly_fire: true")
	}
	if f.RecipientName != "" {
		space()
		sb.WriteString("recipient: " + f.RecipientName)
	}
	if f.DamageToObject {
		space()
		sb.WriteString("damage_to_object: true")
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
	BySource map[string]AggCount[float32]
	ToObject map[string]AggCount[float32]
}

func SummDamageBySource(res *DamageResult, elem *DetailedDamage) *DamageResult {
	if res == nil {
		res = &DamageResult{
			BySource: make(map[string]AggCount[float32]),
			ToObject: make(map[string]AggCount[float32]),
		}
	}

	src := res.BySource[elem.Source]
	src.Count++
	src.Value += elem.Damage
	res.BySource[elem.Source] = src

	obj := res.ToObject[elem.Object]
	obj.Count++
	obj.Value += elem.Damage
	res.ToObject[elem.Object] = obj

	return res
}

type DetailedDamage struct {
	Source string
	Object string
	Damage float32
}

// FilterPlayerDamage суммирует урон игрока по указанным модификаторам
func FilterPlayerDamage(filter *PlayerDamageFilterConfig) Filter[*combat.Damage, *DetailedDamage] {
	return func(line *combat.Damage) (res *DetailedDamage, ok bool) {
		if filter.InitiatorName != "" && line.Initiator.Name != filter.InitiatorName {
			return res, false
		}
		if filter.RecipientName != "" && line.Recipient.Name != filter.RecipientName {
			return res, false
		}

		if filter.DamageToObject {
			if line.Recipient.ObjectName == "" {
				return res, false
			}
			if filter.RecipientName != "" && filter.RecipientName != line.Recipient.Name {
				return res, false
			}
			if filter.RecipientName != line.Recipient.ObjectOwner {
				return res, false
			}
		}

		if filter.FriendlyFire && !line.FriendlyFire {
			return res, false
		}

		if filter.DamageModifiers != nil {
			for _, modifier := range line.DamageModifiers {
				if want, exists := filter.DamageModifiers[modifier]; exists {
					if !want {
						return res, false
					}
				}

			}
		}

		res = &DetailedDamage{Source: line.Source, Object: line.Recipient.ObjectName}
		switch filter.DamageType {
		case DamageTypeTotal:
			res.Damage = line.DamageFull
		case DamageTypeHull:
			res.Damage = line.DamageHull
		case DamageTypeShield:
			res.Damage = line.DamageShield
		}
		return res, true
	}
}
