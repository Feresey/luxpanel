// Code generated by goyacc -v internal/parser/game/ext/y.output -l -p Yacc -o internal/parser/game/lexer.gen.go internal/parser/game/ext/lexer.y. DO NOT EDIT.
package game

import __yyfmt__ "fmt"

import "fmt"

type (
	Int    = int
	String = string
	Float  = float32
	Bool   = bool
)

type Lexer struct {
	line LogLine

	pos    int
	tokens []Token
	err    error
}

func (l *Lexer) Init(toks []Token) {
	l.pos = 0
	l.tokens = toks
	l.err = nil
	l.line = nil
}

func (l *Lexer) Lex(out *YaccSymType) int {
	if l.pos >= len(l.tokens) {
		return 0
	}
	t := l.tokens[l.pos]
	l.pos++
	t.set(out)
	return t.token()
}

func (l *Lexer) Error(errMsg string) {
	l.err = fmt.Errorf("lexer error: %s", errMsg)
}

func (l *Lexer) Res() (LogLine, error) {
	return l.line, l.err
}

func (y *YaccParserImpl) New() *YaccParserImpl {
	return &YaccParserImpl{}
}

type YaccSymType struct {
	yys int
	// Common
	String
	Int
	Float
	Bool
	// Game log
	*ClientAddPlayer
	*ClientPlayerLeave
	*ClientConnected
	*ClientConnectionClosed
	LogLine
}

const GAME = 57346
const INT = 57347
const STRING = 57348
const TIME = 57349
const CLIENT_ADD_PLAYER = 57350
const CLIENT_PLAYER_LEAVE = 57351
const LEAVE = 57352
const CLIENT_CONNECTED = 57353
const CLIENT_CONNECTION_CLOSED = 57354
const EOL = 57355

var YaccToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"GAME",
	"INT",
	"STRING",
	"TIME",
	"CLIENT_ADD_PLAYER",
	"CLIENT_PLAYER_LEAVE",
	"LEAVE",
	"CLIENT_CONNECTED",
	"CLIENT_CONNECTION_CLOSED",
	"EOL",
}

var YaccStatenames = [...]string{}

const YaccEofCode = 1
const YaccErrCode = 2
const YaccInitialStackSize = 16

var YaccExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const YaccPrivate = 57344

const YaccLast = 27

var YaccAct = [...]int8{
	9, 10, 13, 11, 12, 19, 2, 22, 18, 17,
	16, 27, 25, 24, 23, 20, 15, 14, 3, 1,
	4, 8, 7, 6, 21, 5, 26,
}

var YaccPact = [...]int16{
	-1, -1000, 14, -8, -11, -1000, -1000, -1000, -1000, 12,
	11, 4, 3, -1000, 2, -5, 10, -1000, 1, -1000,
	-1000, 9, -1000, 8, 7, 6, -1000, -1000,
}

var YaccPgo = [...]int8{
	0, 26, 25, 24, 23, 22, 21, 20, 19,
}

var YaccR1 = [...]int8{
	0, 8, 7, 7, 7, 7, 1, 1, 3, 3,
	2, 4, 5, 6,
}

var YaccR2 = [...]int8{
	0, 4, 1, 1, 1, 1, 1, 0, 1, 0,
	8, 3, 3, 2,
}

var YaccChk = [...]int16{
	-1000, -8, 7, 4, -7, -2, -4, -5, -6, 8,
	9, 11, 12, 13, 5, 5, 6, 6, 6, 10,
	5, -3, 6, 5, 5, 5, -1, 5,
}

var YaccDef = [...]int8{
	0, -2, 0, 0, 0, 2, 3, 4, 5, 0,
	0, 0, 0, 1, 0, 0, 0, 13, 9, 11,
	12, 0, 8, 0, 0, 7, 10, 6,
}

var YaccTok1 = [...]int8{
	1,
}

var YaccTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13,
}

var YaccTok3 = [...]int8{
	0,
}

var YaccErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	YaccDebug        = 0
	YaccErrorVerbose = false
)

type YaccLexer interface {
	Lex(lval *YaccSymType) int
	Error(s string)
}

type YaccParser interface {
	Parse(YaccLexer) int
	Lookahead() int
}

type YaccParserImpl struct {
	lval  YaccSymType
	stack [YaccInitialStackSize]YaccSymType
	char  int
}

func (p *YaccParserImpl) Lookahead() int {
	return p.char
}

func YaccNewParser() YaccParser {
	return &YaccParserImpl{}
}

const YaccFlag = -1000

func YaccTokname(c int) string {
	if c >= 1 && c-1 < len(YaccToknames) {
		if YaccToknames[c-1] != "" {
			return YaccToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func YaccStatname(s int) string {
	if s >= 0 && s < len(YaccStatenames) {
		if YaccStatenames[s] != "" {
			return YaccStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func YaccErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !YaccErrorVerbose {
		return "syntax error"
	}

	for _, e := range YaccErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + YaccTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(YaccPact[state])
	for tok := TOKSTART; tok-1 < len(YaccToknames); tok++ {
		if n := base + tok; n >= 0 && n < YaccLast && int(YaccChk[int(YaccAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if YaccDef[state] == -2 {
		i := 0
		for YaccExca[i] != -1 || int(YaccExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; YaccExca[i] >= 0; i += 2 {
			tok := int(YaccExca[i])
			if tok < TOKSTART || YaccExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if YaccExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += YaccTokname(tok)
	}
	return res
}

func Yacclex1(lex YaccLexer, lval *YaccSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(YaccTok1[0])
		goto out
	}
	if char < len(YaccTok1) {
		token = int(YaccTok1[char])
		goto out
	}
	if char >= YaccPrivate {
		if char < YaccPrivate+len(YaccTok2) {
			token = int(YaccTok2[char-YaccPrivate])
			goto out
		}
	}
	for i := 0; i < len(YaccTok3); i += 2 {
		token = int(YaccTok3[i+0])
		if token == char {
			token = int(YaccTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(YaccTok2[1]) /* unknown char */
	}
	if YaccDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", YaccTokname(token), uint(char))
	}
	return char, token
}

func YaccParse(Yacclex YaccLexer) int {
	return YaccNewParser().Parse(Yacclex)
}

func (Yaccrcvr *YaccParserImpl) Parse(Yacclex YaccLexer) int {
	var Yaccn int
	var YaccVAL YaccSymType
	var YaccDollar []YaccSymType
	_ = YaccDollar // silence set and not used
	YaccS := Yaccrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Yaccstate := 0
	Yaccrcvr.char = -1
	Yacctoken := -1 // Yaccrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Yaccstate = -1
		Yaccrcvr.char = -1
		Yacctoken = -1
	}()
	Yaccp := -1
	goto Yaccstack

ret0:
	return 0

ret1:
	return 1

Yaccstack:
	/* put a state and value onto the stack */
	if YaccDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", YaccTokname(Yacctoken), YaccStatname(Yaccstate))
	}

	Yaccp++
	if Yaccp >= len(YaccS) {
		nyys := make([]YaccSymType, len(YaccS)*2)
		copy(nyys, YaccS)
		YaccS = nyys
	}
	YaccS[Yaccp] = YaccVAL
	YaccS[Yaccp].yys = Yaccstate

Yaccnewstate:
	Yaccn = int(YaccPact[Yaccstate])
	if Yaccn <= YaccFlag {
		goto Yaccdefault /* simple state */
	}
	if Yaccrcvr.char < 0 {
		Yaccrcvr.char, Yacctoken = Yacclex1(Yacclex, &Yaccrcvr.lval)
	}
	Yaccn += Yacctoken
	if Yaccn < 0 || Yaccn >= YaccLast {
		goto Yaccdefault
	}
	Yaccn = int(YaccAct[Yaccn])
	if int(YaccChk[Yaccn]) == Yacctoken { /* valid shift */
		Yaccrcvr.char = -1
		Yacctoken = -1
		YaccVAL = Yaccrcvr.lval
		Yaccstate = Yaccn
		if Errflag > 0 {
			Errflag--
		}
		goto Yaccstack
	}

Yaccdefault:
	/* default state action */
	Yaccn = int(YaccDef[Yaccstate])
	if Yaccn == -2 {
		if Yaccrcvr.char < 0 {
			Yaccrcvr.char, Yacctoken = Yacclex1(Yacclex, &Yaccrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if YaccExca[xi+0] == -1 && int(YaccExca[xi+1]) == Yaccstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Yaccn = int(YaccExca[xi+0])
			if Yaccn < 0 || Yaccn == Yacctoken {
				break
			}
		}
		Yaccn = int(YaccExca[xi+1])
		if Yaccn < 0 {
			goto ret0
		}
	}
	if Yaccn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Yacclex.Error(YaccErrorMessage(Yaccstate, Yacctoken))
			Nerrs++
			if YaccDebug >= 1 {
				__yyfmt__.Printf("%s", YaccStatname(Yaccstate))
				__yyfmt__.Printf(" saw %s\n", YaccTokname(Yacctoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Yaccp >= 0 {
				Yaccn = int(YaccPact[YaccS[Yaccp].yys]) + YaccErrCode
				if Yaccn >= 0 && Yaccn < YaccLast {
					Yaccstate = int(YaccAct[Yaccn]) /* simulate a shift of "error" */
					if int(YaccChk[Yaccstate]) == YaccErrCode {
						goto Yaccstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if YaccDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", YaccS[Yaccp].yys)
				}
				Yaccp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if YaccDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", YaccTokname(Yacctoken))
			}
			if Yacctoken == YaccEofCode {
				goto ret1
			}
			Yaccrcvr.char = -1
			Yacctoken = -1
			goto Yaccnewstate /* try again in the same state */
		}
	}

	/* reduction by production Yaccn */
	if YaccDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Yaccn, YaccStatname(Yaccstate))
	}

	Yaccnt := Yaccn
	Yaccpt := Yaccp
	_ = Yaccpt // guard against "declared and not used"

	Yaccp -= int(YaccR2[Yaccn])
	// Yaccp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Yaccp+1 >= len(YaccS) {
		nyys := make([]YaccSymType, len(YaccS)*2)
		copy(nyys, YaccS)
		YaccS = nyys
	}
	YaccVAL = YaccS[Yaccp+1]

	/* consult goto table to find next state */
	Yaccn = int(YaccR1[Yaccn])
	Yaccg := int(YaccPgo[Yaccn])
	Yaccj := Yaccg + YaccS[Yaccp].yys + 1

	if Yaccj >= YaccLast {
		Yaccstate = int(YaccAct[Yaccg])
	} else {
		Yaccstate = int(YaccAct[Yaccj])
		if int(YaccChk[Yaccstate]) != -Yaccn {
			Yaccstate = int(YaccAct[Yaccg])
		}
	}
	// dummy call; replaced with literal code
	switch Yaccnt {

	case 1:
		YaccDollar = YaccS[Yaccpt-4 : Yaccpt+1]
		{
			YaccDollar[3].LogLine.setTime(YaccDollar[1].String)
			Yacclex.(*Lexer).line = YaccDollar[3].LogLine
		}
	case 2:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.LogLine = YaccDollar[1].ClientAddPlayer
		}
	case 3:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.LogLine = YaccDollar[1].ClientPlayerLeave
		}
	case 4:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.LogLine = YaccDollar[1].ClientConnected
		}
	case 5:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.LogLine = YaccDollar[1].ClientConnectionClosed
		}
	case 6:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.Int = YaccDollar[1].Int
		}
	case 7:
		YaccDollar = YaccS[Yaccpt-0 : Yaccpt+1]
		{
		}
	case 8:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.String = YaccDollar[1].String
		}
	case 9:
		YaccDollar = YaccS[Yaccpt-0 : Yaccpt+1]
		{
		}
	case 10:
		YaccDollar = YaccS[Yaccpt-8 : Yaccpt+1]
		{
			YaccVAL.ClientAddPlayer = &ClientAddPlayer{
				InGamePlayerID: YaccDollar[2].Int,
				Name:           YaccDollar[3].String,
				ClanTag:        YaccDollar[4].String,
				PlayerID:       YaccDollar[5].Int,
				Status:         YaccDollar[6].Int,
				TeamID:         YaccDollar[7].Int,
				GroupID:        YaccDollar[8].Int,
			}
		}
	case 11:
		YaccDollar = YaccS[Yaccpt-3 : Yaccpt+1]
		{
			YaccVAL.ClientPlayerLeave = &ClientPlayerLeave{
				InGamePlayerID: YaccDollar[2].Int,
			}
		}
	case 12:
		YaccDollar = YaccS[Yaccpt-3 : Yaccpt+1]
		{
			YaccVAL.ClientConnected = &ClientConnected{
				ServerAddr: YaccDollar[2].String,
				MTU:        YaccDollar[3].Int,
			}
		}
	case 13:
		YaccDollar = YaccS[Yaccpt-2 : Yaccpt+1]
		{
			YaccVAL.ClientConnectionClosed = &ClientConnectionClosed{
				Reason: YaccDollar[2].String,
			}
		}
	}
	goto Yaccstack /* stack new state and value */
}
