package main

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/Feresey/luxpanel/config"
	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/mytrace"
	"github.com/Feresey/luxpanel/internal/prettyfmt"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/Feresey/luxpanel/internal/splitter"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type stringCacheEntry struct {
	gen uint64
	s   string
}

type Runtime struct {
	lg       logger.Factory
	app      *fx.App
	splitter *splitter.Splitter
	Data     struct {
		// Разреженный массив: индекс заполняется полным парсингом среза лога при первом обращении к матчу.
		Levels []*splitter.Level
	}
	// Результат ленивого первого прохода (метаданные селектора и байтовые диапазоны).
	discover *splitter.DiscoverResult
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
	t0 := time.Now()
	heap0 := heapAlloc()
	disc, err := r.splitter.DiscoverLevels(ctx, gameLog, combatLog)
	if err != nil {
		return fmt.Errorf("splitter.DiscoverLevels: %w", err)
	}
	logPerfWasm(ctx, r.lg, "discover_levels", t0, heap0,
		"level_count", len(disc.Spans),
		"game_size", prettyfmt.FormatBytes(uint64(len(gameLog))),
		"combat_size", prettyfmt.FormatBytes(uint64(len(combatLog))),
	)
	r.discover = disc
	r.Data.Levels = make([]*splitter.Level, len(disc.Spans))
	r.bumpWasmCache()
	return nil
}

func (r *Runtime) bumpWasmCache() {
	r.wasmCacheGen++
}

func (r *Runtime) discoverPreviewLevels() []*splitter.Level {
	if r.discover == nil {
		return nil
	}
	return r.discover.PreviewLevels
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

func heapAlloc() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.HeapAlloc
}

func logPerfWasm(ctx context.Context, lg logger.Factory, op string, t0 time.Time, heapBefore uint64, keysAndValues ...interface{}) {
	heapAfter := heapAlloc()
	delta := int64(heapAfter) - int64(heapBefore)
	args := []interface{}{
		"op", op,
		"duration", prettyfmt.FormatDuration(time.Since(t0)),
		"heap_delta", prettyfmt.FormatBytesDelta(delta),
		"heap_alloc", prettyfmt.FormatBytes(heapAfter),
	}
	args = append(args, keysAndValues...)
	lg.For(ctx).Infow("perf_wasm", args...)
}
