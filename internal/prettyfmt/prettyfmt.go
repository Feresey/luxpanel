// Package prettyfmt turns byte counts and durations into short human-readable strings (e.g. "2.1 MB", "4.5ms").
package prettyfmt

import (
	"fmt"
	"time"
)

// FormatBytes renders n bytes with binary prefixes (KiB, MiB, GiB).
func FormatBytes(n uint64) string {
	switch {
	case n >= 1<<40:
		return trimFloat(float64(n)/float64(1<<40)) + " TiB"
	case n >= 1<<30:
		return trimFloat(float64(n)/float64(1<<30)) + " GiB"
	case n >= 1<<20:
		return trimFloat(float64(n)/float64(1<<20)) + " MiB"
	case n >= 1<<10:
		return trimFloat(float64(n)/float64(1<<10)) + " KiB"
	default:
		if n == 0 {
			return "0 B"
		}
		return fmt.Sprintf("%d B", n)
	}
}

// FormatBytesDelta renders a signed byte delta (e.g. "+12 KiB", "-2.1 MiB").
func FormatBytesDelta(delta int64) string {
	if delta == 0 {
		return "0 B"
	}
	sign := "+"
	v := delta
	if v < 0 {
		sign = "-"
		v = -v
	}
	return sign + FormatBytes(uint64(v))
}

// FormatDuration renders d in a compact form (e.g. "0µs", "850µs", "4.2ms", "1.25s", "2m30s").
func FormatDuration(d time.Duration) string {
	if d <= 0 {
		return "0µs"
	}
	if d < time.Microsecond {
		return fmt.Sprintf("%dns", d.Nanoseconds())
	}
	if d < time.Millisecond {
		return fmt.Sprintf("%.0fµs", float64(d.Microseconds()))
	}
	if d < time.Second {
		return fmt.Sprintf("%.1fms", float64(d.Nanoseconds())/1e6)
	}
	if d < time.Minute {
		return trimFloat(d.Seconds()) + "s"
	}
	sec := int(d.Round(time.Second).Seconds())
	m := sec / 60
	s := sec % 60
	if m < 60 {
		return fmt.Sprintf("%dm%02ds", m, s)
	}
	h := m / 60
	m %= 60
	return fmt.Sprintf("%dh%02dm", h, m)
}

func trimFloat(x float64) string {
	s := fmt.Sprintf("%.2f", x)
	for len(s) > 1 && s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}
	if len(s) > 0 && s[len(s)-1] == '.' {
		s = s[:len(s)-1]
	}
	return s
}
