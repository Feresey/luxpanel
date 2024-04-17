package parser_test

import (
	_ "embed"
	"strings"
	"testing"
	"time"

	"github.com/Feresey/luxpanel/internal/parser/game"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed testdata/connected.txt
	connectedRaw string
	//go:embed testdata/connection_closed.txt
	connectionClosedRaw string
	//go:embed testdata/add_player.txt
	addPlayerRaw string
	//go:embed testdata/finished.txt
	finishedGameRaw string
	//go:embed testdata/leave_player.txt
	leavePlayerRaw string
)

func TestConnectedUnmarshal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name      string
		raw       string
		want      game.Connected
		wantError bool
	}{
		{
			name: "ok",
			raw:  "21:32:08.064         | client: connected to 185.253.20.243|35008, MTU 1492. setting up session...",
			want: game.Connected{
				LogTime: time.Date(2023, 1, 0, 21, 32, 8, 64000000, time.Local),
			},
		},
		{
			name:      "cutted",
			raw:       "21:32:08.064         | client: connected t",
			wantError: true,
		},
		{
			name:      "empty",
			raw:       "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			var val game.Connected
			err := val.Unmarshal(tt.raw, now)
			if tt.wantError {
				r.Error(err)
				return
			} else {
				r.NoError(err)
			}
			r.Equal(tt.want, val)
		})
	}

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(connectedRaw, "\n")

		for _, line := range lines {
			if line == "" {
				return
			}
			var val game.Connected
			err := val.Unmarshal(line, now)
			r.NoError(err, line)
		}
	})
}

func New[T any](val T) *T { return &val }

func TestAddPlayerUnmarshal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name      string
		raw       string
		want      game.AddPlayer
		wantError bool
	}{
		{
			name: "player on group",
			raw:  "21:32:08.505         | client: ADD_PLAYER 0 (Omega33cz [], 3904235) status 6 team 1 group 5212392",
			want: game.AddPlayer{
				LogTime:          time.Date(2023, 1, 0, 21, 32, 8, 505000000, time.Local),
				SessionPlayerID:  0,
				PlayerName:       "Omega33cz",
				PlayerCorpTag:    "",
				PlayerID:         3904235,
				ConnectionStatus: 6,
				TeamID:           1,
				GroupID:          New(5212392),
			},
		},
		{
			name: "player without group",
			raw:  "19:32:59.001         | client: ADD_PLAYER 0 (Py6Jl [LuX], 2914804) status 6 team 1",
			want: game.AddPlayer{
				LogTime:          time.Date(2023, 1, 0, 19, 32, 59, 1000000, time.Local),
				SessionPlayerID:  0,
				PlayerName:       "Py6Jl",
				PlayerCorpTag:    "LuX",
				PlayerID:         2914804,
				ConnectionStatus: 6,
				TeamID:           1,
			},
		},
		{
			name: "bot",
			raw:  "18:10:24.379         | client: ADD_PLAYER 56 (maksprost, 0) status 4 team 1",
			want: game.AddPlayer{
				LogTime:          time.Date(2023, 1, 0, 18, 10, 24, 379000000, time.Local),
				SessionPlayerID:  56,
				PlayerName:       "maksprost",
				PlayerCorpTag:    "",
				PlayerID:         0,
				ConnectionStatus: 4,
				TeamID:           1,
			},
		},
		{
			name:      "cutted",
			raw:       "18:10:24.379         | client: ADD_PLAYER 56 (maksprost, 0) status",
			wantError: true,
		},
		{
			name:      "empty",
			raw:       "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			var val game.AddPlayer
			err := val.Unmarshal(tt.raw, now)
			if tt.wantError {
				r.Error(err)
				return
			} else {
				r.NoError(err)
			}
			r.Equal(tt.want, val)
		})
	}

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(addPlayerRaw, "\n")

		for _, line := range lines {
			if line == "" {
				return
			}
			var val game.AddPlayer
			err := val.Unmarshal(line, now)
			r.NoError(err, line)
		}
	})
}

func TestFinishedUnmarshal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name      string
		raw       string
		want      game.ConnectionClosed
		wantError bool
	}{
		{
			name: "ok",
			raw:  "21:37:38.024         | client: connection closed. DR_CLIENT_GAME_FINISHED",
			want: game.ConnectionClosed{
				LogTime:     time.Date(2023, 1, 0, 21, 37, 38, 24000000, time.Local),
				CloseReason: game.ConnectionClosedReasonGameFinished,
			},
		},
		{
			name: "connection closed",
			raw:  "21:37:38.024         | client: connection closed. DR_CLIENT_COULD_NOT_CONNECT",
			want: game.ConnectionClosed{
				LogTime:     time.Date(2023, 1, 0, 21, 37, 38, 24000000, time.Local),
				CloseReason: game.ConnectionClosedReasonClientCouldNotConnect,
			},
		},
		{
			name:      "cutted",
			raw:       "21:37:38.024         | client: connection closed",
			wantError: true,
		},
		{
			name:      "empty",
			raw:       "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			var val game.ConnectionClosed
			err := val.Unmarshal(tt.raw, now)
			if tt.wantError {
				r.Error(err)
				return
			} else {
				r.NoError(err)
			}
			r.Equal(tt.want, val)
		})
	}

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(connectionClosedRaw, "\n")

		for _, line := range lines {
			if line == "" {
				return
			}
			var val game.ConnectionClosed
			err := val.Unmarshal(line, now)
			r.NoError(err, line)
		}
	})
}

func TestLeaveUnmarshal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name      string
		raw       string
		want      game.PlayerLeave
		wantError bool
	}{
		{
			name: "ok",
			raw:  "21:37:37.915         | client: player 12 leave game",
			want: game.PlayerLeave{
				LogTime:  time.Date(2023, 1, 0, 21, 37, 37, 915000000, time.Local),
				PlayerID: 12,
			},
		},
		{
			name:      "cutted",
			raw:       "21:37:37.915         | client: player 12 leave gam",
			wantError: true,
		},
		{
			name:      "empty",
			raw:       "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			var val game.PlayerLeave
			err := val.Unmarshal(tt.raw, now)
			if tt.wantError {
				r.Error(err)
				return
			} else {
				r.NoError(err)
			}
			r.Equal(tt.want, val)
		})
	}

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(leavePlayerRaw, "\n")

		for _, line := range lines {
			if line == "" {
				return
			}
			var val game.PlayerLeave
			err := val.Unmarshal(line, now)
			r.NoError(err, line)
		}
	})
}
