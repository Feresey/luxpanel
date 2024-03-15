package parser

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

var (
	ErrUndefinedLineType = errors.New("undefined line type")
	ErrWrongLineFormat   = errors.New("wrong format")
)

const space = `\s+`

const timeFormat = "15:04:05.000"

func parseLogTime(nowTime time.Time) func(string) (time.Time, error) {
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

func parseFloat(s string) (float64, error) {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("strconv.ParseFloat: %w", err)
	}
	return res, nil
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

func ParseField[T any](raw string, fieldName string, convert func(raw string) (res T, err error)) (res T, err error) {
	if res, err = convert(raw); err != nil {
		return res, ParseFieldError{
			Raw:       raw,
			FieldName: fieldName,
			Err:       err,
		}
	}
	return res, nil
}
