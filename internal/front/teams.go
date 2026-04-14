package front

import (
	"slices"
	"strings"

	"github.com/Feresey/luxpanel/internal/splitter"
	"golang.org/x/exp/maps"
)

func SortedNonEmptyTeamIDs(level *splitter.Level) []int {
	if level == nil {
		return nil
	}
	ids := maps.Keys(level.Teams)
	slices.Sort(ids)
	var out []int
	for _, id := range ids {
		if id == 0 {
			continue
		}
		if pl := level.Teams[id]; len(pl) > 0 {
			out = append(out, id)
		}
	}
	return out
}

func AllyTeamIDForTimeline(level *splitter.Level) (int, bool) {
	if level == nil {
		return 0, false
	}
	if level.CombatLog != nil && !level.CombatLog.Start.IsEmpty() && level.CombatLog.Start.LocalClientTeamID != 0 {
		tid := level.CombatLog.Start.LocalClientTeamID
		if _, ok := level.Teams[tid]; ok {
			return tid, true
		}
	}
	return 0, false
}

func ChartTeamSides(level *splitter.Level) (leftID, rightID int, pla, plb []splitter.Player, ok bool) {
	if level == nil || len(level.Teams) == 0 {
		return 0, 0, nil, nil, false
	}
	allyID, okAlly := AllyTeamIDForTimeline(level)
	if !okAlly || allyID == 0 {
		return 0, 0, nil, nil, false
	}
	pla = level.Teams[allyID]
	if len(pla) == 0 {
		return 0, 0, nil, nil, false
	}
	for _, id := range SortedNonEmptyTeamIDs(level) {
		if id == 0 || id == allyID {
			continue
		}
		plb = level.Teams[id]
		if len(plb) > 0 {
			return allyID, id, pla, plb, true
		}
	}
	return 0, 0, nil, nil, false
}

func ChartTeams(level *splitter.Level) (pla, plb []splitter.Player, ok bool) {
	_, _, a, b, ok := ChartTeamSides(level)
	return a, b, ok
}

func LeftRightTeamIDs(level *splitter.Level) (leftID, rightID int, ok bool) {
	l, r, _, _, ok := ChartTeamSides(level)
	return l, r, ok
}

func ResolveAllyEnemyTeamIDs(level *splitter.Level) (allyID, enemyID int, ok bool) {
	allyID, allyOK := AllyTeamIDForTimeline(level)
	leftID, rightID, lrOK := LeftRightTeamIDs(level)
	if !allyOK || !lrOK || allyID == 0 {
		return 0, 0, false
	}
	enemyID = rightID
	if allyID == rightID {
		enemyID = leftID
	}
	if enemyID == 0 || enemyID == allyID {
		return 0, 0, false
	}
	return allyID, enemyID, true
}

func FilterRosterPlayers(players []splitter.Player, includeBots bool) []splitter.Player {
	if includeBots {
		return players
	}
	out := make([]splitter.Player, 0, len(players))
	for _, p := range players {
		if p.PlayerID != 0 {
			out = append(out, p)
		}
	}
	return out
}

func TwoTeamRosters(level *splitter.Level, includeBots bool) (
	pla, plb []splitter.Player,
	nameToTeam map[string]int,
	nameToIdx map[string]int,
	ok bool,
) {
	var okTeams bool
	pla, plb, okTeams = ChartTeams(level)
	if !okTeams {
		return nil, nil, nil, nil, false
	}
	pla = FilterRosterPlayers(pla, includeBots)
	plb = FilterRosterPlayers(plb, includeBots)
	if len(pla) == 0 || len(plb) == 0 {
		return nil, nil, nil, nil, false
	}
	nameToTeam = make(map[string]int)
	nameToIdx = make(map[string]int)
	for i, p := range pla {
		nameToTeam[p.Name] = 0
		nameToIdx[p.Name] = i
	}
	for i, p := range plb {
		nameToTeam[p.Name] = 1
		nameToIdx[p.Name] = i
	}
	return pla, plb, nameToTeam, nameToIdx, true
}

func RosterNameToTeamID(level *splitter.Level) map[string]int {
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
				continue
			}
			out[key] = tid
		}
	}
	return out
}
