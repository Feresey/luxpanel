package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Feresey/luxpanel/config"
	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/mytrace"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/Feresey/luxpanel/internal/service"
	"github.com/Feresey/luxpanel/internal/splitter"
)

func main() {
	var svc *service.Service

	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	app := fx.New(
		fx.NopLogger,
		fx.Supply(
			config.GetConfig(),
			logConfig,
		),
		fx.Provide(
			logger.NewFactory,
			mytrace.NewTraceProvider,
			splitter.NewSplitter,
			service.NewService,
			parser.NewParser,
		),
		fx.Populate(&svc),
	)

	if err := run(app, svc); err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
}

func run(app *fx.App, svc *service.Service) (err error) {
	if err := app.Start(context.Background()); err != nil {
		return fmt.Errorf("fx.Start: %w", err)
	}

	defer func() {
		if stopErr := app.Stop(context.Background()); stopErr != nil && !errors.Is(stopErr, syscall.ENOTTY) {
			err = errors.Join(err, fmt.Errorf("fx.Stop: %w", stopErr))
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := svc.Run(ctx); err != nil {
		return fmt.Errorf("service.Run: %w", err)
	}

	return nil
}
