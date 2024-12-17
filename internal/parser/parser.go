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
	"github.com/Feresey/luxpanel/internal/parser/game"
	"go.opentelemetry.io/otel/trace"
)

var (
	ErrUndefinedLineType = errors.New("{{$packageNameCamel}}: undefined line type")
	ErrWrongLineFormat   = errors.New("{{$packageNameCamel}}: wrong format")
)

func NewParser(lf logger.Factory, tr trace.TracerProvider) *Parser {
	return &Parser{
		lf: lf,
		tr: tr.Tracer("parser"),
	}
}

type Parser struct {
	lf logger.Factory
	tr trace.Tracer
}

func (p *Parser) ParseGameLog(ctx context.Context, r io.Reader) ([]LogLine[game.LogLine], error) {
	parser := game.NewParser()
	return parseLogFile(r, parser.Parse)
}

func (p *Parser) ParseCombatLog(ctx context.Context, r io.Reader) ([]LogLine[combat.LogLine], error) {
	parser := combat.NewParser()
	return parseLogFile(r, parser.Parse)
}

type LogLine[T any] struct {
	Num  int
	Raw  string
	Data T
	Err  error
}

func parseLogFile[T any](r io.Reader, parseLine func(time.Time, string) (T, error)) (res []LogLine[T], err error) {
	rd := bufio.NewReader(r)
	logTime, err := getLogTime(rd)
	if err != nil {
		return nil, fmt.Errorf("getLogTime: %w", err)
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
				return res, nil
			}

			next.Err = fmt.Errorf("read log: %w", err)
			res = append(res, next)
			continue
		}
		if isPrefix {
			next.Err = fmt.Errorf("very long line detected at %d: %s", counter, rawLine)
			res = append(res, next)
		}

		line, err := parseLine(logTime, string(rawLine))
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
