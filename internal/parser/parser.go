package parser

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/Feresey/luxpanel/internal/parser/common"
	"github.com/Feresey/luxpanel/internal/parser/game"
)

var (
	ErrWrongLineFormat = errors.New("{{$packageNameCamel}}: wrong format")
)

func NewParser(lg logger.Factory) *Parser {
	return &Parser{
		lg: lg,
	}
}

type Parser struct {
	lg logger.Factory
}

var (
	shortReConnectToGameSession = regexp.MustCompile(`^Connect to game session`)
	shortReDamage               = regexp.MustCompile(`^Damage`)
)

func matchPrefix(line string, offset int, wantPrefix string) bool {
	return len(line) >= offset+len(wantPrefix) && line[offset:offset+len(wantPrefix)] == wantPrefix
}

func NewGameLogParser() func(string) (game.LogLine, error) {
	p := &common.Parser[game.Token, game.LogLine, game.YaccSymType, game.YaccLexer, game.YaccParser]{T: &game.Tokenizer{}, L: &game.Lexer{}, NewGramma: game.YaccNewParser}
	return func(line string) (game.LogLine, error) {
		switch {
		case matchPrefix(line, 23, "client: ADD_PLAYER"):
		case matchPrefix(line, 23, "client: connected to"):
		case matchPrefix(line, 23, "client: connection closed"):
		case matchPrefix(line, 23, "client: player"):
		default:
			return nil, nil
		}
		return p.Parse(line)
	}
}

func NewCombatLogParser() func(string) (combat.LogLine, error) {
	p := &common.Parser[combat.Token, combat.LogLine, combat.YaccSymType, combat.YaccLexer, combat.YaccParser]{T: &combat.Tokenizer{}, L: &combat.Lexer{}, NewGramma: combat.YaccNewParser}
	return func(line string) (combat.LogLine, error) {
		switch {
		case matchPrefix(line, 23, "======= Connect to game session"):
		case matchPrefix(line, 23, "Damage"):
		case matchPrefix(line, 23, "Gameplay"):
		case matchPrefix(line, 23, "Heal"):
		case matchPrefix(line, 23, "Killed"):
		case matchPrefix(line, 26, "Participant"):
		case matchPrefix(line, 23, "Reward"):
		case matchPrefix(line, 23, "======= Start"):
		default:
			return nil, nil
		}
		return p.Parse(line)
	}
}

func (p *Parser) ParseGameLog(ctx context.Context, r io.Reader) (time.Time, []LogLine[game.LogLine], error) {
	return parseLogFile(ctx, r, p.lg, NewGameLogParser())
}

func (p *Parser) ParseCombatLog(ctx context.Context, r io.Reader) (time.Time, []LogLine[combat.LogLine], error) {
	return parseLogFile(ctx, r, p.lg, NewCombatLogParser())
}

type LogLine[T any] struct {
	Num  int
	Raw  string
	Data T
	Err  error
}

func parseLogFile[T any](ctx context.Context, r io.Reader, lg logger.Factory, parseLine func(string) (T, error)) (logTime time.Time, res []LogLine[T], err error) {
	startTime := time.Now()
	lg.For(ctx).Debugw("start parse")
	defer func() {
		lg.For(ctx).Debugw("end parse", "total_time", time.Since(startTime))
	}()

	rd := bufio.NewReader(r)
	logTime, err = getLogTime(rd)
	if err != nil {
		return logTime, nil, fmt.Errorf("getLogTime: %w", err)
	}

	// time parse offset
	for counter := 3; ; counter++ {
		rawLineBytes, isPrefix, err := rd.ReadLine()
		rawLine := string(rawLineBytes)

		next := LogLine[T]{
			Num: counter,
			Raw: rawLine,
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				return logTime, res, nil
			}

			next.Err = fmt.Errorf("read log: %w", err)
			res = append(res, next)
			continue
		}
		if isPrefix {
			next.Err = fmt.Errorf("very long line detected at %d: %s", counter, rawLine)
			res = append(res, next)
		}

		line, err := parseLine(rawLine)
		next.Data = line
		if err != nil {
			next.Err = fmt.Errorf("gramma.Parse: %w", err)
		}

		res = append(res, next)
	}
}

var firstLogLineRe = regexp.MustCompile(`--- Date: (\d\d\d\d-\d\d-\d\d)`)

const (
	firstLineReDate = iota + 1
	firstLineReTotal
)

func getLogTime(rd *bufio.Reader) (time.Time, error) {
	rawLine, isPrefix, err := rd.ReadLine()
	if err != nil {
		return time.Time{}, fmt.Errorf("read line: %w", err)
	}
	if string(rawLine) != "" || isPrefix {
		return time.Time{}, fmt.Errorf("first line should be empty: %q", rawLine)
	}
	rawLine, _, err = rd.ReadLine()
	if err != nil {
		return time.Time{}, fmt.Errorf("read line: %w", err)
	}
	matches := firstLogLineRe.FindStringSubmatch(string(rawLine))
	if len(matches) != firstLineReTotal {
		return time.Time{}, fmt.Errorf("%w: %q", ErrWrongLineFormat, string(rawLine))
	}

	res, err := time.Parse("2006-01-02", matches[firstLineReDate])
	if err != nil {
		return time.Time{}, fmt.Errorf("parse time: %s: %w", rawLine, err)
	}

	return res, nil
}
