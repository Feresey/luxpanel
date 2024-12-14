package gramma

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const debugTokenizer = true

func init() {
	yyDebug = 4
	yyErrorVerbose = true
}

const timeFormat = "15:04:05.000"

func ParseTime(nowTime time.Time, raw string) (time.Time, error) {
	res, err := time.Parse(timeFormat, raw)
	if err != nil {
		return res, fmt.Errorf("time.Parse(%s): %w", timeFormat, err)
	}
	res = time.Date(
		nowTime.Year(), nowTime.Month(), nowTime.Day(),
		res.Hour(), res.Minute(), res.Second(), res.Nanosecond(),
		nowTime.Location(),
	)
	if res.Before(nowTime) {
		return res.AddDate(0, 0, 1), nil
	}
	return res, nil
}

func parseDamageModifiers(raw string) (res DamageModifiersMap, err error) {
	parts := strings.Split(raw, "|")
	if len(parts) == 0 {
		return res, fmt.Errorf("wrong parts number")
	}
	res = make(DamageModifiersMap, len(parts))
	for _, p := range parts {
		res[DamageModifier(p)] = struct{}{}
	}
	return res, nil
}

type ParseTokenError struct {
	TokType int
	Err     error
	Raw     string
}

func (e ParseTokenError) Error() string {
	return fmt.Sprintf("parse field (raw: %s, tok: %s, err: %v)", e.Raw, getTokStrName(e.TokType), e.Err)
}

type LexerError struct {
	LineIsNotFinished bool
	
}

var ErrLineIsNotFinished = errors.New("line is not finished")

type Token interface {
	Set(*yySymType)
	Token() int
	String() string
}

type StrTok string

func (s StrTok) Set(out *yySymType) { out.String = string(s) }
func (StrTok) Token() int           { return STRING }
func (s StrTok) String() string     { return string(s) }

type IntTok int

func (s IntTok) Set(out *yySymType) { out.Int = Int(s) }
func (IntTok) Token() int           { return INT }
func (s IntTok) String() string     { return strconv.Itoa(int(s)) }

type FloatTok float32

func (s FloatTok) Set(out *yySymType) { out.Float = Float(s) }
func (FloatTok) Token() int           { return FLOAT }
func (s FloatTok) String() string     { return strconv.FormatFloat(float64(s), 'f', -1, 32) }

type TimeTok Time

func (s TimeTok) Set(out *yySymType) { out.Time = Time(s) }
func (TimeTok) Token() int           { return TIME }
func (s TimeTok) String() string     { return time.Time(s).String() }

type VoidTok int

func (a VoidTok) Token() int   { return int(a) }
func (VoidTok) Set(*yySymType) {}
func (VoidTok) String() string { return "" }

type AnyVal struct {
	token int
	value Token
}

func NewAnyVal(token int, value Token) AnyVal {
	return AnyVal{
		token: token,
		value: value,
	}
}

func (t AnyVal) Token() int         { return t.token }
func (t AnyVal) Set(out *yySymType) { t.value.Set(out) }
func (t AnyVal) String() string     { return t.value.String() }

func getTokStrName(tok int) string {
	if tok > yyPrivate {
		return yyTokname(tok - yyPrivate + yyErrCode)
	} else {
		return string(byte(tok))
	}
}

func ShowTok(t Token) string {
	var sb strings.Builder
	if tok := t.Token(); tok > yyPrivate {
		sb.WriteString(yyTokname(tok - yyPrivate + yyErrCode))
	} else {
		sb.WriteByte(byte(tok))
	}
	sb.WriteString(": ")
	sb.WriteString(t.String())
	return sb.String()
}

type lexer struct {
	res    LogLine
	pos    int
	tokens []Token
	err    error
}

func (l *lexer) reset(toks []Token) {
	l.res = LogLine{}
	l.pos = 0
	l.tokens = toks
	l.err = nil
}

func (l *lexer) Lex(out *yySymType) int {
	if l.pos >= len(l.tokens) {
		return yyEofCode
	}
	t := l.tokens[l.pos]
	l.pos++
	t.Set(out)
	return t.Token()
}

func (l *lexer) Error(errMsg string) {
	l.err = fmt.Errorf("lexer error: %s", errMsg)
}

type Parser struct {
	t tokenizer
	l lexer
	p yyParser
}

func NewParser() *Parser {
	return &Parser{
		p: yyNewParser(),
	}
}

func (p *Parser) Parse(logTime time.Time, line string) (*LogLine, error) {
	toks, err := p.t.Parse(logTime, line)
	if err != nil {
		return nil, fmt.Errorf("tokenizer.Parse: %w", err)
	}
	if len(toks) == 0 {
		return nil, nil
	}

	p.l.reset(toks)

	staus := p.p.Parse(&p.l)
	if staus != 0 {
		return nil, &GrammaError{
			Err:    p.l.err,
			Tokens: toks,
		}
	}

	return &p.l.res, nil
}

type GrammaError struct {
	Err    error
	Tokens []Token
}

func (e GrammaError) Error() string {
	return e.Err.Error()
}
