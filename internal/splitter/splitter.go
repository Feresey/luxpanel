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
	combatLevels, _ := s.GetCombatLogLevels(ctx, combatLog)

	gameLevels := slices.Collect(gameLevelsIt)

	if len(gameLevels) != len(combatLevels) {
		mx := max(len(gameLevels), len(combatLevels))
		s.lg.For(ctx).Infow("levels count mismatch", "combat", len(combatLevels), "game", len(gameLevels))
		for i := range mx {
			if i < len(gameLevels) {
				gm := gameLevels[i].Result
				s.lg.For(ctx).Infow("game log time", "number", i, "start", gm.StartGameplay.Time, "end", gm.FinishGameplay.Time)
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

	logTime, combatLogLines, err := s.parser.ParseCombatLog(ctx, combatLog)
	if err != nil {
		return logTime, nil, nil, fmt.Errorf("parser.ParseCombatLog(game): %w", err)
	}

	logTime, gameLogLines, err := s.parser.ParseGameLog(ctx, gameLog)
	if err != nil {
		return logTime, nil, nil, fmt.Errorf("parser.ParseGameLog(game): %w", err)
	}

	return logTime, gameLogLines, combatLogLines, nil
}

func (s *Splitter) makeLevel(ctx context.Context, logTime time.Time, gameLevel *GameLogLevel, combatLevel *CombatLogLevel) (*Level, error) {
	ctx, span := s.tr.Start(ctx, "makeLevel")
	defer span.End()

	lvl := &Level{
		GameLog:   gameLevel,
		CombatLog: combatLevel,
	}

	if gameLevel != nil {
		lvl.StartLevelTime = gameLevel.StartGameplay.GetTime(logTime)
	}
	if combatLevel != nil {
		if lvl.StartLevelTime.After(combatLevel.Start.GetTime(logTime)) {
			lvl.StartLevelTime = combatLevel.Start.GetTime(logTime)
		}
		if lvl.StartLevelTime.After(combatLevel.Connect.GetTime(logTime)) {
			lvl.StartLevelTime = combatLevel.Connect.GetTime(logTime)
		}
	}
	lvl.EndLevelTime = gameLevel.FinishGameplay.GetTime(logTime)
	if cmbt := combatLevel.Finished; lvl.EndLevelTime.Before(cmbt.GetTime(logTime)) {
		lvl.EndLevelTime = cmbt.GetTime(logTime)
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
	return len(g.Lines) == 0
}

type CombatLogLevel struct {
	logTime time.Time
	Lines   []combat.LogLine

	Connect  combat.ConnectToGameSession
	Start    combat.Start
	Damage   []*combat.Damage
	Heal     []*combat.Heal
	Kill     []*combat.Kill
	Finished combat.Finished
}

func (g *CombatLogLevel) String() string {
	if len(g.Lines) == 0 {
		return "empty game level"
	}

	var sb strings.Builder

	sb.WriteString("game level: ")
	fmt.Fprintf(&sb, "time: %s ", g.Lines[0].GetTime(g.logTime))
	fmt.Fprintf(&sb, "lines: %d ", len(g.Lines))
	fmt.Fprintf(&sb, "connect(session_id): %d ", g.Connect.SessionID)
	fmt.Fprintf(&sb, "game_mode: %s map_name: %s ", g.Start.GameMode, g.Start.MapName)
	fmt.Fprintf(&sb, "finished: %s reason: %s game_time: %f", g.Finished.FinishReason, g.Finished.WinReason, g.Finished.GameTime)

	return sb.String()
}

func (g *CombatLogLevel) IsEmpty() bool {
	return len(g.Lines) == 0 || (g.Connect.IsEmpty() && g.Start.IsEmpty() && g.Finished.IsEmpty())
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

func (s *Splitter) GetCombatLogLevels(ctx context.Context, lines []parser.LogLine[combat.LogLine]) (res []*CombatLogLevel, errs []error) {
	ctx, span := s.tr.Start(ctx, "GetCombatLogLevels")
	defer span.End()

	currLevel := new(CombatLogLevel)
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
		case *combat.ConnectToGameSession:
			if !currLevel.Connect.IsEmpty() && currLevel.Connect.SessionID != line.SessionID {
				res = append(res, currLevel)
				currLevel = new(CombatLogLevel)
			}
			currLevel.Connect = *line
		case *combat.Start:
			if !currLevel.Start.IsEmpty() {
				res = append(res, currLevel)
				currLevel = new(CombatLogLevel)
			}
			currLevel.Start = *line
		case *combat.Damage:
			currLevel.Damage = append(currLevel.Damage, line)
		case *combat.Heal:
			currLevel.Heal = append(currLevel.Heal, line)
		case *combat.Kill:
			currLevel.Kill = append(currLevel.Kill, line)
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
