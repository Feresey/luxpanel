package formatter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/maps"

	"github.com/Feresey/sclogparser/internal/logger"
	"github.com/Feresey/sclogparser/internal/parser"
)

var ErrLogsCorrupted = errors.New("logs corrupted")

type Formatter struct {
	lg logger.Factory
	tr trace.Tracer
}

func NewFormatter(lg logger.Factory, tr trace.Tracer) *Formatter {
	return &Formatter{lg: lg, tr: tr}
}

func (f *Formatter) SplitLevels(ctx context.Context, gameLog []parser.GameLogLine, combatLog []parser.CombatLogLine) ([]*Level, error) {
	ctx, span := f.tr.Start(ctx, "SplitLevels")
	defer span.End()

	gameLevels := f.GetGameLogLevels(ctx, gameLog)
	combatLevels := f.GetCombatLogLevels(ctx, combatLog)
	// if len(gameLevels) != len(combatLevels) {
	// 	mx := max(len(gameLevels), len(combatLevels))
	// 	f.lg.For(ctx).Debugw("levels count mismatch", "combat", len(combatLevels), "game", len(gameLevels))
	// 	for i := range mx {
	// 		var gm *GameLogLevel
	// 		var cm *CombatLogLevel
	// 		if i < len(gameLevels) {
	// 			gm = gameLevels[i]
	// 		}
	// 		if i < len(combatLevels) {
	// 			cm = combatLevels[i]
	// 		}

	// 		var startTime, endTime time.Time
	// 		if gm != nil {
	// 			startTime = gm.StartGameplay.LogTime
	// 		}
	// 		if cm != nil && startTime.After(cm.Connect.LogTime) {
	// 			startTime = cm.Connect.LogTime
	// 		}

	// 		if gm != nil {
	// 			endTime = gm.FinishGameplay.LogTime
	// 		}
	// 		if cm != nil && endTime.Before(cm.Finished.LogTime) {
	// 			endTime = cm.Finished.LogTime
	// 		}
	// 		f.lg.For(ctx).Debugw("game", "start", startTime, "end", endTime)
	// 	}
	// 	return nil, fmt.Errorf("%w: levels count mismatch: game logs: %d, combat logs: %d", ErrLogsCorrupted, len(gameLevels), len(combatLevels))
	// }

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

		lvl, err := f.makeLevel(ctx, gm, cm)
		if err != nil {
			return nil, fmt.Errorf("makeLevel: %w", err)
		}
		levels = append(levels, lvl)
	}

	f.lg.For(ctx).Debugw("got games", "count", len(levels))
	return levels, nil
}

func (f *Formatter) makeLevel(ctx context.Context, gameLevel *GameLogLevel, combatLevel *CombatLogLevel) (*Level, error) {
	ctx, span := f.tr.Start(ctx, "makeLevel")
	defer span.End()

	lvl := &Level{
		GameLog:   gameLevel,
		CombatLog: combatLevel,
	}

	if gameLevel != nil {
		lvl.StartLevelTime = gameLevel.StartGameplay.LogTime
	}
	if combatLevel != nil && lvl.StartLevelTime.After(combatLevel.Start.LogTime) {
		lvl.StartLevelTime = combatLevel.Start.LogTime
	}
	if gameLevel != nil {
		lvl.EndLevelTime = gameLevel.FinishGameplay.LogTime
	}
	if combatLevel != nil && lvl.EndLevelTime.Before(combatLevel.Finished.LogTime) {
		lvl.EndLevelTime = combatLevel.Finished.LogTime
	}

	playerTeamsMap := make(map[int]map[string]Player)
	for _, logPlayer := range gameLevel.AddPlayer {
		player := Player{
			ID:      logPlayer.PlayerID,
			Name:    logPlayer.PlayerName,
			CorpTag: logPlayer.PlayerCorp,
			TeamID:  logPlayer.TeamID,
		}
		teamMap := playerTeamsMap[logPlayer.TeamID]
		if teamMap == nil {
			teamMap = make(map[string]Player)
			playerTeamsMap[logPlayer.TeamID] = teamMap
		}

		teamMap[player.String()] = player
	}

	f.lg.For(ctx).Debugw("detected teams", "team_ids", maps.Keys(playerTeamsMap))
	lvl.Players = make(map[int][]Player, len(playerTeamsMap))
	for teamID, team := range playerTeamsMap {
		players := maps.Values(team)
		f.lg.For(ctx).Debugw("detected team players", "team_id", teamID, "players", players)
		lvl.Players[teamID] = players
	}

	return lvl, nil
}

type Player struct {
	ID      int
	Name    string
	CorpTag string
	TeamID  int
}

func (p Player) String() string {
	return fmt.Sprintf("%s [%s] (id: %d) (in_game_id: %d)", p.Name, p.CorpTag, p.ID, p.TeamID)
}

type Level struct {
	StartLevelTime time.Time
	EndLevelTime   time.Time
	Players        map[int][]Player

	GameLog   *GameLogLevel
	CombatLog *CombatLogLevel
}

// TODO make timeline func for level

type GameLogLevel struct {
	Lines []parser.GameLogLine

	StartGameplay  *parser.GameLogLineConnected
	AddPlayer      []*parser.GameLogLineAddPlayer
	LeavePlayer    []*parser.GameLogLinePlayerLeave
	FinishGameplay *parser.GameLogLineFinished
}

type CombatLogLevel struct {
	Lines []parser.CombatLogLine

	Connect  *parser.CombatLogLineConnectToGameSession
	Start    *parser.CombatLogLineStartGameplay
	Damage   []*parser.CombatLogLineDamage
	Heal     []*parser.CombatLogLineHeal
	Kill     []*parser.CombatLogLineKill
	Finished *parser.CombatLogLineGameFinished
}

func (f *Formatter) GetGameLogLevels(ctx context.Context, lines []parser.GameLogLine) (res []*GameLogLevel) {
	ctx, span := f.tr.Start(ctx, "GetGameLogLevels")
	defer span.End()

	currLevel := new(GameLogLevel)

	for _, line := range lines {
		f.lg.For(ctx).Debugw("processing line", "line", line)

		currLevel.Lines = append(currLevel.Lines, line)
		switch line := line.(type) {
		case *parser.GameLogLineConnected:
			currLevel.StartGameplay = line
		case *parser.GameLogLineAddPlayer:
			currLevel.AddPlayer = append(currLevel.AddPlayer, line)
		case *parser.GameLogLineFinished:
			currLevel.FinishGameplay = line
			res = append(res, currLevel)
			currLevel = new(GameLogLevel)
		case *parser.GameLogLinePlayerLeave:
			currLevel.LeavePlayer = append(currLevel.LeavePlayer, line)
		case *parser.GameLogLineNetStat: // Я хз что с этим делать
		}
	}

	f.lg.For(ctx).Debugw("got levels", "count", len(res))
	return res
}
func (f *Formatter) GetCombatLogLevels(ctx context.Context, lines []parser.CombatLogLine) (res []*CombatLogLevel) {
	ctx, span := f.tr.Start(ctx, "GetCombatLogLevels")
	defer span.End()

	currLevel := new(CombatLogLevel)

	for _, line := range lines {
		currLevel.Lines = append(currLevel.Lines, line)
		switch line := line.(type) {
		case *parser.CombatLogLineConnectToGameSession:
			currLevel.Connect = line
		case *parser.CombatLogLineStartGameplay:
			currLevel.Start = line
		case *parser.CombatLogLineDamage:
			currLevel.Damage = append(currLevel.Damage, line)
		case *parser.CombatLogLineHeal:
			currLevel.Heal = append(currLevel.Heal, line)
		case *parser.CombatLogLineKill:
			currLevel.Kill = append(currLevel.Kill, line)
		case *parser.CombatLogLineGameFinished:
			currLevel.Finished = line
			res = append(res, currLevel)
			currLevel = new(CombatLogLevel)
		}
	}

	f.lg.For(ctx).Debugw("got levels", "count", len(res))
	return res
}
