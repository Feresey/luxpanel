// Generated code. DO NOT EDIT

package gramma

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

type Tokenizer struct {
	nowTime time.Time
	tokens []Token
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
		t.tokval(StrTok(t.data[t.prev:t.p]))
	}
	action setInt {
		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(fmt.Errorf("strconv.Atoi: %w", parseErr))
			fbreak;
		}
		t.tokval(IntTok(temp.Int))
	}
	action setFloat {
		parsed, parseErr := strconv.ParseFloat(t.data[t.prev:t.p], 32)
		if parseErr != nil {
			t.err(fmt.Errorf("strconv.ParseFloat: %w", parseErr))
			fbreak;
		}
		t.tokval(FloatTok(parsed))
	}
	action setTime {
		temp.Time, parseErr = t.parseTime(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(fmt.Errorf("parseTime: %w", parseErr))
			fbreak;
		}
		t.tokval(TimeTok(temp.Time))
	}
	action start {
		fmt.Printf("start: %s\n", t.state)
		t.prev = t.p
	}

	action error {
		t.err(t.state)
	}

	action sym { t.tok(int(t.data[t.p]))}

	Time = (digit {2} ':' digit {2} ':' digit {2} '.' digit {3}) >start %setTime;
	Combat = ('CMBT   |') >start @{ t.tok(COMBAT) };
	Game = ('GAME   |') >start @{ t.tok(GAME) };


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
	Strings = string (' ' string)* >start %setString;
	Source = (
				string 
					>start %{t.tokval(NewAnyVal(SOURCE, StrTok(t.data[t.prev:t.p])))}
				|
				('(' string
					>start %{t.tokval(NewAnyVal(SOURCE, StrTok(t.data[t.prev:t.p])))}
				')') |
				zlen %{t.tok(SOURCE)}
			);
	DamageModifier = [A-Z_]+ >start %setString;
	Int = '-'?[0-9]+ >start %setInt;
	Float = '-'?[0-9]+'.'?[0-9]*;
	EOL = '\n' %{t.tok(EOL)};

	Object = (String | (String '(' >sym String ')' >sym)) '|' >sym Int;
	Killer = (String | (String '\t' >sym String)) '|' >sym Int;
	DamageModifiers = DamageModifier ('|' >sym DamageModifier);
	FriendlyFire = '<FriendlyFire>' @{t.tok(FRIENDLY_FIRE)};
	ParticipationModifier = '<' ('buff'|'debuff'|'heal') >start %{t.tokval(StrTok(t.data[t.prev:t.p]))} '>';
	ParticipationModifiers = ParticipationModifier (' ' ParticipationModifier)*;

	damageLine = Damage
		space+ Object ' -> ' Object ' '
		Float ' (h:' Float ' s:' Float ') ' Source? ' ' DamageModifiers (' ' FriendlyFire)?;
	healLine = Heal
		' '+ Object ' ->' ' '+ Object ' ' Float Source?;
	killLine = Kill space Object ';' space+ 'killer ' String '|' Int ' ' String (' ' FriendlyFire)?;
	participationLine = Participant ' '+ String '\t ' String? ' '* '\t' ' '? (
		'totalDamage ' Float '; mostDamageWith ' "'" Source "';" ' '?
	)?
	(ParticipationModifiers ' '?)?
	FriendlyFire?;
	startLine = Start ' ' (
		("gameplay '" >{t.tokval(StrTok("gameplay"))} String "'") |
		("PVE mission '" >{t.tokval(StrTok("PVE mission"))} String "'")
		) "map '" String "'"
		(', local client team ' Int '=======')?;
	finishedLine = Finish ' Winner team: ' Int Source ". Finish reason: '" Strings "'. Actual game time " Float ' sec';
	connectedLine = '======= Connect to game session ' >start %{t.tok(CONNECT_TO_GAME_SESSION_PREFIX)} Int ' =======';
	rewardLine = Reward ' '+ String (' ' String ' '*)? '\t' ' '* Int ' '
	('effective points' | 'experience' | 'karma' | 'credits') ' '*
	'for ' (any - '\n') >start %setString;

	clientAddPlayerLine = 'client: ADD_PLAYER ' >start %{t.tok(CLIENT_ADD_PLAYER)}
		Int '('String ' [' String '], ' Int ') status ' Int ' team ' Int (' group ' Int)?;
	clientPlayerLeaveLine = 'client: player '>start %{t.tok(CLIENT_PLAYER_LEAVE)} Int ' leave game';
	clientConnectedLine = 'client: connected to ' >start %{t.tok(CLIENT_CONNECTED)}
		(digit{1,3}('.'digit{1,3}){3} '|' digit+) >start %setString ', MTU ' Int
		'. setting up session...';
	clientConnectionClosed = 'client: connection closed. ' >start %{t.tok(CLIENT_CONNECTION_CLOSED)} String;

	main := Time ' '+
	(
		(Combat ' '+
			(
				damageLine |
				healLine |
				killLine |
				participationLine |
				startLine |
				finishedLine |
				connectedLine |
				rewardLine
			)
		)
		|
		(Game ' '+
			(
				clientAddPlayerLine |
				clientPlayerLeaveLine |
				clientConnectedLine |
				clientConnectionClosed
			)
		)
	)
		' '* EOL;
}%%

func (t *Tokenizer) Parse(nowTime time.Time, data string) ([]Token, error) {
	var parseErr error
	var temp yySymType

	t.state = state{}
	%%write init;

	t.prev = 0
	t.pe = len(data)
	t.nowTime = nowTime
	t.data = data

	t.tokens = t.tokens[:0]
	t.errors = nil

	%%write exec;

	fmt.Printf("exited: %s\n", t.state)
	if t.p != t.pe {
		t.err(fmt.Errorf("line is not finished: %s|||||%s", t.data[:t.p], t.data[t.p:t.pe]))
	}
	return t.tokens, errors.Join(t.errors...)
}

func (t *Tokenizer) parseTime(s string) (time.Time, error) {
	return ParseTime(t.nowTime, s)
}

func (t *Tokenizer) tokval(token Token) {
	t.tokens = append(t.tokens, token)
}

func (t *Tokenizer) tok(tok int) {
	t.tokens = append(t.tokens, AnyTok(tok))
}

func (t *Tokenizer) err(err error) {
	t.errors = append(t.errors, err)
}