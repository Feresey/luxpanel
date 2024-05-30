// Generated code. DO NOT EDIT

package main

import (
    "fmt"
    "strconv"
    "time"
)

%%{
    machine logparser;
    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;

    Time = digit {2} ':' digit {2} ':' digit {2} '.' digit {3};
    Name = [a-zA-Z0-9_/]+;
    Int = '-'?[0-9]+;
    Float = '-'?[0-9]+'.'?[0-9]*;
}%%

type lexer struct {
    nowTime time.Time
    res CombatLine

    errors []string

    // It must be an array containting the data to process.
    data string

    // Data end pointer. This should be initialized to p plus the data length on every run of the machine. In Java and Ruby code this should be initialized to the data length.
    pe int

    // Data pointer. In C/D code this variable is expected to be a pointer to the character data to process.
    // It should be initialized to the beginning of the data block on every run of the machine.
    // In Java and Ruby it is used as an offset to data and must be an integer.
    // In this case it should be initialized to zero on every run of the machine.
    p int


    // This must be a pointer to character data. In Java and Ruby code this must be an integer. See Section 6.3 for more information.
    ts int

    // Also a pointer to character data.
    te int

    // This must be an integer value. It is a variable sometimes used by scanner code to keep track of the most recent successful pattern match.
    act int

    // Current state. This must be an integer and it should persist across invocations of the machine when the data is broken into blocks that are processed independently. This variable may be modified from outside the execution loop, but not from within.
    cs int

    // This must be an integer value and will be used as an offset to stack, giving the next available spot on the top of the stack.
    top int
}

func newLexer(data string, nowTime time.Time) *lexer {
    lex := &lexer{
        nowTime: nowTime,
        data: data,
        pe: len(data),
    }
    %% write init;
    return lex
}

func (lex *lexer) Lex(out *yySymType) int {
    eof := lex.pe
    var tok int
    %%{
        main := |*
            ' CMBT   | ' => {tok = COMBAT; fbreak;};
            '=======' => {tok = EQ_DELIM; fbreak;};
            "'" => {tok = int('\''); fbreak;};
            "," => {tok = int(','); fbreak;};
            ";" => {tok = int(';'); fbreak;};
            "\t" => {tok = int('\t'); fbreak;};
            '(' => {tok = int('('); fbreak;};
            '|' => {tok = int('|'); fbreak;};
            ')' => {tok = int(')'); fbreak;};
            'Connect to game session' => {tok = CONNECT_TO_GAME_SESSION_PREFIX; fbreak;};
            'Gameplay finished. Winner team: ' => {tok = GAMEPLAY_FINISHED; fbreak;};
            '. Finish reason: ' => {tok = FINISH_REASON; fbreak;};
            '. Actual game time ' => {tok = ACTUAL_GAME_TIME; fbreak;};
            'local client team' => {tok = LOCAL_CLIENT_TEAM; fbreak;};
            'Start' => {tok = START; fbreak;};
            'Damage' => {tok = DAMAGE; fbreak;};
            'Killed' => {tok = KILL; fbreak;};
            'Participant' => {tok = PARTICIPANT; fbreak;};
            'Heal' => {tok = HEAL; fbreak;};
            '->' => {tok = ARROW; fbreak;};
            'h:' => {tok = DAMAGE_HULL_START; fbreak;};
            's:' => {tok = DAMAGE_SHIELD_START; fbreak;};
            '<FriendlyFire>' => {tok = FRIENDLY_FIRE; out.Bool = true; fbreak;};
            Time => {
                tok = TIME;
                // fmt.Printf("time:%q\n", lex.data[lex.ts:lex.te])
                res, err := lex.parseTime(lex.data[lex.ts:lex.te])
                if err != nil {
                    lex.Error(err.Error())
                    fbreak;
                }
                out.Time = res
                fbreak;
            };
            Int => {
                // fmt.Printf("int: %q\n", lex.data[lex.ts:lex.te])
                res, err := strconv.Atoi(lex.data[lex.ts:lex.te])
                if err != nil {
                    lex.Error(err.Error())
                    fbreak;
                }
                out.Int = res
                tok = INT;
                fbreak;
            };
            Float => {
                // fmt.Printf("float: %q\n", lex.data[lex.ts:lex.te])
                res, err := strconv.ParseFloat(lex.data[lex.ts:lex.te], 32)
                if err != nil {
                    lex.Error(err.Error())
                    fbreak;
                }
                out.Float = float32(res)
                tok = FLOAT;
                fbreak;
            };
            Name => {
                tok = STRING;
                // fmt.Printf("name: %q\n", lex.data[lex.ts:lex.te])
                out.String = lex.data[lex.ts:lex.te]
                fbreak;
            };
            '\n' => {tok = EOL; fbreak;};
            space;
        *|;

         write exec;
    }%%

    return tok;
}

func (lex *lexer) Error(e string) {
    lex.errors = append(lex.errors, e)
}

func (lex *lexer) Err() error {
    if len(lex.errors) != 0 {
        return fmt.Errorf("%v", lex.errors)
    }
    return nil
}


const timeFormat = "15:04:05.000"

func (lex *lexer) parseTime(s string) (time.Time, error) {
	t, err := time.Parse(timeFormat, s)
	if err != nil {
		return t, fmt.Errorf("time.Parse(%s): %w", timeFormat, err)
	}
	res := time.Date(lex.nowTime.Year(), lex.nowTime.Month(), lex.nowTime.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	if res.Before(lex.nowTime) {
		return res.AddDate(0, 0, 1), nil
	}
	return res, nil
}