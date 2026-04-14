package front

import (
	"math"
	"sort"
	"strings"
	"time"

	"github.com/Feresey/luxpanel/internal/splitter"
)

type LifePoint struct {
	TimeSec float64 `json:"t"`
	Allies  int     `json:"allies"`
	Enemies int     `json:"enemies"`
	Player  int     `json:"player,omitempty"`
}

type LifeTeamPoint struct {
	TimeSec float64 `json:"t"`
	Alive   int     `json:"v"`
}

type LifeTeamSeries struct {
	TeamID int             `json:"team_id"`
	Label  string          `json:"label"`
	Ally   bool            `json:"ally,omitempty"`
	Points []LifeTeamPoint `json:"points"`
}

type IntensityPoint struct {
	TimeSec   float64 `json:"t"`
	AllyDPS   float64 `json:"ally"`
	EnemyDPS  float64 `json:"enemy"`
	PlayerOut float64 `json:"player_out,omitempty"`
	PlayerIn  float64 `json:"player_in,omitempty"`
}

// BattleInsightResult is JSON for battle insight widget (team labels optional, filled in gojs if needed).
type BattleInsightResult struct {
	EndSec         float64          `json:"end_sec"`
	AllyTeamID     int              `json:"ally_team_id,omitempty"`
	EnemyTeamID    int              `json:"enemy_team_id,omitempty"`
	AllyTeamLabel  string           `json:"ally_team_label"`
	EnemyTeamLabel string           `json:"enemy_team_label"`
	GodGiftSec     float64          `json:"godgift_sec,omitempty"`
	GodGiftLabel   string           `json:"godgift_label,omitempty"`
	GodGiftSide    string           `json:"-"` // allies | enemies — localized in cmd/gojs
	FocusedPlayer  string           `json:"focused_player,omitempty"`
	LifeTeams      []LifeTeamSeries `json:"life_teams,omitempty"`
	Life           []LifePoint      `json:"life"`
	Intensity      []IntensityPoint `json:"intensity"`
}

const intensitySampleStepSec = 1.0
const intensityWindowSec = 10.0
const maxIntensitySamples = 12000

type lifeEvent struct {
	t      time.Time
	kind   int // 0 spawn, 1 death
	name   string
	teamID int
}

// BuildBattleInsight computes life/intensity series for optional time window and focus player.
func BuildBattleInsight(level *splitter.Level, timeRangeJSON string, focusedPlayer string) BattleInsightResult {
	if level == nil {
		return BattleInsightResult{}
	}
	lo, hi := ClampTimeRange(level, timeRangeJSON)
	span := LevelSpanSeconds(level)
	focusedPlayer = strings.TrimSpace(focusedPlayer)
	focusKey := strings.ToLower(focusedPlayer)
	hasFocus := focusKey != ""
	allyID, ok := AllyTeamIDForTimeline(level)
	if !ok {
		return BattleInsightResult{EndSec: span}
	}
	enemyID := 0
	for tid := range level.Teams {
		if tid != 0 && tid != allyID {
			enemyID = tid
			break
		}
	}
	t0 := level.StartLevelTime
	nameTeam := RosterNameToTeamID(level)
	nameCanon := make(map[string]string)
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
	if hasFocus {
		if canon, ok := nameCanon[focusKey]; ok {
			focusedPlayer = canon
		}
	}
	teamIDs := make([]int, 0, len(level.Teams))
	for tid := range level.Teams {
		if tid != 0 {
			teamIDs = append(teamIDs, tid)
		}
	}
	sort.Ints(teamIDs)
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
			events = append(events, lifeEvent{t: t, kind: 0, name: name, teamID: tid})
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
			t := k.GetTime(t0)
			if t.Before(t0) {
				t = t0
			}
			victimTeam, inRoster := nameTeam[strings.ToLower(victim)]
			if !inRoster {
				continue
			}
			if killer != "" && victim != "" && victimTeam != 0 {
				events = append(events, lifeEvent{t: t, kind: 1, name: victim, teamID: victimTeam})
			}
		}
	}
	lifeEventOrder := map[int]int{1: 0, 0: 1}
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
	teamAlive := make(map[int]map[string]bool, len(teamIDs))
	for _, tid := range teamIDs {
		teamAlive[tid] = make(map[string]bool)
	}
	countTeam := func(tid int) int {
		roster := level.Teams[tid]
		alive := teamAlive[tid]
		n := 0
		for _, p := range roster {
			if alive[strings.TrimSpace(p.Name)] {
				n++
			}
		}
		return n
	}
	teamSeriesFull := make(map[int][]LifeTeamPoint, len(teamIDs))
	for _, tid := range teamIDs {
		teamSeriesFull[tid] = []LifeTeamPoint{{TimeSec: 0, Alive: 0}}
	}
	playerAlive := false
	if hasFocus {
		for _, ev := range events {
			if strings.ToLower(strings.TrimSpace(ev.name)) != focusKey {
				continue
			}
			if ev.kind == 1 {
				playerAlive = true
			}
			break
		}
	}
	var fullLife []LifePoint
	fullLife = append(fullLife, LifePoint{TimeSec: 0, Allies: 0, Enemies: 0, Player: 0})
	for _, ev := range events {
		switch ev.kind {
		case 0:
			teamAlive[ev.teamID][ev.name] = true
		case 1:
			teamAlive[ev.teamID][ev.name] = false
		}
		if hasFocus && strings.ToLower(strings.TrimSpace(ev.name)) == focusKey {
			playerAlive = ev.kind == 0
		}
		allyCount := countTeam(allyID)
		enemyCount := 0
		for _, tid := range teamIDs {
			if tid != allyID {
				enemyCount += countTeam(tid)
			}
		}
		sec := ev.t.Sub(t0).Seconds()
		pv := 0
		if playerAlive {
			pv = 1
		}
		fullLife = append(fullLife, LifePoint{
			TimeSec: sec,
			Allies:  allyCount,
			Enemies: enemyCount,
			Player:  pv,
		})
		for _, tid := range teamIDs {
			teamSeriesFull[tid] = append(teamSeriesFull[tid], LifeTeamPoint{TimeSec: sec, Alive: countTeam(tid)})
		}
	}
	lifeOut := clipLifeSeries(fullLife, lo, hi)
	lifeTeams := make([]LifeTeamSeries, 0, len(teamIDs))
	for _, tid := range teamIDs {
		lifeTeams = append(lifeTeams, LifeTeamSeries{
			TeamID: tid,
			Label:  "",
			Ally:   tid == allyID,
			Points: clipLifeTeamSeries(teamSeriesFull[tid], lo, hi),
		})
	}
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
			} else if st != 0 {
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
	godSec, godSide, hasGod := detectGodGiftMoment(level, t0, nameTeam, allyID)
	res := BattleInsightResult{
		EndSec:        span,
		AllyTeamID:    allyID,
		EnemyTeamID:   enemyID,
		FocusedPlayer: focusedPlayer,
		LifeTeams:     lifeTeams,
		Life:          lifeOut,
		Intensity:     intensityOut,
	}
	if hasGod {
		res.GodGiftSec = godSec
		res.GodGiftSide = godSide
	}
	return res
}

func detectGodGiftMoment(level *splitter.Level, t0 time.Time, nameTeam map[string]int, allyID int) (sec float64, side string, ok bool) {
	if level == nil || level.CombatLog == nil || len(level.CombatLog.Spell) == 0 {
		return 0, "", false
	}
	for _, sp := range level.CombatLog.Spell {
		if sp == nil || !sp.GodGift {
			continue
		}
		t := sp.GetTime(t0)
		if t.Before(t0) {
			t = t0
		}
		s := t.Sub(t0).Seconds()
		for _, targetRaw := range sp.Targets {
			target := strings.TrimSpace(targetRaw)
			if target == "" {
				continue
			}
			tid := nameTeam[strings.ToLower(target)]
			if tid == allyID {
				return s, "allies", true
			}
			if tid != 0 && tid != allyID {
				return s, "enemies", true
			}
		}
	}
	return 0, "", false
}

func clipLifeTeamSeries(full []LifeTeamPoint, lo, hi float64) []LifeTeamPoint {
	if len(full) == 0 {
		return nil
	}
	last := full[0]
	for _, p := range full {
		if p.TimeSec <= lo {
			last = p
		} else {
			break
		}
	}
	out := []LifeTeamPoint{{TimeSec: lo, Alive: last.Alive}}
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
	lastPt := out[len(out)-1]
	if lastPt.TimeSec < hi-1e-9 {
		out = append(out, LifeTeamPoint{TimeSec: hi, Alive: lastPt.Alive})
	}
	return out
}

func clipLifeSeries(full []LifePoint, lo, hi float64) []LifePoint {
	if len(full) == 0 {
		return nil
	}
	var last LifePoint
	for _, p := range full {
		if p.TimeSec <= lo {
			last = p
		} else {
			break
		}
	}
	out := []LifePoint{{TimeSec: lo, Allies: last.Allies, Enemies: last.Enemies, Player: last.Player}}
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
		out = append(out, LifePoint{TimeSec: hi, Allies: lastPt.Allies, Enemies: lastPt.Enemies, Player: lastPt.Player})
	}
	return out
}

type dmgEv struct {
	t   time.Time
	amt float64
	tid int
}

type dmgAmtEv struct {
	t   time.Time
	amt float64
}

func sampleIntensity(t0 time.Time, dmgList []dmgEv, playerOutList, playerInList []dmgAmtEv, lo, hi float64) []IntensityPoint {
	if hi < lo {
		lo, hi = hi, lo
	}
	if math.IsNaN(lo) || math.IsInf(lo, 0) || math.IsNaN(hi) || math.IsInf(hi, 0) {
		return nil
	}
	if hi-lo < 1e-9 {
		return nil
	}
	step := intensitySampleStepSec
	span := hi - lo
	if span/step > maxIntensitySamples {
		step = math.Ceil(span / maxIntensitySamples)
		if step < intensitySampleStepSec {
			step = intensitySampleStepSec
		}
	}
	winDur := time.Duration(float64(time.Second) * intensityWindowSec)
	est := int(math.Ceil((hi-lo)/step)) + 1
	if est < 0 {
		est = 0
	}
	if est > maxIntensitySamples+1 {
		est = maxIntensitySamples + 1
	}
	out := make([]IntensityPoint, 0, est)
	la, ra := 0, 0
	sumAlly, sumEnemy := 0.0, 0.0
	lPO, rPO := 0, 0
	sumPO := 0.0
	lPI, rPI := 0, 0
	sumPI := 0.0
	for t := lo; t <= hi+1e-9; t += step {
		curTime := t0.Add(time.Duration(t * float64(time.Second)))
		winLoTime := curTime.Add(-winDur)
		for ra < len(dmgList) && !dmgList[ra].t.After(curTime) {
			if dmgList[ra].tid == 0 {
				sumAlly += dmgList[ra].amt
			} else {
				sumEnemy += dmgList[ra].amt
			}
			ra++
		}
		for la < ra && !dmgList[la].t.After(winLoTime) {
			if dmgList[la].tid == 0 {
				sumAlly -= dmgList[la].amt
			} else {
				sumEnemy -= dmgList[la].amt
			}
			la++
		}
		for rPO < len(playerOutList) && !playerOutList[rPO].t.After(curTime) {
			sumPO += playerOutList[rPO].amt
			rPO++
		}
		for lPO < rPO && !playerOutList[lPO].t.After(winLoTime) {
			sumPO -= playerOutList[lPO].amt
			lPO++
		}
		for rPI < len(playerInList) && !playerInList[rPI].t.After(curTime) {
			sumPI += playerInList[rPI].amt
			rPI++
		}
		for lPI < rPI && !playerInList[lPI].t.After(winLoTime) {
			sumPI -= playerInList[lPI].amt
			lPI++
		}
		out = append(out, IntensityPoint{
			TimeSec:   t,
			AllyDPS:   sumAlly / intensityWindowSec,
			EnemyDPS:  sumEnemy / intensityWindowSec,
			PlayerOut: sumPO / intensityWindowSec,
			PlayerIn:  sumPI / intensityWindowSec,
		})
	}
	return out
}
