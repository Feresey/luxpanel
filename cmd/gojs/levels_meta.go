package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"slices"

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
	WatcherActive        bool     `json:"watcher_active"`
	WatcherConnectSec    float64  `json:"watcher_connect_sec"`
	WatcherDisconnectSec float64  `json:"watcher_disconnect_sec"`
	WatcherNames         []string `json:"watcher_names,omitempty"`
}

func (r *Runtime) marshalLevelsMetaJSON(ctx context.Context, levels []*splitter.Level) (string, error) {
	if len(levels) == 0 {
		r.lg.For(ctx).Debugw("marshalLevelsMetaJSON", "levels", 0, "empty", true)
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
		wa, wc, wd, wnames := watcherPresence(lvl)
		m.WatcherActive = wa
		m.WatcherConnectSec = wc
		m.WatcherDisconnectSec = wd
		m.WatcherNames = wnames
		m.Label = formatLevelLabel(m)
		out = append(out, m)
	}
	b, err := json.Marshal(out)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalLevelsMetaJSON", "err", err, "levels", len(levels))
		return "", fmt.Errorf("marshal levels meta: %w", err)
	}
	s := string(b)
	r.lg.For(ctx).Debugw("marshalLevelsMetaJSON", "levels", len(levels), "out_len", len(s))
	return s, nil
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

// watcherPresence: team 0 в ростере — модератор/наблюдатель; время по game.log ADD_PLAYER team 0 и leave / конец сессии.
func watcherPresence(level *splitter.Level) (active bool, connectSec, disconnectSec float64, names []string) {
	if level == nil {
		return false, 0, 0, nil
	}
	t0 := level.StartLevelTime
	if t0.IsZero() {
		return false, 0, 0, nil
	}
	team0, ok := level.Teams[0]
	if !ok || len(team0) == 0 {
		return false, 0, 0, nil
	}
	active = true
	nameSeen := make(map[string]struct{})
	for _, p := range team0 {
		if n := strings.TrimSpace(p.Name); n != "" {
			if _, dup := nameSeen[n]; !dup {
				nameSeen[n] = struct{}{}
				names = append(names, n)
			}
		}
	}
	slices.Sort(names)

	connectSec = -1
	disconnectSec = -1
	gl := level.GameLog
	if gl != nil {
		for _, ap := range gl.AddPlayer {
			if ap == nil || ap.TeamID != 0 {
				continue
			}
			t := ap.GetTime(t0)
			if t.Before(t0) {
				t = t0
			}
			sec := t.Sub(t0).Seconds()
			if connectSec < 0 || sec < connectSec {
				connectSec = sec
			}
		}
		sess := make(map[int]struct{})
		for _, p := range team0 {
			sess[p.SessionPlayerID] = struct{}{}
		}
		for _, lv := range gl.LeavePlayer {
			if lv == nil {
				continue
			}
			if _, ok := sess[lv.InGamePlayerID]; !ok {
				continue
			}
			t := lv.GetTime(t0)
			if t.Before(t0) {
				t = t0
			}
			sec := t.Sub(t0).Seconds()
			if disconnectSec < 0 || sec > disconnectSec {
				disconnectSec = sec
			}
		}
	}
	if connectSec < 0 {
		connectSec = 0
	}
	if disconnectSec < 0 {
		if gl != nil && gl.FinishGameplay != nil {
			t := gl.FinishGameplay.GetTime(t0)
			if !t.IsZero() && !t.Before(t0) {
				disconnectSec = t.Sub(t0).Seconds()
			}
		}
	}
	if disconnectSec < 0 {
		disconnectSec = connectSec
	}
	return active, connectSec, disconnectSec, names
}
