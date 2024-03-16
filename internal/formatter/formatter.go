package formatter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Feresey/sclogparser/internal/logger"
	"github.com/Feresey/sclogparser/internal/parser"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/maps"
)

var ErrLogsCorrupted = errors.New("logs corrupted")

type Formatter struct {
	lg logger.Factory
	tr trace.Tracer
}

func (f *Formatter) SplitLevels(ctx context.Context, gameLog []parser.GameLogLine, combatLog []parser.CombatLogLine) ([]*Level, error) {
	ctx, span := f.tr.Start(ctx, "SplitLevels")
	defer span.End()

	gameLevels := f.getGameLogLevels(ctx, gameLog)
	combatLevels := f.getCombatLogLevels(ctx, combatLog)
	if len(gameLevels) != len(combatLevels) {
		return nil, fmt.Errorf("%w: levels count mismatch: combat: %d, combat logsw: %d", ErrLogsCorrupted, len(gameLevels), len(combatLevels))
	}

	levels := make([]*Level, 0, len(gameLevels))
	for i := range len(gameLevels) {
		lvl := &Level{
			GameLog:   gameLevels[i],
			CombatLog: combatLevels[i],
		}
		levels = append(levels, lvl)
	}

	return levels, nil
}

func (f *Formatter) makeLevel(ctx context.Context, gameLevel *GameLogLevel, combatLevel *CombatLogLevel) (*Level, error) {
	ctx, span := f.tr.Start(ctx, "makeLevel")
	defer span.End()

	lvl := &Level{
		GameLog:   gameLevel,
		CombatLog: combatLevel,
	}

	lvl.StartLevelTime = gameLevel.StartGameplay.LogTime
	if lvl.StartLevelTime.After(combatLevel.Start.LogTime) {
		lvl.StartLevelTime = combatLevel.Start.LogTime
	}
	lvl.EndLevelTime = gameLevel.FinishGameplay.LogTime
	if lvl.EndLevelTime.Before(combatLevel.Finished.LogTime) {
		lvl.EndLevelTime = combatLevel.Finished.LogTime
	}

	playersMap := make(map[string]Player)
	for _, logPlayer := range gameLevel.AddPlayer {
		player := Player{
			ID:      logPlayer.PlayerID,
			Name:    logPlayer.PlayerName,
			CorpTag: logPlayer.PlayerCorp,
		}
		playersMap[player.String()] = player
	}
	lvl.Players = maps.Values(playersMap)

	return lvl, nil
}

type Player struct {
	ID      int
	Name    string
	CorpTag string
}

func (p Player) String() string {
	return fmt.Sprintf("%d_%s_%s", p.ID, p.Name, p.CorpTag)
}

type Level struct {
	StartLevelTime time.Time
	EndLevelTime   time.Time
	Players        []Player

	GameLog   *GameLogLevel
	CombatLog *CombatLogLevel
}

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

func (f *Formatter) getGameLogLevels(ctx context.Context, lines []parser.GameLogLine) (res []*GameLogLevel) {
	ctx, span := f.tr.Start(ctx, "getGameLogLevels")
	defer span.End()

	currLevel := new(GameLogLevel)

	for _, line := range lines {
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

	res = append(res, currLevel)

	return res
}
func (f *Formatter) getCombatLogLevels(ctx context.Context, lines []parser.CombatLogLine) (res []*CombatLogLevel) {
	ctx, span := f.tr.Start(ctx, "getCombatLogLevels")
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

	res = append(res, currLevel)
	return res
}
