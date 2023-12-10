package parser

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"time"
)

type Parser struct {
	nowDate   time.Time
	combatLog *bufio.Reader
	gameLog   *bufio.Reader
}

func New(combatLog io.Reader, gameLog io.Reader, nowDate time.Time) *Parser {
	return &Parser{
		combatLog: bufio.NewReaderSize(combatLog, 10<<20),
		gameLog:   bufio.NewReaderSize(gameLog, 1<<20),
		nowDate:   nowDate,
	}
}

type CombatLogData struct {
	connect *CombatLineConnectToGameSession
	start   *CombatLineStartGameplay
	damage  []CombatLineDamage
	heal    []CombatLineHeal
	kill    []CombatLineKill
}

type GameLogLevel struct {
	combat CombatLogData
}

func (p *Parser) ParseLevel() (level *GameLogLevel, err error) {
	for {
		rawLine, isPrefix, err := p.combatLog.ReadLine()
		if err != nil {
			return nil, fmt.Errorf("read combat log: %w", err)
		}
		if isPrefix {
			return nil, fmt.Errorf("very long line detected")
		}

		combatLine, err := ParseCombatLogLine(rawLine, p.nowDate)
		if err != nil {
			if errors.Is(err, ErrUndefinedLineType) {
				// TODO может лучше есть идея?
				continue
			}
			return nil, err
		}

		switch combatLine.(type) {
		case *CombatLineConnectToGameSession:
		case *CombatLineStartGameplay:
		case *CombatLineDamage:
		case *CombatLineHeal:
		case *CombatLineKill:
		}
	}
}
