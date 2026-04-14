package front

import (
	"strings"

	"slices"

	"github.com/Feresey/luxpanel/internal/splitter"
)

type LevelMeta struct {
	Index                int      `json:"index"`
	GameMode             string   `json:"game_mode"`
	MapName              string   `json:"map_name"`
	SessionID            int      `json:"session_id"`
	GameTimeSec          float32  `json:"game_time_sec"`
	StartTime            string   `json:"start_time"`
	Label                string   `json:"label"`
	WatcherActive        bool     `json:"watcher_active"`
	WatcherConnectSec    float64  `json:"watcher_connect_sec"`
	WatcherDisconnectSec float64  `json:"watcher_disconnect_sec"`
	WatcherNames         []string `json:"watcher_names,omitempty"`
}

func WatcherPresence(level *splitter.Level) (active bool, connectSec, disconnectSec float64, names []string) {
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

func BuildLevelMetaSlice(levels []*splitter.Level) []LevelMeta {
	if len(levels) == 0 {
		return nil
	}
	out := make([]LevelMeta, 0, len(levels))
	for i, lvl := range levels {
		m := LevelMeta{Index: i}
		if lvl == nil {
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
		wa, wc, wd, wnames := WatcherPresence(lvl)
		m.WatcherActive = wa
		m.WatcherConnectSec = wc
		m.WatcherDisconnectSec = wd
		m.WatcherNames = wnames
		out = append(out, m)
	}
	return out
}
