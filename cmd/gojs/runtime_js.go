//go:build js

package main

import (
	"context"
	"fmt"
	"runtime/pprof"
	"slices"
	"strconv"
	"strings"
	"syscall/js"
	"time"

	"github.com/Feresey/luxpanel/internal/splitter"
)

func wasmCacheKey(parts ...string) string {
	return strings.Join(parts, "\x1e")
}

func (r *Runtime) RegisterJSBindings(ctx context.Context) {
	r.register("parseFiles", func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			return js.Error{Value: js.ValueOf(fmt.Sprintf("2 arguments required, %d got", len(args)))}
		}
		if args[0].Type() != js.TypeString || args[1].Type() != js.TypeString {
			return js.Error{Value: js.ValueOf(fmt.Sprintf("string arguments expected, but got: %s %s", args[0].Type(), args[1].Type()))}
		}
		gameLog := args[0].String()
		combatLog := args[1].String()
		if err := r.LoadFiles(ctx, gameLog, combatLog); err != nil {
			r.lg.For(ctx).Errorw("parseFiles failed", "err", err)
			return js.Error{Value: js.ValueOf(fmt.Sprintf("LoadFiles: %v", err))}
		}
		r.lg.For(ctx).Infow("parseFiles done", "levels", len(r.Data.Levels), "game_log_len", len(gameLog), "combat_log_len", len(combatLog))
		return nil
	})

	r.register("getLevelsMetaJSON", func(this js.Value, args []js.Value) any {
		const k = "levelsMeta"
		if s, ok := r.wasmCacheGet(k); ok {
			return s
		}
		s, err := r.marshalLevelsMetaJSON(ctx, r.discoverPreviewLevels())
		if err != nil {
			return "[]"
		}
		r.wasmCacheSet(k, s)
		return s
	})

	r.register("getTeamPanelsRosterJSON", func(this js.Value, args []js.Value) any {
		level, ok := r.levelFromArgs(ctx, args, 1)
		if !ok {
			return "null"
		}
		idx := args[0].Int()
		k := wasmCacheKey("roster", strconv.Itoa(idx))
		if s, ok := r.wasmCacheGet(k); ok {
			return s
		}
		s, err := r.marshalTeamPanelsRosterJSON(ctx, level)
		if err != nil {
			return "null"
		}
		r.wasmCacheSet(k, s)
		return s
	})

	r.register("getMatchSummaryJSON", func(this js.Value, args []js.Value) any {
		level, ok := r.levelFromArgs(ctx, args, 1)
		if !ok {
			return "null"
		}
		idx := args[0].Int()
		k := wasmCacheKey("summary", strconv.Itoa(idx))
		if s, ok := r.wasmCacheGet(k); ok {
			return s
		}
		s, err := r.marshalMatchSummaryJSON(ctx, level)
		if err != nil {
			return "null"
		}
		r.wasmCacheSet(k, s)
		return s
	})

	r.register("getDamageChartsJSON", func(this js.Value, args []js.Value) any {
		level, ok := r.levelFromArgs(ctx, args, 1)
		if !ok {
			return "null"
		}
		idx := args[0].Int()
		timeRangeJSON := ""
		if len(args) >= 2 && args[1].Type() == js.TypeString {
			timeRangeJSON = args[1].String()
		}
		k := wasmCacheKey("charts", strconv.Itoa(idx), "damage", timeRangeJSON)
		if s, ok := r.wasmCacheGet(k); ok {
			return s
		}
		s, err := r.marshalChartsJSON(ctx, level, "damage", timeRangeJSON)
		if err != nil {
			return "null"
		}
		r.wasmCacheSet(k, s)
		return s
	})

	r.register("getChartsJSON", func(this js.Value, args []js.Value) any {
		level, ok := r.levelFromArgs(ctx, args, 2)
		if !ok {
			return "null"
		}
		idx := args[0].Int()
		mode := "damage"
		if args[1].Type() == js.TypeString {
			mode = args[1].String()
		}
		timeRangeJSON := ""
		if len(args) >= 3 && args[2].Type() == js.TypeString {
			timeRangeJSON = args[2].String()
		}
		k := wasmCacheKey("charts", strconv.Itoa(idx), mode, timeRangeJSON)
		if s, ok := r.wasmCacheGet(k); ok {
			return s
		}
		s, err := r.marshalChartsJSON(ctx, level, mode, timeRangeJSON)
		if err != nil {
			return "null"
		}
		r.wasmCacheSet(k, s)
		return s
	})

	r.register("getChartMatricesJSON", func(this js.Value, args []js.Value) any {
		level, ok := r.levelFromArgs(ctx, args, 2)
		if !ok {
			return "null"
		}
		idx := args[0].Int()
		mode := "damage"
		if args[1].Type() == js.TypeString {
			mode = args[1].String()
		}
		timeRangeJSON := ""
		if len(args) >= 3 && args[2].Type() == js.TypeString {
			timeRangeJSON = args[2].String()
		}
		optsJSON := ""
		if len(args) >= 4 && args[3].Type() == js.TypeString {
			optsJSON = args[3].String()
		}
		k := wasmCacheKey("matrices", strconv.Itoa(idx), mode, timeRangeJSON, optsJSON)
		if s, ok := r.wasmCacheGet(k); ok {
			return s
		}
		s, err := r.marshalChartMatricesJSON(ctx, level, mode, timeRangeJSON, optsJSON)
		if err != nil {
			return "null"
		}
		r.wasmCacheSet(k, s)
		return s
	})

	r.register("getTimelineJSON", func(this js.Value, args []js.Value) any {
		level, ok := r.levelFromArgs(ctx, args, 1)
		if !ok {
			return "null"
		}
		idx := args[0].Int()
		focusPlayer := ""
		if len(args) >= 2 && args[1].Type() == js.TypeString {
			focusPlayer = args[1].String()
		}
		k := wasmCacheKey("timeline", strconv.Itoa(idx), focusPlayer)
		if s, ok := r.wasmCacheGet(k); ok {
			return s
		}
		s, err := r.marshalTimelineJSON(ctx, level, focusPlayer)
		if err != nil {
			return "null"
		}
		r.wasmCacheSet(k, s)
		return s
	})

	r.register("getBattleInsightJSON", func(this js.Value, args []js.Value) any {
		level, ok := r.levelFromArgs(ctx, args, 1)
		if !ok {
			return "null"
		}
		idx := args[0].Int()
		timeRangeJSON := ""
		if len(args) >= 2 && args[1].Type() == js.TypeString {
			timeRangeJSON = args[1].String()
		}
		focusPlayer := ""
		if len(args) >= 3 && args[2].Type() == js.TypeString {
			focusPlayer = args[2].String()
		}
		k := wasmCacheKey("battle", strconv.Itoa(idx), timeRangeJSON, focusPlayer)
		if s, ok := r.wasmCacheGet(k); ok {
			return s
		}
		s, err := r.marshalBattleInsightJSON(ctx, level, timeRangeJSON, focusPlayer)
		if err != nil {
			return "null"
		}
		r.wasmCacheSet(k, s)
		return s
	})

	r.register("getCombatLogLinesJSON", func(this js.Value, args []js.Value) any {
		level, ok := r.levelFromArgs(ctx, args, 1)
		if !ok {
			return "[]"
		}
		timeRangeJSON := ""
		if len(args) >= 2 && args[1].Type() == js.TypeString {
			timeRangeJSON = args[1].String()
		}
		s, err := r.marshalCombatLogLinesJSON(ctx, level, timeRangeJSON)
		if err != nil {
			return "[]"
		}
		return s
	})

	r.register("getDamageFilterMetaJSON", func(this js.Value, args []js.Value) any {
		level, ok := r.levelFromArgs(ctx, args, 1)
		if !ok {
			return "null"
		}
		idx := args[0].Int()
		initiator := ""
		if len(args) >= 2 && args[1].Type() == js.TypeString {
			initiator = args[1].String()
		}
		timeRangeJSON := ""
		if len(args) >= 3 && args[2].Type() == js.TypeString {
			timeRangeJSON = args[2].String()
		}
		k := wasmCacheKey("dfmeta", strconv.Itoa(idx), initiator, timeRangeJSON)
		if s, hit := r.wasmCacheGet(k); hit {
			return s
		}

		humans, bots := humanAndBotPlayerNames(level)
		h := mapKeysSorted(humans)
		b := mapKeysSorted(bots)
		r.lg.For(ctx).Infow("damage filter players", "humans_count", len(h), "bots_count", len(b), "humans", h, "bots", b, "initiator", initiator)

		s, err := r.marshalDamageFilterMetaJSON(ctx, level, initiator, timeRangeJSON)
		if err != nil {
			return "null"
		}
		r.wasmCacheSet(k, s)
		return s
	})

	r.register("getDamageTableJSON", func(this js.Value, args []js.Value) any {
		level, ok := r.levelFromArgs(ctx, args, 2)
		if !ok {
			return "null"
		}
		filterJSON := ""
		if args[1].Type() == js.TypeString {
			filterJSON = args[1].String()
		}
		s, err := r.marshalDamageTableJSON(ctx, level, filterJSON)
		if err != nil {
			return "null"
		}
		return s
	})

	r.register("getDamageDefaultTableJSON", func(this js.Value, args []js.Value) any {
		level, ok := r.levelFromArgs(ctx, args, 2)
		if !ok {
			return "null"
		}
		filterJSON := ""
		if args[1].Type() == js.TypeString {
			filterJSON = args[1].String()
		}
		s, err := r.marshalDamageDefaultFiltersTableJSON(ctx, level, filterJSON)
		if err != nil {
			return "null"
		}
		return s
	})

	r.register("profile", func(this js.Value, args []js.Value) any {
		var buf strings.Builder
		if err := pprof.WriteHeapProfile(&buf); err != nil {
			return fmt.Errorf("write profile: %w", err)
		}
		return buf.String()
	})
}

func (r *Runtime) register(name string, fn func(this js.Value, args []js.Value) any) {
	js.Global().Set(name, js.FuncOf(fn))
}

func (r *Runtime) levelFromArgs(ctx context.Context, args []js.Value, minArgs int) (*splitter.Level, bool) {
	if len(args) < minArgs {
		return nil, false
	}
	idx := args[0].Int()
	if idx < 0 || idx >= len(r.Data.Levels) {
		return nil, false
	}
	if r.Data.Levels[idx] != nil {
		return r.Data.Levels[idx], true
	}
	if r.discover == nil {
		return nil, false
	}
	t0 := time.Now()
	heap0 := heapAlloc()
	lvl, err := r.splitter.HydrateLevel(ctx, r.discover, idx)
	if err != nil {
		r.lg.For(ctx).Errorw("HydrateLevel failed", "idx", idx, "err", err)
		return nil, false
	}
	logPerfWasm(ctx, r.lg, "hydrate_level", t0, heap0, "idx", idx)
	r.Data.Levels[idx] = lvl
	return lvl, true
}

func mapKeysSorted(m map[string]struct{}) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	slices.Sort(out)
	return out
}

