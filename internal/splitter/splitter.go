package splitter

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"slices"
	"strings"
	"time"

	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/maps"

	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/Feresey/luxpanel/internal/parser/combat"
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

	gameLog, combatLog, err := s.parseFiles(ctx, fs)
	if err != nil {
		return nil, fmt.Errorf("parseFiles: %w", err)
	}

	gameLevels := s.GetGameLogLevels(ctx, gameLog)
	combatLevels := s.GetCombatLogLevels(ctx, combatLog)
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

func (s *Splitter) parseFiles(ctx context.Context, fs fs.FS) ([]game.LogLine, []combat.LogLine, error) {
	ctx, span := s.tr.Start(ctx, "parseFiles")
	defer span.End()

	combatLog, err := fs.Open("combat.log")
	if err != nil {
		return nil, nil, fmt.Errorf("fs.Open(combat.log): %w", err)
	}
	defer combatLog.Close()

	gameLog, err := fs.Open("game.log")
	if err != nil {
		return nil, nil, fmt.Errorf("fs.Open(game.log): %w", err)
	}
	defer gameLog.Close()

	combatLogLines, err := s.parser.ParseCombatLog(ctx, combatLog)
	if err != nil {
		return nil, nil, fmt.Errorf("parser.ParseCombatLog: %w", err)
	}
	gameLogLines, err := s.parser.ParseGameLog(ctx, gameLog)
	if err != nil {
		return nil, nil, fmt.Errorf("parser.ParseGameLog: %w", err)
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
		lvl.StartLevelTime = gameLevel.StartGameplay.LogTime
	}
	if combatLevel != nil {
		if lvl.StartLevelTime.After(combatLevel.Start.GetLogTime()) {
			lvl.StartLevelTime = combatLevel.Start.GetLogTime()
		}
		if lvl.StartLevelTime.After(combatLevel.Connect.GetLogTime()) {
			lvl.StartLevelTime = combatLevel.Connect.GetLogTime()
		}
	}
	lvl.EndLevelTime = gameLevel.FinishGameplay.GetLogTime()
	if cmbt := combatLevel.Finished.GetLogTime(); lvl.EndLevelTime.Before(cmbt) {
		lvl.EndLevelTime = cmbt
	}

	playerTeamsMap := make(map[int]map[int]Player)
	for _, logPlayer := range gameLevel.AddPlayer {
		player := Player{
			PlayerID:        logPlayer.PlayerID,
			SessionPlayerID: logPlayer.SessionPlayerID,
			Name:            logPlayer.PlayerName,
			CorpTag:         logPlayer.PlayerCorpTag,
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

	StartGameplay  *game.Connected
	AddPlayer      []*game.AddPlayer
	LeavePlayer    []*game.PlayerLeave
	FinishGameplay *game.ConnectionClosed
}

func (g *GameLogLevel) IsEmpty() bool {
	return len(g.Lines) == 0
}

type CombatLogLevel struct {
	Lines []combat.LogLine

	Connect  *combat.ConnectToGameSession
	Start    *combat.StartGameplay
	Damage   []*combat.Damage
	Heal     []*combat.Heal
	Kill     []*combat.Kill
	Finished *combat.FinishedGameplay
}

func (g *CombatLogLevel) String() string {
	if len(g.Lines) == 0 {
		return "empty combat level"
	}

	var sb strings.Builder

	sb.WriteString("combat level: ")
	fmt.Fprintf(&sb, "time: %s ", g.Lines[0].GetLogTime())
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

func (s *Splitter) GetGameLogLevels(ctx context.Context, lines []game.LogLine) (res []*GameLogLevel) {
	ctx, span := s.tr.Start(ctx, "GetGameLogLevels")
	defer span.End()

	currLevel := new(GameLogLevel)

	for _, line := range lines {
		currLevel.Lines = append(currLevel.Lines, line)
		switch line := line.(type) {
		case *game.Connected:
			if currLevel.StartGameplay != nil {
				s.lg.For(ctx).Warnw("start gameplay twice", "prev", currLevel.StartGameplay, "next", line,
					"start", currLevel.StartGameplay, "end", currLevel.FinishGameplay)
				res = append(res, currLevel)
				currLevel = new(GameLogLevel)
			}
			currLevel.StartGameplay = line
		case *game.AddPlayer:
			currLevel.AddPlayer = append(currLevel.AddPlayer, line)
		case *game.ConnectionClosed:
			currLevel.FinishGameplay = line
			if line.CloseReason == game.ConnectionClosedReasonClientCouldNotConnect {
				s.lg.For(ctx).Infow("detected could not connect log", "line", line)
				currLevel = new(GameLogLevel)
				continue
			}
			res = append(res, currLevel)
			currLevel = new(GameLogLevel)
		case *game.PlayerLeave:
			currLevel.LeavePlayer = append(currLevel.LeavePlayer, line)
		}
	}

	if !currLevel.IsEmpty() {
		res = append(res, currLevel)
	}

	s.lg.For(ctx).Infow("got game log levels", "count", len(res))
	return res
}

func (s *Splitter) GetCombatLogLevels(ctx context.Context, lines []combat.LogLine) (res []*CombatLogLevel) {
	ctx, span := s.tr.Start(ctx, "GetCombatLogLevels")
	defer span.End()

	currLevel := new(CombatLogLevel)

	for _, line := range lines {
		currLevel.Lines = append(currLevel.Lines, line)
		switch line := line.(type) {
		case *combat.ConnectToGameSession:
			if currLevel.Connect != nil && currLevel.Connect.SessionID != line.SessionID {
				res = append(res, currLevel)
				currLevel = new(CombatLogLevel)
			}
			currLevel.Connect = line
		case *combat.StartGameplay:
			if currLevel.Start != nil {
				res = append(res, currLevel)
				currLevel = new(CombatLogLevel)
			}
			currLevel.Start = line
		case *combat.Damage:
			currLevel.Damage = append(currLevel.Damage, line)
		case *combat.Heal:
			currLevel.Heal = append(currLevel.Heal, line)
		case *combat.Kill:
			currLevel.Kill = append(currLevel.Kill, line)
		case *combat.FinishedGameplay:
			currLevel.Finished = line
		}
	}

	if !currLevel.IsEmpty() {
		s.lg.For(ctx).Debugw("level", "level", currLevel.String())
		res = append(res, currLevel)
	}

	s.lg.For(ctx).Infow("got combat log levels", "count", len(res))
	return res
}
