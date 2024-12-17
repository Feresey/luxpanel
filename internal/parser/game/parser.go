package game

import (
	"fmt"
	"strings"
	"time"
)

type GrammaError struct {
	Err    error
	tokens []token
}

func (e *GrammaError) Error() string {
	var sb strings.Builder

	sb.WriteString("grammar error: ")
	sb.WriteString(e.Err.Error())
	sb.WriteString(". tokens:")
	for _, tok := range e.tokens {
		sb.WriteByte(' ')
		sb.WriteString(showTok(tok))
	}

	return sb.String()
}

func (e *GrammaError) Unwrap() error { return e.Err }

func showTok(t token) string {
	var sb strings.Builder
	if tok := t.Token(); tok > yyPrivate {
		sb.WriteString(yyTokname(tok - yyPrivate + yyErrCode))
	} else {
		sb.WriteByte(byte(tok))
	}
	if _, ok := t.(voidTok); !ok {
		sb.WriteByte('(')
		sb.WriteString(t.String())
		sb.WriteByte(')')
	}
	return sb.String()
}

type lexer struct {
	line LogLine

	pos    int
	tokens []token
	err    error
}

func (l *lexer) reset(toks []token) {
	l.pos = 0
	l.tokens = toks
	l.err = nil
	l.line = nil
}

func (l *lexer) Lex(out *yySymType) int {
	if l.pos >= len(l.tokens) {
		return 0
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
	return &Parser{}
}

func (p *Parser) Parse(logTime time.Time, line string) (LogLine, error) {
	toks, err := p.t.Parse(logTime, line)
	if err != nil {
		return nil, fmt.Errorf("tokenizer.Parse: %w", err)
	}
	if len(toks) == 0 {
		return nil, nil
	}

	p.l.reset(toks)
	p.p = yyNewParser()

	staus := p.p.Parse(&p.l)
	if staus != 0 {
		return nil, &GrammaError{
			Err:    p.l.err,
			tokens: toks,
		}
	}

	return p.l.line, nil
}
