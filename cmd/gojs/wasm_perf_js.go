//go:build js

package main

import (
	"context"
	"syscall/js"
	"time"

	"github.com/Feresey/luxpanel/internal/prettyfmt"
)

func exportStringOutLen(v any) uint64 {
	if s, ok := v.(string); ok {
		return uint64(len(s))
	}
	return 0
}

func exportOutHasErr(v any) bool {
	if v == nil {
		return false
	}
	if _, ok := v.(js.Error); ok {
		return true
	}
	if e, ok := v.(error); ok && e != nil {
		return true
	}
	return false
}

// jsExportPerfEnd logs perf_wasm for a global JS→WASM entry point; used from defer with &ret.
func (r *Runtime) jsExportPerfEnd(ctx context.Context, op string, t0 time.Time, heap0 uint64, ret *any, extra ...any) {
	kv := append(append([]any{}, extra...),
		"out_size", prettyfmt.FormatBytes(exportStringOutLen(*ret)),
		"has_err", exportOutHasErr(*ret),
	)
	logWasmPerf(ctx, r.lg, op, t0, heap0, kv...)
}
