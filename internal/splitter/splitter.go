package splitter

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"iter"
	"os"
	"slices"
	"strings"
	"time"

	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/maps"

	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/Feresey/luxpanel/internal/parser/common"
	"github.com/Feresey/luxpanel/internal/parser/game"
)

var ErrLogsCorrupted = errors.New("logs corrupted")

type Splitter struct {
	lg     logger.Factory
	tr     trace.Tracer
	parser *parser.Parser
}

type lineRange struct {
	Start int
	End   int
}

type matchRange struct {
	Game   lineRange
	Combat lineRange
}

func NewSplitter(lg logger.Factory, tr trace.TracerProvider, parser *parser.Parser) *Splitter {
	return &Splitter{lg: lg, tr: tr.Tracer("splitter"), parser: parser}
}

func (s *Splitter) SplitLevels(ctx context.Context, fs fs.FS) ([]*Level, error) {
	ctx, span := s.tr.Start(ctx, "SplitLevels")
	defer span.End()

	logTime, matches, err := s.scanMatchRanges(ctx, fs)
	if err != nil {
		return nil, fmt.Errorf("scanMatchRanges: %w", err)
	}
	gameLevels, combatLevels, err := s.parseLevelsByRanges(ctx, fs, logTime, matches)
	if err != nil {
		return nil, fmt.Errorf("parseLevelsByRanges: %w", err)
	}

	levels := make([]*Level, 0, len(gameLevels))
	mx := len(matches)
	for i := range mx {
		var gm *GameLogLevel
		var cm *CombatLogLevel
		if i < len(gameLevels) {
			gm = gameLevels[i]
		}
		if i < len(combatLevels) {
			cm = combatLevels[i]
		}

		lvl, err := s.makeLevel(ctx, logTime, gm, cm)
		if err != nil {
			return nil, fmt.Errorf("makeLevel: %w", err)
		}
		levels = append(levels, lvl)
	}

	s.lg.For(ctx).Debugw("got games", "count", len(levels))
	return levels, nil
}

func (s *Splitter) scanMatchRanges(ctx context.Context, fs fs.FS) (time.Time, []matchRange, error) {
	ctx, span := s.tr.Start(ctx, "scanMatchRanges")
	defer span.End()

	combatLog, err := fs.Open("combat.log")
	if err != nil {
		return time.Time{}, nil, fmt.Errorf("fs.Open(combat.log): %w", err)
	}
	defer combatLog.Close()

	gameLog, err := fs.Open("game.log")
	if err != nil {
		return time.Time{}, nil, fmt.Errorf("fs.Open(game.log): %w", err)
	}
	defer gameLog.Close()
	logTime, combatRanges, err := s.scanCombatRanges(ctx, combatLog)
	if err != nil {
		return time.Time{}, nil, fmt.Errorf("scanCombatRanges: %w", err)
	}
	logTime, gameRanges, err := s.scanGameRanges(ctx, gameLog)
	if err != nil {
		return time.Time{}, nil, fmt.Errorf("scanGameRanges: %w", err)
	}
	if len(gameRanges) != len(combatRanges) {
		mx := max(len(gameRanges), len(combatRanges))
		s.lg.For(ctx).Infow("levels count mismatch", "combat", len(combatRanges), "game", len(gameRanges))
		for i := range mx {
			if i < len(gameRanges) {
				gm := gameRanges[i]
				s.lg.For(ctx).Infow("game log range", "number", i, "start_line", gm.Start, "end_line", gm.End)
			}
			if i < len(combatRanges) {
				cm := combatRanges[i]
				s.lg.For(ctx).Infow("combat log range", "number", i, "start_line", cm.Start, "end_line", cm.End)
			}
		}
		return time.Time{}, nil, fmt.Errorf("%w: levels count mismatch: game logs: %d, combat logs: %d", ErrLogsCorrupted, len(gameRanges), len(combatRanges))
	}
	matches := make([]matchRange, 0, len(gameRanges))
	for i := range len(gameRanges) {
		m := matchRange{Game: gameRanges[i], Combat: combatRanges[i]}
		s.lg.For(ctx).Infow(
			"match_offsets_detected",
			"match_index", i,
			"game_start_line", m.Game.Start,
			"game_end_line", m.Game.End,
			"combat_start_line", m.Combat.Start,
			"combat_end_line", m.Combat.End,
		)
		matches = append(matches, m)
	}
	return logTime, matches, nil
}

func (s *Splitter) scanGameRanges(ctx context.Context, r fs.File) (time.Time, []lineRange, error) {
	var (
		ranges    []lineRange
		curStart  int
		lastLine  int
	)
	logTime, err := s.parser.WalkGameLog(ctx, r, func(line parser.LogLine[game.LogLine]) error {
		lastLine = line.Num
		if line.Data == nil {
			return nil
		}
		switch v := line.Data.(type) {
		case *game.ClientConnected:
			if curStart == 0 {
				curStart = line.Num
			} else if line.Num > curStart {
				ranges = append(ranges, lineRange{Start: curStart, End: line.Num - 1})
				curStart = line.Num
			}
		case *game.ClientConnectionClosed:
			if v.Reason == ConnectionClosedReasonClientCouldNotConnect {
				curStart = 0
				return nil
			}
			if curStart == 0 {
				curStart = line.Num
			}
			ranges = append(ranges, lineRange{Start: curStart, End: line.Num})
			curStart = 0
		}
		return nil
	})
	if err != nil {
		return time.Time{}, nil, err
	}
	if curStart != 0 && lastLine >= curStart {
		ranges = append(ranges, lineRange{Start: curStart, End: lastLine})
	}
	return logTime, ranges, nil
}

func (s *Splitter) scanCombatRanges(ctx context.Context, r fs.File) (time.Time, []lineRange, error) {
	var (
		ranges        []lineRange
		curStart      int
		lastLine      int
		curSessionID  int
		hasConnect    bool
		hasStart      bool
		hasFinished   bool
	)
	pushIfNonEmpty := func(endLine int) {
		if curStart == 0 || endLine < curStart {
			return
		}
		if hasConnect || hasStart || hasFinished {
			ranges = append(ranges, lineRange{Start: curStart, End: endLine})
		}
		curStart = 0
		curSessionID = 0
		hasConnect = false
		hasStart = false
		hasFinished = false
	}
	logTime, err := s.parser.WalkCombatLog(ctx, r, func(line parser.LogLine[combat.LogLine]) error {
		lastLine = line.Num
		if line.Data == nil {
			return nil
		}
		if curStart == 0 {
			curStart = line.Num
		}
		switch v := line.Data.(type) {
		case *combat.ConnectToGameSession:
			if hasConnect && curSessionID != v.SessionID {
				pushIfNonEmpty(line.Num - 1)
				curStart = line.Num
			}
			hasConnect = true
			curSessionID = v.SessionID
		case *combat.Start:
			if hasStart {
				pushIfNonEmpty(line.Num - 1)
				curStart = line.Num
			}
			hasStart = true
		case *combat.Finished:
			hasFinished = true
		}
		return nil
	})
	if err != nil {
		return time.Time{}, nil, err
	}
	pushIfNonEmpty(lastLine)
	return logTime, ranges, nil
}

func levelIndexForLine(ranges []lineRange, idx int, lineNum int) int {
	for idx < len(ranges) && lineNum > ranges[idx].End {
		idx++
	}
	return idx
}

func (s *Splitter) parseLevelsByRanges(
	ctx context.Context,
	fsys fs.FS,
	logTime time.Time,
	matches []matchRange,
) ([]*GameLogLevel, []*CombatLogLevel, error) {
	ctx, span := s.tr.Start(ctx, "parseLevelsByRanges")
	defer span.End()

	gameLevels := make([]*GameLogLevel, len(matches))
	for i := range gameLevels {
		gameLevels[i] = new(GameLogLevel)
	}
	combatLevels := make([]*CombatLogLevel, len(matches))
	for i := range combatLevels {
		combatLevels[i] = &CombatLogLevel{logTime: logTime}
	}
	gameRanges := make([]lineRange, 0, len(matches))
	combatRanges := make([]lineRange, 0, len(matches))
	for _, m := range matches {
		gameRanges = append(gameRanges, m.Game)
		combatRanges = append(combatRanges, m.Combat)
	}

	gameLog, err := fsys.Open("game.log")
	if err != nil {
		return nil, nil, fmt.Errorf("fs.Open(game.log): %w", err)
	}
	defer gameLog.Close()

	gameIdx := 0
	_, err = s.parser.WalkGameLogFull(ctx, gameLog, func(line parser.LogLine[game.LogLine]) error {
		gameIdx = levelIndexForLine(gameRanges, gameIdx, line.Num)
		if gameIdx >= len(gameRanges) {
			return nil
		}
		rg := gameRanges[gameIdx]
		if line.Num < rg.Start || line.Num > rg.End || line.Data == nil {
			return nil
		}
		lvl := gameLevels[gameIdx]
		switch v := line.Data.(type) {
		case *game.ClientConnected:
			if lvl.StartGameplay == nil {
				lvl.StartGameplay = v
			}
		case *game.ClientAddPlayer:
			lvl.AddPlayer = append(lvl.AddPlayer, v)
		case *game.ClientConnectionClosed:
			lvl.FinishGameplay = v
		case *game.ClientPlayerLeave:
			lvl.LeavePlayer = append(lvl.LeavePlayer, v)
		}
		return nil
	})
	if err != nil {
		return nil, nil, fmt.Errorf("WalkGameLog: %w", err)
	}

	combatLog, err := fsys.Open("combat.log")
	if err != nil {
		return nil, nil, fmt.Errorf("fs.Open(combat.log): %w", err)
	}
	defer combatLog.Close()
	combatIdx := 0
	_, err = s.parser.WalkCombatLogFull(ctx, combatLog, func(line parser.LogLine[combat.LogLine]) error {
		combatIdx = levelIndexForLine(combatRanges, combatIdx, line.Num)
		if combatIdx >= len(combatRanges) {
			return nil
		}
		rg := combatRanges[combatIdx]
		if line.Num < rg.Start || line.Num > rg.End || line.Data == nil {
			return nil
		}
		lvl := combatLevels[combatIdx]
		switch v := line.Data.(type) {
		case *combat.ConnectToGameSession:
			lvl.Connect = *v
		case *combat.Start:
			lvl.Start = *v
		case *combat.Damage:
			lvl.Damage = append(lvl.Damage, v)
		case *combat.Heal:
			lvl.Heal = append(lvl.Heal, v)
		case *combat.Kill:
			lvl.Kill = append(lvl.Kill, v)
		case *combat.Spawn:
			lvl.Spawn = append(lvl.Spawn, v)
		case *combat.Reward:
			lvl.Reward = append(lvl.Reward, v)
		case *combat.Spell:
			lvl.Spell = append(lvl.Spell, v)
		case *combat.Finished:
			lvl.Finished = *v
		}
		return nil
	})
	if err != nil {
		return nil, nil, fmt.Errorf("WalkCombatLog: %w", err)
	}
	for i := range len(matches) {
		gr := matches[i].Game
		cr := matches[i].Combat
		gl := gameLevels[i]
		cl := combatLevels[i]
		addPlayers := 0
		hasGameStart := false
		hasGameFinish := false
		if gl != nil {
			addPlayers = len(gl.AddPlayer)
			hasGameStart = gl.StartGameplay != nil
			hasGameFinish = gl.FinishGameplay != nil
		}
		damageN, healN, killN, spawnN, rewardN, spellN := 0, 0, 0, 0, 0, 0
		hasCombatConnect := false
		hasCombatStart := false
		hasCombatFinish := false
		if cl != nil {
			damageN = len(cl.Damage)
			healN = len(cl.Heal)
			killN = len(cl.Kill)
			spawnN = len(cl.Spawn)
			rewardN = len(cl.Reward)
			spellN = len(cl.Spell)
			hasCombatConnect = !cl.Connect.IsEmpty()
			hasCombatStart = !cl.Start.IsEmpty()
			hasCombatFinish = !cl.Finished.IsEmpty()
		}
		s.lg.For(ctx).Infow(
			"match_offsets_parse_summary",
			"match_index", i,
			"game_start_line", gr.Start,
			"game_end_line", gr.End,
			"combat_start_line", cr.Start,
			"combat_end_line", cr.End,
			"game_has_start", hasGameStart,
			"game_has_finish", hasGameFinish,
			"game_add_players", addPlayers,
			"combat_has_connect", hasCombatConnect,
			"combat_has_start", hasCombatStart,
			"combat_has_finish", hasCombatFinish,
			"combat_damage", damageN,
			"combat_heal", healN,
			"combat_kill", killN,
			"combat_spawn", spawnN,
			"combat_reward", rewardN,
			"combat_spell", spellN,
		)
	}
	return gameLevels, combatLevels, nil
}

func (s *Splitter) makeLevel(ctx context.Context, logTime time.Time, gameLevel *GameLogLevel, combatLevel *CombatLogLevel) (*Level, error) {
	ctx, span := s.tr.Start(ctx, "makeLevel")
	defer span.End()

	lvl := &Level{
		GameLog:   gameLevel,
		CombatLog: combatLevel,
	}

	// Start: earliest known non-zero marker among game/combat start-like events.
	startCandidates := []time.Time{}
	if gameLevel != nil && gameLevel.StartGameplay != nil {
		if t := gameLevel.StartGameplay.GetTime(logTime); !t.IsZero() {
			startCandidates = append(startCandidates, t)
		}
	}
	if combatLevel != nil {
		if t := combatLevel.Start.GetTime(logTime); !t.IsZero() {
			startCandidates = append(startCandidates, t)
		}
		if t := combatLevel.Connect.GetTime(logTime); !t.IsZero() {
			startCandidates = append(startCandidates, t)
		}
		if t := earliestCombatEventTime(combatLevel, logTime); !t.IsZero() {
			startCandidates = append(startCandidates, t)
		}
	}
	if len(startCandidates) > 0 {
		lvl.StartLevelTime = startCandidates[0]
		for i := 1; i < len(startCandidates); i++ {
			if startCandidates[i].Before(lvl.StartLevelTime) {
				lvl.StartLevelTime = startCandidates[i]
			}
		}
	}

	// End: latest known non-zero marker among finish and real combat activity.
	endCandidates := []time.Time{}
	if gameLevel != nil && gameLevel.FinishGameplay != nil {
		if t := gameLevel.FinishGameplay.GetTime(logTime); !t.IsZero() {
			endCandidates = append(endCandidates, t)
		}
	}
	if combatLevel != nil {
		if t := combatLevel.Finished.GetTime(logTime); !t.IsZero() {
			endCandidates = append(endCandidates, t)
		}
		if t := latestCombatEventTime(combatLevel, logTime); !t.IsZero() {
			endCandidates = append(endCandidates, t)
		}
	}
	if len(endCandidates) > 0 {
		lvl.EndLevelTime = endCandidates[0]
		for i := 1; i < len(endCandidates); i++ {
			if endCandidates[i].After(lvl.EndLevelTime) {
				lvl.EndLevelTime = endCandidates[i]
			}
		}
	}
	if !lvl.StartLevelTime.IsZero() && !lvl.EndLevelTime.IsZero() && lvl.EndLevelTime.Before(lvl.StartLevelTime) {
		lvl.EndLevelTime = lvl.StartLevelTime
	}

	playerTeamsMap := make(map[int]map[int]Player)
	for _, logPlayer := range gameLevel.AddPlayer {
		player := Player{
			PlayerID:        logPlayer.PlayerID,
			SessionPlayerID: logPlayer.InGamePlayerID,
			Name:            logPlayer.Name,
			CorpTag:         logPlayer.ClanTag,
			TeamID:          logPlayer.TeamID,
		}
		teamMap := playerTeamsMap[logPlayer.TeamID]
		if teamMap == nil {
			teamMap = make(map[int]Player)
			playerTeamsMap[logPlayer.TeamID] = teamMap
		}

		teamMap[player.SessionPlayerID] = player
	}

	teams := maps.Keys(playerTeamsMap)
	slices.Sort(teams)
	if os.Getenv("SHOW_WATCHERS") == "" {
		if teams[0] == 0 {
			teams = teams[1:]
		}
	}
	s.lg.For(ctx).Debugw("level", "level", lvl.CombatLog.Start, "team_ids", teams)

	lvl.Teams = make(map[int][]Player, len(playerTeamsMap))
	for teamID, team := range playerTeamsMap {
		players := maps.Values(team)
		playerNames := make([]string, 0, len(players))
		for _, p := range players {
			playerNames = append(playerNames, p.String())
		}
		slices.Sort(playerNames)
		if os.Getenv("SHOW_WATCHERS") != "" && teamID == 0 {
			s.lg.For(ctx).Debugw("team players", "team_id", teamID, "players", playerNames)
		}
		lvl.Teams[teamID] = players
	}

	if len(lvl.Teams) == 0 {
		return nil, fmt.Errorf("no teams found")
	}

	return lvl, nil
}

func earliestCombatEventTime(level *CombatLogLevel, logTime time.Time) time.Time {
	if level == nil {
		return time.Time{}
	}
	best := time.Time{}
	setMin := func(t time.Time) {
		if t.IsZero() {
			return
		}
		if best.IsZero() || t.Before(best) {
			best = t
		}
	}
	for _, x := range level.Damage {
		if x != nil {
			setMin(x.GetTime(logTime))
		}
	}
	for _, x := range level.Heal {
		if x != nil {
			setMin(x.GetTime(logTime))
		}
	}
	for _, x := range level.Kill {
		if x != nil {
			setMin(x.GetTime(logTime))
		}
	}
	for _, x := range level.Spawn {
		if x != nil {
			setMin(x.GetTime(logTime))
		}
	}
	for _, x := range level.Reward {
		if x != nil {
			setMin(x.GetTime(logTime))
		}
	}
	for _, x := range level.Spell {
		if x != nil {
			setMin(x.GetTime(logTime))
		}
	}
	return best
}

func latestCombatEventTime(level *CombatLogLevel, logTime time.Time) time.Time {
	if level == nil {
		return time.Time{}
	}
	best := time.Time{}
	setMax := func(t time.Time) {
		if t.IsZero() {
			return
		}
		if best.IsZero() || t.After(best) {
			best = t
		}
	}
	for _, x := range level.Damage {
		if x != nil {
			setMax(x.GetTime(logTime))
		}
	}
	for _, x := range level.Heal {
		if x != nil {
			setMax(x.GetTime(logTime))
		}
	}
	for _, x := range level.Kill {
		if x != nil {
			setMax(x.GetTime(logTime))
		}
	}
	for _, x := range level.Spawn {
		if x != nil {
			setMax(x.GetTime(logTime))
		}
	}
	for _, x := range level.Reward {
		if x != nil {
			setMax(x.GetTime(logTime))
		}
	}
	for _, x := range level.Spell {
		if x != nil {
			setMax(x.GetTime(logTime))
		}
	}
	return best
}

type Player struct {
	PlayerID        int
	SessionPlayerID int
	Name            string
	CorpTag         string
	TeamID          int
}

func (p Player) String() string {
	return fmt.Sprintf("%s [%s] (id: %d)", p.Name, p.CorpTag, p.PlayerID)
}

type Level struct {
	StartLevelTime time.Time
	EndLevelTime   time.Time
	Teams          map[int][]Player

	GameLog   *GameLogLevel
	CombatLog *CombatLogLevel
}

// TODO make timeline func for level

type GameLogLevel struct {
	Lines []game.LogLine

	StartGameplay  *game.ClientConnected
	AddPlayer      []*game.ClientAddPlayer
	LeavePlayer    []*game.ClientPlayerLeave
	FinishGameplay *game.ClientConnectionClosed
}

func (g *GameLogLevel) IsEmpty() bool {
	return g == nil || (g.StartGameplay == nil && g.FinishGameplay == nil && len(g.AddPlayer) == 0 && len(g.LeavePlayer) == 0)
}

type CombatLogLevel struct {
	logTime time.Time
	// LogLines — распарсенные записи combat.log в порядке файла (только успешно распознанные строки).
	LogLines []combat.LogLine

	Connect  combat.ConnectToGameSession
	Start    combat.Start
	Damage   []*combat.Damage
	Heal     []*combat.Heal
	Kill     []*combat.Kill
	Spawn    []*combat.Spawn
	Reward   []*combat.Reward
	Spell    []*combat.Spell
	Finished combat.Finished
}

func (g *CombatLogLevel) String() string {
	if len(g.LogLines) == 0 {
		return "empty game level"
	}

	var sb strings.Builder

	sb.WriteString("game level: ")
	fmt.Fprintf(&sb, "time: %s ", g.LogLines[0].GetTime(g.logTime))
	fmt.Fprintf(&sb, "lines: %d ", len(g.LogLines))
	fmt.Fprintf(&sb, "connect(session_id): %d ", g.Connect.SessionID)
	fmt.Fprintf(&sb, "game_mode: %s map_name: %s ", g.Start.GameMode, g.Start.MapName)
	fmt.Fprintf(&sb, "finished: %s reason: %s game_time: %f", g.Finished.FinishReason, g.Finished.WinReason, g.Finished.GameTime)

	return sb.String()
}

func (g *CombatLogLevel) IsEmpty() bool {
	return g == nil || (g.Connect.IsEmpty() && g.Start.IsEmpty() && g.Finished.IsEmpty())
}

// LogAnchorTime is the calendar date from the combat log header; used with line clocks in GetTime.
func (g *CombatLogLevel) LogAnchorTime() time.Time {
	if g == nil {
		return time.Time{}
	}
	return g.logTime
}

const (
	ConnectionClosedReasonGameFinished          = "DR_CLIENT_GAME_FINISHED"
	ConnectionClosedReasonClientCouldNotConnect = "DR_CLIENT_COULD_NOT_CONNECT"
	ConnectionClosedReasonQuit                  = "DR_CLIENT_QUIT"
	ConnectionClosedReasonServerTransfer        = "DR_CLIENT_SERVER_TRANSFER"
	ConnectionClosedReasonDockSpaceStation      = "DR_CLIENT_DOCK_SPACE_STATION"
	ConnectionClosedReasonReturnSpaceStation    = "DR_CLIENT_RETURN_SPACE_STATION"
)

type Result[T any] struct {
	Err    error
	Result T
}

func (s *Splitter) GetGameLogLevels(ctx context.Context, lines []parser.LogLine[game.LogLine]) iter.Seq[*Result[*GameLogLevel]] {
	ctx, span := s.tr.Start(ctx, "GetGameLogLevels")
	defer span.End()

	return func(yield func(*Result[*GameLogLevel]) bool) {
		var errs []error
		currLevel := new(GameLogLevel)

		pushLevel := func() bool {
			next := yield(&Result[*GameLogLevel]{Err: errors.Join(errs...), Result: currLevel})
			currLevel = new(GameLogLevel)
			errs = nil
			return next
		}
		for _, line := range lines {
			var cutErr *common.LineIsNotFinishedError
			if line.Err != nil && !errors.As(line.Err, &cutErr) {
				errs = append(errs, line.Err)
			}
			if line.Data == nil {
				continue
			}

			currLevel.Lines = append(currLevel.Lines, line.Data)
			switch line := line.Data.(type) {
			case *game.ClientConnected:
				if currLevel.StartGameplay != nil {
					s.lg.For(ctx).Warnw("start gameplay twice", "prev", currLevel.StartGameplay, "next", line,
						"start", currLevel.StartGameplay, "end", currLevel.FinishGameplay)
					if !pushLevel() {
						return
					}
				}
				currLevel.StartGameplay = line
			case *game.ClientAddPlayer:
				currLevel.AddPlayer = append(currLevel.AddPlayer, line)
			case *game.ClientConnectionClosed:
				currLevel.FinishGameplay = line
				if line.Reason == ConnectionClosedReasonClientCouldNotConnect {
					s.lg.For(ctx).Infow("detected could not connect log", "line", line)
					currLevel = new(GameLogLevel)
					continue
				}
				if !pushLevel() {
					return
				}
			case *game.ClientPlayerLeave:
				currLevel.LeavePlayer = append(currLevel.LeavePlayer, line)
			}
		}

		if !currLevel.IsEmpty() {
			pushLevel()
		}
	}
}

func (s *Splitter) GetCombatLogLevels(ctx context.Context, logTime time.Time, lines []parser.LogLine[combat.LogLine]) (res []*CombatLogLevel, errs []error) {
	ctx, span := s.tr.Start(ctx, "GetCombatLogLevels")
	defer span.End()

	newLevel := func() *CombatLogLevel {
		l := new(CombatLogLevel)
		l.logTime = logTime
		return l
	}
	currLevel := newLevel()
	for _, line := range lines {
		var cutErr *common.LineIsNotFinishedError
		if line.Err != nil && !errors.As(line.Err, &cutErr) {
			errs = append(errs, line.Err)
		}
		if line.Data == nil {
			continue
		}
		currLevel.LogLines = append(currLevel.LogLines, line.Data)
		switch line := line.Data.(type) {
		case *combat.ConnectToGameSession:
			if !currLevel.Connect.IsEmpty() && currLevel.Connect.SessionID != line.SessionID {
				res = append(res, currLevel)
				currLevel = newLevel()
			}
			currLevel.Connect = *line
		case *combat.Start:
			if !currLevel.Start.IsEmpty() {
				res = append(res, currLevel)
				currLevel = newLevel()
			}
			currLevel.Start = *line
		case *combat.Damage:
			currLevel.Damage = append(currLevel.Damage, line)
		case *combat.Heal:
			currLevel.Heal = append(currLevel.Heal, line)
		case *combat.Kill:
			currLevel.Kill = append(currLevel.Kill, line)
		case *combat.Spawn:
			currLevel.Spawn = append(currLevel.Spawn, line)
		case *combat.Reward:
			currLevel.Reward = append(currLevel.Reward, line)
		case *combat.Spell:
			currLevel.Spell = append(currLevel.Spell, line)
		case *combat.Finished:
			currLevel.Finished = *line
		}
	}

	if !currLevel.IsEmpty() {
		s.lg.For(ctx).Debugw("level", "level", currLevel.String())
		res = append(res, currLevel)
	}

	s.lg.For(ctx).Infow("got combat log levels", "count", len(res))
	return res, errs
}
