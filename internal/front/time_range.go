package front

import (
	"encoding/json"
	"math"
	"strings"
	"time"

	"github.com/Feresey/luxpanel/internal/splitter"
)

func LevelSpanSeconds(level *splitter.Level) float64 {
	if level == nil {
		return 0
	}
	if level.StartLevelTime.IsZero() || level.EndLevelTime.IsZero() {
		return 0
	}
	d := level.EndLevelTime.Sub(level.StartLevelTime)
	if d < 0 {
		return 0
	}
	span := float64(d) / float64(time.Second)
	if math.IsNaN(span) || math.IsInf(span, 0) || span < 0 {
		return 0
	}
	const hardMaxMatchSpanSec = 12 * 60 * 60
	if span > hardMaxMatchSpanSec {
		return hardMaxMatchSpanSec
	}
	return span
}

func TimeInRangeFromStart(t, t0 time.Time, lo, hi float64) bool {
	from := t0.Add(time.Duration(lo * float64(time.Second)))
	to := t0.Add(time.Duration(hi * float64(time.Second)))
	return !t.Before(from) && !t.After(to)
}

func ClampTimeRange(level *splitter.Level, jsonStr string) (lo, hi float64) {
	span := LevelSpanSeconds(level)
	lo, hi = 0, span
	s := strings.TrimSpace(jsonStr)
	if s == "" || s == "null" {
		return lo, hi
	}
	var tr struct {
		TimeFromSec *float64 `json:"time_from_sec"`
		TimeToSec   *float64 `json:"time_to_sec"`
	}
	if err := json.Unmarshal([]byte(s), &tr); err != nil {
		return lo, hi
	}
	if tr.TimeFromSec != nil {
		lo = *tr.TimeFromSec
	}
	if tr.TimeToSec != nil {
		hi = *tr.TimeToSec
	}
	if math.IsNaN(lo) || math.IsInf(lo, 0) {
		lo = 0
	}
	if math.IsNaN(hi) || math.IsInf(hi, 0) {
		hi = span
	}
	if lo < 0 {
		lo = 0
	}
	if hi > span {
		hi = span
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	return lo, hi
}

func ApplyTimeBoundsToReq(level *splitter.Level, from, to *float64) (lo, hi float64) {
	span := LevelSpanSeconds(level)
	lo, hi = 0, span
	if from != nil {
		lo = *from
	}
	if to != nil {
		hi = *to
	}
	if math.IsNaN(lo) || math.IsInf(lo, 0) {
		lo = 0
	}
	if math.IsNaN(hi) || math.IsInf(hi, 0) {
		hi = span
	}
	if lo < 0 {
		lo = 0
	}
	if hi > span {
		hi = span
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	return lo, hi
}
