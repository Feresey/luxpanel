package parser

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"time"

	"go.uber.org/zap"
)

type Parser struct {
	nowDate           time.Time
	combatLog         *bufio.Reader
	nextGameCombatLog *CombatLogLineConnectToGameSession
	gameLog           *bufio.Reader
	nextGameGameLog   *GameLogLineConnected
	lg                *zap.SugaredLogger
}

func New(combatLog io.Reader, gameLog io.Reader, nowDate time.Time) *Parser {
	return &Parser{
		combatLog: bufio.NewReaderSize(combatLog, 10<<20),
		gameLog:   bufio.NewReaderSize(gameLog, 1<<20),
		nowDate:   nowDate,
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

func (p *Parser) ParseLevel() (level *Level, err error) {
	level = &Level{}
	for level.Combat.Finished == nil || level.Game.Finished == nil {
		if level.Game.Finished == nil {
			if err := p.loadCombatLogLine(level); err != nil && !errors.Is(err, ErrUndefinedLineType) {
				return nil, fmt.Errorf("parser.loadCombatLogLine: %w", err)
			}
		}

		if level.Game.Finished == nil {
			if err := p.loadGameLogLine(level); err != nil && !errors.Is(err, ErrUndefinedLineType) {
				return nil, fmt.Errorf("parser.loadGameLogLine: %w", err)
			}
		}
	}

	return level, nil
}

func (p *Parser) loadCombatLogLine(level *Level) error {
	if p.nextGameCombatLog == nil {
		level.Combat.Connect = p.nextGameCombatLog
		p.nextGameCombatLog = nil
		return nil
	}

	rawLine, isPrefix, err := p.combatLog.ReadLine()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return fmt.Errorf("read log: %w", err)
	}
	if isPrefix {
		return fmt.Errorf("very long line detected")
	}

	combatLogLine, err := ParseCombatLogLine(rawLine, p.nowDate)
	if err != nil {
		return err
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
		return fmt.Errorf("%w: %T", ErrUndefinedLineType, line)
	}

	return nil
}

func (p *Parser) loadGameLogLine(level *Level) error {
	if p.nextGameGameLog == nil {
		level.Game.Connected = p.nextGameGameLog
		p.nextGameGameLog = nil
		return nil
	}

	rawLine, isPrefix, err := p.gameLog.ReadLine()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return fmt.Errorf("read log: %w", err)
	}
	if isPrefix {
		return fmt.Errorf("very long line detected")
	}

	gameLogLine, err := ParseGameLogLine(rawLine, p.nowDate)
	if err != nil {
		return ErrUndefinedLineType
	}

	switch line := gameLogLine.(type) {
	case *GameLogLineConnected:
		level.Game.Connected = line
	case *GameLogLineAddPlayer:
		level.Game.AddedPlayers = append(level.Game.AddedPlayers, *line)
	case *GameLogLineFinished:
		level.Game.Finished = line
	case *GameLogLinePlayerLeave:
		level.Game.LeavedPlayers = append(level.Game.LeavedPlayers, *line)
	default:
		return fmt.Errorf("%w: %T", ErrUndefinedLineType, line)
	}

	return nil
}
