package service

import (
	"strings"

	"github.com/Feresey/luxpanel/internal/parser/combat"
)

type PlayerHealFilterConfig struct {
	InitiatorName string `json:"initiator_name,omitempty"`
	RecipientName string `json:"recipient_name,omitempty"`
	HealToObject  bool   `json:"Heal_to_object,omitempty"`
}

func (f *PlayerHealFilterConfig) String() string {
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
	if f.HealToObject {
		space()
		sb.WriteString("heal_to_object: true")
	}

	return sb.String()
}

type HealResult struct {
	BySource map[string]AggCount[float32]
	ToObject map[string]AggCount[float32]
}

func SummHealBySource(res *HealResult, elem *DetailedHeal) *HealResult {
	if res == nil {
		res = &HealResult{
			BySource: make(map[string]AggCount[float32]),
			ToObject: make(map[string]AggCount[float32]),
		}
	}

	src := res.BySource[elem.Source]
	src.Count++
	src.Value += elem.Heal
	res.BySource[elem.Source] = src

	obj := res.ToObject[elem.Object]
	obj.Count++
	obj.Value += elem.Heal
	res.ToObject[elem.Object] = obj

	return res
}

type DetailedHeal struct {
	Source string
	Object string
	Heal   float32
}

// FilterPlayerHeal суммирует урон игрока по указанным модификаторам
func FilterPlayerHeal(filter *PlayerHealFilterConfig) Filter[*combat.Heal, *DetailedHeal] {
	return func(line *combat.Heal) (res *DetailedHeal, ok bool) {
		if filter.InitiatorName != "" && line.Initiator != filter.InitiatorName {
			return res, false
		}
		if filter.RecipientName != "" && line.Recipient != filter.RecipientName {
			return res, false
		}

		if !filter.HealToObject && line.Recipient == "" {
			return res, false
		}

		if filter.HealToObject {
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

		res = &DetailedHeal{Source: line.ActionSource, Object: line.ObjectName, Heal: line.Heal}
		return res, true
	}
}

// type PlayerHeal struct {
// 	Player string

// 	Total       float32
// 	HullTotal   float32
// 	ShieldTotal float32

// 	ByModifiers map[combat.HealModifier][]*combat.Heal
// 	BySource    map[string][]*combat.Heal

// 	ToPlayer        map[string][]*combat.Heal
// 	ToObject        map[string][]*combat.Heal
// 	ToPlayerObjects map[string][]*combat.Heal

// 	FriendlyFireTotal    float32
// 	FriendlyFireByPlayer map[string][]*combat.Heal
// }

// func addMap[K comparable, V any, M ~map[K][]V](m M, key K, val V) {
// 	m[key] = append(m[key], val)
// }

// func AggregatePlayerHeal() Aggregator[map[string]PlayerHeal, *combat.Heal] {
// 	return func(agg map[string]PlayerHeal, line *combat.Heal) map[string]PlayerHeal {
// 		if agg == nil {
// 			agg = make(map[string]PlayerHeal)
// 		}

// 		playerData, ok := agg[line.Initiator]
// 		if !ok {
// 			playerData = PlayerHeal{
// 				Player:               line.Initiator,
// 				ByModifiers:          make(map[combat.HealModifier][]*combat.Heal),
// 				BySource:             make(map[string][]*combat.Heal),
// 				ToPlayer:             make(map[string][]*combat.Heal),
// 				ToObject:             make(map[string][]*combat.Heal),
// 				ToPlayerObjects:      make(map[string][]*combat.Heal),
// 				FriendlyFireByPlayer: make(map[string][]*combat.Heal),
// 			}
// 			agg[line.Initiator] = playerData
// 		}

// 		playerData.Total += line.HealTotal
// 		playerData.HullTotal += line.HealHull
// 		playerData.ShieldTotal += line.HealShield
// 		if line.FriendlyFire {
// 			playerData.FriendlyFireTotal += line.HealTotal
// 			addMap(playerData.FriendlyFireByPlayer, line.Recipient, line)
// 		}
// 		for modifier := range line.HealModifiers {
// 			addMap(playerData.ByModifiers, modifier, line)
// 		}
// 		addMap(playerData.BySource, line.ActionSource, line)
// 		addMap(playerData.ToObject, line.ObjectName, line)
// 		addMap(playerData.ToPlayerObjects, line.ObjectOwner+" "+line.ObjectName, line)

// 		return agg
// 	}
// }
