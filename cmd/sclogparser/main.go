package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Feresey/sclogparser/cmd/sclogparser/config"
	"github.com/Feresey/sclogparser/internal/formatter"
	"github.com/Feresey/sclogparser/internal/logger"
	"github.com/Feresey/sclogparser/internal/parser"
	"github.com/Feresey/sclogparser/internal/service"
	"github.com/Feresey/sclogparser/internal/trace"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cfg := config.GetConfig()
	lc := zap.NewDevelopmentConfig()
	lc.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	lg, sync, err := logger.NewFactory(lc)
	if err != nil {
		panic(err)
	}
	defer sync()

	tp, shutdown, err := trace.NewTraceProvider(ctx, "sclogparser")
	if err != nil {
		panic(err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := shutdown(ctx); err != nil {
			lg.For(ctx).Errorw("shutting down tracer", "err", err)
		}
	}()

	fr := formatter.NewFormatter(lg, tp.Tracer("formatter"))
	pr := parser.NewParser(lg, tp.Tracer("parser"))

	svc := service.NewService(cfg, lg, tp.Tracer("service"), fr, pr)

	if err := svc.Run(ctx); err != nil {
		panic(fmt.Sprintf("service.Run: %v", err))
	}
}
