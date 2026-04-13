package prettyfmt

import (
	"testing"
	"time"
)

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		n    uint64
		want string
	}{
		{0, "0 B"},
		{500, "500 B"},
		{1024, "1 KiB"},
		{2048, "2 KiB"},
		{3 * 1024 * 1024, "3 MiB"},
	}
	for _, tt := range tests {
		if got := FormatBytes(tt.n); got != tt.want {
			t.Errorf("FormatBytes(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}

func TestFormatBytesDelta(t *testing.T) {
	if got := FormatBytesDelta(2 * 1024 * 1024); got != "+2 MiB" {
		t.Errorf("got %q", got)
	}
	if got := FormatBytesDelta(-1536); got != "-1.5 KiB" {
		t.Errorf("got %q", got)
	}
}

func TestFormatDuration(t *testing.T) {
	if got := FormatDuration(500 * time.Microsecond); got != "500µs" {
		t.Errorf("got %q", got)
	}
	if got := FormatDuration(time.Millisecond / 2); got != "0.5ms" {
		t.Errorf("got %q", got)
	}
	if got := FormatDuration(1500 * time.Millisecond); got != "1.5s" {
		t.Errorf("got %q", got)
	}
}
