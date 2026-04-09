package splitter

import (
	"context"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

// matchSpanSeconds mirrors cmd/gojs levelSpanSeconds for logging (same empty charts if 0).
func matchSpanSeconds(lvl *Level) float64 {
	if lvl == nil {
		return 0
	}
	if lvl.StartLevelTime.IsZero() || lvl.EndLevelTime.IsZero() {
		return 0
	}
	d := lvl.EndLevelTime.Sub(lvl.StartLevelTime)
	if d < 0 {
		return 0
	}
	span := float64(d) / float64(time.Second)
	if math.IsNaN(span) || math.IsInf(span, 0) || span < 0 {
		return 0
	}
	const hardMaxMatchSpanSec = 12 * 60 * 60
	if span > hardMaxMatchSpanSec {
		return hardMaxMatchSpanSec
	}
	return span
}

func sortedNonEmptyTeamIDsForDiag(lvl *Level) []int {
	if lvl == nil {
		return nil
	}
	var ids []int
	for id, pl := range lvl.Teams {
		if id == 0 {
			continue
		}
		if len(pl) > 0 {
			ids = append(ids, id)
		}
	}
	slices.Sort(ids)
	return ids
}

// chartTeamSidesDiagnostic mirrors cmd/gojs chartTeamSides (без импорта main).
func chartTeamSidesDiagnostic(lvl *Level) (allyID, enemyID int, nAlly, nEnemy int, ok bool) {
	if lvl == nil || len(lvl.Teams) == 0 {
		return 0, 0, 0, 0, false
	}
	var allyOK bool
	if lvl.CombatLog != nil && !lvl.CombatLog.Start.IsEmpty() && lvl.CombatLog.Start.LocalClientTeamID != 0 {
		tid := lvl.CombatLog.Start.LocalClientTeamID
		if _, ok := lvl.Teams[tid]; ok {
			allyID = tid
			allyOK = true
		}
	}
	if !allyOK || allyID == 0 {
		return 0, 0, 0, 0, false
	}
	pla := lvl.Teams[allyID]
	if len(pla) == 0 {
		return 0, 0, 0, 0, false
	}
	for _, id := range sortedNonEmptyTeamIDsForDiag(lvl) {
		if id == allyID {
			continue
		}
		plb := lvl.Teams[id]
		if len(plb) > 0 {
			return allyID, id, len(pla), len(plb), true
		}
	}
	return 0, 0, 0, 0, false
}

func humanPlayers(pl []Player) []Player {
	out := make([]Player, 0, len(pl))
	for _, p := range pl {
		if p.PlayerID != 0 {
			out = append(out, p)
		}
	}
	return out
}

// damageEventsForCharts mirrors buildDamageCharts filters (люди, две стороны, кросс-сторона, окно времени).
func damageEventsForCharts(lvl *Level, logTime time.Time) (eligible int, sidesOK bool, reason string) {
	if lvl == nil || lvl.CombatLog == nil {
		return 0, false, "nil_level_or_combat"
	}
	allyID, enemyID, _, _, ok := chartTeamSidesDiagnostic(lvl)
	if !ok {
		return 0, false, "chart_team_sides_fail"
	}
	pla := humanPlayers(lvl.Teams[allyID])
	plb := humanPlayers(lvl.Teams[enemyID])
	if len(pla) == 0 || len(plb) == 0 {
		return 0, true, "empty_human_roster_on_side"
	}
	nameToTeam := make(map[string]int)
	for _, p := range pla {
		nameToTeam[strings.ToLower(strings.TrimSpace(p.Name))] = 0
	}
	for _, p := range plb {
		nameToTeam[strings.ToLower(strings.TrimSpace(p.Name))] = 1
	}
	t0 := lvl.StartLevelTime
	span := matchSpanSeconds(lvl)
	hi := span
	for _, dmg := range lvl.CombatLog.Damage {
		if dmg == nil || dmg.IsEmpty() {
			continue
		}
		if dmg.Initiator.Name == "" || dmg.Recipient.Name == "" {
			continue
		}
		sn := strings.ToLower(strings.TrimSpace(dmg.Initiator.Name))
		tgn := strings.ToLower(strings.TrimSpace(dmg.Recipient.Name))
		st, okS := nameToTeam[sn]
		tt, okT := nameToTeam[tgn]
		if !okS || !okT || st == tt {
			continue
		}
		t := dmg.GetTime(logTime)
		if t.Before(t0) {
			t = t0
		}
		sec := t.Sub(t0).Seconds()
		if sec < 0 || sec > hi {
			continue
		}
		eligible++
	}
	if eligible == 0 && len(lvl.CombatLog.Damage) > 0 {
		return 0, true, "no_damage_passes_roster_time_filters"
	}
	return eligible, true, ""
}

func teamSizesString(lvl *Level) string {
	if lvl == nil || len(lvl.Teams) == 0 {
		return ""
	}
	ids := mapsKeysSorted(lvl.Teams)
	var b strings.Builder
	for i, id := range ids {
		if i > 0 {
			b.WriteByte('|')
		}
		pl := lvl.Teams[id]
		h := 0
		for _, p := range pl {
			if p.PlayerID != 0 {
				h++
			}
		}
		b.WriteString(strconv.Itoa(id))
		b.WriteString(":n=")
		b.WriteString(strconv.Itoa(len(pl)))
		b.WriteString(",humans=")
		b.WriteString(strconv.Itoa(h))
	}
	return b.String()
}

func mapsKeysSorted(m map[int][]Player) []int {
	ids := make([]int, 0, len(m))
	for k := range m {
		ids = append(ids, k)
	}
	slices.Sort(ids)
	return ids
}

// LogMatchParseStats пишет статистику по матчу сразу после сборки Level (для отладки пустых графиков).
func (s *Splitter) LogMatchParseStats(ctx context.Context, matchIdx, nMatches int, logTime time.Time, lvl *Level) {
	if lvl == nil {
		return
	}
	last := nMatches > 0 && matchIdx == nMatches-1
	cl := lvl.CombatLog
	span := matchSpanSeconds(lvl)
	eligible, sidesOK, chartReason := damageEventsForCharts(lvl, logTime)
	allyID, enemyID, nAlly, nEnemy, _ := chartTeamSidesDiagnostic(lvl)
	hAlly := 0
	hEnemy := 0
	if allyID != 0 {
		hAlly = len(humanPlayers(lvl.Teams[allyID]))
	}
	if enemyID != 0 {
		hEnemy = len(humanPlayers(lvl.Teams[enemyID]))
	}

	var st, en string
	if !lvl.StartLevelTime.IsZero() {
		st = lvl.StartLevelTime.UTC().Format(time.RFC3339Nano)
	}
	if !lvl.EndLevelTime.IsZero() {
		en = lvl.EndLevelTime.UTC().Format(time.RFC3339Nano)
	}

	dmgN, healN, killN, spawnN, spellN := 0, 0, 0, 0, 0
	var connEmpty, startEmpty, finEmpty bool
	if cl != nil {
		dmgN = len(cl.Damage)
		healN = len(cl.Heal)
		killN = len(cl.Kill)
		spawnN = len(cl.Spawn)
		spellN = len(cl.Spell)
		connEmpty = cl.Connect.IsEmpty()
		startEmpty = cl.Start.IsEmpty()
		finEmpty = cl.Finished.IsEmpty()
	}

	gl := lvl.GameLog
	gameStart := gl != nil && gl.StartGameplay != nil
	gameFinish := gl != nil && gl.FinishGameplay != nil
	addN := 0
	if gl != nil {
		addN = len(gl.AddPlayer)
	}

	localTeam := 0
	if cl != nil && !cl.Start.IsEmpty() {
		localTeam = cl.Start.LocalClientTeamID
	}

	lg := s.lg.For(ctx)
	lg.Infow("match_parse_stats",
		"match_index", matchIdx,
		"is_last_match", last,
		"damage_lines", dmgN,
		"heal_lines", healN,
		"kill_lines", killN,
		"spawn_lines", spawnN,
		"spell_lines", spellN,
		"game_has_start_gameplay", gameStart,
		"game_has_finish_gameplay", gameFinish,
		"game_add_player_n", addN,
		"combat_connect_empty", connEmpty,
		"combat_start_empty", startEmpty,
		"combat_finished_empty", finEmpty,
		"local_client_team_id", localTeam,
		"start_level_utc", st,
		"end_level_utc", en,
		"span_sec", span,
		"chart_team_sides_ok", sidesOK,
		"ally_team_id", allyID,
		"enemy_team_id", enemyID,
		"ally_roster_n", nAlly,
		"enemy_roster_n", nEnemy,
		"ally_humans_n", hAlly,
		"enemy_humans_n", hEnemy,
		"damage_chart_eligible_n", eligible,
		"chart_filter_reason", chartReason,
		"team_sizes", teamSizesString(lvl),
	)

	if last && dmgN > 0 && eligible == 0 {
		lg.Warnw("last_match_charts_likely_empty",
			"damage_lines", dmgN,
			"span_sec", span,
			"chart_team_sides_ok", sidesOK,
			"chart_filter_reason", chartReason,
			"local_client_team_id", localTeam,
		)
	}
}
