package main

import (
	"context"
	"encoding/json"
	"sort"
	"strings"

	"github.com/Feresey/luxpanel/internal/splitter"
)

type lifePoint struct {
	TimeSec float64 `json:"t"`
	Allies  int     `json:"allies"`
	Enemies int     `json:"enemies"`
}

type intensityPoint struct {
	TimeSec  float64 `json:"t"`
	AllyDPS  float64 `json:"ally"`  // средний урон за 10 с (кросс-команда), со стороны союзников
	EnemyDPS float64 `json:"enemy"`
}

type battleInsightResult struct {
	EndSec         float64          `json:"end_sec"`
	AllyTeamLabel  string           `json:"ally_team_label"`
	EnemyTeamLabel string           `json:"enemy_team_label"`
	Life           []lifePoint      `json:"life"`
	Intensity      []intensityPoint `json:"intensity"`
}

const intensitySampleStepSec = 0.25
const intensityWindowSec = 10.0

type lifeEvent struct {
	t    float64
	kind int // 0 spawn ally, 1 spawn enemy, 2 ally death, 3 enemy death
	name string
}

func (r *Runtime) marshalBattleInsightJSON(ctx context.Context, level *splitter.Level, timeRangeJSON string) (string, error) {
	if level == nil {
		b, err := json.Marshal(battleInsightResult{})
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	lo, hi := clampTimeRange(level, timeRangeJSON)
	span := levelSpanSeconds(level)

	allyID, allyOk := allyTeamIDForTimeline(level)
	id0, id1, twoOk := leftRightTeamIDs(level)
	if !allyOk || !twoOk || allyID == 0 {
		res := battleInsightResult{EndSec: span, AllyTeamLabel: "Союзники", EnemyTeamLabel: "Противники"}
		b, err := json.Marshal(res)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	enemyID := id1
	if allyID == id1 {
		enemyID = id0
	}
	if enemyID == allyID {
		res := battleInsightResult{EndSec: span, AllyTeamLabel: "Союзники", EnemyTeamLabel: "Противники"}
		b, err := json.Marshal(res)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	t0 := level.StartLevelTime
	nameTeam := rosterNameToTeamID(level)

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
			sec := t.Sub(t0).Seconds()
			if tid == allyID {
				events = append(events, lifeEvent{t: sec, kind: 0, name: name})
			} else if tid == enemyID {
				events = append(events, lifeEvent{t: sec, kind: 1, name: name})
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
			sec := t.Sub(t0).Seconds()
			victimTeam, inRoster := nameTeam[strings.ToLower(victim)]
			if !inRoster {
				continue
			}
			allyDeath := victim != "" && len(nameTeam) > 0 && victimTeam == allyID

			if allyDeath {
				events = append(events, lifeEvent{t: sec, kind: 2, name: victim})
				continue
			}
			if killer != "" && victim != "" {
				events = append(events, lifeEvent{t: sec, kind: 3, name: victim})
			}
		}
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].t != events[j].t {
			return events[i].t < events[j].t
		}
		return events[i].kind < events[j].kind
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

	var fullLife []lifePoint
	fullLife = append(fullLife, lifePoint{TimeSec: 0, Allies: 0, Enemies: 0})

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
		fullLife = append(fullLife, lifePoint{
			TimeSec: ev.t,
			Allies:  countAlly(allyAlive),
			Enemies: countEnemy(enemyAlive),
		})
	}

	lifeOut := clipLifeSeries(fullLife, lo, hi)

	// Интенсивность: скользящее среднее урона за 10 с (кросс-командный урон, как в графиках)
	var dmgList []dmgEv
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
			sec := t.Sub(t0).Seconds()
			var teamMark int
			if st == allyID {
				teamMark = 0
			} else if st == enemyID {
				teamMark = 1
			} else {
				continue
			}
			dmgList = append(dmgList, dmgEv{t: sec, amt: amount, tid: teamMark})
		}
	}
	sort.Slice(dmgList, func(i, j int) bool {
		return dmgList[i].t < dmgList[j].t
	})

	intensityOut := sampleIntensity(dmgList, lo, hi)

	res := battleInsightResult{
		EndSec:         span,
		AllyTeamLabel:  "Союзники",
		EnemyTeamLabel: "Противники",
		Life:           lifeOut,
		Intensity:      intensityOut,
	}
	b, err := json.Marshal(res)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalBattleInsightJSON", "err", err)
		return "", err
	}
	return string(b), nil
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
	out := []lifePoint{{TimeSec: lo, Allies: last.Allies, Enemies: last.Enemies}}
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
		out = append(out, lifePoint{TimeSec: hi, Allies: lastPt.Allies, Enemies: lastPt.Enemies})
	}
	return out
}

type dmgEv struct {
	t   float64
	amt float64
	tid int // 0 = союзная команда, 1 = вражеская (инициатор)
}

func sampleIntensity(dmgList []dmgEv, lo, hi float64) []intensityPoint {
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
	sumAlly := 0.0
	sumEnemy := 0.0
	for t := lo; t <= hi+1e-9; t += step {
		winLo := t - intensityWindowSec
		for head < len(dmgList) && dmgList[head].t <= winLo {
			if dmgList[head].tid == 0 {
				sumAlly -= dmgList[head].amt
			} else {
				sumEnemy -= dmgList[head].amt
			}
			head++
		}
		for tail < len(dmgList) && dmgList[tail].t <= t {
			if dmgList[tail].tid == 0 {
				sumAlly += dmgList[tail].amt
			} else {
				sumEnemy += dmgList[tail].amt
			}
			tail++
		}
		out = append(out, intensityPoint{
			TimeSec:  t,
			AllyDPS:  sumAlly / intensityWindowSec,
			EnemyDPS: sumEnemy / intensityWindowSec,
		})
	}
	return out
}
