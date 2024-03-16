package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Feresey/sclogparser/cmd/sclogparser/config"
	"github.com/Feresey/sclogparser/internal/formatter"
	"github.com/Feresey/sclogparser/internal/logger"
	"github.com/Feresey/sclogparser/internal/parser"
	"github.com/Feresey/sclogparser/internal/service"
	"github.com/Feresey/sclogparser/internal/trace"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cfg := config.GetConfig()

	lg, sync, err := logger.NewFactory(zap.NewDevelopmentConfig())
	if err != nil {
		panic(err)
	}
	defer sync()

	tp, err := trace.NewTraceProvider(ctx)

	fr := formatter.NewFormatter(lg, tp.Tracer("formatter"))
	pr := parser.NewParser(lg, tp.Tracer("parser"))

	svc := service.NewService(cfg, lg, tp.Tracer("service"), fr, pr)

	if err := svc.Run(ctx); err != nil {
		panic(fmt.Sprintf("service.Run: %v", err))
	}
}
