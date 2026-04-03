package damagefilters

import (
	"strings"

	"github.com/Feresey/luxpanel/internal/parser/combat"
)

type DamageType int

const (
	DamageTypeTotal DamageType = iota
	DamageTypeHull
	DamageTypeShield
)

type DamageModifiersMap map[combat.DamageModifier]bool

type PlayerDamageFilterConfig struct {
	InitiatorName   string
	RecipientName   string
	DamageToObject  bool
	DamageType      DamageType
	DamageModifiers DamageModifiersMap
	FriendlyFire    bool
	Weapon          string
}

type DetailedDamage struct {
	Source string
	Object string
	Damage float32
}

// Filter logic is intentionally kept aligned with internal/service/damage.go.
func (filter *PlayerDamageFilterConfig) Filter(line *combat.Damage) (res *DetailedDamage, ok bool) {
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
		for wantModifier, shouldBe := range filter.DamageModifiers {
			var exists bool
			for _, modifier := range line.DamageModifiers {
				if modifier == wantModifier {
					exists = true
					break
				}
			}
			if exists != shouldBe {
				return res, false
			}
		}
	}

	if filter.Weapon != "" && line.Source != filter.Weapon {
		return res, false
	}

	res = &DetailedDamage{Source: line.Source, Object: line.Recipient.ObjectName}
	switch filter.DamageType {
	case DamageTypeHull:
		res.Damage = line.DamageHull
	case DamageTypeShield:
		res.Damage = line.DamageShield
	default:
		res.Damage = line.DamageFull
	}
	return res, true
}

func ParseDamageType(s string) DamageType {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "hull":
		return DamageTypeHull
	case "shield":
		return DamageTypeShield
	default:
		return DamageTypeTotal
	}
}

func DefaultModifierFilters() []DamageModifiersMap {
	return []DamageModifiersMap{
		{},
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

