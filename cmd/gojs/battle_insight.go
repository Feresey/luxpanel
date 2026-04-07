package main

import (
	"context"
	"encoding/json"
	"sort"
	"strings"
	"time"

	"github.com/Feresey/luxpanel/internal/splitter"
)

type lifePoint struct {
	TimeSec float64 `json:"t"`
	Allies  int     `json:"allies"`
	Enemies int     `json:"enemies"`
	Player  int     `json:"player,omitempty"` // 0/1: жив ли выбранный игрок в момент времени
}

type intensityPoint struct {
	TimeSec   float64 `json:"t"`
	AllyDPS   float64 `json:"ally"` // средний урон за 10 с (кросс-команда), со стороны союзников
	EnemyDPS  float64 `json:"enemy"`
	PlayerOut float64 `json:"player_out,omitempty"` // исходящий урон выбранного игрока в сек
	PlayerIn  float64 `json:"player_in,omitempty"`  // входящий урон по выбранному игроку в сек
}

type battleInsightResult struct {
	EndSec         float64          `json:"end_sec"`
	AllyTeamID     int              `json:"ally_team_id,omitempty"`
	EnemyTeamID    int              `json:"enemy_team_id,omitempty"`
	AllyTeamLabel  string           `json:"ally_team_label"`
	EnemyTeamLabel string           `json:"enemy_team_label"`
	FocusedPlayer  string           `json:"focused_player,omitempty"`
	Life           []lifePoint      `json:"life"`
	Intensity      []intensityPoint `json:"intensity"`
}

const intensitySampleStepSec = 0.25
const intensityWindowSec = 10.0

type lifeEvent struct {
	t    time.Time
	kind int // 0 spawn ally, 1 spawn enemy, 2 ally death, 3 enemy death
	name string
}

func (r *Runtime) marshalBattleInsightJSON(ctx context.Context, level *splitter.Level, timeRangeJSON string, focusedPlayer string) (string, error) {
	if level == nil {
		b, err := json.Marshal(battleInsightResult{})
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	lo, hi := clampTimeRange(level, timeRangeJSON)
	span := levelSpanSeconds(level)
	focusedPlayer = strings.TrimSpace(focusedPlayer)
	focusKey := strings.ToLower(focusedPlayer)
	hasFocus := focusKey != ""

	allyID, enemyID, ok := resolveAllyEnemyTeamIDs(level)
	if !ok {
		res := battleInsightResult{EndSec: span, AllyTeamLabel: "Союзники", EnemyTeamLabel: "Противники"}
		b, err := json.Marshal(res)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	t0 := level.StartLevelTime
	nameTeam := rosterNameToTeamID(level)
	nameCanon := make(map[string]string)
	if level != nil {
		for _, pl := range level.Teams {
			for _, p := range pl {
				n := strings.TrimSpace(p.Name)
				if n == "" {
					continue
				}
				k := strings.ToLower(n)
				if _, ok := nameCanon[k]; !ok {
					nameCanon[k] = n
				}
			}
		}
	}
	if hasFocus {
		if canon, ok := nameCanon[focusKey]; ok {
			focusedPlayer = canon
		}
	}

	var events []lifeEvent

	if level.CombatLog != nil {
		for _, sp := range level.CombatLog.Spawn {
			if sp == nil {
				continue
			}
			name := strings.TrimSpace(sp.Name)
			if name == "" {
				continue
			}
			tid, ok := nameTeam[strings.ToLower(name)]
			if !ok || tid == 0 {
				continue
			}
			t := sp.GetTime(t0)
			if t.Before(t0) {
				t = t0
			}
			if tid == allyID {
				r.lg.For(ctx).Debugw(
					"life_event_source_line",
					"line_type", "spawn",
					"time", sp.Time.Time,
					"time_sec", t.Sub(t0).Seconds(),
					"kind", "ally_spawn",
					"player", name,
					"ship", strings.TrimSpace(sp.Ship),
					"team_id", tid,
				)
				events = append(events, lifeEvent{t: t, kind: 0, name: name})
			} else if tid == enemyID {
				r.lg.For(ctx).Debugw(
					"life_event_source_line",
					"line_type", "spawn",
					"time", sp.Time.Time,
					"time_sec", t.Sub(t0).Seconds(),
					"kind", "enemy_spawn",
					"player", name,
					"ship", strings.TrimSpace(sp.Ship),
					"team_id", tid,
				)
				events = append(events, lifeEvent{t: t, kind: 1, name: name})
			}
		}

		for _, k := range level.CombatLog.Kill {
			if k == nil || k.IsEmpty() {
				continue
			}
			if strings.TrimSpace(k.Killed.Name) == "" {
				continue
			}
			killer := strings.TrimSpace(k.Killer.Name)
			victim := killVictimDisplayName(k)
			t := k.GetTime(t0)
			if t.Before(t0) {
				t = t0
			}
			victimTeam, inRoster := nameTeam[strings.ToLower(victim)]
			if !inRoster {
				continue
			}
			allyDeath := victim != "" && len(nameTeam) > 0 && victimTeam == allyID

			if allyDeath {
				r.lg.For(ctx).Debugw(
					"life_event_source_line",
					"line_type", "kill",
					"time", k.Time.Time,
					"time_sec", t.Sub(t0).Seconds(),
					"kind", "ally_death",
					"killer", killer,
					"victim", victim,
					"victim_team_id", victimTeam,
				)
				events = append(events, lifeEvent{t: t, kind: 2, name: victim})
				continue
			}
			if killer != "" && victim != "" {
				r.lg.For(ctx).Debugw(
					"life_event_source_line",
					"line_type", "kill",
					"time", k.Time.Time,
					"time_sec", t.Sub(t0).Seconds(),
					"kind", "enemy_death",
					"killer", killer,
					"victim", victim,
					"victim_team_id", victimTeam,
				)
				events = append(events, lifeEvent{t: t, kind: 3, name: victim})
			}
		}
	}

	// При одинаковом времени сначала применяем смерти, потом спавны.
	// Иначе на метке "смерть" линия может визуально идти вверх из-за спавнов в ту же секунду.
	lifeEventOrder := map[int]int{
		2: 0, // ally death
		3: 1, // enemy death
		0: 2, // ally spawn
		1: 3, // enemy spawn
	}
	sort.Slice(events, func(i, j int) bool {
		if !events[i].t.Equal(events[j].t) {
			return events[i].t.Before(events[j].t)
		}
		pi := lifeEventOrder[events[i].kind]
		pj := lifeEventOrder[events[j].kind]
		if pi != pj {
			return pi < pj
		}
		return events[i].name < events[j].name
	})

	allyRoster := level.Teams[allyID]
	enemyRoster := level.Teams[enemyID]

	countAlly := func(m map[string]bool) int {
		n := 0
		for _, p := range allyRoster {
			if m[strings.TrimSpace(p.Name)] {
				n++
			}
		}
		return n
	}
	countEnemy := func(m map[string]bool) int {
		n := 0
		for _, p := range enemyRoster {
			if m[strings.TrimSpace(p.Name)] {
				n++
			}
		}
		return n
	}

	allyAlive := make(map[string]bool)
	enemyAlive := make(map[string]bool)
	playerAlive := false
	if hasFocus {
		// Если первый найденный эвент игрока — смерть, значит до него игрок был жив,
		// даже если в логе нет явного spawn (или он не сопоставился по нику).
		for _, ev := range events {
			if strings.ToLower(strings.TrimSpace(ev.name)) != focusKey {
				continue
			}
			if ev.kind == 2 || ev.kind == 3 {
				playerAlive = true
			}
			break
		}
	}

	var fullLife []lifePoint
	fullLife = append(fullLife, lifePoint{TimeSec: 0, Allies: 0, Enemies: 0, Player: 0})

	for _, ev := range events {
		switch ev.kind {
		case 0:
			allyAlive[ev.name] = true
		case 1:
			enemyAlive[ev.name] = true
		case 2:
			allyAlive[ev.name] = false
		case 3:
			enemyAlive[ev.name] = false
		}
		if hasFocus && strings.ToLower(strings.TrimSpace(ev.name)) == focusKey {
			playerAlive = ev.kind == 0 || ev.kind == 1
		}
		fullLife = append(fullLife, lifePoint{
			TimeSec: ev.t.Sub(t0).Seconds(),
			Allies:  countAlly(allyAlive),
			Enemies: countEnemy(enemyAlive),
			Player: func() int {
				if playerAlive {
					return 1
				}
				return 0
			}(),
		})
	}

	lifeOut := clipLifeSeries(fullLife, lo, hi)

	// Интенсивность: скользящее среднее урона за 10 с (кросс-командный урон, как в графиках)
	var dmgList []dmgEv
	var playerOutList []dmgAmtEv
	var playerInList []dmgAmtEv
	if level.CombatLog != nil {
		for _, dmg := range level.CombatLog.Damage {
			if dmg == nil || dmg.IsEmpty() {
				continue
			}
			srcName := strings.TrimSpace(dmg.Initiator.Name)
			tgtName := strings.TrimSpace(dmg.Recipient.Name)
			st, okS := nameTeam[strings.ToLower(srcName)]
			tt, okT := nameTeam[strings.ToLower(tgtName)]
			if !okS || !okT || st == tt {
				continue
			}
			amount := float64(dmg.DamageFull)
			t := dmg.GetTime(t0)
			if t.Before(t0) {
				t = t0
			}
			var teamMark int
			if st == allyID {
				teamMark = 0
			} else if st == enemyID {
				teamMark = 1
			} else {
				continue
			}
			dmgList = append(dmgList, dmgEv{t: t, amt: amount, tid: teamMark})
			if hasFocus && strings.EqualFold(srcName, focusedPlayer) {
				playerOutList = append(playerOutList, dmgAmtEv{t: t, amt: amount})
			}
			if hasFocus && strings.EqualFold(tgtName, focusedPlayer) {
				playerInList = append(playerInList, dmgAmtEv{t: t, amt: amount})
			}
		}
	}
	sort.Slice(dmgList, func(i, j int) bool {
		return dmgList[i].t.Before(dmgList[j].t)
	})

	intensityOut := sampleIntensity(t0, dmgList, playerOutList, playerInList, lo, hi)

	res := battleInsightResult{
		EndSec:         span,
		AllyTeamID:     allyID,
		EnemyTeamID:    enemyID,
		AllyTeamLabel:  "Союзники",
		EnemyTeamLabel: "Противники",
		FocusedPlayer:  focusedPlayer,
		Life:           lifeOut,
		Intensity:      intensityOut,
	}
	b, err := json.Marshal(res)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalBattleInsightJSON", "err", err)
		return "", err
	}
	s := string(b)
	r.lg.For(ctx).Debugw(
		"marshalBattleInsightJSON",
		"focus", focusedPlayer,
		"range_lo_sec", lo,
		"range_hi_sec", hi,
		"life_points", len(lifeOut),
		"intensity_points", len(intensityOut),
		"raw_damage_events", len(dmgList),
		"player_out_events", len(playerOutList),
		"player_in_events", len(playerInList),
		"out_len", len(s),
	)
	return s, nil
}

func clipLifeSeries(full []lifePoint, lo, hi float64) []lifePoint {
	if len(full) == 0 {
		return nil
	}
	var last lifePoint
	for _, p := range full {
		if p.TimeSec <= lo {
			last = p
		} else {
			break
		}
	}
	out := []lifePoint{{TimeSec: lo, Allies: last.Allies, Enemies: last.Enemies, Player: last.Player}}
	for _, p := range full {
		if p.TimeSec <= lo {
			continue
		}
		if p.TimeSec > hi {
			break
		}
		if len(out) > 0 && p.TimeSec == out[len(out)-1].TimeSec {
			out[len(out)-1] = p
			continue
		}
		out = append(out, p)
	}
	if len(out) == 0 {
		return out
	}
	lastPt := out[len(out)-1]
	if lastPt.TimeSec < hi-1e-9 {
		out = append(out, lifePoint{TimeSec: hi, Allies: lastPt.Allies, Enemies: lastPt.Enemies, Player: lastPt.Player})
	}
	return out
}

type dmgEv struct {
	t   time.Time
	amt float64
	tid int // 0 = союзная команда, 1 = вражеская (инициатор)
}

type dmgAmtEv struct {
	t   time.Time
	amt float64
}

func sampleIntensity(t0 time.Time, dmgList []dmgEv, playerOutList, playerInList []dmgAmtEv, lo, hi float64) []intensityPoint {
	if hi < lo {
		lo, hi = hi, lo
	}
	if hi-lo < 1e-9 {
		return nil
	}
	step := intensitySampleStepSec
	maxPts := int((hi-lo)/step) + 2
	if maxPts > 12000 {
		step = (hi - lo) / 8000
		if step < 0.05 {
			step = 0.05
		}
	}

	var out []intensityPoint
	head := 0
	tail := 0
	headPO := 0
	tailPO := 0
	headPI := 0
	tailPI := 0
	sumAlly := 0.0
	sumEnemy := 0.0
	sumPO := 0.0
	sumPI := 0.0
	for t := lo; t <= hi+1e-9; t += step {
		curTime := t0.Add(time.Duration(t * float64(time.Second)))
		winLoTime := curTime.Add(-time.Duration(intensityWindowSec * float64(time.Second)))
		for head < len(dmgList) && (dmgList[head].t.Before(winLoTime) || dmgList[head].t.Equal(winLoTime)) {
			if dmgList[head].tid == 0 {
				sumAlly -= dmgList[head].amt
			} else {
				sumEnemy -= dmgList[head].amt
			}
			head++
		}
		for tail < len(dmgList) && (dmgList[tail].t.Before(curTime) || dmgList[tail].t.Equal(curTime)) {
			if dmgList[tail].tid == 0 {
				sumAlly += dmgList[tail].amt
			} else {
				sumEnemy += dmgList[tail].amt
			}
			tail++
		}
		for headPO < len(playerOutList) && (playerOutList[headPO].t.Before(winLoTime) || playerOutList[headPO].t.Equal(winLoTime)) {
			sumPO -= playerOutList[headPO].amt
			headPO++
		}
		for tailPO < len(playerOutList) && (playerOutList[tailPO].t.Before(curTime) || playerOutList[tailPO].t.Equal(curTime)) {
			sumPO += playerOutList[tailPO].amt
			tailPO++
		}
		for headPI < len(playerInList) && (playerInList[headPI].t.Before(winLoTime) || playerInList[headPI].t.Equal(winLoTime)) {
			sumPI -= playerInList[headPI].amt
			headPI++
		}
		for tailPI < len(playerInList) && (playerInList[tailPI].t.Before(curTime) || playerInList[tailPI].t.Equal(curTime)) {
			sumPI += playerInList[tailPI].amt
			tailPI++
		}
		out = append(out, intensityPoint{
			TimeSec:   t,
			AllyDPS:   sumAlly / intensityWindowSec,
			EnemyDPS:  sumEnemy / intensityWindowSec,
			PlayerOut: sumPO / intensityWindowSec,
			PlayerIn:  sumPI / intensityWindowSec,
		})
	}
	return out
}
