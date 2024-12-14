package parser_test

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"os"
	"slices"
	"testing"

	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/Feresey/luxpanel/internal/parser/gramma"
	"go.opentelemetry.io/otel/trace/noop"
)

//go:embed testdata
var parserFS embed.FS

func (s *Suite) TestParseGameLog() {
	r := s.Require()
	ctx := context.Background()

	gameLog, err := parserFS.Open("testdata/game.log")
	r.NoError(err)

	it, err := s.parser.ParseLogFile(ctx, gameLog)
	r.NoError(err)

	var res []gramma.LogLine
	it(func(line parser.LogLine) bool {
		fmt.Fprintln(os.Stderr, line.Num, line.Raw)
		if err := line.Err; err != nil {
			var parseErr *gramma.GrammaError
			if errors.As(err, &parseErr) {
				fmt.Fprintln(os.Stderr, parseErr, len(parseErr.Tokens))
				for _, tok := range parseErr.Tokens {
					fmt.Fprintf(os.Stderr, "%s\n", gramma.ShowTok(tok))
				}
			}
			if errors.Is(err, gramma.ErrLineIsNotFinished) {
				return true
			}
			r.NoError(err)
			return false
		}
		if line.Data != nil {
			res = append(res, *line.Data)
		}
		return true
	})

	r.Equal(154005, len(res))
}

func BenchmarkParseGameLog(b *testing.B) {
	p := parser.NewParser(logger.NewNop(), noop.NewTracerProvider())

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gameLog, err := parserFS.Open("testdata/game.log")
			if err != nil {
				b.Errorf("open: %v", err)
				b.FailNow()
			}

			it, err := p.ParseLogFile(context.TODO(), gameLog)
			if err != nil {
				b.Errorf("parse: %v", err)
				b.FailNow()
			}
			_ = slices.Collect(it)
		}
	})
}

func (s *Suite) TestParseCombatLog() {
	r := s.Require()
	ctx := context.Background()

	gameLog, err := parserFS.Open("testdata/combat.log")
	r.NoError(err)

	it, err := s.parser.ParseLogFile(ctx, gameLog)
	r.NoError(err)

	var res []gramma.LogLine
	it(func(line parser.LogLine) bool {
		if line.Err != nil {
			r.NoError(err)
			return false
		}
		res = append(res, *line.Data)
		return true
	})

	r.Equal(154005, len(res))
}

func BenchmarkCombatGameLog(b *testing.B) {
	p := parser.NewParser(logger.NewNop(), noop.NewTracerProvider())

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gameLog, err := parserFS.Open("testdata/combat.log")
			if err != nil {
				b.Errorf("open: %v", err)
				b.FailNow()
			}

			it, err := p.ParseLogFile(context.TODO(), gameLog)
			if err != nil {
				b.Errorf("parse: %v", err)
				b.FailNow()
			}
			_ = slices.Collect(it)
		}
	})
}
