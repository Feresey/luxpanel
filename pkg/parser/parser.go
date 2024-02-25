package parser

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"time"
)

type Parser struct {
	sessionStartTime time.Time
	combatLog        *bufio.Reader
	gameLog          *bufio.Reader
}

func New(combatLog io.Reader, gameLog io.Reader, sessionStartTime time.Time) *Parser {
	return &Parser{
		combatLog:        bufio.NewReaderSize(combatLog, 10<<20),
		gameLog:          bufio.NewReaderSize(gameLog, 1<<20),
		sessionStartTime: sessionStartTime,
	}
}

type CombatLogData struct {
	Connect  *CombatLogLineConnectToGameSession
	Start    *CombatLogLineStartGameplay
	Damage   []CombatLogLineDamage
	Heal     []CombatLogLineHeal
	Kill     []CombatLogLineKill
	Finished *CombatLogLineGameFinished
}

type GameLogData struct {
	Connected     *GameLogLineConnected
	AddedPlayers  []GameLogLineAddPlayer
	LeavedPlayers []GameLogLinePlayerLeave
	Finished      *GameLogLineFinished
}

type Level struct {
	Game   GameLogData
	Combat CombatLogData
}

// ParseLevel читает один уровень из логов
func (p *Parser) ParseLevel(
	ctx context.Context,
	stopLevelAfterTime time.Time,
) (level *Level, err error) {
	level = &Level{}

	// TODO add log
	for {
		line, err := p.loadCombatLogLine(level)
		if err != nil {
			if errors.Is(err, io.EOF) {
				// TODO add log
				break
			}
			if !errors.Is(err, ErrUndefinedLineType) {
				return nil, fmt.Errorf("parser.loadCombatLogLine: %w", err)
			}
		}
		if line.Type() == CombatLogLineTypeGameplayFinished {
			// TODO add log
			break
		}
	}

	for {
		line, err := p.loadGameLogLine(level)
		if err != nil {
			if errors.Is(err, io.EOF) {
				// TODO add log
				break
			}
			if !errors.Is(err, ErrUndefinedLineType) {
				return nil, fmt.Errorf("parser.loadGameLogLine: %w", err)
			}
		}
		if line.Type() == GameLogLineTypeGameFinished {
			// TODO add log
			break
		}
	}
	// TODO add log

	return level, nil
}

func (p *Parser) loadCombatLogLine(level *Level) (logLine CombatLogLine, err error) {
	rawLine, isPrefix, err := p.combatLog.ReadLine()
	if err != nil {
		return nil, fmt.Errorf("read log: %w", err)
	}
	if isPrefix {
		// TODO add log
		return nil, fmt.Errorf("very long line detected: %s", rawLine)
	}

	combatLogLine, err := ParseCombatLogLine(rawLine, p.sessionStartTime)
	if err != nil {
		return nil, fmt.Errorf("ParseCombatLogLine: %w", err)
	}

	switch line := combatLogLine.(type) {
	case *CombatLogLineConnectToGameSession:
		level.Combat.Connect = line
	case *CombatLogLineStartGameplay:
		level.Combat.Start = line
	case *CombatLogLineDamage:
		level.Combat.Damage = append(level.Combat.Damage, *line)
	case *CombatLogLineHeal:
		level.Combat.Heal = append(level.Combat.Heal, *line)
	case *CombatLogLineKill:
		level.Combat.Kill = append(level.Combat.Kill, *line)
	case *CombatLogLineGameFinished:
		level.Combat.Finished = line
	default:
		return nil, fmt.Errorf("%w: %T", ErrUndefinedLineType, line)
	}

	return combatLogLine, nil
}

func (p *Parser) loadGameLogLine(level *Level) (logLine GameLogLine, err error) {
	rawLine, isPrefix, err := p.gameLog.ReadLine()
	if err != nil {
		return nil, fmt.Errorf("read log: %w", err)
	}
	if isPrefix {
		// TODO add log
		return nil, fmt.Errorf("very long line detected: %s", rawLine)
	}

	gameLogLine, err := ParseGameLogLine(rawLine, p.sessionStartTime)
	if err != nil {
		return nil, fmt.Errorf("ParseGameLogLine: %w", err)
	}

	switch line := gameLogLine.(type) {
	case *GameLogLineConnected:
		level.Game.Connected = line
	case *GameLogLineAddPlayer:
		level.Game.AddedPlayers = append(level.Game.AddedPlayers, *line)
	case *GameLogLineFinished:
		level.Game.Finished = line
	case *GameLogLineNetStat: // Я хз что с этим делать
	case *GameLogLinePlayerLeave:
		level.Game.LeavedPlayers = append(level.Game.LeavedPlayers, *line)
	default:
		return line, fmt.Errorf("%w: %T", ErrUndefinedLineType, line)
	}

	return gameLogLine, nil
}
