package formatter_test

import (
	"context"
	"time"

	"github.com/Feresey/sclogparser/internal/formatter"
	"github.com/Feresey/sclogparser/internal/parser"
)

// func (s *Suite) TestFormatter() {
// 	r := s.Require()
// 	ctx := context.Background()
// 	now := time.Date(2024, time.April, 0, 0, 0, 0, 0, time.Local)

// 	// tests := []struct {
// 	// 	gameLines   []*parser.GameLogLine
// 	// 	combatLines []*parser.CombatLogLine

// 	// 	wantLevels []*formatter.Level
// 	// }{
// 	// 	{

// 	// 	},
// 	// }

// 	levels, err := s.formatter.SplitLevels(ctx,
// 		[]parser.GameLogLine{
// 			&parser.GameLogLineConnected{LogTime: now},
// 			&parser.GameLogLineAddPlayer{LogTime: now.Add(1)},
// 			&parser.GameLogLinePlayerLeave{LogTime: now.Add(2)},
// 			&parser.GameLogLineFinished{LogTime: now.Add(3)},
// 		},
// 		[]parser.CombatLogLine{
// 			&parser.CombatLogLineConnectToGameSession{LogTime: now},
// 			&parser.CombatLogLineStartGameplay{LogTime: now.Add(1)},
// 			&parser.CombatLogLineDamage{LogTime: now.Add(2)},
// 			&parser.CombatLogLineHeal{LogTime: now.Add(3)},
// 			&parser.CombatLogLineKill{LogTime: now.Add(4)},
// 			&parser.CombatLogLineGameFinished{LogTime: now.Add(5)},
// 		},
// 	)
// 	r.NoError(err)

// 	wantLevels := []*formatter.Level{
// 		{
// 			StartLevelTime: now,
// 			EndLevelTime:   now.Add(5),
// 			Players: map[int][]formatter.Player{
// 				1: []formatter.Player{},
// 			},
// 		},
// 	}

// 	r.Equal(wantLevels, levels)
// }

func (s *Suite) TestGameLogLevels() {
	r := s.Require()
	ctx := context.Background()
	now := time.Date(2024, time.April, 0, 0, 0, 0, 0, time.Local)

	levels := s.formatter.GetGameLogLevels(ctx, []parser.GameLogLine{
		&parser.GameLogLineConnected{LogTime: now},
		&parser.GameLogLineAddPlayer{LogTime: now.Add(1)},
		&parser.GameLogLinePlayerLeave{LogTime: now.Add(2)},
		&parser.GameLogLineFinished{LogTime: now.Add(3)},
	})

	r.Len(levels, 1)
	r.Equal([]*formatter.GameLogLevel{{
		Lines: []parser.GameLogLine{
			&parser.GameLogLineConnected{LogTime: now},
			&parser.GameLogLineAddPlayer{LogTime: now.Add(1)},
			&parser.GameLogLinePlayerLeave{LogTime: now.Add(2)},
			&parser.GameLogLineFinished{LogTime: now.Add(3)},
		},
		StartGameplay: &parser.GameLogLineConnected{LogTime: now},
		AddPlayer: []*parser.GameLogLineAddPlayer{
			{LogTime: now.Add(1)},
		},
		LeavePlayer: []*parser.GameLogLinePlayerLeave{
			{LogTime: now.Add(2)},
		},
		FinishGameplay: &parser.GameLogLineFinished{LogTime: now.Add(3)},
	}}, levels)
}
