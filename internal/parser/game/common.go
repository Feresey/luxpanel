// DO NOT EDIT. This file was auto-generated

package game

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type GameLineType string

func (s GameLineType) String() string { return string(s) }

const (
	PlayerLeaveLineType      = "PlayerLeave"
	ConnectionClosedLineType = "ConnectionClosed"
	ConnectedLineType        = "Connected"
	AddPlayerLineType        = "AddPlayer"
)

type LogLine interface {
	Type() GameLineType
	Time() time.Time
	Unmarshal(raw string, now time.Time) error
}

var errWrongLineFormat = errors.New("Game: wrong format")

func ParseLogLine(raw string, now time.Time) (line LogLine, matchedToRegexp bool, err error) {
	switch {
	case rePlayerLeave.MatchString(raw):
		line = &PlayerLeave{}
	case reConnectionClosed.MatchString(raw):
		line = &ConnectionClosed{}
	case reConnected.MatchString(raw):
		line = &Connected{}
	case reAddPlayer.MatchString(raw):
		line = &AddPlayer{}
	default:
		return nil, false, nil
	}

	if err := line.Unmarshal(raw, now); err != nil {
		return nil, true, fmt.Errorf("line matched to regex(%s), but failed to parse: %s: %s", line.Type().String(), raw, err.Error())
	}

	return line, true, nil
}

const timeFormat = "15:04:05.000"

func parseTime(nowTime time.Time) func(string) (time.Time, error) {
	return func(s string) (time.Time, error) {
		t, err := time.Parse(timeFormat, s)
		if err != nil {
			return t, fmt.Errorf("time.Parse(%s): %w", timeFormat, err)
		}
		res := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
		if res.Before(nowTime) {
			return res.AddDate(0, 0, 1), nil
		}
		return res, nil
	}
}

func parseSeconds(s string) (time.Duration, error) {
	return time.ParseDuration(s + "s")
}

func parseBool(s string) (bool, error) {
	return s != "", nil
}

func parseOptional[T any](parseFunc func(string) (T, error)) func(string) (*T, error) {
	return func(s string) (res *T, err error) {
		if s == "" {
			return nil, nil
		}
		r, err := parseFunc(s)
		if err != nil {
			return nil, err
		}
		return &r, nil
	}
}

func parseFloat(s string) (float32, error) {
	res, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, fmt.Errorf("strconv.ParseFloat: %w", err)
	}
	return float32(res), nil
}

type ParseFieldError struct {
	Raw       string
	FieldName string
	Err       error
}

func (e ParseFieldError) Unwrap() error {
	return e.Err
}

func (p ParseFieldError) Error() string {
	return fmt.Sprintf("parse %s: (%s) %v", p.FieldName, p.Raw, p.Err)
}

func parseField[T any](raw string, fieldName string, convert func(raw string) (res T, err error)) (res T, err error) {
	if res, err = convert(raw); err != nil {
		return res, ParseFieldError{
			Raw:       raw,
			FieldName: fieldName,
			Err:       err,
		}
	}
	return res, nil
}

type ConnectionClosedReason string

const (
	ConnectionClosedReasonGameFinished          ConnectionClosedReason = "DR_CLIENT_GAME_FINISHED"
	ConnectionClosedReasonClientCouldNotConnect ConnectionClosedReason = "DR_CLIENT_COULD_NOT_CONNECT"
	ConnectionClosedReasonQuit                  ConnectionClosedReason = "DR_CLIENT_QUIT"
)

func (c ConnectionClosedReason) Validate() error {
	switch c {
	case ConnectionClosedReasonGameFinished:
	case ConnectionClosedReasonClientCouldNotConnect:
	case ConnectionClosedReasonQuit:
	default:
		return fmt.Errorf("undefined connection closed reason: %q", c)
	}
	return nil
}

func parseConnectionClosedReason(s string) (res ConnectionClosedReason, err error) {
	res = ConnectionClosedReason(s)
	if err := res.Validate(); err != nil {
		return res, err
	}
	return res, nil
}
