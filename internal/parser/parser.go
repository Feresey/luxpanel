package parser

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
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

func matchPrefix(line string, offset int, wantPrefix string) bool {
	return len(line) >= offset+len(wantPrefix) && line[offset:offset+len(wantPrefix)] == wantPrefix
}

// NewLazyGameLogParser is the lightweight first pass for game.log (same prefixes as full parser).
func NewLazyGameLogParser() func(string) (game.LogLine, error) {
	return NewGameLogParser()
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
		case matchPrefix(line, 23, "Spawn"):
		case matchPrefix(line, 23, "Spell"):
		default:
			return nil, nil
		}
		return p.Parse(line)
	}
}

// NewLazyCombatLogParser is a lightweight parser for the first pass (match boundaries only).
func NewLazyCombatLogParser() func(string) (combat.LogLine, error) {
	p := &common.Parser[combat.Token, combat.LogLine, combat.YaccSymType, combat.YaccLexer, combat.YaccParser]{T: &combat.Tokenizer{}, L: &combat.Lexer{}, NewGramma: combat.YaccNewParser}
	return func(line string) (combat.LogLine, error) {
		switch {
		case matchPrefix(line, 23, "======= Connect to game session"):
		case matchPrefix(line, 23, "======= Start"):
		case matchPrefix(line, 23, "Gameplay"):
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

func (p *Parser) WalkCombatLog(ctx context.Context, r io.Reader, sink func(LogLine[combat.LogLine]) error) (time.Time, error) {
	return parseLogFileStream(ctx, r, p.lg, NewLazyCombatLogParser(), sink)
}

// WalkGameLogFull parses all supported game.log records.
func (p *Parser) WalkGameLogFull(ctx context.Context, r io.Reader, sink func(LogLine[game.LogLine]) error) (time.Time, error) {
	return parseLogFileStream(ctx, r, p.lg, NewGameLogParser(), sink)
}

// WalkCombatLogFull parses all supported combat.log records.
func (p *Parser) WalkCombatLogFull(ctx context.Context, r io.Reader, sink func(LogLine[combat.LogLine]) error) (time.Time, error) {
	return parseLogFileStream(ctx, r, p.lg, NewCombatLogParser(), sink)
}

// WalkGameLogLazyString runs the lightweight game.log pass over a full file string with byte offsets.
func (p *Parser) WalkGameLogLazyString(ctx context.Context, s string, sink func(LogLine[game.LogLine]) error) (time.Time, error) {
	return ParseLogString(ctx, s, p.lg, NewLazyGameLogParser(), sink)
}

// WalkCombatLogLazyString runs the lightweight combat.log pass over a full file string with byte offsets.
func (p *Parser) WalkCombatLogLazyString(ctx context.Context, s string, sink func(LogLine[combat.LogLine]) error) (time.Time, error) {
	return ParseLogString(ctx, s, p.lg, NewLazyCombatLogParser(), sink)
}

type LogLine[T any] struct {
	Num        int
	ByteOffset int // start of this line in the source string; 0 if unknown (e.g. streamed Reader)
	ByteEnd    int // byte after line terminator in the source; 0 if unknown
	Raw        string
	Data       T
	Err        error
}

func parseLogFile[T any](ctx context.Context, r io.Reader, lg logger.Factory, parseLine func(string) (T, error)) (logTime time.Time, res []LogLine[T], err error) {
	errSink := func(next LogLine[T]) error {
		res = append(res, next)
		return nil
	}
	logTime, err = parseLogFileStream(ctx, r, lg, parseLine, errSink)
	if err != nil {
		return logTime, nil, err
	}
	return logTime, res, nil
}

func parseLogFileStream[T any](
	ctx context.Context,
	r io.Reader,
	lg logger.Factory,
	parseLine func(string) (T, error),
	sink func(LogLine[T]) error,
) (logTime time.Time, err error) {
	startTime := time.Now()
	lg.For(ctx).Debugw("start parse")
	defer func() {
		lg.For(ctx).Debugw("end parse", "total_time", time.Since(startTime))
	}()

	rd := bufio.NewReaderSize(r, 1<<20)
	logTime, err = getLogTime(rd)
	if err != nil {
		return logTime, fmt.Errorf("getLogTime: %w", err)
	}

	for counter := 3; ; counter++ {
		rawLineBytes, isPrefix, err := rd.ReadLine()
		rawLine := string(rawLineBytes)

		next := LogLine[T]{
			Num: counter,
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				return logTime, nil
			}

			next.Err = fmt.Errorf("read log: %w", err)
			if sinkErr := sink(next); sinkErr != nil {
				return logTime, sinkErr
			}
			continue
		}
		if isPrefix {
			next.Err = fmt.Errorf("very long line detected at %d: %s", counter, rawLine)
			if sinkErr := sink(next); sinkErr != nil {
				return logTime, sinkErr
			}
		}

		line, perr := parseLine(rawLine)
		next.Data = line
		next.Raw = rawLine
		if perr != nil {
			next.Err = fmt.Errorf("gramma.Parse: %w", perr)
		}

		if sinkErr := sink(next); sinkErr != nil {
			return logTime, sinkErr
		}
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

// ParseLogHeader parses the standard two-line SC log preamble and returns the calendar date and
// the byte offset in s where line 3 (first body line) begins.
func ParseLogHeader(s string) (logTime time.Time, bodyOffset int, err error) {
	if len(s) == 0 {
		return time.Time{}, 0, fmt.Errorf("empty log")
	}
	i := 0
	// line 1
	j := strings.IndexByte(s[i:], '\n')
	if j < 0 {
		return time.Time{}, 0, fmt.Errorf("log: missing first newline")
	}
	line1 := s[i : i+j]
	if len(line1) > 0 && line1[len(line1)-1] == '\r' {
		line1 = line1[:len(line1)-1]
	}
	if line1 != "" {
		return time.Time{}, 0, fmt.Errorf("first line should be empty: %q", s[i:i+j])
	}
	i += j + 1
	// line 2
	j = strings.IndexByte(s[i:], '\n')
	if j < 0 {
		return time.Time{}, 0, fmt.Errorf("log: missing second newline")
	}
	line2 := s[i : i+j]
	if len(line2) > 0 && line2[len(line2)-1] == '\r' {
		line2 = line2[:len(line2)-1]
	}
	matches := firstLogLineRe.FindStringSubmatch(line2)
	if len(matches) != firstLineReTotal {
		return time.Time{}, 0, fmt.Errorf("%w: %q", ErrWrongLineFormat, line2)
	}
	logTime, err = time.Parse("2006-01-02", matches[firstLineReDate])
	if err != nil {
		return time.Time{}, 0, fmt.Errorf("parse time: %q: %w", line2, err)
	}
	bodyOffset = i + j + 1
	return logTime, bodyOffset, nil
}

// ParseLogString scans a full log string, invoking sink for each line with byte offsets into s.
func ParseLogString[T any](ctx context.Context, s string, lg logger.Factory, parseLine func(string) (T, error), sink func(LogLine[T]) error) (logTime time.Time, err error) {
	startTime := time.Now()
	var lineCount int
	lg.For(ctx).Debugw("start parse string", "source_bytes", len(s))
	defer func() {
		lg.For(ctx).Debugw("end parse string",
			"total_time", time.Since(startTime),
			"source_bytes", len(s),
			"body_lines", lineCount,
		)
	}()

	logTime, off, err := ParseLogHeader(s)
	if err != nil {
		return logTime, fmt.Errorf("ParseLogHeader: %w", err)
	}

	for counter := 3; off < len(s); counter++ {
		if err := ctx.Err(); err != nil {
			return logTime, err
		}
		lineCount++

		lineStart := off
		j := strings.IndexByte(s[off:], '\n')
		var rawLine string
		var lineEnd int
		if j < 0 {
			rest := s[off:]
			if len(rest) > 0 && rest[len(rest)-1] == '\r' {
				rawLine = rest[:len(rest)-1]
			} else {
				rawLine = rest
			}
			lineEnd = len(s)
			off = len(s)
		} else {
			end := off + j
			seg := s[off:end]
			if len(seg) > 0 && seg[len(seg)-1] == '\r' {
				rawLine = seg[:len(seg)-1]
			} else {
				rawLine = seg
			}
			lineEnd = end + 1
			off = lineEnd
		}

		next := LogLine[T]{
			Num:        counter,
			ByteOffset: lineStart,
			ByteEnd:    lineEnd,
			Raw:        rawLine,
		}

		line, perr := parseLine(rawLine)
		next.Data = line
		if perr != nil {
			next.Err = fmt.Errorf("gramma.Parse: %w", perr)
		}

		if sinkErr := sink(next); sinkErr != nil {
			return logTime, sinkErr
		}
	}

	return logTime, nil
}
