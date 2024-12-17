package game

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const debugTokenizer = false

func init() {
	yyErrorVerbose = true
	yyDebug = 0
}

var ErrLineIsNotFinished = errors.New("line is not finished")

const timeFormat = "15:04:05.000"

func parseTime(nowTime time.Time, raw string) (Time, error) {
	res, err := time.Parse(timeFormat, raw)
	if err != nil {
		return Time(res), fmt.Errorf("time.Parse(%s): %w", timeFormat, err)
	}
	res = time.Date(
		nowTime.Year(), nowTime.Month(), nowTime.Day(),
		res.Hour(), res.Minute(), res.Second(), res.Nanosecond(),
		nowTime.Location(),
	)
	if res.Before(nowTime) {
		return Time(res.AddDate(0, 0, 1)), nil
	}
	return Time(res), nil
}


type ParseTokenError struct {
	TokType string
	Err     error
	Raw     string
}

func (e ParseTokenError) Error() string {
	return fmt.Sprintf("parse field (raw: %s, tok: %s, err: %v)", e.Raw, e.TokType, e.Err)
}

type token interface {
	Set(*yySymType)
	Token() int
	String() string
}

type strTok string

func (s strTok) Set(out *yySymType) { out.String = string(s) }
func (strTok) Token() int           { return STRING }
func (s strTok) String() string     { return string(s) }

type intTok int

func (s intTok) Set(out *yySymType) { out.Int = Int(s) }
func (intTok) Token() int           { return INT }
func (s intTok) String() string     { return strconv.Itoa(int(s)) }

type timeTok Time

func (s timeTok) Set(out *yySymType) { out.Time = Time(s) }
func (timeTok) Token() int           { return TIME }
func (s timeTok) String() string     { return time.Time(s).String() }

type voidTok int

func (a voidTok) Token() int   { return int(a) }
func (voidTok) Set(*yySymType) {}
func (voidTok) String() string { return "" }

type anyVal struct {
	token int
	value token
}

func newAnyVal(token int, value token) anyVal {
	return anyVal{
		token: token,
		value: value,
	}
}

func (t anyVal) Token() int         { return t.token }
func (t anyVal) Set(out *yySymType) { t.value.Set(out) }
func (t anyVal) String() string     { return t.value.String() }
