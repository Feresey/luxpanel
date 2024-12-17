// Generated code. DO NOT EDIT

package combat

import (
	"fmt"
	"strconv"
	"time"
	"errors"
)

%%{
	machine logparser;
	access t.;
	variable p t.p;
	variable pe t.pe;
	variable eof t.pe;
	variable data t.data;
	write data;
}%%

type state struct {
	data string

	prev int

	// Data end pointer. This should be initialized to p plus the data length on every run of the machine. In Java and Ruby code this should be initialized to the data length.
	pe int

	// Data pointer. In C/D code this variable is expected to be a pointer to the character data to process.
	// It should be initialized to the beginning of the data block on every run of the machine.
	// In Java and Ruby it is used as an offset to data and must be an integer.
	// In this case it should be initialized to zero on every run of the machine.
	p int

	// Token start
	ts int
	// Token end
	te int

	// This must be an integer value. It is a variable sometimes used by scanner code to keep track of the most recent successful pattern match.
	act int

	// Current state.
	// This must be an integer and it should persist across invocations of the machine when the data is broken into blocks that are processed independently.
	// This variable may be modified from outside the execution loop, but not from within.
	cs int

	// This must be an integer value and will be used as an offset to stack, giving the next available spot on the top of the stack.
	top int
}

func (e state) String() string {
	return fmt.Sprintf("state: %q p=%d, pe=%d, ts=%d, te=%d, act=%d, cs=%d, top=%d", e.data[:e.p], e.p, e.pe, e.ts, e.te, e.act, e.cs, e.top)
}

func (e state) Error() string {
	return e.String()
}

type tokenizer struct {
	nowTime time.Time
	tokens []token
	errors []error
	state
}

%%{
	action setCombat {
		t.tok(COMBAT)
	}
	action setDamage {
		t.tok(DAMAGE)
	}

	action setString {
		if debugTokenizer {
			fmt.Printf("set string: %s %s\n", t.data[t.prev:t.p],t.state)
		}
		t.tokval(strTok(t.data[t.prev:t.p]))
	}
	action setInt {
		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: INT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			fbreak;
		}
		t.tokval(intTok(temp.Int))
	}
	action setFloat {
		Float, parseErr = strconv.ParseFloat(t.data[t.prev:t.p], 32)
		if parseErr != nil {
			t.err(ParseTokenError{TokType: FLOAT, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.ParseFloat: %w", parseErr)})
			fbreak;
		}
		t.tokval(floatTok(float32(Float)))
	}
	action setTime {
		temp.Time, parseErr = parseTime(t.nowTime, t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: TIME, Raw: t.data[t.prev:t.p], Err: fmt.Errorf("parseTime: %w", parseErr)})
			fbreak;
		}
		t.tokval(timeTok(temp.Time))
	}
	action start {
		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	}

	action error {
		t.err(t.state)
	}

	action sym { t.tok(int(t.data[t.p]))}

	action eol { t.tok(EOL) }

	Time = (digit {2} ':' digit {2} ':' digit {2} '.' digit {3}) >start %setTime;
	Combat = ('CMBT' space+ '|') >start @{ t.tok(COMBAT) };
	Game = (space+ '|') >start @{ t.tok(GAME) };

	Damage = ('Damage') >start @{ t.tok(DAMAGE) };
	Heal = ('Heal') >start @{ t.tok(HEAL) };
	Kill = ('Killed') >start @{ t.tok(KILL) };
	Participant = ('Participant') >start @{ t.tok(PARTICIPANT) };
	Start = ('======= Start') >start @{ t.tok(START) };
	Finish = ('Gameplay finished.') >start @{ t.tok(GAMEPLAY_FINISHED) };
	Reward = ('Reward') >start @{ t.tok(REWARD) };

	string = [a-zA-Z0-9_/]+;
	value = '(' string ')';
	String = (string | value) >start %setString;
	Strings = (string (' ' string)*) >start %setString;
	Source = (string | value) >start %{t.tokval(newAnyVal(SOURCE, strTok(t.data[t.prev:t.p])))};
	DamageModifier = [A-Z_]+ >start %{t.tokval(newAnyVal(DAMAGE_MODIFIER, strTok(t.data[t.prev:t.p])))};
	Int = ('-'?[0-9]+) >start %setInt;
	Float = ('-'?[0-9]+'.'?[0-9]*) >start %setFloat;

	Object = String ('(' >sym String ')' >sym)? '|' >sym Int;
	Arrow = '->' @{t.tok(ARROW)};
	Killed = (String '\t' >sym)? space* Object;
	DamageModifiers = DamageModifier ('|' >sym DamageModifier)*;
	FriendlyFire = '<FriendlyFire>' @{t.tok(FRIENDLY_FIRE)};
	Rocket = 'Rocket' >{t.tok(ROCKET)} space+ Int;
	ParticipationModifier = '<' ('buff'|'debuff'|'heal') >start %{t.tokval(newAnyVal(PARTICIPATION_MODIFIER, strTok(t.data[t.prev:t.p])))} '>';
	ParticipationModifiers = ParticipationModifier (' ' ParticipationModifier)* %{t.tok(PARTICIPATION_MODIFIERS_END)};

	damageLine = Damage
		space+ Object space+ Arrow space+ Object space+
		Float space+ '(h:' Float space+ 's:' Float ') ' Source? ' ' DamageModifiers (space+ FriendlyFire)? (space+ Rocket)?;
	healLine = Heal
		' '+ Object space+ Arrow space+ Object space+ Float space+ Source;
	killLine = Kill space+ Killed ';' >sym space+ 'killer ' String '|' >sym Int ' ' Source? (' ' FriendlyFire)?;
	participationLine = Participant ' '+ String '\t ' String? ' '* '\t' ' '? (
		'totalDamage ' Float '; mostDamageWith ' "'" Source "';" ' '?
	)?
	(ParticipationModifiers ' '?)?
	FriendlyFire?;
	startLine = Start ' ' (
		("gameplay " >{t.tokval(strTok("gameplay"))} "'" String "'") |
		("PVE mission " >{t.tokval(strTok("PVE mission"))} "'" String "'")
		) " map '" String "'"
		(', local client team ' Int)? space* '=======';
	finishedLine = Finish ' Winner team: ' Int '(' String "). Finish reason: '" Strings "'. Actual game time " Float ' sec';
	connectedLine = '======= Connect to game session ' >start %{t.tok(CONNECT_TO_GAME_SESSION_PREFIX)} Int ' =======';
	rewardLine = Reward space+ String (' '+ String)? space* '\t' space* Int space+
	('effective points' | 'experience' | 'karma' | 'credits') >start %setString space+
	'for ' >start %{t.tokval(strTok(t.data[t.p:]))} (any - ('\r'|'\n'))*;

	main := Time ' '+ Combat ' '+
		(
			damageLine |
			healLine |
			killLine |
			participationLine |
			startLine |
			finishedLine |
			connectedLine |
			rewardLine
		) space* %eol;
}%%

func (t *tokenizer) Parse(nowTime time.Time, data string) ([]token, error) {
	var parseErr error
	var temp yySymType
	var Float float64

	t.state = state{}
	%%write init;

	t.prev = 0
	t.pe = len(data)
	t.nowTime = nowTime
	t.data = data

	t.tokens = t.tokens[:0]
	t.errors = nil

	%%write exec;

	if debugTokenizer {
		fmt.Printf("exited: %s\n", t.state)
	}
	if t.p != t.pe || (len(t.tokens) > 0 && t.tokens[len(t.tokens)-1] != voidTok(EOL)) {
		t.err(fmt.Errorf("%w: %q %q", ErrLineIsNotFinished, t.data[:t.p], t.data[t.p:t.pe]))
	}

	return t.tokens, errors.Join(t.errors...)
}

func (t *tokenizer) tokval(token token) {
	t.tokens = append(t.tokens, token)
}

func (t *tokenizer) tok(tok int) {
	t.tokens = append(t.tokens, voidTok(tok))
}

func (t *tokenizer) err(err error) {
	t.errors = append(t.errors, err)
}