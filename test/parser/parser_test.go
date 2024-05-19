package parser_test

import (
	"context"
	"embed"
	"testing"

	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace/noop"
)

//go:embed testdata
var parserFS embed.FS

func (s *Suite) TestParseGameLog() {
	r := s.Require()
	ctx := context.Background()

	gameLog, err := parserFS.Open("testdata/game.log")
	r.NoError(err)

	res, err := s.parser.ParseGameLog(ctx, gameLog)
	r.NoError(err)

	r.Equal(434, len(res))
}

func BenchmarkParseGameLog(b *testing.B) {
	r := require.New(b)

	p := parser.NewParser(logger.NewNop(), noop.NewTracerProvider())

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gameLog, err := parserFS.Open("testdata/game.log")
			r.NoError(err)

			p.ParseGameLog(context.TODO(), gameLog)
		}
	})
}

func (s *Suite) TestParseCombatLog() {
	r := s.Require()
	ctx := context.Background()

	gameLog, err := parserFS.Open("testdata/combat.log")
	r.NoError(err)

	res, err := s.parser.ParseCombatLog(ctx, gameLog)
	r.NoError(err)

	r.Equal(154005, len(res))
}

func BenchmarkCombatGameLog(b *testing.B) {
	r := require.New(b)

	p := parser.NewParser(logger.NewNop(), noop.NewTracerProvider())

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gameLog, err := parserFS.Open("testdata/combat.log")
			r.NoError(err)

			p.ParseCombatLog(context.TODO(), gameLog)
		}
	})
}
