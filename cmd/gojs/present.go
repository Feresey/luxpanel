//go:build js

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Feresey/luxpanel/internal/front"
	"github.com/Feresey/luxpanel/internal/prettyfmt"
	"github.com/Feresey/luxpanel/internal/splitter"
)

func (r *Runtime) formatGameDuration(sec float32) string {
	if sec <= 0 {
		return ""
	}
	if sec >= 60 {
		return r.Trf("go_duration_min", map[string]interface{}{
			"Min": fmt.Sprintf("%.1f", float64(sec)/60),
		})
	}
	return r.Trf("go_duration_sec", map[string]interface{}{
		"Sec": fmt.Sprintf("%.0f", sec),
	})
}

func (r *Runtime) formatLevelLabel(m front.LevelMeta) string {
	mode := m.GameMode
	if mode == "" {
		mode = "?"
	}
	mapName := m.MapName
	if mapName == "" {
		mapName = "?"
	}
	line := r.Trf("go_level_label", map[string]interface{}{
		"N":    m.Index + 1,
		"Mode": mode,
		"Map":  mapName,
	})
	if d := r.formatGameDuration(m.GameTimeSec); d != "" {
		line += " (" + d + ")"
	} else if m.StartTime != "" {
		line += " (" + m.StartTime + ")"
	}
	return line
}

func (r *Runtime) marshalLevelsMetaJSON(ctx context.Context, levels []*splitter.Level) (string, error) {
	if len(levels) == 0 {
		r.lg.For(ctx).Debugw("marshalLevelsMetaJSON", "levels", 0, "empty", true)
		return "[]", nil
	}
	raw := front.BuildLevelMetaSlice(levels)
	for i := range raw {
		if levels[i] == nil {
			raw[i].Label = r.Trf("go_level_match_fallback", map[string]interface{}{"N": i + 1})
			continue
		}
		raw[i].Label = r.formatLevelLabel(raw[i])
	}
	b, err := json.Marshal(raw)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalLevelsMetaJSON", "err", err, "levels", len(levels))
		return "", fmt.Errorf("marshal levels meta: %w", err)
	}
	s := string(b)
	r.lg.For(ctx).Debugw("marshalLevelsMetaJSON", "levels", len(levels), "out_len", len(s))
	return s, nil
}

func (r *Runtime) chartMetricWord(metric string) string {
	switch metric {
	case "heal":
		return r.Tr("go_chart_metric_heal")
	case "kill":
		return r.Tr("go_chart_metric_kill")
	default:
		return r.Tr("go_chart_metric_damage")
	}
}

func (r *Runtime) localizeChartCurves(ch [][]front.ChartCurve) {
	for ti := range ch {
		for ci := range ch[ti] {
			c := &ch[ti][ci]
			if c.CurveKind == "avg" {
				word := r.chartMetricWord(c.Metric)
				c.Name = r.Trf("go_chart_curve_avg", map[string]interface{}{"Metric": word})
			}
		}
	}
}

func (r *Runtime) marshalChartsJSON(ctx context.Context, level *splitter.Level, mode string, timeRangeJSON string) (s string, err error) {
	t0 := time.Now()
	heap0 := heapAlloc()
	defer func() {
		ob := 0
		if err == nil {
			ob = len(s)
		}
		logWasmPerf(ctx, r.lg, "marshal_charts_json", t0, heap0,
			"mode", mode,
			"out_size", prettyfmt.FormatBytes(uint64(ob)),
			"has_err", err != nil,
		)
	}()
	lo, hi := front.ClampTimeRange(level, timeRangeJSON)
	ch := front.BuildChartsForMode(level, mode, lo, hi)
	r.localizeChartCurves(ch)
	var b []byte
	b, err = json.Marshal(ch)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalChartsJSON", "err", err, "mode", mode)
		return "", fmt.Errorf("marshal charts: %w", err)
	}
	s = string(b)
	r.lg.For(ctx).Debugw("marshalChartsJSON", "mode", mode, "time_lo", lo, "time_hi", hi, "out_len", len(s))
	return s, nil
}

func (r *Runtime) marshalChartMatricesJSON(ctx context.Context, level *splitter.Level, mode string, timeRangeJSON, optsJSON string) (result string, err error) {
	t0 := time.Now()
	heap0 := heapAlloc()
	defer func() {
		logWasmPerf(ctx, r.lg, "marshal_chart_matrices", t0, heap0,
			"mode", mode,
			"out_size", prettyfmt.FormatBytes(uint64(len(result))),
			"has_err", err != nil,
			"opts_size", prettyfmt.FormatBytes(uint64(len(optsJSON))),
		)
	}()
	if level == nil {
		var b []byte
		b, err = json.Marshal(front.ChartMatricesResult{})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalChartMatricesJSON", "err", err, "nil_level", true)
			return "", err
		}
		result = string(b)
		return result, nil
	}
	lo, hi := front.ClampTimeRange(level, timeRangeJSON)
	opts := front.ParseChartQueryOptions(optsJSON)
	res := front.BuildChartMatrices(level, mode, lo, hi, opts)
	var b []byte
	b, err = json.Marshal(res)
	if err != nil {
		return "", err
	}
	result = string(b)
	r.lg.For(ctx).Debugw("marshalChartMatricesJSON", "out_len", len(result))
	return result, nil
}

func (r *Runtime) marshalTeamPanelsRosterJSON(_ context.Context, level *splitter.Level) (string, error) {
	roster := front.BuildTeamPanelsRoster(level)
	if roster == nil {
		return "null", nil
	}
	b, err := json.Marshal(roster)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type matchSummaryTeam struct {
	TeamID int     `json:"team_id"`
	Name   string  `json:"name"`
	Role   string  `json:"role"`
	Damage float64 `json:"damage"`
	Heal   float64 `json:"heal"`
	Kills  int     `json:"kills"`
}

type matchSummaryResult struct {
	MapName       string             `json:"map_name,omitempty"`
	GameMode      string             `json:"game_mode,omitempty"`
	StartTime     string             `json:"start_time,omitempty"`
	DurationSec   float64            `json:"duration_sec"`
	TeamLeftID    int                `json:"team_left_id,omitempty"`
	TeamRightID   int                `json:"team_right_id,omitempty"`
	TeamLeftName  string             `json:"team_left_name,omitempty"`
	TeamRightName string             `json:"team_right_name,omitempty"`
	DamageLeft    float64            `json:"damage_left"`
	DamageRight   float64            `json:"damage_right"`
	HealLeft      float64            `json:"heal_left"`
	HealRight     float64            `json:"heal_right"`
	KillsLeft     int                `json:"kills_left"`
	KillsRight    int                `json:"kills_right"`
	Teams         []matchSummaryTeam `json:"teams,omitempty"`
}

func (r *Runtime) marshalMatchSummaryJSON(ctx context.Context, level *splitter.Level) (string, error) {
	if level == nil {
		b, err := json.Marshal(matchSummaryResult{})
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	c := front.BuildMatchSummary(level)
	out := matchSummaryResult{
		MapName:     c.MapName,
		GameMode:    c.GameMode,
		StartTime:   c.StartTime,
		DurationSec: c.DurationSec,
		TeamLeftID:  c.TeamLeftID,
		TeamRightID: c.TeamRightID,
		DamageLeft:  c.DamageLeft,
		DamageRight: c.DamageRight,
		HealLeft:    c.HealLeft,
		HealRight:   c.HealRight,
		KillsLeft:   c.KillsLeft,
		KillsRight:  c.KillsRight,
	}
	if c.TeamLeftID != 0 {
		out.TeamLeftName = r.Trf("go_team_name", map[string]interface{}{"N": c.TeamLeftID})
	}
	if c.TeamRightID != 0 {
		out.TeamRightName = r.Trf("go_team_name", map[string]interface{}{"N": c.TeamRightID})
	}
	allyID, _ := front.AllyTeamIDForTimeline(level)
	for _, t := range c.TeamAggs {
		role := r.Tr("team_label_enemies")
		if t.TeamID == allyID {
			role = r.Tr("team_label_allies")
		} else if c.MultiTeamMode {
			role = r.Trf("go_match_role_enemies_team", map[string]interface{}{"Team": strconv.Itoa(t.TeamID)})
		}
		side := ""
		if t.SpawnID == 1 {
			side = r.Tr("go_match_spawn_left")
		} else if t.SpawnID == 2 {
			side = r.Tr("go_match_spawn_right")
		}
		out.Teams = append(out.Teams, matchSummaryTeam{
			TeamID: t.TeamID,
			Name:   r.Trf("go_team_name", map[string]interface{}{"N": t.TeamID}),
			Role:   role + side,
			Damage: t.Damage,
			Heal:   t.Heal,
			Kills:  t.Kills,
		})
	}
	b, err := json.Marshal(out)
	if err != nil {
		return "", err
	}
	r.lg.For(ctx).Debugw("marshalMatchSummaryJSON", "left", c.TeamLeftID, "right", c.TeamRightID, "duration", out.DurationSec, "out_len", len(b))
	return string(b), nil
}

func (r *Runtime) marshalBattleInsightJSON(ctx context.Context, level *splitter.Level, timeRangeJSON string, focusPlayer string) (string, error) {
	if level == nil {
		b, err := json.Marshal(front.BattleInsightResult{})
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	res := front.BuildBattleInsight(level, timeRangeJSON, focusPlayer)
	switch res.GodGiftSide {
	case "allies":
		res.GodGiftLabel = r.Tr("go_godgift_allies")
		r.lg.For(ctx).Infow("godgift_detected", "time_sec", res.GodGiftSec, "label", res.GodGiftLabel)
	case "enemies":
		res.GodGiftLabel = r.Tr("go_godgift_enemies")
		r.lg.For(ctx).Infow("godgift_detected", "time_sec", res.GodGiftSec, "label", res.GodGiftLabel)
	}
	b, err := json.Marshal(res)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalBattleInsightJSON", "err", err)
		return "", err
	}
	return string(b), nil
}

func (r *Runtime) marshalTimelineJSON(ctx context.Context, level *splitter.Level, focusedPlayer string) (string, error) {
	if level == nil {
		b, err := json.Marshal(front.TimelineResult{})
		if err != nil {
			r.lg.For(ctx).Errorw("marshalTimelineJSON", "err", err, "nil_level", true)
			return "", err
		}
		return string(b), nil
	}
	res := front.BuildTimeline(level, focusedPlayer)
	b, err := json.Marshal(res)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalTimelineJSON", "err", err)
		return "", err
	}
	s := string(b)
	var cSpawn, cEnemySpawn, cKill, cDeath, cOther int
	for _, m := range res.Markers {
		switch m.Kind {
		case "spawn":
			cSpawn++
		case "enemy_spawn":
			cEnemySpawn++
		case "kill":
			cKill++
		case "death":
			cDeath++
		default:
			cOther++
		}
	}
	r.lg.For(ctx).Debugw(
		"marshalTimelineJSON",
		"markers", len(res.Markers),
		"spawn", cSpawn,
		"enemy_spawn", cEnemySpawn,
		"kill", cKill,
		"death", cDeath,
		"other", cOther,
		"end_sec", res.EndSec,
		"out_len", len(s),
	)
	return s, nil
}

func (r *Runtime) marshalDamageFilterMetaJSON(ctx context.Context, level *splitter.Level, initiator string, timeRangeJSON string) (string, error) {
	lo, hi := front.ClampTimeRange(level, timeRangeJSON)
	meta := front.CollectDamageFilterMeta(level, initiator, lo, hi)
	b, err := json.Marshal(meta)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalDamageFilterMetaJSON", "err", err, "initiator", initiator)
		return "", err
	}
	s := string(b)
	r.lg.For(ctx).Debugw("marshalDamageFilterMetaJSON", "initiator", initiator, "time_lo", lo, "time_hi", hi,
		"players", len(meta.Players), "recipients", len(meta.Recipients), "weapons", len(meta.Weapons), "out_len", len(s))
	return s, nil
}

func (r *Runtime) marshalDamageTableJSON(ctx context.Context, level *splitter.Level, filterJSON string) (string, error) {
	res := front.BuildDamageTableResult(level, filterJSON)
	if res.Row.IncomingPerspective && res.Row.Source != "" {
		res.Row.Source = res.Row.Source + r.Tr("df_source_incoming_suffix")
	}
	b, err := json.Marshal(res)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalDamageTableJSON", "err", err)
		return "", err
	}
	s := string(b)
	r.lg.For(ctx).Debugw("marshalDamageTableJSON", "hits", res.Row.Summary.Hits, "events", len(res.Events), "out_len", len(s))
	return s, nil
}

func (r *Runtime) marshalDamageDefaultFiltersTableJSON(ctx context.Context, level *splitter.Level, filterJSON string) (string, error) {
	res := front.BuildDamageDefaultTableResult(level, filterJSON)
	b, err := json.Marshal(res)
	if err != nil {
		r.lg.For(ctx).Errorw("marshalDamageDefaultFiltersTableJSON", "err", err)
		return "", err
	}
	s := string(b)
	r.lg.For(ctx).Debugw("marshalDamageDefaultFiltersTableJSON", "rows", len(res.Rows), "out_len", len(s))
	return s, nil
}
