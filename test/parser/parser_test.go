package parser_test

import (
	"context"
	"embed"
)

//go:embed testdata/parser
var parserFS embed.FS

func (s *Suite) TestParseGameLog() {
	r := s.Require()
	ctx := context.Background()

	gameLog, err := parserFS.Open("testdata/parser/game.log")
	r.NoError(err)

	res, err := s.parser.ParseGameLog(ctx, gameLog)
	r.NoError(err)

	r.Equal(3728, len(res))
}

func (s *Suite) TestParseCombatLog() {
	r := s.Require()
	ctx := context.Background()

	gameLog, err := parserFS.Open("testdata/parser/combat.log")
	r.NoError(err)

	res, err := s.parser.ParseCombatLog(ctx, gameLog)
	r.NoError(err)

	r.Equal(116614, len(res))
}
