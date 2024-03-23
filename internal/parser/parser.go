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

type logLineParser[T any] func(string, time.Time) (T, bool, error)

func parseLog[T any](
	ctx context.Context,
	p *Parser,
	r io.Reader,
	parse logLineParser[T],
) ([]T, error) {
	var res []T

	rd := bufio.NewReader(r)
	logTime, err := p.getLogTime(ctx, rd)
	if err != nil {
		return nil, fmt.Errorf("getLogTime: %w", err)
	}

	for {
		rawLine, isPrefix, err := rd.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				p.lf.For(ctx).Debugw("EOF", "total_lines", len(res))
				break
			}
			return nil, fmt.Errorf("read log: %w", err)
		}
		if isPrefix {
			p.lf.For(ctx).Warnw("very long line. possible log corruption", "line", rawLine)
			return nil, fmt.Errorf("very long line detected: %s", rawLine)
		}

		line, matched, err := parse(string(rawLine), logTime)
		if !matched {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("parse: %w", err)
		}
		res = append(res, line)
	}

	return res, nil
}

func (p *Parser) ParseGameLog(ctx context.Context, r io.Reader) ([]game.LogLine, error) {
	ctx, trace := p.tr.Start(ctx, "ParseGameLog")
	defer trace.End()

	return parseLog(ctx, p, r, game.ParseLogLine)
}

func (p *Parser) ParseCombatLog(ctx context.Context, r io.Reader) ([]combat.LogLine, error) {
	ctx, trace := p.tr.Start(ctx, "ParseCombatLog")
	defer trace.End()

	return parseLog(ctx, p, r, combat.ParseLogLine)
}

var firstLogLineRe = regexp.MustCompile(`--- Date: (\d\d\d\d-\d\d-\d\d)`)

const (
	firstLineReDate = iota + 1
	firstLineReTotal
)

func (p *Parser) getLogTime(ctx context.Context, rd *bufio.Reader) (time.Time, error) {
	ctx, span := p.tr.Start(ctx, "getLogTime")
	defer span.End()

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

	p.lf.For(ctx).Debugw("got log time", "log_time", res)
	return res, nil
}
