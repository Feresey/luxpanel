package splitter

import (
	"context"
	"errors"
	"testing"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/stretchr/testify/require"
)

// Regression: byte span for a combat level must not include the first line of the next level
// (e.g. a second Gameplay finished), otherwise HydrateLevel sees multiple levels in one slice.
func TestCollectCombatLogLevels_twoFinishedExclusiveSpans(t *testing.T) {
	ctx := context.Background()
	lg := logger.NewNop()
	p := parser.NewParser(lg)
	s := NewSplitter(lg, trace.NewNoopTracerProvider(), p)

	logTime := time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC)
	f1 := combat.Finished{
		Time:         combat.Time{Time: "10:00:00.000"},
		WinnerTeamID: 1,
		WinReason:    "A",
		FinishReason: "r1",
		GameTime:     1,
	}
	f2 := combat.Finished{
		Time:         combat.Time{Time: "10:01:00.000"},
		WinnerTeamID: 2,
		WinReason:    "B",
		FinishReason: "r2",
		GameTime:     2,
	}

	lines := []parser.LogLine[combat.LogLine]{
		{ByteOffset: 0, ByteEnd: 10, Data: &f1},
		{ByteOffset: 10, ByteEnd: 25, Data: &f2},
	}

	levels, spans, parseErrs := s.collectCombatLogLevels(ctx, logTime, lines)
	require.NoError(t, errors.Join(parseErrs...))
	require.Len(t, levels, 2)
	require.Len(t, spans, 2)
	require.Equal(t, 0, spans[0].Start)
	require.Equal(t, 10, spans[0].End)
	require.Equal(t, 10, spans[1].Start)
	require.Equal(t, 25, spans[1].End)
}
