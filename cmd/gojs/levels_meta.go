
package main
import (
	"encoding/json"
	"fmt"

	"github.com/Feresey/luxpanel/internal/splitter"
)

// levelMeta is returned to the browser for match dropdown labels.
type levelMeta struct {
	Index        int     `json:"index"`
	GameMode     string  `json:"game_mode"`
	MapName      string  `json:"map_name"`
	SessionID    int     `json:"session_id"`
	GameTimeSec  float32 `json:"game_time_sec"`
	StartTime    string  `json:"start_time"`
	Label        string  `json:"label"`
}

func marshalLevelsMetaJSON(levels []*splitter.Level) (string, error) {
	if len(levels) == 0 {
		return "[]", nil
	}
	out := make([]levelMeta, 0, len(levels))
	for i, lvl := range levels {
		m := levelMeta{Index: i}
		if lvl == nil {
			m.Label = fmt.Sprintf("Match %d", i+1)
			out = append(out, m)
			continue
		}
		if !lvl.StartLevelTime.IsZero() {
			m.StartTime = lvl.StartLevelTime.Format("2006-01-02 15:04")
		}
		if cl := lvl.CombatLog; cl != nil {
			m.GameMode = cl.Start.GameMode
			m.MapName = cl.Start.MapName
			m.SessionID = cl.Connect.SessionID
			m.GameTimeSec = cl.Finished.GameTime
		}
		m.Label = formatLevelLabel(m)
		out = append(out, m)
	}
	b, err := json.Marshal(out)
	if err != nil {
		return "", fmt.Errorf("marshal levels meta: %w", err)
	}
	return string(b), nil
}

func formatLevelLabel(m levelMeta) string {
	mode := m.GameMode
	if mode == "" {
		mode = "?"
	}
	mapName := m.MapName
	if mapName == "" {
		mapName = "?"
	}
	line := fmt.Sprintf("Match %d: %s — %s", m.Index+1, mode, mapName)
	if d := formatGameDuration(m.GameTimeSec); d != "" {
		line += " (" + d + ")"
	} else if m.StartTime != "" {
		line += " (" + m.StartTime + ")"
	}
	return line
}

func formatGameDuration(sec float32) string {
	if sec <= 0 {
		return ""
	}
	if sec >= 60 {
		return fmt.Sprintf("%.1f min", float64(sec)/60)
	}
	return fmt.Sprintf("%.0f s", sec)
}
