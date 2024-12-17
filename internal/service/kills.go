package service

import (
	"strings"

	"github.com/Feresey/luxpanel/internal/parser/combat"
)

type PlayerKillsFilterConfig struct {
	Killer        string `json:"initiator_name,omitempty"`
	Killed        string `json:"recipient_name,omitempty"`
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

	if f.Killer != "" {
		space()
		sb.WriteString("initiator: " + f.Killer)
	}
	if f.FriendlyFire {
		space()
		sb.WriteString("friendly_fire: true")
	}
	if f.Killed != "" {
		space()
		sb.WriteString("recipient: " + f.Killed)
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
		if filter.Killer != "" && line.Killer.Name != filter.Killer {
			return res, false
		}
		if filter.Killed != "" && line.Killed.Name != filter.Killed {
			return res, false
		}

		if !filter.DestroyObject && line.Killed.ObjectName == "" {
			return res, false
		}

		if filter.DestroyObject {
			if line.Killed.ObjectName == "" {
				return res, false
			}
			if filter.Killed != "" && filter.Killed != line.Killer.Name {
				return res, false
			}
			if filter.Killed != line.Killed.ObjectOwner {
				return res, false
			}
		}

		if filter.FriendlyFire && !line.FriendlyFire {
			return res, false
		}

		res = &DetailedKill{Source: line.Source, Target: line.Killed.Name}
		return res, true
	}
}
