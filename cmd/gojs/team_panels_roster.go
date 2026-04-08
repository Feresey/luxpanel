package main

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/Feresey/luxpanel/internal/splitter"
)

// teamPanelsRoster — имена в порядке левой/правой панели графиков (Team 1 / Team 2).
type teamPanelsRoster struct {
	Left  []string `json:"left"`
	Right []string `json:"right"`
}

func playerNamesTrimmed(players []splitter.Player) []string {
	out := make([]string, 0, len(players))
	for _, p := range players {
		n := strings.TrimSpace(p.Name)
		if n != "" {
			out = append(out, n)
		}
	}
	return out
}

func (r *Runtime) marshalTeamPanelsRosterJSON(_ context.Context, level *splitter.Level) (string, error) {
	if level == nil {
		return "null", nil
	}
	_, _, pla, plb, ok := chartTeamSides(level)
	if !ok {
		return "null", nil
	}
	b, err := json.Marshal(teamPanelsRoster{
		Left:  playerNamesTrimmed(pla),
		Right: playerNamesTrimmed(plb),
	})
	if err != nil {
		return "", err
	}
	return string(b), nil
}
