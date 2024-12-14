package main

import (
	"context"
	"fmt"

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

type Runtime struct {
	app      *fx.App
	splitter *splitter.Splitter
	Data     struct {
		Levels []*splitter.Level
	}
}

func NewRuntime(ctx context.Context) *Runtime {
	var res Runtime

	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logConfig.OutputPaths = []string{"stderr"}

	res.app = fx.New(
		fx.NopLogger,
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
		fx.Populate(&res.splitter),
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
	levels, err := r.splitter.SplitLevels(ctx, memfs.FS{
		"game.log":   memfs.File(gameLog),
		"combat.log": memfs.File(combatLog),
	})
	if err != nil {
		return fmt.Errorf("splitter.SplitLevels: %w", err)
	}
	r.Data.Levels = levels
	return nil
}
