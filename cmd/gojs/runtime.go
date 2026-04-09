package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Feresey/luxpanel/config"
	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/mytrace"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/Feresey/luxpanel/internal/splitter"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/spiretechnology/go-memfs"
)

type stringCacheEntry struct {
	gen uint64
	s   string
}

type Runtime struct {
	lg       logger.Factory
	app      *fx.App
	splitter *splitter.Splitter
	logTime  time.Time
	logFS    memfs.FS
	ranges   []splitter.MatchRange
	metaHints []splitter.MatchMeta
	Data     struct {
		Levels []*splitter.Level
	}
	// Ответы WASM для одинаковых аргументов; сброс при LoadFiles (wasmCacheGen++).
	wasmCacheGen  uint64
	wasmJSONCache map[string]stringCacheEntry
}

func NewRuntime(ctx context.Context) *Runtime {
	var res Runtime

	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logConfig.OutputPaths = []string{"stderr"}

	res.app = fx.New(
		fx.Supply(
			&config.TraceConfig{
				ServiceName: "luxpanel",
				Enabled:     false,
				Addr:        "",
			},
			logConfig,
		),
		fx.Provide(
			logger.NewFactory,
			mytrace.NewTraceProvider,
			splitter.NewSplitter,
			parser.NewParser,
		),
		fx.Populate(
			&res.splitter,
			&res.lg,
		),
	)

	return &res
}

func (r *Runtime) Start(ctx context.Context) error {
	if err := r.app.Start(ctx); err != nil {
		return fmt.Errorf("fx.Start: %w", err)
	}
	return nil
}

func (r *Runtime) Stop(ctx context.Context) error {
	if err := r.app.Stop(ctx); err != nil {
		return fmt.Errorf("fx.Stop: %w", err)
	}
	return nil
}

func (r *Runtime) LoadFiles(ctx context.Context, gameLog, combatLog string) error {
	fs := memfs.FS{
		"game.log":   memfs.File(gameLog),
		"combat.log": memfs.File(combatLog),
	}
	logTime, ranges, err := r.splitter.ScanMatchRanges(ctx, fs)
	if err != nil {
		return fmt.Errorf("splitter.ScanMatchRanges: %w", err)
	}
	r.logFS = fs
	r.logTime = logTime
	r.ranges = ranges
	metaHints, err := r.splitter.BuildMatchMetaByRanges(ctx, fs, logTime, ranges)
	if err != nil {
		return fmt.Errorf("splitter.BuildMatchMetaByRanges: %w", err)
	}
	r.metaHints = metaHints
	r.Data.Levels = make([]*splitter.Level, len(r.ranges))
	r.bumpWasmCache()
	return nil
}

func (r *Runtime) ensureLevelParsed(ctx context.Context, idx int) (*splitter.Level, error) {
	if idx < 0 || idx >= len(r.Data.Levels) {
		return nil, fmt.Errorf("level index out of range: %d", idx)
	}
	if r.Data.Levels[idx] != nil {
		return r.Data.Levels[idx], nil
	}
	if len(r.ranges) == 0 {
		return nil, fmt.Errorf("no match ranges")
	}
	lvl, err := r.splitter.ParseLevelByRange(ctx, r.logFS, r.logTime, r.ranges[idx])
	if err != nil {
		return nil, fmt.Errorf("splitter.ParseLevelByRange[%d]: %w", idx, err)
	}
	r.Data.Levels[idx] = lvl
	return lvl, nil
}

func (r *Runtime) bumpWasmCache() {
	r.wasmCacheGen++
}

func (r *Runtime) wasmCacheGet(key string) (string, bool) {
	if r.wasmJSONCache == nil {
		return "", false
	}
	e, ok := r.wasmJSONCache[key]
	if !ok || e.gen != r.wasmCacheGen {
		return "", false
	}
	return e.s, true
}

func (r *Runtime) wasmCacheSet(key, val string) {
	if r.wasmJSONCache == nil {
		r.wasmJSONCache = make(map[string]stringCacheEntry)
	}
	r.wasmJSONCache[key] = stringCacheEntry{gen: r.wasmCacheGen, s: val}
}
