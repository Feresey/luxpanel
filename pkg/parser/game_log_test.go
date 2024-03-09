package parser

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	//go:embed testdata/game/connected.txt
	connectedRaw string
	//go:embed testdata/game/add_player.txt
	addPlayerRaw string
	//go:embed testdata/game/finished.txt
	finishedGameRaw string
	//go:embed testdata/game/leave_player.txt
	leavePlayerRaw string
)

func TestConnectedUnmarshal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name      string
		raw       string
		want      GameLogLineConnected
		wantError bool
	}{
		{
			name: "ok",
			raw:  "21:32:08.064         | client: connected to 185.253.20.243|35008, MTU 1492. setting up session...",
			want: GameLogLineConnected{
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
			var val GameLogLineConnected

			err := val.Unmarshal([]byte(tt.raw), now)
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
			var val GameLogLineConnected
			err := val.Unmarshal([]byte(line), now)
			r.NoError(err)
		}
	})
}

func TestAddPlayerUnmarshal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name      string
		raw       string
		want      GameLogLineAddPlayer
		wantError bool
	}{
		{
			name: "player on group",
			raw:  "21:32:08.505         | client: ADD_PLAYER 0 (Omega33cz [], 3904235) status 6 team 1 group 5212392",
			want: GameLogLineAddPlayer{
				LogTime:            time.Date(2023, 1, 0, 21, 32, 8, 505000000, time.Local),
				SessionPlayerID: 0,
				PlayerName:      "Omega33cz",
				PlayerCorp:      "",
				PlayerID:        3904235,
				Status:          6,
				TeamID:          1,
				GroupID:         5212392,
			},
		},
		{
			name: "player without group",
			raw:  "19:32:59.001         | client: ADD_PLAYER 0 (Py6Jl [LuX], 2914804) status 6 team 1",
			want: GameLogLineAddPlayer{
				LogTime:            time.Date(2023, 1, 0, 19, 32, 59, 1000000, time.Local),
				SessionPlayerID: 0,
				PlayerName:      "Py6Jl",
				PlayerCorp:      "LuX",
				PlayerID:        2914804,
				Status:          6,
				TeamID:          1,
				GroupID:         0,
			},
		},
		{
			name: "bot",
			raw:  "18:10:24.379         | client: ADD_PLAYER 56 (maksprost, 0) status 4 team 1",
			want: GameLogLineAddPlayer{
				LogTime:            time.Date(2023, 1, 0, 18, 10, 24, 379000000, time.Local),
				SessionPlayerID: 56,
				PlayerName:      "maksprost",
				PlayerCorp:      "",
				PlayerID:        0,
				Status:          4,
				TeamID:          1,
				GroupID:         0,
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
			var val GameLogLineAddPlayer

			err := val.Unmarshal([]byte(tt.raw), now)
			if tt.wantError {
				r.Error(err)
				return
			} else {
				r.NoError(err)
			}

			r.Equal(tt.want, val)
		})
	}

	println(gameLineAddPlayer)

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(addPlayerRaw, "\n")
		for _, line := range lines {
			if line == "" {
				return
			}
			var val GameLogLineAddPlayer
			err := val.Unmarshal([]byte(line), now)
			r.NoError(err)
		}
	})
}

func TestFinishedUnmarshal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name      string
		raw       string
		want      GameLogLineFinished
		wantError bool
	}{
		{
			name: "ok",
			raw:  "21:37:38.024         | client: connection closed. DR_CLIENT_GAME_FINISHED",
			want: GameLogLineFinished{
				LogTime: time.Date(2023, 1, 0, 21, 37, 38, 24000000, time.Local),
			},
		},
		{
			name:      "cutted",
			raw:       "21:37:38.024         | client: connection closed. DR_CLIENT_GAME_FINISHE",
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
			var val GameLogLineFinished

			err := val.Unmarshal([]byte(tt.raw), now)
			if tt.wantError {
				r.Error(err)
				return
			} else {
				r.NoError(err)
			}

			r.Equal(tt.want, val)
		})
	}

	println(gameLineFinished)

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(finishedGameRaw, "\n")
		for _, line := range lines {
			if line == "" {
				return
			}
			var val GameLogLineFinished
			err := val.Unmarshal([]byte(line), now)
			r.NoError(err)
		}
	})
}

func TestLeaveUnmarshal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name      string
		raw       string
		want      GameLogLinePlayerLeave
		wantError bool
	}{
		{
			name: "ok",
			raw:  "21:37:37.915         | client: player 12 leave game",
			want: GameLogLinePlayerLeave{
				LogTime: time.Date(2023, 1, 0, 21, 37, 37, 915000000, time.Local),
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
			var val GameLogLinePlayerLeave

			err := val.Unmarshal([]byte(tt.raw), now)
			if tt.wantError {
				r.Error(err)
				return
			} else {
				r.NoError(err)
			}

			r.Equal(tt.want, val)
		})
	}

	println(gameLinePlayerLeave)

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(leavePlayerRaw, "\n")
		for _, line := range lines {
			if line == "" {
				return
			}
			var val GameLogLinePlayerLeave
			err := val.Unmarshal([]byte(line), now)
			r.NoError(err)
		}
	})
}

//go:embed testdata/parser/game.log
var gameLog []byte

func TestParseGameLog(t *testing.T) {
	r := require.New(t)

	now := time.Date(2023, time.November, 11, 22, 55, 47, 688000000, time.Local)

	rd := bufio.NewReader(bytes.NewReader(gameLog))
	for {
		rawLine, _, err := rd.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		r.NoError(err)

		line, err := ParseGameLogLine(rawLine, now)
		if errors.Is(err, ErrUndefinedLineType) {
			continue
		}
		r.NoError(err)
		_ = line
	}
}
