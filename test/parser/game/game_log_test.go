package parser_test

import (
	_ "embed"
	"encoding/json"
	"strings"
	"testing"

	"github.com/Feresey/luxpanel/internal/parser"
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

type testData[Out any] struct {
	name      string
	input     string
	want      Out
	wantError bool
}

func runTests[Out any](t *testing.T, tests []testData[Out], rawTests string, runFunc func(*testing.T, string) (Out, error)) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			r := require.New(t)

			got, err := runFunc(t, tt.input)
			if tt.wantError {
				raw, _ := json.Marshal(got)
				r.Error(err, "res: %s, err: %v", raw, err)
				return
			}
			r.NoError(err)
			r.Equal(tt.want, got)
		})
	}

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(rawTests, "\n")

		for _, line := range lines {
			if line == "" {
				return
			}
			res, err := runFunc(t, line)
			r.NoError(err)
			r.NotEmpty(res)
		}
	})
}

func TestConnected(t *testing.T) {
	tests := []testData[game.LogLine]{
		{
			name:  "ok",
			input: "21:32:08.064         | client: connected to 185.253.20.243|35008, MTU 1492. setting up session...",
			want: &game.ClientConnected{
				Time:       game.Time{Time: "21:32:08.064"},
				ServerAddr: "185.253.20.243|35008",
				MTU:        1492,
			},
		},
		{
			name:      "cutted",
			input:     "21:32:08.064         | client: connected to",
			wantError: true,
		},
		{
			name:      "empty",
			input:     "",
			wantError: false,
		},
	}

	parse := parser.NewGameLogParser()
	runTests(t, tests, connectedRaw, func(t *testing.T, raw string) (game.LogLine, error) {
		return parse(raw)
	})
}

func TestAddPlayer(t *testing.T) {
	tests := []testData[game.LogLine]{
		{
			name:  "player on group",
			input: "21:32:08.505         | client: ADD_PLAYER 0 (Omega33cz [], 3904235) status 6 team 1 group 5212392",
			want: &game.ClientAddPlayer{
				Time:           game.Time{Time: "21:32:08.505"},
				InGamePlayerID: 0,
				Name:           "Omega33cz",
				ClanTag:        "",
				PlayerID:       3904235,
				Status:         6,
				TeamID:         1,
				GroupID:        5212392,
			},
		},
		{
			name:  "player without group",
			input: "19:32:59.001         | client: ADD_PLAYER 0 (Py6Jl [LuX], 2914804) status 6 team 1",
			want: &game.ClientAddPlayer{
				Time:           game.Time{Time: "19:32:59.001"},
				InGamePlayerID: 0,
				Name:           "Py6Jl",
				ClanTag:        "LuX",
				PlayerID:       2914804,
				Status:         6,
				TeamID:         1,
			},
		},
		{
			name:  "bot",
			input: "18:10:24.379         | client: ADD_PLAYER 56 (maksprost, 0) status 4 team 1",
			want: &game.ClientAddPlayer{
				Time:           game.Time{Time: "18:10:24.379"},
				InGamePlayerID: 56,
				Name:           "maksprost",
				ClanTag:        "",
				PlayerID:       0,
				Status:         4,
				TeamID:         1,
			},
		},
		{
			name:      "cutted",
			input:     "18:10:24.379         | client: ADD_PLAYER 56 (maksprost, 0) status",
			wantError: true,
		},
		{
			name:      "empty",
			input:     "",
			wantError: false,
		},
	}

	parse := parser.NewGameLogParser()
	runTests(t, tests, addPlayerRaw, func(t *testing.T, raw string) (game.LogLine, error) {
		return parse(raw)
	})
}

func TestFinished(t *testing.T) {
	tests := []testData[game.LogLine]{
		{
			name:  "ok",
			input: "21:37:38.024         | client: connection closed. DR_CLIENT_GAME_FINISHED",
			want: &game.ClientConnectionClosed{
				Time:   game.Time{Time: "21:37:38.024"},
				Reason: "DR_CLIENT_GAME_FINISHED",
			},
		},
		{
			name:  "connection closed",
			input: "21:37:38.024         | client: connection closed. DR_CLIENT_COULD_NOT_CONNECT",
			want: &game.ClientConnectionClosed{
				Time:   game.Time{Time: "21:37:38.024"},
				Reason: "DR_CLIENT_COULD_NOT_CONNECT",
			},
		},
		{
			name:      "cutted",
			input:     "21:37:38.024         | client: connection closed",
			wantError: true,
		},
		{
			name:      "empty",
			input:     "",
			wantError: false,
		},
	}

	parse := parser.NewGameLogParser()
	runTests(t, tests, connectionClosedRaw, func(t *testing.T, raw string) (game.LogLine, error) {
		return parse(raw)
	})
}

func TestLeave(t *testing.T) {
	tests := []testData[game.LogLine]{
		{
			name:  "ok",
			input: "21:37:37.915         | client: player 12 leave game",
			want: &game.ClientPlayerLeave{
				Time:           game.Time{Time: "21:37:37.915"},
				InGamePlayerID: 12,
			},
		},
		{
			name:      "cutted",
			input:     "21:37:37.915         | client: player 12 lea",
			wantError: true,
		},
		{
			name:      "empty",
			input:     "",
			wantError: false,
		},
	}

	parse := parser.NewGameLogParser()
	runTests(t, tests, leavePlayerRaw, func(t *testing.T, raw string) (game.LogLine, error) {
		return parse(raw)
	})
}
