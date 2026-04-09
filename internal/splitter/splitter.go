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

func NewSplitter(lg logger.Factory, tr trace.TracerProvider, parser *parser.Parser) *Splitter {
	return &Splitter{lg: lg, tr: tr.Tracer("splitter"), parser: parser}
}

func (s *Splitter) SplitLevels(ctx context.Context, fs fs.FS) ([]*Level, error) {
	ctx, span := s.tr.Start(ctx, "SplitLevels")
	defer span.End()

	logTime, gameLog, combatLog, err := s.parseFiles(ctx, fs)
	if err != nil {
		return nil, fmt.Errorf("parseFiles: %w", err)
	}

	gameLevelsIt := s.GetGameLogLevels(ctx, gameLog)
	combatLevels, _ := s.GetCombatLogLevels(ctx, logTime, combatLog)

	gameLevels := slices.Collect(gameLevelsIt)

	if len(gameLevels) != len(combatLevels) {
		mx := max(len(gameLevels), len(combatLevels))
		s.lg.For(ctx).Infow("levels count mismatch", "combat", len(combatLevels), "game", len(gameLevels))
		for i := range mx {
			if i < len(gameLevels) {
				gm := gameLevels[i].Result
				var gs, ge time.Time
				if gm != nil {
					if gm.StartGameplay != nil {
						gs = gm.StartGameplay.GetTime(logTime)
					}
					if gm.FinishGameplay != nil {
						ge = gm.FinishGameplay.GetTime(logTime)
					}
				}
				s.lg.For(ctx).Infow("game log time", "number", i, "start", gs, "end", ge)
			}
			if i < len(combatLevels) {
				cm := combatLevels[i]
				s.lg.For(ctx).Infow("combat log time", "number", i, "start", cm.Start.Time, "end", cm.Finished.Time)
			}
		}
		return nil, fmt.Errorf("%w: levels count mismatch: game logs: %d, combat logs: %d", ErrLogsCorrupted, len(gameLevels), len(combatLevels))
	}

	levels := make([]*Level, 0, len(gameLevels))
	mx := max(len(gameLevels), len(combatLevels))
	for i := range mx {
		var gm *GameLogLevel
		var cm *CombatLogLevel
		if i < len(gameLevels) {
			gm = gameLevels[i].Result
		}
		if i < len(combatLevels) {
			cm = combatLevels[i]
		}

		lvl, err := s.makeLevel(ctx, logTime, gm, cm)
		if err != nil {
			return nil, fmt.Errorf("makeLevel: %w", err)
		}
		s.LogMatchParseStats(ctx, i, mx, logTime, lvl)
		levels = append(levels, lvl)
	}

	s.lg.For(ctx).Debugw("got games", "count", len(levels))
	return levels, nil
}

func (s *Splitter) parseFiles(ctx context.Context, fs fs.FS) (
	logTime time.Time,
	game []parser.LogLine[game.LogLine],
	combat []parser.LogLine[combat.LogLine],
	err error,
) {
	ctx, span := s.tr.Start(ctx, "parseFiles")
	defer span.End()

	combatLog, err := fs.Open("combat.log")
	if err != nil {
		return logTime, nil, nil, fmt.Errorf("fs.Open(combat.log): %w", err)
	}
	defer combatLog.Close()

	gameLog, err := fs.Open("game.log")
	if err != nil {
		return logTime, nil, nil, fmt.Errorf("fs.Open(game.log): %w", err)
	}
	defer gameLog.Close()

	_, combatLines, err := s.parser.ParseCombatLog(ctx, combatLog)
	if err != nil {
		return logTime, nil, nil, fmt.Errorf("parser.ParseCombatLog: %w", err)
	}

	logTime, gameLines, err := s.parser.ParseGameLog(ctx, gameLog)
	if err != nil {
		return logTime, nil, nil, fmt.Errorf("parser.ParseGameLog: %w", err)
	}

	return logTime, gameLines, combatLines, nil
}

func (s *Splitter) makeLevel(ctx context.Context, logTime time.Time, gameLevel *GameLogLevel, combatLevel *CombatLogLevel) (*Level, error) {
	ctx, span := s.tr.Start(ctx, "makeLevel")
	defer span.End()

	lvl := &Level{
		GameLog:   gameLevel,
		CombatLog: combatLevel,
	}

	if gameLevel != nil && gameLevel.StartGameplay != nil {
		lvl.StartLevelTime = gameLevel.StartGameplay.GetTime(logTime)
	}
	if combatLevel != nil {
		for _, t := range []time.Time{
			combatLevel.Connect.GetTime(logTime),
			combatLevel.Start.GetTime(logTime),
		} {
			if t.IsZero() {
				continue
			}
			if lvl.StartLevelTime.IsZero() || t.Before(lvl.StartLevelTime) {
				lvl.StartLevelTime = t
			}
		}
		if lvl.StartLevelTime.IsZero() {
			if et := earliestCombatEventTime(combatLevel, logTime); !et.IsZero() {
				lvl.StartLevelTime = et
			}
		}
	}
	// День из заголовка лога — лучше, чем нулевое время (Unix -62135596800000 в JSON).
	if lvl.StartLevelTime.IsZero() && !logTime.IsZero() {
		lvl.StartLevelTime = time.Date(
			logTime.Year(), logTime.Month(), logTime.Day(),
			0, 0, 0, 0, logTime.Location(),
		)
	}
	lvl.EndLevelTime = gameLevel.FinishGameplay.GetTime(logTime)
	if cmbt := combatLevel.Finished; lvl.EndLevelTime.Before(cmbt.GetTime(logTime)) {
		lvl.EndLevelTime = cmbt.GetTime(logTime)
	}
	// Незавершённый матч: нет finish в game/combat — иначе EndLevelTime=0, span=0 и
	// timeInRangeFromStart(..., lo=0, hi=0) отбрасывает весь урон/отхил/киллы.
	if lt := latestCombatEventTime(combatLevel, logTime); !lt.IsZero() && lvl.EndLevelTime.Before(lt) {
		lvl.EndLevelTime = lt
	}
	if !lvl.StartLevelTime.IsZero() && (lvl.EndLevelTime.IsZero() || lvl.EndLevelTime.Before(lvl.StartLevelTime)) {
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
