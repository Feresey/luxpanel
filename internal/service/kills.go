package service

import (
	"strings"

	"github.com/Feresey/luxpanel/internal/parser/combat"
)

type PlayerKillsFilterConfig struct {
	InitiatorName string `json:"initiator_name,omitempty"`
	RecipientName string `json:"recipient_name,omitempty"`
	DestroyObject bool   `json:"destroy_object,omitempty"`
	FriendlyFire  bool   `json:"friendly_fire,omitempty"`
}

func (f *PlayerKillsFilterConfig) String() string {
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
	if f.DestroyObject {
		space()
		sb.WriteString("Kills_to_object: true")
	}

	return sb.String()
}

type KillsResult struct {
	BySource map[string]AggCount[int]
	ToObject map[string]AggCount[int]
}

func SummKillsBySource(res *KillsResult, elem *DetailedKill) *KillsResult {
	if res == nil {
		res = &KillsResult{
			BySource: make(map[string]AggCount[int]),
			ToObject: make(map[string]AggCount[int]),
		}
	}

	src := res.BySource[elem.Source]
	src.Count++
	res.BySource[elem.Source] = src

	obj := res.ToObject[elem.Target]
	obj.Count++
	res.ToObject[elem.Target] = obj

	return res
}

type DetailedKill struct {
	Source string
	Target string
}

// FilterPlayerKills суммирует урон игрока по указанным модификаторам
func FilterPlayerKills(filter *PlayerKillsFilterConfig) Filter[*combat.Kill, *DetailedKill] {
	return func(line *combat.Kill) (res *DetailedKill, ok bool) {
		if filter.InitiatorName != "" && line.Initiator != filter.InitiatorName {
			return res, false
		}
		if filter.RecipientName != "" && line.RecipientName != filter.RecipientName {
			return res, false
		}

		if filter.DestroyObject {
			if line.RecipientObjectName == "" {
				return res, false
			}
			if filter.RecipientName != "" && filter.RecipientName != line.RecipientName {
				return res, false
			}
			if filter.RecipientName != line.RecipientObjectOwner {
				return res, false
			}
		}

		if filter.FriendlyFire && !line.FriendlyFire {
			return res, false
		}

		res = &DetailedKill{Source: line.ActionSource, Target: line.RecipientObjectName}
		return res, true
	}
}
