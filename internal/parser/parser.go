package parser

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"regexp"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/Feresey/sclogparser/internal/logger"
)

func NewParser(lf logger.Factory, tr trace.Tracer) *Parser {
	return &Parser{
		lf: lf,
		tr: tr,
	}
}

type Parser struct {
	lf logger.Factory
	tr trace.Tracer
}

func parseLog[T any](ctx context.Context, p *Parser, r io.Reader, parse func(context.Context, *bufio.Reader, time.Time) (T, error)) ([]T, error) {
	var res []T

	rd := bufio.NewReader(r)
	logTime, err := p.getLogTime(ctx, rd)
	if err != nil {
		return nil, fmt.Errorf("getLogTime: %w", err)
	}

	for {
		line, err := parse(ctx, rd, logTime)
		if err != nil {
			if !errors.Is(err, ErrUndefinedLineType) {
				if errors.Is(err, io.EOF) {
					p.lf.For(ctx).Debugw("EOF", "total_lines", len(res))
					break
				}
				return nil, fmt.Errorf("parse: %w", err)
			}
			continue
		}
		res = append(res, line)
	}

	return res, nil
}

func (p *Parser) ParseGameLog(ctx context.Context, r io.Reader) ([]GameLogLine, error) {
	ctx, trace := p.tr.Start(ctx, "ParseGameLog")
	defer trace.End()

	return parseLog(ctx, p, r, p.loadGameLogLine)
}

func (p *Parser) ParseCombatLog(ctx context.Context, r io.Reader) ([]CombatLogLine, error) {
	ctx, trace := p.tr.Start(ctx, "ParseCombatLog")
	defer trace.End()

	return parseLog(ctx, p, r, p.loadCombatLogLine)
}

func (p *Parser) loadCombatLogLine(ctx context.Context, r *bufio.Reader, startTime time.Time) (logLine CombatLogLine, err error) {
	ctx, span := p.tr.Start(ctx, "loadCombatLogLine")
	defer span.End()

	rawLine, isPrefix, err := r.ReadLine()
	if err != nil {
		return nil, fmt.Errorf("read log: %w", err)
	}
	if isPrefix {
		p.lf.For(ctx).Warnw("very long line. possible log corruption", "line", rawLine)
		return nil, fmt.Errorf("very long line detected: %s", rawLine)
	}

	line, err := ParseCombatLogLine(rawLine, startTime)
	if err != nil {
		return nil, fmt.Errorf("ParseCombatLogLine: %w", err)
	}
	return line, nil
}

func (p *Parser) loadGameLogLine(ctx context.Context, r *bufio.Reader, startTime time.Time) (logLine GameLogLine, err error) {
	ctx, span := p.tr.Start(ctx, "loadGameLogLine")
	defer span.End()

	rawLine, isPrefix, err := r.ReadLine()
	if err != nil {
		return nil, fmt.Errorf("read log: %w", err)
	}
	if isPrefix {
		p.lf.For(ctx).Warnw("very long line. possible log corruption", "line", rawLine)
		return nil, fmt.Errorf("very long line detected: %s", rawLine)
	}

	line, err := ParseGameLogLine(rawLine, startTime)
	if err != nil {
		return nil, fmt.Errorf("ParseGameLogLine: %w", err)
	}

	return line, nil
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
