package front

import (
	"math"
	"sort"
	"strings"
	"time"

	"slices"

	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/Feresey/luxpanel/internal/splitter"
)

type TimelineMarker struct {
	Kind         string   `json:"kind"`
	TimeSec      float64  `json:"time_sec"`
	TeamID       int      `json:"team_id,omitempty"`
	KillerTeamID int      `json:"killer_team_id,omitempty"`
	VictimTeamID int      `json:"victim_team_id,omitempty"`
	Label        string   `json:"label,omitempty"`
	Player       string   `json:"player,omitempty"`
	Killer       string   `json:"killer,omitempty"`
	Victim       string   `json:"victim,omitempty"`
	KillerShip   string   `json:"killer_ship,omitempty"`
	VictimShip   string   `json:"victim_ship,omitempty"`
	PlayerShip   string   `json:"player_ship,omitempty"`
	Weapon       string   `json:"weapon,omitempty"`
	Assists      []string `json:"assists,omitempty"`
}

type TimelineResult struct {
	StartSec    float64          `json:"start_sec"`
	EndSec      float64          `json:"end_sec"`
	StartUnixMs int64            `json:"start_unix_ms,omitempty"`
	AllyTeamID  int              `json:"ally_team_id,omitempty"`
	EnemyTeamID int              `json:"enemy_team_id,omitempty"`
	Markers     []TimelineMarker `json:"markers"`
}

// BuildTimeline assembles markers for the given level and optional focus player (case-insensitive).
func BuildTimeline(level *splitter.Level, focusedPlayer string) TimelineResult {
	if level == nil {
		return TimelineResult{}
	}
	t0 := level.StartLevelTime
	span := LevelSpanSeconds(level)
	focusKey := strings.ToLower(strings.TrimSpace(focusedPlayer))
	if focusKey != "" {
		if roster := RosterNameToTeamID(level); len(roster) > 0 {
			for n := range roster {
				if n == focusKey {
					focusKey = n
					break
				}
			}
		}
	}
	var markers []TimelineMarker
	allyTeamID, enemyTeamID, teamsOK := ResolveAllyEnemyTeamIDs(level)
	if level.CombatLog != nil {
		nameTeam := RosterNameToTeamID(level)
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
			kind, ok := SpawnKindForTeam(level, tid)
			if !ok {
				continue
			}
			t := sp.GetTime(t0)
			if t.Before(t0) {
				t = t0
			}
			sec := t.Sub(t0).Seconds()
			markers = append(markers, TimelineMarker{
				Kind:       kind,
				TimeSec:    sec,
				TeamID:     tid,
				Player:     name,
				Label:      name,
				PlayerShip: strings.TrimSpace(sp.Ship),
			})
		}
		for _, k := range level.CombatLog.Kill {
			if k == nil || k.IsEmpty() {
				continue
			}
			if strings.TrimSpace(k.Killed.Name) == "" {
				continue
			}
			killer := strings.TrimSpace(k.Killer.Name)
			victim := KillVictimDisplayName(k)
			killerShip := strings.TrimSpace(k.Killer.ObjectName)
			victimShip := strings.TrimSpace(k.Killed.ObjectName)
			t := k.GetTime(t0)
			if t.Before(t0) {
				t = t0
			}
			sec := t.Sub(t0).Seconds()
			victimTeam, inRoster := nameTeam[strings.ToLower(victim)]
			killerTeam := 0
			if killer != "" {
				killerTeam = nameTeam[strings.ToLower(killer)]
			}
			if !inRoster {
				continue
			}
			allyDeath := victim != "" && teamsOK && len(nameTeam) > 0 && victimTeam == allyTeamID
			weapon := strings.TrimSpace(k.Source)
			assists := AssistNamesForKill(level, t0, sec, victim, killer)
			if allyDeath {
				lbl := victim
				if killer != "" {
					lbl = killer + " → " + victim
				}
				markers = append(markers, TimelineMarker{
					Kind:         "death",
					TimeSec:      sec,
					TeamID:       victimTeam,
					KillerTeamID: killerTeam,
					VictimTeamID: victimTeam,
					Player:       victim,
					Killer:       killer,
					Victim:       victim,
					KillerShip:   killerShip,
					VictimShip:   victimShip,
					PlayerShip:   victimShip,
					Weapon:       weapon,
					Assists:      assists,
					Label:        lbl,
				})
				continue
			}
			if killer != "" {
				markers = append(markers, TimelineMarker{
					Kind:         "kill",
					TimeSec:      sec,
					TeamID:       killerTeam,
					KillerTeamID: killerTeam,
					VictimTeamID: victimTeam,
					Killer:       killer,
					Victim:       victim,
					KillerShip:   killerShip,
					VictimShip:   victimShip,
					Weapon:       weapon,
					Assists:      assists,
					Label:        killer + " → " + victim,
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
	if focusKey != "" {
		var filtered []TimelineMarker
		matches := func(s string) bool {
			return strings.ToLower(strings.TrimSpace(s)) == focusKey
		}
		for _, m := range markers {
			switch m.Kind {
			case "spawn", "enemy_spawn":
				if matches(m.Player) {
					filtered = append(filtered, m)
				}
			case "death":
				if matches(m.Victim) || matches(m.Player) {
					filtered = append(filtered, m)
				}
			case "kill":
				if matches(m.Killer) || matches(m.Victim) {
					filtered = append(filtered, m)
					continue
				}
				keep := false
				for _, a := range m.Assists {
					if matches(a) {
						keep = true
						break
					}
				}
				if keep {
					filtered = append(filtered, m)
				}
			default:
				if matches(m.Player) || matches(m.Killer) || matches(m.Victim) {
					filtered = append(filtered, m)
				}
			}
		}
		markers = filtered
	}
	res := TimelineResult{
		StartSec:    0,
		EndSec:      span,
		AllyTeamID:  allyTeamID,
		EnemyTeamID: enemyTeamID,
		Markers:     markers,
	}
	if !t0.IsZero() {
		res.StartUnixMs = t0.UnixMilli()
	}
	return res
}

// SpawnKindForTeam maps team id to spawn marker kind for UI.
func SpawnKindForTeam(level *splitter.Level, teamID int) (kind string, ok bool) {
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
	left, right, ok := LeftRightTeamIDs(level)
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

// KillVictimDisplayName prefers player nickname; for drones FF lines owner is in ObjectOwner.
func KillVictimDisplayName(k *combat.Kill) string {
	if k == nil {
		return ""
	}
	n := strings.TrimSpace(k.Killed.Name)
	if n != "" {
		return n
	}
	return strings.TrimSpace(k.Killed.ObjectOwner)
}

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

// AssistNamesForKill finds Reward lines with matching assist victim near kill time.
func AssistNamesForKill(level *splitter.Level, t0 time.Time, killSec float64, victim, killer string) []string {
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
