package logger

import "context"

func NewNop() Factory {
	return nopFactory{}
}

type nopFactory struct{}
type nopLogger struct{}

func (nopFactory) For(context.Context) Logger { return nopLogger{} }
func (nopFactory) With(fields ...any) Factory { return nopFactory{} }

func (nopLogger) With(vals ...any) Logger                         { return nopLogger{} }
func (nopLogger) Sync() error                                     { return nil }
func (nopLogger) Debugw(msg string, keysAndValues ...interface{}) {}
func (nopLogger) Infow(msg string, keysAndValues ...interface{})  {}
func (nopLogger) Warnw(msg string, keysAndValues ...interface{})  {}
func (nopLogger) Errorw(msg string, keysAndValues ...interface{}) {}
