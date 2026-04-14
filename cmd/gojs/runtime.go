package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Feresey/luxpanel/config"
	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/mytrace"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/Feresey/luxpanel/internal/prettyfmt"
	"github.com/Feresey/luxpanel/internal/splitter"
	"github.com/nicksnyder/go-i18n/v2/i18n"
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

	// UI strings (YAML via embed, WASM only).
	i18nBundle   *i18n.Bundle
	uiMessageIDs []string
	// Язык UI ("ru" | "en"), синхронизирован с кнопкой в браузере через setUILanguage.
	uiLang string
}

// normalizeUILangTag приводит тег к "ru" или "en".
func normalizeUILangTag(tag string) string {
	s := strings.TrimSpace(strings.ToLower(tag))
	if s == "" {
		return "en"
	}
	if strings.HasPrefix(s, "ru") {
		return "ru"
	}
	return "en"
}

// SetUILang задаёт язык ответов JSON и сбрасывает кеш WASM.
func (r *Runtime) SetUILang(tag string) {
	r.uiLang = normalizeUILangTag(tag)
	r.bumpWasmCache()
}

// UILang текущий язык UI ("ru" | "en").
func (r *Runtime) UILang() string {
	return normalizeUILangTag(r.uiLang)
}

// Tr возвращает локализованную строку по id сообщения (go-i18n + r.UILang()).
func (r *Runtime) Tr(messageID string) string {
	if r == nil || r.i18nBundle == nil || messageID == "" {
		return messageID
	}
	loc := i18n.NewLocalizer(r.i18nBundle, r.UILang())
	s, err := loc.Localize(&i18n.LocalizeConfig{MessageID: messageID})
	if err != nil || s == "" {
		return messageID
	}
	return s
}

// Trf — как Tr, с подстановкой полей шаблона go-i18n (TemplateData).
func (r *Runtime) Trf(messageID string, data map[string]interface{}) string {
	if r == nil || r.i18nBundle == nil || messageID == "" {
		return messageID
	}
	loc := i18n.NewLocalizer(r.i18nBundle, r.UILang())
	s, err := loc.Localize(&i18n.LocalizeConfig{MessageID: messageID, TemplateData: data})
	if err != nil || s == "" {
		return messageID
	}
	return s
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
	logWasmPerf(ctx, r.lg, "discover_levels", t0, heap0,
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
