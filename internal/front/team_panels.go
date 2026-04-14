package front

import (
	"strings"

	"github.com/Feresey/luxpanel/internal/splitter"
)

// TeamPanelsRoster holds player names for left/right chart panels.
type TeamPanelsRoster struct {
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

// BuildTeamPanelsRoster returns roster names for chart panels or nil if layout unknown.
func BuildTeamPanelsRoster(level *splitter.Level) *TeamPanelsRoster {
	if level == nil {
		return nil
	}
	_, _, pla, plb, ok := ChartTeamSides(level)
	if !ok {
		return nil
	}
	return &TeamPanelsRoster{
		Left:  playerNamesTrimmed(pla),
		Right: playerNamesTrimmed(plb),
	}
}
