package game

import (
	"strconv"
	"strings"
)

const debugTokenizer = false

func init() {
	YaccErrorVerbose = true
	YaccDebug = 0
}

type Token interface {
	set(*YaccSymType)
	token() int
	show() string
	String() string
}

type strTok string

func (s strTok) set(out *YaccSymType) { out.String = string(s) }
func (strTok) token() int           { return STRING }
func (s strTok) show() string       { return string(s) }
func (s strTok) String() string     { return showTok(s) }

type intTok int

func (s intTok) set(out *YaccSymType) { out.Int = Int(s) }
func (intTok) token() int           { return INT }
func (s intTok) show() string       { return strconv.Itoa(int(s)) }
func (s intTok) String() string     { return showTok(s) }

type voidTok int

func (a voidTok) token() int     { return int(a) }
func (voidTok) set(*YaccSymType)   {}
func (a voidTok) show() string   { return "" }
func (s voidTok) String() string { return showTok(s) }

type anyVal struct {
	tok   int
	value Token
}

func newAnyVal(token int, value Token) anyVal {
	return anyVal{
		tok:   token,
		value: value,
	}
}

func (t anyVal) token() int         { return t.tok }
func (t anyVal) set(out *YaccSymType) { t.value.set(out) }
func (t anyVal) show() string       { return t.value.show() }
func (t anyVal) String() string     { return showTok(t) }

func showTok(t Token) string {
	var sb strings.Builder
	if tok := t.token(); tok > YaccPrivate {
		sb.WriteString(YaccTokname(tok - YaccPrivate + YaccErrCode))
	} else {
		sb.WriteString(strconv.QuoteRune(rune(tok)))
	}
	if _, ok := t.(voidTok); !ok {
		sb.WriteByte('(')
		sb.WriteString(t.show())
		sb.WriteByte(')')
	}
	return sb.String()
}
