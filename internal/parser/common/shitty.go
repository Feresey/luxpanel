package common

import (
	"fmt"
)

type tokenizer[T any] interface {
	Parse(string) ([]T, error)
}

type lexer[T, R, Y any] interface {
	Init([]T)
	Res() (R, error)
	yaccLexer[Y]
}

type yaccLexer[T any] interface {
	Lex(lval *T) int
	Error(s string)
}

type grammar[T any] interface {
	Parse(T) int
}

type Parser[
	T fmt.Stringer,
	R, Y any,
	Yacc yaccLexer[Y],
	G grammar[Yacc],
] struct {
	T         tokenizer[T]
	L         lexer[T, R, Y]
	NewGramma func() G
}

func (p *Parser[T, R, Y, Yacc, G]) Parse(line string) (res R, err error) {
	toks, err := p.T.Parse(line)
	if err != nil {
		return res, fmt.Errorf("tokenizer.Parse: %w", err)
	}
	if len(toks) == 0 {
		return res, nil
	}

	p.L.Init(toks)

	parser := p.NewGramma()
	status := parser.Parse(p.L.(Yacc))
	res, err = p.L.Res()
	if status != 0 {
		return res, &GrammaError[T]{
			Err:    err,
			Tokens: toks,
		}
	}

	return res, nil
}
