package parser_test

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"iter"
	"maps"
	"os"
	"strings"
	"testing"

	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/Feresey/luxpanel/internal/parser/game"
	"go.opentelemetry.io/otel/trace/noop"
)

//go:embed testdata
var parserFS embed.FS

func (s *Suite) TestParseGameLog() {
	r := s.Require()
	ctx := context.Background()

	gameLog, err := parserFS.Open("testdata/game.log")
	r.NoError(err)

	it, err := s.parser.ParseGameLog(ctx, gameLog)
	r.NoError(err)

	var res []game.LogLine
	for _, line := range it {
		// fmt.Fprintln(os.Stderr, line.Num, line.Raw)
		if err := line.Err; err != nil {
			if errors.Is(err, game.ErrLineIsNotFinished) {
				continue
			}
			var parseErr *game.GrammaError
			if errors.As(err, &parseErr) {
				fmt.Fprintln(os.Stderr, strings.Repeat("=", 150))
				fmt.Fprintln(os.Stderr, line.Num, line.Raw)
				fmt.Fprintln(os.Stderr, parseErr)
				fmt.Fprintln(os.Stderr, strings.Repeat("=", 150))
			}
			r.NoError(err)
		}
		if line.Data != nil {
			// fmt.Fprintln(os.Stderr, line.Num, line.Raw)
			res = append(res, line.Data)
		}
	}

	resIt := iter.Seq[string](func(yield func(typ string) bool) {
		for _, line := range res {
			if !yield(fmt.Sprintf("%T", line)) {
				return
			}
		}
	})
	counts := countByKeys(resIt)

	var sum int
	for i := range maps.Values(counts) {
		sum += i
	}

	r.Equal(434, len(res), "counts: %d: %+v", sum, counts)
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

			res, err := p.ParseGameLog(context.TODO(), gameLog)
			if err != nil {
				b.Errorf("parse: %v", err)
				b.FailNow()
			}
			if len(res) == 0 {
				b.FailNow()
			}
		}
	})
}

func (s *Suite) TestParseCombatLog() {
	r := s.Require()
	ctx := context.Background()

	gameLog, err := parserFS.Open("testdata/combat.log")
	r.NoError(err)

	it, err := s.parser.ParseCombatLog(ctx, gameLog)
	r.NoError(err)

	var res []combat.LogLine
	for _, line := range it {
		if err := line.Err; err != nil {
			if errors.Is(err, combat.ErrLineIsNotFinished) {
				continue
			}
			var parseErr *combat.GrammaError
			if errors.As(err, &parseErr) {
				fmt.Fprintln(os.Stderr, strings.Repeat("=", 150))
				fmt.Fprintln(os.Stderr, line.Num, line.Raw)
				fmt.Fprintln(os.Stderr, parseErr)
				fmt.Fprintln(os.Stderr, strings.Repeat("=", 150))
			}
			r.NoError(err)
		}
		if line.Data != nil {
			res = append(res, line.Data)
		}
	}

	resIt := iter.Seq[string](func(yield func(typ string) bool) {
		for _, line := range res {
			if !yield(fmt.Sprintf("%T", line)) {
				return
			}
		}
	})
	counts := countByKeys(resIt)

	var sum int
	for i := range maps.Values(counts) {
		sum += i
	}

	// TODO
	// r.Equal(154005, len(res), "counts: %d: %+v", sum, counts)
}

func countByKeys[K comparable](seq iter.Seq[K]) map[K]int {
	m := make(map[K]int)
	Count(m, seq)

	return m
}

func Count[Map ~map[K]int, K comparable](m Map, seq iter.Seq[K]) {
	for k := range seq {
		m[k]++
	}
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

			res, err := p.ParseCombatLog(context.TODO(), gameLog)
			if err != nil {
				b.Errorf("parse: %v", err)
				b.FailNow()
			}
			if len(res) == 0 {
				b.FailNow()
			}
		}
	})
}
