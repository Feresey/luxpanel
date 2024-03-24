package service

import (
	"slices"
	"strconv"
	"strings"

	"github.com/Feresey/luxpanel/internal/parser/combat"
	"golang.org/x/exp/maps"
)

//go:generate stringer -type=DamageType
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
	BySource map[string]float32
}

func SummDamageBySource(res *DamageResult, elem *DamageWithSource) *DamageResult {
	if res == nil {
		res = new(DamageResult)
		res.BySource = make(map[string]float32)
	}

	res.BySource[elem.Source] += elem.Damage
	return res
}

type DamageWithSource struct {
	Source string
	Object string
	Damage float32
}

// FilterPlayerDamage суммирует урон игрока по указанным модификаторам
func FilterPlayerDamage(filter *PlayerDamageFilterConfig) Filter[*combat.Damage, *DamageWithSource] {
	return func(line *combat.Damage) (res *DamageWithSource, ok bool) {
		if filter.InitiatorName != "" && line.Initiator != filter.InitiatorName {
			return res, false
		}
		if filter.RecipientName != "" && line.Recipient != filter.RecipientName {
			return res, false
		}

		if filter.DamageToObject {
			if line.ObjectName == "" {
				return res, false
			}
			if filter.RecipientName != "" && filter.RecipientName != line.Recipient {
				return res, false
			}
			if filter.RecipientName != line.ObjectOwner {
				return res, false
			}
		}

		if filter.FriendlyFire && !line.FriendlyFire {
			return res, false
		}

		if filter.DamageModifiers != nil {
			for wantModifier, exists := range filter.DamageModifiers {
				if _, ok := line.DamageModifiers[wantModifier]; ok != exists {
					return res, false
				}
			}
		}

		res = &DamageWithSource{Source: line.ActionSource, Object: line.ObjectName}
		switch filter.DamageType {
		case DamageTypeTotal:
			res.Damage = line.DamageTotal
		case DamageTypeHull:
			res.Damage = line.DamageHull
		case DamageTypeShield:
			res.Damage = line.DamageShield
		}
		return res, false
	}
}

// type PlayerDamage struct {
// 	Player string

// 	Total       float32
// 	HullTotal   float32
// 	ShieldTotal float32

// 	ByModifiers map[combat.DamageModifier][]*combat.Damage
// 	BySource    map[string][]*combat.Damage

// 	ToPlayer        map[string][]*combat.Damage
// 	ToObject        map[string][]*combat.Damage
// 	ToPlayerObjects map[string][]*combat.Damage

// 	FriendlyFireTotal    float32
// 	FriendlyFireByPlayer map[string][]*combat.Damage
// }

// func addMap[K comparable, V any, M ~map[K][]V](m M, key K, val V) {
// 	m[key] = append(m[key], val)
// }

// func AggregatePlayerDamage() Aggregator[map[string]PlayerDamage, *combat.Damage] {
// 	return func(agg map[string]PlayerDamage, line *combat.Damage) map[string]PlayerDamage {
// 		if agg == nil {
// 			agg = make(map[string]PlayerDamage)
// 		}

// 		playerData, ok := agg[line.Initiator]
// 		if !ok {
// 			playerData = PlayerDamage{
// 				Player:               line.Initiator,
// 				ByModifiers:          make(map[combat.DamageModifier][]*combat.Damage),
// 				BySource:             make(map[string][]*combat.Damage),
// 				ToPlayer:             make(map[string][]*combat.Damage),
// 				ToObject:             make(map[string][]*combat.Damage),
// 				ToPlayerObjects:      make(map[string][]*combat.Damage),
// 				FriendlyFireByPlayer: make(map[string][]*combat.Damage),
// 			}
// 			agg[line.Initiator] = playerData
// 		}

// 		playerData.Total += line.DamageTotal
// 		playerData.HullTotal += line.DamageHull
// 		playerData.ShieldTotal += line.DamageShield
// 		if line.FriendlyFire {
// 			playerData.FriendlyFireTotal += line.DamageTotal
// 			addMap(playerData.FriendlyFireByPlayer, line.Recipient, line)
// 		}
// 		for modifier := range line.DamageModifiers {
// 			addMap(playerData.ByModifiers, modifier, line)
// 		}
// 		addMap(playerData.BySource, line.ActionSource, line)
// 		addMap(playerData.ToObject, line.ObjectName, line)
// 		addMap(playerData.ToPlayerObjects, line.ObjectOwner+" "+line.ObjectName, line)

// 		return agg
// 	}
// }
