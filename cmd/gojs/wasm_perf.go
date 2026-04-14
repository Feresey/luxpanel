package main

import (
	"context"
	"runtime"
	"time"

	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/prettyfmt"
)

func heapAlloc() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.HeapAlloc
}

// logWasmPerf logs duration, heap delta, and overall heap/runtime memory from MemStats via prettyfmt (stderr / dev console).
func logWasmPerf(ctx context.Context, lg logger.Factory, op string, t0 time.Time, heapBefore uint64, keysAndValues ...interface{}) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	heapAfter := m.HeapAlloc
	delta := int64(heapAfter) - int64(heapBefore)
	elapsed := time.Since(t0)
	args := []interface{}{
		"op", op,
		"duration", prettyfmt.FormatDuration(elapsed),
		"heap_delta", prettyfmt.FormatBytesDelta(delta),
		"heap_alloc_before", prettyfmt.FormatBytes(heapBefore),
		"heap_alloc", prettyfmt.FormatBytes(heapAfter),
		"heap_sys", prettyfmt.FormatBytes(m.HeapSys),
		"heap_inuse", prettyfmt.FormatBytes(m.HeapInuse),
		"heap_idle", prettyfmt.FormatBytes(m.HeapIdle),
		"sys", prettyfmt.FormatBytes(m.Sys),
		"total_alloc", prettyfmt.FormatBytes(m.TotalAlloc),
	}
	args = append(args, keysAndValues...)
	lg.For(ctx).Infow("perf_wasm", args...)
}
