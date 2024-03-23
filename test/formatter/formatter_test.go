package formatter_test

import (
	"context"
	"time"

	"github.com/Feresey/luxpanel/internal/parser/game"
	"github.com/Feresey/luxpanel/internal/splitter"
)

// func (s *Suite) TestFormatter() {
// 	r := s.Require()
// 	ctx := context.Background()
// 	now := time.Date(2024, time.April, 0, 0, 0, 0, 0, time.Local)

// 	// tests := []struct {
// 	// 	gameLines   []*game.
// 	// 	combatLines []*parser.CombatLogLine

// 	// 	wantLevels []*splitter.Level
// 	// }{
// 	// 	{

// 	// 	},
// 	// }

// 	levels, err := s.splitter.SplitLevels(ctx,
// 		[]game.{
// 			&game.Connected{LogTime: now},
// 			&game.AddPlayer{LogTime: now.Add(1)},
// 			&game.PlayerLeave{LogTime: now.Add(2)},
// 			&game.ConnectionClosed{LogTime: now.Add(3)},
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

// 	wantLevels := []*splitter.Level{
// 		{
// 			StartLevelTime: now,
// 			EndLevelTime:   now.Add(5),
// 			Players: map[int][]splitter.Player{
// 				1: []splitter.Player{},
// 			},
// 		},
// 	}

// 	r.Equal(wantLevels, levels)
// }

func (s *Suite) TestGameLogLevels() {
	r := s.Require()
	ctx := context.Background()
	now := time.Date(2024, time.April, 0, 0, 0, 0, 0, time.Local)

	levels := s.splitter.GetGameLogLevels(ctx, []game.LogLine{
		&game.Connected{LogTime: now},
		&game.AddPlayer{LogTime: now.Add(1)},
		&game.PlayerLeave{LogTime: now.Add(2)},
		&game.ConnectionClosed{LogTime: now.Add(3)},
	})

	r.Len(levels, 1)
	r.Equal([]*splitter.GameLogLevel{{
		Lines: []game.LogLine{
			&game.Connected{LogTime: now},
			&game.AddPlayer{LogTime: now.Add(1)},
			&game.PlayerLeave{LogTime: now.Add(2)},
			&game.ConnectionClosed{LogTime: now.Add(3)},
		},
		StartGameplay: &game.Connected{LogTime: now},
		AddPlayer: []*game.AddPlayer{
			{LogTime: now.Add(1)},
		},
		LeavePlayer: []*game.PlayerLeave{
			{LogTime: now.Add(2)},
		},
		FinishGameplay: &game.ConnectionClosed{LogTime: now.Add(3)},
	}}, levels)
}
