// Generated code. DO NOT EDIT

package game

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
		t.tokval(strTok(t.data[t.prev:t.p]))
	}
	action setInt {
		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: "INT", Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			fbreak;
		}
		t.tokval(intTok(temp.Int))
	}
	action setFloat {
		parsed, parseErr := strconv.ParseFloat(t.data[t.prev:t.p], 32)
		if parseErr != nil {
			t.err(ParseTokenError{TokType: "FLOAT", Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.ParseFloat: %w", parseErr)})
			fbreak;
		}
		t.tokval(floatTok(parsed))
	}
	action setTime {
		temp.Time, parseErr = parseTime(t.nowTime, t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(ParseTokenError{TokType: "TIME", Raw: t.data[t.prev:t.p], Err: fmt.Errorf("parseTime: %w", parseErr)})
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
	Game = (space+ '|') >start @{ t.tok(GAME) };

	string = [a-zA-Z0-9_/]+;
	value = '(' string ')';
	String = (string | value) >start %setString;

	Int = ('-'?[0-9]+) >start %setInt;

	clientAddPlayerLine = 'client: ADD_PLAYER ' >start %{t.tok(CLIENT_ADD_PLAYER)}
		Int ' (' String (' [' String? ']')? ', ' Int ') status ' Int ' team ' Int (' group ' Int)?;
	clientPlayerLeaveLine = 'client: player '>start %{t.tok(CLIENT_PLAYER_LEAVE)} Int ' leave game' @{t.tok(LEAVE)};
	clientConnectedLine = 'client: connected to ' >start %{t.tok(CLIENT_CONNECTED)}
		(digit{1,3}('.'digit{1,3}){3} '|' digit+) >start %setString ', MTU ' Int
		'. setting up session...';
	clientConnectionClosed = 'client: connection closed. ' >start %{t.tok(CLIENT_CONNECTION_CLOSED)} String;

	main := Time ' '+ Game ' '+
		(
			clientAddPlayerLine |
			clientPlayerLeaveLine |
			clientConnectedLine |
			clientConnectionClosed
		) space* %eol;
}%%

func (t *tokenizer) Parse(nowTime time.Time, data string) ([]token, error) {
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