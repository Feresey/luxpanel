package logger

import (
	"context"
	"fmt"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type Factory interface {
	For(context.Context) Logger
	With(fields ...any) Factory
}

type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
}

type FactoryImpl struct {
	lg *otelzap.SugaredLogger
}

func NewFactory(lc zap.Config, opts ...zap.Option) (lf Factory, sync context.CancelFunc, err error) {
	lg, err := lc.Build(opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("build logger")
	}
	ozlog := otelzap.New(lg,
		otelzap.WithCaller(true),
		otelzap.WithMinLevel(zap.InfoLevel),
		otelzap.WithTraceIDField(true),
		otelzap.WithCallerDepth(1),
	).WithOptions(zap.AddCallerSkip(1)).
		Sugar()
	return &FactoryImpl{lg: ozlog}, func() { _ = lg.Sync() }, nil
}

func (f *FactoryImpl) With(fields ...any) Factory {
	return &FactoryImpl{
		lg: f.lg.With(fields...),
	}
}

func (f *FactoryImpl) For(ctx context.Context) Logger {
	return &LoggerImpl{
		lg: f.lg.Ctx(ctx),
	}
}

type LoggerImpl struct {
	lg otelzap.SugaredLoggerWithCtx
}

func (l *LoggerImpl) Debugw(message string, args ...interface{}) {
	l.lg.Debugw(message, args...)
}

func (l *LoggerImpl) Infow(message string, args ...interface{}) {
	l.lg.Infow(message, args...)
}

func (l *LoggerImpl) Warnw(message string, args ...interface{}) {
	l.lg.Warnw(message, args...)
}

func (l *LoggerImpl) Errorw(message string, args ...interface{}) {
	l.lg.Errorw(message, args...)
}
