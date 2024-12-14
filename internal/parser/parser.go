package parser

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"iter"
	"regexp"
	"time"

	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/parser/gramma"
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

type LogLine struct {
	Num  int
	Raw  string
	Data *gramma.LogLine
	Err  error
}

func (p *Parser) ParseLogFile(ctx context.Context, r io.Reader) (iter.Seq[LogLine], error) {
	ctx, trace := p.tr.Start(ctx, "ParseLogFile")
	defer trace.End()

	rd := bufio.NewReader(r)
	logTime, err := p.getLogTime(ctx, rd)
	if err != nil {
		return nil, fmt.Errorf("getLogTime: %w", err)
	}

	return func(yield func(LogLine) bool) {
		gp := gramma.NewParser()

		for counter := 0; ; counter++ {
			rawLineBytes, isPrefix, err := rd.ReadLine()
			rawLine := string(rawLineBytes)

			next := LogLine{
				Num: counter,
				Raw: rawLine,
			}
			if err != nil {
				if errors.Is(err, io.EOF) {
					p.lf.For(ctx).Debugw("EOF", "total_lines", counter)
					return
				}

				next.Err = fmt.Errorf("read log: %w", err)

				if !yield(next) {
					return
				}
			}
			if isPrefix {
				p.lf.For(ctx).Warnw("very long line. possible log corruption", "line", rawLine)
				next.Err = fmt.Errorf("very long line detected: %s", rawLine)
				if !yield(next) {
					return
				}
			}

			line, err := gp.Parse(logTime, string(rawLine))
			next.Data = line
			if err != nil {
				next.Err = fmt.Errorf("gramma.Parse: %w", err)
			}

			if !yield(next) {
				return
			}
		}
	}, nil
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
