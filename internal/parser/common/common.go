package common

import (
	"fmt"
	"strings"
	"time"
)

const timeFormat = "15:04:05.000"

func ParseTime(logTime time.Time, raw string) time.Time {
	res, err := time.Parse(timeFormat, raw)
	if err != nil {
		return res
	}
	res = time.Date(
		logTime.Year(), logTime.Month(), logTime.Day(),
		res.Hour(), res.Minute(), res.Second(), res.Nanosecond(),
		logTime.Location(),
	)
	if res.Before(logTime) {
		return res.AddDate(0, 0, 1)
	}
	return res
}

type ParseTokenError struct {
	TokType string
	Err     error
	Raw     string
}

func (e ParseTokenError) Error() string {
	return fmt.Sprintf("parse field (raw: %s, tok: %s, err: %v)", e.Raw, e.TokType, e.Err)
}

type LineIsNotFinishedError struct {
	Parsed string
	Rest   string
}

func (l LineIsNotFinishedError) Error() string {
	return fmt.Sprintf("line is not finished: %q %q", l.Parsed, l.Rest)
}

type GrammaError[T fmt.Stringer] struct {
	Err    error
	Tokens []T
}

func (e *GrammaError[T]) Error() string {
	var sb strings.Builder

	sb.WriteString("grammar error: ")
	sb.WriteString(e.Err.Error())
	sb.WriteString(". tokens:")
	for _, tok := range e.Tokens {
		sb.WriteByte(' ')
		sb.WriteString(tok.String())
	}

	return sb.String()
}

func (e *GrammaError[T]) Unwrap() error { return e.Err }
