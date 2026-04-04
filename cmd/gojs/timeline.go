package main

import (
	"context"
	"encoding/json"
	"math"
	"sort"
	"strings"
	"time"

	"slices"

	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/Feresey/luxpanel/internal/splitter"
)

type timelineMarker struct {
	Kind       string  `json:"kind"`
	TimeSec    float64 `json:"time_sec"`
	Label      string  `json:"label,omitempty"`
	Player     string  `json:"player,omitempty"`
	Killer     string  `json:"killer,omitempty"`
	Victim     string  `json:"victim,omitempty"`
	KillerShip string   `json:"killer_ship,omitempty"`
	VictimShip string   `json:"victim_ship,omitempty"`
	PlayerShip string   `json:"player_ship,omitempty"`
	Weapon     string   `json:"weapon,omitempty"`
	Assists    []string `json:"assists,omitempty"`
}

type timelineResult struct {
	StartSec float64          `json:"start_sec"`
	EndSec   float64          `json:"end_sec"`
	Markers  []timelineMarker `json:"markers"`
}

func (r *Runtime) marshalTimelineJSON(ctx context.Context, level *splitter.Level) (string, error) {
	if level == nil {
		b, err := json.Marshal(timelineResult{})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalTimelineJSON", "err", err, "nil_level", true)
			return "", err
		}
		s := string(b)
		r.lg.For(ctx).Debugw("marshalTimelineJSON", "nil_level", true, "out_len", len(s))
		return s, nil
	}

	t0 := level.StartLevelTime
	span := levelSpanSeconds(level)

	var markers []timelineMarker

	if level.CombatLog != nil {
		nameTeam := rosterNameToTeamID(level)
		for _, sp := range level.CombatLog.Spawn {
			if sp == nil {
				continue
			}
			name := strings.TrimSpace(sp.Name)
			if name == "" {
				continue
			}
			tid, inRoster := nameTeam[strings.ToLower(name)]
			if !inRoster {
				continue
			}
			if tid == 0 {
				continue
			}
			kind, ok := spawnKindForTeam(level, tid)
			if !ok {
				continue
			}
			t := sp.GetTime(t0)
			if t.Before(t0) {
				t = t0
			}
			sec := t.Sub(t0).Seconds()
			markers = append(markers, timelineMarker{
				Kind:       kind,
				TimeSec:    sec,
				Player:     name,
				Label:      name,
				PlayerShip: strings.TrimSpace(sp.Ship),
			})
		}

		allyTeamID, allyTeamOk := allyTeamIDForTimeline(level)
		for _, k := range level.CombatLog.Kill {
			if k == nil || k.IsEmpty() {
				continue
			}
			// Модули/дроны/объекты: в строке нет ника игрока в Killed.Name — только ObjectOwner (см. combat.Kill в тестах).
			if strings.TrimSpace(k.Killed.Name) == "" {
				continue
			}
			killer := strings.TrimSpace(k.Killer.Name)
			victim := killVictimDisplayName(k)
			killerShip := strings.TrimSpace(k.Killer.ObjectName)
			victimShip := strings.TrimSpace(k.Killed.ObjectName)
			t := k.GetTime(t0)
			if t.Before(t0) {
				t = t0
			}
			sec := t.Sub(t0).Seconds()
			victimTeam, inRoster := nameTeam[strings.ToLower(victim)]
			// NPC/объекты на карте: ник в логе не совпадает с ростером матча.
			if !inRoster {
				continue
			}
			allyDeath := victim != "" && allyTeamOk && len(nameTeam) > 0 && victimTeam == allyTeamID

			weapon := strings.TrimSpace(k.Source)
			assists := assistNamesForKill(level, t0, sec, victim, killer)

			if allyDeath {
				// Одно событие: смерть союзника (кто кого — как в убийстве). Отдельный «kill» не ставим — для врага он и есть смерть.
				lbl := victim
				if killer != "" {
					lbl = killer + " → " + victim
				}
				markers = append(markers, timelineMarker{
					Kind:       "death",
					TimeSec:    sec,
					Player:     victim,
					Killer:     killer,
					Victim:     victim,
					KillerShip: killerShip,
					VictimShip: victimShip,
					PlayerShip: victimShip,
					Weapon:     weapon,
					Assists:    assists,
					Label:      lbl,
				})
				continue
			}
			if killer != "" {
				markers = append(markers, timelineMarker{
					Kind:       "kill",
					TimeSec:    sec,
					Killer:     killer,
					Victim:     victim,
					KillerShip: killerShip,
					VictimShip: victimShip,
					Weapon:     weapon,
					Assists:    assists,
					Label:      killer + " → " + victim,
				})
			}
		}
	}

	sort.Slice(markers, func(i, j int) bool {
		if markers[i].TimeSec != markers[j].TimeSec {
			return markers[i].TimeSec < markers[j].TimeSec
		}
		return markers[i].Kind < markers[j].Kind
	})

	res := timelineResult{
		StartSec: 0,
		EndSec:   span,
		Markers:  markers,
	}
	b, err := json.Marshal(res)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalTimelineJSON", "err", err, "markers", len(markers))
		return "", err
	}
	s := string(b)
	r.lg.For(ctx).Debugw("marshalTimelineJSON", "markers", len(markers), "end_sec", span, "out_len", len(s))
	return s, nil
}

// leftRightTeamIDs совпадает с chartTeamSides: слева id 1 (если есть ростер), справа id 2, иначе общий fallback.
func leftRightTeamIDs(level *splitter.Level) (leftID, rightID int, ok bool) {
	l, r, _, _, ok := chartTeamSides(level)
	return l, r, ok
}

// spawnKindForTeam: тима 1 → spawn (как левый график), тима 2 → enemy_spawn; иначе две первые ненулевые команды по id.
func spawnKindForTeam(level *splitter.Level, teamID int) (kind string, ok bool) {
	if teamID == 0 {
		return "", false
	}
	t1, ok1 := level.Teams[1]
	t2, ok2 := level.Teams[2]
	if ok1 && len(t1) > 0 && ok2 && len(t2) > 0 {
		if teamID == 1 {
			return "spawn", true
		}
		if teamID == 2 {
			return "enemy_spawn", true
		}
		return "", false
	}
	left, right, ok := leftRightTeamIDs(level)
	if !ok {
		return "spawn", true
	}
	if teamID == left {
		return "spawn", true
	}
	if teamID == right {
		return "enemy_spawn", true
	}
	return "", false
}

// killVictimDisplayName prefers player nickname; for drones / FF lines the owner is in ObjectOwner.
func killVictimDisplayName(k *combat.Kill) string {
	if k == nil {
		return ""
	}
	n := strings.TrimSpace(k.Killed.Name)
	if n != "" {
		return n
	}
	return strings.TrimSpace(k.Killed.ObjectOwner)
}

// rosterNameToTeamID maps lowercased trimmed nickname → TeamID (combat ники могут отличаться регистром).
func rosterNameToTeamID(level *splitter.Level) map[string]int {
	if level == nil || len(level.Teams) == 0 {
		return nil
	}
	out := make(map[string]int)
	for tid, players := range level.Teams {
		if tid == 0 {
			continue
		}
		for _, p := range players {
			n := strings.TrimSpace(p.Name)
			if n == "" {
				continue
			}
			key := strings.ToLower(n)
			if _, exists := out[key]; exists {
				continue // одинаковый ник на двух командах — оставляем первое сопоставление
			}
			out[key] = tid
		}
	}
	return out
}

// allyTeamIDForTimeline uses LocalClientTeamID from combat Start when present; otherwise first sorted team (как левая диаграмма).
func allyTeamIDForTimeline(level *splitter.Level) (int, bool) {
	if level == nil {
		return 0, false
	}
	if level.CombatLog != nil && !level.CombatLog.Start.IsEmpty() && level.CombatLog.Start.LocalClientTeamID != 0 {
		tid := level.CombatLog.Start.LocalClientTeamID
		if _, ok := level.Teams[tid]; ok {
			return tid, true
		}
	}
	left, _, ok := leftRightTeamIDs(level)
	if !ok {
		return 0, false
	}
	return left, true
}

// assistVictimFromReason: хвост "… assist to kill VictimNick" в Reward.Reason.
func assistVictimFromReason(reason string) (string, bool) {
	if reason == "" {
		return "", false
	}
	lo := strings.ToLower(reason)
	sub := "assist to kill"
	i := strings.Index(lo, sub)
	if i < 0 {
		return "", false
	}
	rest := strings.TrimSpace(reason[i+len(sub):])
	fields := strings.Fields(rest)
	if len(fields) == 0 {
		return "", false
	}
	return fields[0], true
}

// assistNamesForKill: Reward с тем же жертвой в reason и близким временем; Recipient — ассистент (не киллер).
func assistNamesForKill(level *splitter.Level, t0 time.Time, killSec float64, victim, killer string) []string {
	if level == nil || level.CombatLog == nil || len(level.CombatLog.Reward) == 0 {
		return nil
	}
	victim = strings.TrimSpace(victim)
	killer = strings.TrimSpace(killer)
	if victim == "" {
		return nil
	}
	killLower := strings.ToLower(killer)
	seen := make(map[string]struct{})
	var out []string
	const win = 0.2
	for _, r := range level.CombatLog.Reward {
		if r == nil || r.IsEmpty() {
			continue
		}
		vic, ok := assistVictimFromReason(r.Reason)
		if !ok || !strings.EqualFold(vic, victim) {
			continue
		}
		rt := r.GetTime(t0)
		if rt.IsZero() {
			continue
		}
		if rt.Before(t0) {
			rt = t0
		}
		rsec := rt.Sub(t0).Seconds()
		if math.Abs(rsec-killSec) > win {
			continue
		}
		rec := strings.TrimSpace(r.Recipient)
		if rec == "" || strings.ToLower(rec) == killLower {
			continue
		}
		kl := strings.ToLower(rec)
		if _, dup := seen[kl]; dup {
			continue
		}
		seen[kl] = struct{}{}
		out = append(out, rec)
	}
	if len(out) == 0 {
		return nil
	}
	slices.Sort(out)
	return out
}
