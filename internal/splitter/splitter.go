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
	"github.com/Feresey/luxpanel/internal/parser/gramma"
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

	gameLog, combatLog, err := s.parseFiles(ctx, fs)
	if err != nil {
		return nil, fmt.Errorf("parseFiles: %w", err)
	}

	gameLevels, gameErrs := s.GetGameLogLevels(ctx, gameLog)
	if err := errors.Join(gameErrs...); err != nil {
		return nil, fmt.Errorf("GetGameLogLevels: %w", err)
	}
	combatLevels, combatErrs := s.GetCombatLogLevels(ctx, combatLog)
	if err := errors.Join(combatErrs...); err != nil {
		return nil, fmt.Errorf("GetCombatLogLevels: %w", err)
	}
	if len(gameLevels) != len(combatLevels) {
		mx := max(len(gameLevels), len(combatLevels))
		s.lg.For(ctx).Infow("levels count mismatch", "combat", len(combatLevels), "game", len(gameLevels))
		for i := range mx {
			if i < len(gameLevels) {
				gm := gameLevels[i]
				s.lg.For(ctx).Infow("game log time", "number", i, "start", gm.StartGameplay, "end", gm.FinishGameplay)
			}
			if i < len(combatLevels) {
				cm := combatLevels[i]
				s.lg.For(ctx).Infow("combat log time", "number", i, "start", cm.Start, "end", cm.Finished)
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
			gm = gameLevels[i]
		}
		if i < len(combatLevels) {
			cm = combatLevels[i]
		}

		lvl, err := s.makeLevel(ctx, gm, cm)
		if err != nil {
			return nil, fmt.Errorf("makeLevel: %w", err)
		}
		levels = append(levels, lvl)
	}

	s.lg.For(ctx).Debugw("got games", "count", len(levels))
	return levels, nil
}

func (s *Splitter) parseFiles(ctx context.Context, fs fs.FS) (game, gramma iter.Seq[parser.LogLine], err error) {
	ctx, span := s.tr.Start(ctx, "parseFiles")
	defer span.End()

	combatLog, err := fs.Open("gramma.log")
	if err != nil {
		return nil, nil, fmt.Errorf("fs.Open(gramma.log): %w", err)
	}
	defer combatLog.Close()

	gameLog, err := fs.Open("gramma.log")
	if err != nil {
		return nil, nil, fmt.Errorf("fs.Open(gramma.log): %w", err)
	}
	defer gameLog.Close()

	combatLogLines, err := s.parser.ParseLogFile(ctx, combatLog)
	if err != nil {
		return nil, nil, fmt.Errorf("parser.ParseLogFile(gramma): %w", err)
	}
	gameLogLines, err := s.parser.ParseLogFile(ctx, gameLog)
	if err != nil {
		return nil, nil, fmt.Errorf("parser.ParseLogFile(game): %w", err)
	}

	return gameLogLines, combatLogLines, nil
}

func (s *Splitter) makeLevel(ctx context.Context, gameLevel *GameLogLevel, combatLevel *CombatLogLevel) (*Level, error) {
	ctx, span := s.tr.Start(ctx, "makeLevel")
	defer span.End()

	lvl := &Level{
		GameLog:   gameLevel,
		CombatLog: combatLevel,
	}

	if gameLevel != nil {
		lvl.StartLevelTime = gameLevel.StartGameplay.Time
	}
	if combatLevel != nil {
		if lvl.StartLevelTime.After(GetCombatLineTime(combatLevel.Start)) {
			lvl.StartLevelTime = GetCombatLineTime(combatLevel.Start)
		}
		if lvl.StartLevelTime.After(GetCombatLineTime(combatLevel.Connect)) {
			lvl.StartLevelTime = GetCombatLineTime(combatLevel.Connect)
		}
	}
	lvl.EndLevelTime = GetGameLineTime(gameLevel.FinishGameplay)
	if cmbt := GetCombatLineTime(combatLevel.Finished); lvl.EndLevelTime.Before(cmbt) {
		lvl.EndLevelTime = cmbt
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
	Lines []gramma.GameLine

	StartGameplay  *gramma.ClientConnected
	AddPlayer      []*gramma.ClientAddPlayer
	LeavePlayer    []*gramma.ClientPlayerLeave
	FinishGameplay *gramma.ClientConnectionClosed
}

func (g *GameLogLevel) IsEmpty() bool {
	return len(g.Lines) == 0
}

type CombatLogLevel struct {
	Lines []gramma.CombatLine

	Connect  *gramma.ConnectToGameSession
	Start    *gramma.Start
	Damage   []*gramma.Damage
	Heal     []*gramma.Heal
	Kill     []*gramma.Kill
	Finished *gramma.Finished
}

func GetGameLineTime(line gramma.GameLine) time.Time {
	if line == nil {
		return time.Time{}
	}
	switch line := line.(type) {
	case gramma.ClientAddPlayer:
		return line.Time
	case gramma.ClientPlayerLeave:
		return line.Time
	case gramma.ClientConnected:
		return line.Time
	case gramma.ClientConnectionClosed:
		return line.Time
	default:
		return time.Time{}
	}
}

func GetCombatLineTime(line gramma.CombatLine) time.Time {
	if line == nil {
		return time.Time{}
	}
	switch line := line.(type) {
	case gramma.ConnectToGameSession:
		return line.Time
	case gramma.Start:
		return line.Time
	case gramma.Finished:
		return line.Time
	case gramma.Reward:
		return line.Time
	case gramma.Damage:
		return line.Time
	case gramma.Heal:
		return line.Time
	case gramma.Kill:
		return line.Time
	case gramma.Participant:
		return line.Time
	default:
		return time.Time{}
	}
}

func (g *CombatLogLevel) String() string {
	if len(g.Lines) == 0 {
		return "empty gramma level"
	}

	var sb strings.Builder

	sb.WriteString("gramma level: ")
	fmt.Fprintf(&sb, "time: %s ", GetCombatLineTime(g.Lines[0]))
	fmt.Fprintf(&sb, "lines: %d ", len(g.Lines))
	if g.Connect != nil {
		fmt.Fprintf(&sb, "connect(session_id): %d ", g.Connect.SessionID)
	}
	if g.Start != nil {
		fmt.Fprintf(&sb, "game_mode: %s map_name: %s ", g.Start.GameMode, g.Start.MapName)
	}
	if g.Finished != nil {
		fmt.Fprintf(&sb, "finished: %s reason: %s game_time: %s", g.Finished.FinishReason, g.Finished.WinReason, g.Finished.GameTime)
	}

	return sb.String()
}

func (g *CombatLogLevel) IsEmpty() bool {
	return len(g.Lines) == 0 || (g.Connect == nil && g.Start == nil && g.Finished == nil)
}

const (
	ConnectionClosedReasonGameFinished          = "DR_CLIENT_GAME_FINISHED"
	ConnectionClosedReasonClientCouldNotConnect = "DR_CLIENT_COULD_NOT_CONNECT"
	ConnectionClosedReasonQuit                  = "DR_CLIENT_QUIT"
	ConnectionClosedReasonServerTransfer        = "DR_CLIENT_SERVER_TRANSFER"
	ConnectionClosedReasonDockSpaceStation      = "DR_CLIENT_DOCK_SPACE_STATION"
	ConnectionClosedReasonReturnSpaceStation    = "DR_CLIENT_RETURN_SPACE_STATION"
)

func (s *Splitter) GetGameLogLevels(ctx context.Context, lines iter.Seq[parser.LogLine]) (res []*GameLogLevel, errors []error) {
	ctx, span := s.tr.Start(ctx, "GetGameLogLevels")
	defer span.End()

	currLevel := new(GameLogLevel)
	lines(func(line parser.LogLine) bool {
		if line.Err != nil {
			errors = append(errors, line.Err)
			return true
		}
		currLevel.Lines = append(currLevel.Lines, line.Data.Game)
		switch line := line.Data.Game.(type) {
		case *gramma.ClientConnected:
			if currLevel.StartGameplay != nil {
				s.lg.For(ctx).Warnw("start gameplay twice", "prev", currLevel.StartGameplay, "next", line,
					"start", currLevel.StartGameplay, "end", currLevel.FinishGameplay)
				res = append(res, currLevel)
				currLevel = new(GameLogLevel)
			}
			currLevel.StartGameplay = line
		case *gramma.ClientAddPlayer:
			currLevel.AddPlayer = append(currLevel.AddPlayer, line)
		case *gramma.ClientConnectionClosed:
			currLevel.FinishGameplay = line
			if line.Reason == ConnectionClosedReasonClientCouldNotConnect {
				s.lg.For(ctx).Infow("detected could not connect log", "line", line)
				currLevel = new(GameLogLevel)
				return true
			}
			res = append(res, currLevel)
			currLevel = new(GameLogLevel)
		case *gramma.ClientPlayerLeave:
			currLevel.LeavePlayer = append(currLevel.LeavePlayer, line)
		}
		return true
	})

	if !currLevel.IsEmpty() {
		res = append(res, currLevel)
	}

	s.lg.For(ctx).Infow("got game log levels", "count", len(res))
	return res, errors
}

func (s *Splitter) GetCombatLogLevels(ctx context.Context, lines iter.Seq[parser.LogLine]) (res []*CombatLogLevel, errors []error) {
	ctx, span := s.tr.Start(ctx, "GetCombatLogLevels")
	defer span.End()

	currLevel := new(CombatLogLevel)
	lines(func(line parser.LogLine) bool {
		if line.Err != nil {
			errors = append(errors, line.Err)
			return true
		}
		currLevel.Lines = append(currLevel.Lines, line.Data.Combat)
		switch line := line.Data.Combat.(type) {
		case *gramma.ConnectToGameSession:
			if currLevel.Connect != nil && currLevel.Connect.SessionID != line.SessionID {
				res = append(res, currLevel)
				currLevel = new(CombatLogLevel)
			}
			currLevel.Connect = line
		case *gramma.Start:
			if currLevel.Start != nil {
				res = append(res, currLevel)
				currLevel = new(CombatLogLevel)
			}
			currLevel.Start = line
		case *gramma.Damage:
			currLevel.Damage = append(currLevel.Damage, line)
		case *gramma.Heal:
			currLevel.Heal = append(currLevel.Heal, line)
		case *gramma.Kill:
			currLevel.Kill = append(currLevel.Kill, line)
		case *gramma.Finished:
			currLevel.Finished = line
		}
		return true
	})

	if !currLevel.IsEmpty() {
		s.lg.For(ctx).Debugw("level", "level", currLevel.String())
		res = append(res, currLevel)
	}

	s.lg.For(ctx).Infow("got gramma log levels", "count", len(res))
	return res, errors
}
