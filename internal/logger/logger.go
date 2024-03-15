package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

type Factory interface {
	For(context.Context) Logger
	With(fields ...any) Factory
}

type Logger interface {
	With(vals ...any) Logger
	Sync() error
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
}

type FactoryImpl struct {
	lg *zap.SugaredLogger
}

func NewFactory(lc zap.Config, opts ...zap.Option) (lf Factory, closer context.CancelFunc, err error) {
	lg, err := lc.Build(opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("build logger")
	}
	return &FactoryImpl{lg: lg.Sugar()}, func() { _ = lg.Sync() }, nil
}

func (f *FactoryImpl) With(fields ...any) Factory {
	return &FactoryImpl{lg: f.lg.With(fields...)}
}

func (f *FactoryImpl) For(ctx context.Context) Logger {
	return &LoggerImpl{lg: f.lg}
}

type LoggerImpl struct {
	lg *zap.SugaredLogger
}

func (l *LoggerImpl) With(vals ...any) Logger {
	return &LoggerImpl{lg: l.lg.With(vals...)}
}

func (l *LoggerImpl) Sync() error {
	return l.lg.Sync()
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
