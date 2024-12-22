// Code generated by goyacc -v internal/parser/combat/ext/y.output -l -p Yacc -o internal/parser/combat/lexer.gen.go internal/parser/combat/ext/lexer.y. DO NOT EDIT.
package combat

import __yyfmt__ "fmt"

import (
	"fmt"
)

type (
	Int     = int
	String  = string
	Strings = []string
	Float   = float32
	Bool    = bool
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
	Strings
	Int
	Float
	Bool
	// Combat log
	*Damage
	DamageModifiers
	*Object
	*Heal
	*Kill
	*Participant
	ParticipationModifiers
	*ConnectToGameSession
	*Start
	*Finished
	*Reward
	LogLine
}

const INT = 57346
const STRING = 57347
const TIME = 57348
const COMBAT = 57349
const ARROW = 57350
const FLOAT = 57351
const SOURCE = 57352
const FRIENDLY_FIRE = 57353
const DAMAGE = 57354
const DAMAGE_MODIFIER = 57355
const ROCKET = 57356
const HEAL = 57357
const KILL = 57358
const PARTICIPANT = 57359
const PARTICIPATION_MODIFIER = 57360
const PARTICIPATION_MODIFIERS_END = 57361
const CONNECT_TO_GAME_SESSION_PREFIX = 57362
const START = 57363
const GAMEPLAY_FINISHED = 57364
const REWARD = 57365
const EOL = 57366

var YaccToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"INT",
	"STRING",
	"TIME",
	"COMBAT",
	"ARROW",
	"FLOAT",
	"SOURCE",
	"FRIENDLY_FIRE",
	"DAMAGE",
	"DAMAGE_MODIFIER",
	"ROCKET",
	"HEAL",
	"KILL",
	"PARTICIPANT",
	"PARTICIPATION_MODIFIER",
	"PARTICIPATION_MODIFIERS_END",
	"CONNECT_TO_GAME_SESSION_PREFIX",
	"START",
	"GAMEPLAY_FINISHED",
	"REWARD",
	"EOL",
	"'|'",
	"'('",
	"')'",
	"';'",
	"'\\t'",
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

const YaccLast = 90

var YaccAct = [...]int8{
	70, 59, 22, 34, 35, 13, 57, 38, 14, 15,
	16, 37, 71, 20, 17, 18, 19, 24, 27, 34,
	35, 68, 21, 71, 84, 76, 82, 85, 80, 71,
	72, 73, 60, 69, 62, 67, 45, 65, 58, 48,
	49, 50, 56, 52, 36, 33, 3, 2, 75, 66,
	54, 53, 23, 47, 44, 42, 41, 40, 31, 29,
	28, 26, 74, 86, 77, 64, 55, 46, 32, 30,
	1, 4, 43, 11, 10, 9, 63, 12, 78, 61,
	81, 51, 8, 39, 7, 6, 83, 79, 5, 25,
}

var YaccPact = [...]int16{
	41, -1000, 39, -7, -2, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 47, 47, 56, 55, 54, 65, 53,
	64, -1000, 37, -6, 36, -17, -22, -1000, 52, 51,
	50, 49, -1000, 47, 63, 48, 47, 47, 47, 34,
	-1000, 46, 45, 62, -1000, 33, -1000, -21, 29, 22,
	-1000, -1000, 24, 61, 28, 44, 26, -4, 23, 18,
	-1000, 12, -1000, -1000, -1000, -1000, 43, 16, 60, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 22, -1000, 15, 1,
	-1000, 10, 14, -1000, 59, -1000, -1000,
}

var YaccPgo = [...]int8{
	0, 2, 89, 1, 88, 87, 0, 86, 85, 84,
	83, 82, 81, 79, 77, 76, 75, 74, 73, 72,
	71, 70,
}

var YaccR1 = [...]int8{
	0, 21, 20, 20, 20, 20, 20, 20, 20, 20,
	1, 1, 6, 6, 7, 7, 3, 3, 4, 5,
	5, 8, 9, 2, 2, 13, 13, 13, 10, 10,
	11, 12, 12, 15, 15, 16, 14, 17, 19, 19,
	18,
}

var YaccR2 = [...]int8{
	0, 4, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 6, 1, 0, 2, 0, 1, 0, 11, 1,
	3, 6, 6, 3, 1, 2, 2, 0, 1, 0,
	6, 2, 0, 1, 0, 5, 2, 5, 1, 0,
	6,
}

var YaccChk = [...]int16{
	-1000, -21, 6, 7, -20, -4, -8, -9, -11, -16,
	-17, -18, -14, 12, 15, 16, 17, 21, 22, 23,
	20, 24, -1, 5, -1, -2, 5, -1, 5, 5,
	4, 5, 4, 8, 25, 26, 8, 28, 29, -10,
	5, 5, 5, -19, 5, -1, 4, 5, -1, -1,
	-1, -12, 9, 5, 5, 4, 9, 27, 9, -3,
	10, -13, 10, -15, 4, 9, 5, 9, 25, 10,
	-6, 11, 18, 19, -6, 5, 9, 4, -3, -5,
	13, -6, 25, -7, 14, 13, 4,
}

var YaccDef = [...]int8{
	0, -2, 0, 0, 0, 2, 3, 4, 5, 6,
	7, 8, 9, 0, 0, 0, 0, 0, 0, 0,
	0, 1, 0, 0, 0, 0, 0, 24, 29, 0,
	0, 39, 36, 0, 0, 0, 0, 0, 0, 32,
	28, 0, 0, 0, 38, 0, 10, 0, 0, 17,
	23, 27, 0, 34, 0, 0, 0, 0, 0, 13,
	16, 13, 31, 35, 33, 37, 0, 0, 0, 21,
	22, 12, 25, 26, 30, 40, 17, 11, 0, 13,
	19, 15, 0, 18, 0, 20, 14,
}

var YaccTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 29,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	26, 27, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 28,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 25,
}

var YaccTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24,
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
			YaccVAL.LogLine = YaccDollar[1].Damage
		}
	case 3:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.LogLine = YaccDollar[1].Heal
		}
	case 4:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.LogLine = YaccDollar[1].Kill
		}
	case 5:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.LogLine = YaccDollar[1].Participant
		}
	case 6:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.LogLine = YaccDollar[1].Start
		}
	case 7:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.LogLine = YaccDollar[1].Finished
		}
	case 8:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.LogLine = YaccDollar[1].Reward
		}
	case 9:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.LogLine = YaccDollar[1].ConnectToGameSession
		}
	case 10:
		YaccDollar = YaccS[Yaccpt-3 : Yaccpt+1]
		{
			YaccVAL.Object = &Object{
				Name:     YaccDollar[1].String,
				ObjectID: YaccDollar[3].Int,
			}
		}
	case 11:
		YaccDollar = YaccS[Yaccpt-6 : Yaccpt+1]
		{
			YaccVAL.Object = &Object{
				PlayerObject: PlayerObject{
					ObjectName:  YaccDollar[1].String,
					ObjectOwner: YaccDollar[3].String,
				},
				ObjectID: YaccDollar[6].Int,
			}
		}
	case 12:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.Bool = true
		}
	case 13:
		YaccDollar = YaccS[Yaccpt-0 : Yaccpt+1]
		{
			YaccVAL.Bool = false
		}
	case 14:
		YaccDollar = YaccS[Yaccpt-2 : Yaccpt+1]
		{
			YaccVAL.Int = YaccDollar[2].Int
		}
	case 15:
		YaccDollar = YaccS[Yaccpt-0 : Yaccpt+1]
		{
			YaccVAL.Int = 0
		}
	case 16:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.String = YaccDollar[1].String
		}
	case 17:
		YaccDollar = YaccS[Yaccpt-0 : Yaccpt+1]
		{
			YaccVAL.String = ""
		}
	case 18:
		YaccDollar = YaccS[Yaccpt-11 : Yaccpt+1]
		{
			YaccVAL.Damage = &Damage{
				Initiator:       *YaccDollar[2].Object,
				Recipient:       *YaccDollar[4].Object,
				DamageFull:      YaccDollar[5].Float,
				DamageHull:      YaccDollar[6].Float,
				DamageShield:    YaccDollar[7].Float,
				Source:          YaccDollar[8].String,
				DamageModifiers: YaccDollar[9].DamageModifiers,
				FriendlyFire:    YaccDollar[10].Bool,
				Rocket:          YaccDollar[11].Int,
			}
		}
	case 19:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.DamageModifiers = []DamageModifier{DamageModifier(YaccDollar[1].String)}
		}
	case 20:
		YaccDollar = YaccS[Yaccpt-3 : Yaccpt+1]
		{
			YaccVAL.DamageModifiers = append(YaccVAL.DamageModifiers, DamageModifier(YaccDollar[3].String))
		}
	case 21:
		YaccDollar = YaccS[Yaccpt-6 : Yaccpt+1]
		{
			YaccVAL.Heal = &Heal{
				Initiator: *YaccDollar[2].Object,
				Recipient: *YaccDollar[4].Object,
				Heal:      YaccDollar[5].Float,
				Source:    YaccDollar[6].String,
			}
		}
	case 22:
		YaccDollar = YaccS[Yaccpt-6 : Yaccpt+1]
		{
			YaccVAL.Kill = &Kill{
				Killed:       *YaccDollar[2].Object,
				Killer:       *YaccDollar[4].Object,
				Source:       YaccDollar[5].String,
				FriendlyFire: YaccDollar[6].Bool,
			}
		}
	case 23:
		YaccDollar = YaccS[Yaccpt-3 : Yaccpt+1]
		{
			YaccVAL.Object = &Object{
				Name: YaccDollar[1].String,
				PlayerObject: PlayerObject{
					ObjectName: YaccDollar[3].Object.Name,
				},
				ObjectID: YaccDollar[3].Object.ObjectID,
			}
		}
	case 24:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.Object = YaccDollar[1].Object
		}
	case 25:
		YaccDollar = YaccS[Yaccpt-2 : Yaccpt+1]
		{
			YaccVAL.ParticipationModifiers = append(YaccDollar[1].ParticipationModifiers, ParticipationModifier(YaccDollar[2].String))
		}
	case 26:
		YaccDollar = YaccS[Yaccpt-2 : Yaccpt+1]
		{
			YaccVAL.ParticipationModifiers = YaccDollar[1].ParticipationModifiers
		}
	case 27:
		YaccDollar = YaccS[Yaccpt-0 : Yaccpt+1]
		{
			YaccVAL.ParticipationModifiers = nil
		}
	case 28:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.String = YaccDollar[1].String
		}
	case 29:
		YaccDollar = YaccS[Yaccpt-0 : Yaccpt+1]
		{
			YaccVAL.String = ""
		}
	case 30:
		YaccDollar = YaccS[Yaccpt-6 : Yaccpt+1]
		{
			YaccVAL.Participant = &Participant{
				Name:           YaccDollar[2].String,
				Ship:           YaccDollar[3].String,
				Damage:         YaccDollar[4].Participant.Damage,
				MostDamageWith: YaccDollar[4].Participant.MostDamageWith,
				Modifiers:      YaccDollar[5].ParticipationModifiers,
				FriendlyFire:   YaccDollar[6].Bool,
			}
		}
	case 31:
		YaccDollar = YaccS[Yaccpt-2 : Yaccpt+1]
		{
			YaccVAL.Participant = &Participant{
				Damage:         YaccDollar[1].Float,
				MostDamageWith: YaccDollar[2].String,
			}
		}
	case 32:
		YaccDollar = YaccS[Yaccpt-0 : Yaccpt+1]
		{
			YaccVAL.Participant = &Participant{
				Damage:         0,
				MostDamageWith: "",
			}
		}
	case 33:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.Int = YaccDollar[1].Int
		}
	case 34:
		YaccDollar = YaccS[Yaccpt-0 : Yaccpt+1]
		{
			YaccVAL.Int = 0
		}
	case 35:
		YaccDollar = YaccS[Yaccpt-5 : Yaccpt+1]
		{
			YaccVAL.Start = &Start{
				What:              YaccDollar[2].String,
				GameMode:          YaccDollar[3].String,
				MapName:           YaccDollar[4].String,
				LocalClientTeamID: YaccDollar[5].Int,
			}
		}
	case 36:
		YaccDollar = YaccS[Yaccpt-2 : Yaccpt+1]
		{
			YaccVAL.ConnectToGameSession = &ConnectToGameSession{
				SessionID: YaccDollar[2].Int,
			}
		}
	case 37:
		YaccDollar = YaccS[Yaccpt-5 : Yaccpt+1]
		{
			YaccVAL.Finished = &Finished{
				WinnerTeamID: YaccDollar[2].Int,
				WinReason:    YaccDollar[3].String,
				FinishReason: YaccDollar[4].String,
				GameTime:     YaccDollar[5].Float,
			}
		}
	case 38:
		YaccDollar = YaccS[Yaccpt-1 : Yaccpt+1]
		{
			YaccVAL.String = YaccDollar[1].String
		}
	case 39:
		YaccDollar = YaccS[Yaccpt-0 : Yaccpt+1]
		{
			YaccVAL.String = ""
		}
	case 40:
		YaccDollar = YaccS[Yaccpt-6 : Yaccpt+1]
		{
			YaccVAL.Reward = &Reward{
				Recipient:  YaccDollar[2].String,
				Ship:       YaccDollar[3].String,
				Reward:     YaccDollar[4].Int,
				RewardType: YaccDollar[5].String,
				Reason:     YaccDollar[6].String,
			}
		}
	}
	goto Yaccstack /* stack new state and value */
}