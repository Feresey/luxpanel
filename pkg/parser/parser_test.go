package parser

import (
	"embed"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/parser
var parserFS embed.FS

func TestParser(t *testing.T) {
	testParser(t)
}

func BenchmarkParser(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			testParser(b)
		}
	})
}

func testParser(tb testing.TB) {
	tb.Helper()
	r := require.New(tb)

	combatLog, err := parserFS.Open("testdata/parser/combat.log")
	r.NoError(err)

	gameLog, err := parserFS.Open("testdata/parser/game.log")
	r.NoError(err)

	now := time.Date(2023, time.November, 11, 22, 55, 47, 688000000, time.Local)

	p := New(combatLog, gameLog, now)

	var levels int

	for {
		lvl, err := p.ParseLevel()
		r.NoError(err)
		if lvl.Combat.Finished == nil || lvl.Game.Finished == nil {
			break
		}
		levels++
	}

	r.Equal(12, levels)
}
