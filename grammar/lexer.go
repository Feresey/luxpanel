
//line ./lexer.go.rl:1
// Generated code. DO NOT EDIT

package main

import (
    "fmt"
    "strconv"
    "time"
)


//line lexer.go:15
const logparser_start int = 32
const logparser_first_final int = 32
const logparser_error int = 0

const logparser_en_main int = 32


//line ./lexer.go.rl:22


type lexer struct {
    nowTime time.Time
    res []*LogLine[*CombatLine]

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
    
//line lexer.go:66
	{
	 lex.cs = logparser_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line ./lexer.go.rl:64
    return lex
}

func (lex *lexer) Lex(out *yySymType) int {
    eof := lex.pe
    tok := 0
    
//line lexer.go:82
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
	case 32:
		goto st_case_32
	case 0:
		goto st_case_0
	case 33:
		goto st_case_33
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	}
	goto st_out
tr0:
//line ./lexer.go.rl:121
( lex.p) = ( lex.te) - 1

	goto st32
tr8:
//line ./lexer.go.rl:83
 lex.te = ( lex.p)+1
{tok = COMBAT; {( lex.p)++;  lex.cs = 32; goto _out }}
	goto st32
tr11:
//line ./lexer.go.rl:85
 lex.te = ( lex.p)+1
{tok = ARROW; {( lex.p)++;  lex.cs = 32; goto _out }}
	goto st32
tr12:
//line ./lexer.go.rl:92
( lex.p) = ( lex.te) - 1
{
                // fmt.Printf("int: %q\n", lex.data[lex.ts:lex.te])
                res, err := strconv.Atoi(lex.data[lex.ts:lex.te])
                if err != nil {
                    lex.Error(err.Error())
                    {( lex.p)++;  lex.cs = 32; goto _out }
                }
                out.Int = res
                tok = INT;
                {( lex.p)++;  lex.cs = 32; goto _out }
            }
	goto st32
tr21:
//line ./lexer.go.rl:72
 lex.te = ( lex.p)+1
{
                tok = TIME;
                // fmt.Printf("time:%q\n", lex.data[lex.ts:lex.te])
                res, err := lex.parseTime(lex.data[lex.ts:lex.te])
                if err != nil {
                    lex.Error(err.Error())
                    {( lex.p)++;  lex.cs = 32; goto _out }
                }
                out.Time = res
                {( lex.p)++;  lex.cs = 32; goto _out }
            }
	goto st32
tr34:
//line ./lexer.go.rl:91
 lex.te = ( lex.p)+1
{tok = FRIENDLY_FIRE; out.Bool = true; {( lex.p)++;  lex.cs = 32; goto _out }}
	goto st32
tr35:
//line ./lexer.go.rl:121
 lex.te = ( lex.p)+1

	goto st32
tr36:
//line ./lexer.go.rl:120
 lex.te = ( lex.p)+1
{tok = EOL; {( lex.p)++;  lex.cs = 32; goto _out }}
	goto st32
tr38:
//line ./lexer.go.rl:86
 lex.te = ( lex.p)+1
{tok = LBRACE; {( lex.p)++;  lex.cs = 32; goto _out }}
	goto st32
tr39:
//line ./lexer.go.rl:88
 lex.te = ( lex.p)+1
{tok = RBRACE; {( lex.p)++;  lex.cs = 32; goto _out }}
	goto st32
tr47:
//line ./lexer.go.rl:87
 lex.te = ( lex.p)+1
{tok = VSLASH; {( lex.p)++;  lex.cs = 32; goto _out }}
	goto st32
tr48:
//line ./lexer.go.rl:121
 lex.te = ( lex.p)
( lex.p)--

	goto st32
tr50:
//line ./lexer.go.rl:92
 lex.te = ( lex.p)
( lex.p)--
{
                // fmt.Printf("int: %q\n", lex.data[lex.ts:lex.te])
                res, err := strconv.Atoi(lex.data[lex.ts:lex.te])
                if err != nil {
                    lex.Error(err.Error())
                    {( lex.p)++;  lex.cs = 32; goto _out }
                }
                out.Int = res
                tok = INT;
                {( lex.p)++;  lex.cs = 32; goto _out }
            }
	goto st32
tr52:
//line ./lexer.go.rl:103
 lex.te = ( lex.p)
( lex.p)--
{
                // fmt.Printf("float: %q\n", lex.data[lex.ts:lex.te])
                res, err := strconv.ParseFloat(lex.data[lex.ts:lex.te], 32)
                if err != nil {
                    lex.Error(err.Error())
                    {( lex.p)++;  lex.cs = 32; goto _out }
                }
                out.Float = float32(res)
                tok = FLOAT;
                {( lex.p)++;  lex.cs = 32; goto _out }
            }
	goto st32
tr53:
//line NONE:1
	switch  lex.act {
	case 3:
	{( lex.p) = ( lex.te) - 1
tok = DAMAGE; {( lex.p)++;  lex.cs = 32; goto _out }}
	case 13:
	{( lex.p) = ( lex.te) - 1

                tok = STRING;
                // fmt.Printf("name: %q\n", lex.data[lex.ts:lex.te])
                out.String = lex.data[lex.ts:lex.te]
                {( lex.p)++;  lex.cs = 32; goto _out }
            }
	}
	
	goto st32
tr57:
//line ./lexer.go.rl:114
 lex.te = ( lex.p)
( lex.p)--
{
                tok = STRING;
                // fmt.Printf("name: %q\n", lex.data[lex.ts:lex.te])
                out.String = lex.data[lex.ts:lex.te]
                {( lex.p)++;  lex.cs = 32; goto _out }
            }
	goto st32
tr63:
//line ./lexer.go.rl:89
 lex.te = ( lex.p)+1
{tok = DAMAGE_HULL_START; {( lex.p)++;  lex.cs = 32; goto _out }}
	goto st32
tr64:
//line ./lexer.go.rl:90
 lex.te = ( lex.p)+1
{tok = DAMAGE_SHIELD_START; {( lex.p)++;  lex.cs = 32; goto _out }}
	goto st32
	st32:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof32
		}
	st_case_32:
//line NONE:1
 lex.ts = ( lex.p)

//line lexer.go:346
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr36
		case 32:
			goto tr37
		case 40:
			goto tr38
		case 41:
			goto tr39
		case 45:
			goto st9
		case 47:
			goto tr41
		case 60:
			goto st19
		case 68:
			goto st40
		case 95:
			goto tr41
		case 104:
			goto st45
		case 115:
			goto st46
		case 124:
			goto tr47
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr35
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr41
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr41
			}
		default:
			goto st37
		}
		goto st0
st_case_0:
	st0:
		 lex.cs = 0
		goto _out
tr37:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st33
	st33:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof33
		}
	st_case_33:
//line lexer.go:405
		if  lex.data[( lex.p)] == 67 {
			goto st1
		}
		goto tr48
	st1:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		if  lex.data[( lex.p)] == 77 {
			goto st2
		}
		goto tr0
	st2:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		if  lex.data[( lex.p)] == 66 {
			goto st3
		}
		goto tr0
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		if  lex.data[( lex.p)] == 84 {
			goto st4
		}
		goto tr0
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		if  lex.data[( lex.p)] == 32 {
			goto st5
		}
		goto tr0
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		if  lex.data[( lex.p)] == 32 {
			goto st6
		}
		goto tr0
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		if  lex.data[( lex.p)] == 32 {
			goto st7
		}
		goto tr0
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		if  lex.data[( lex.p)] == 124 {
			goto st8
		}
		goto tr0
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		if  lex.data[( lex.p)] == 32 {
			goto tr8
		}
		goto tr0
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		if  lex.data[( lex.p)] == 62 {
			goto tr11
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st34
		}
		goto st0
	st34:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		if  lex.data[( lex.p)] == 46 {
			goto st35
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st34
		}
		goto tr50
	st35:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st35
		}
		goto tr52
tr41:
//line NONE:1
 lex.te = ( lex.p)+1

//line ./lexer.go.rl:114
 lex.act = 13;
	goto st36
tr62:
//line NONE:1
 lex.te = ( lex.p)+1

//line ./lexer.go.rl:84
 lex.act = 3;
	goto st36
	st36:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof36
		}
	st_case_36:
//line lexer.go:534
		if  lex.data[( lex.p)] == 95 {
			goto tr41
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr41
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr41
			}
		default:
			goto tr41
		}
		goto tr53
	st37:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch  lex.data[( lex.p)] {
		case 46:
			goto st35
		case 47:
			goto tr41
		case 95:
			goto tr41
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr54
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr41
			}
		default:
			goto tr41
		}
		goto tr50
tr54:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st38
	st38:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof38
		}
	st_case_38:
//line lexer.go:587
		switch  lex.data[( lex.p)] {
		case 46:
			goto st35
		case 47:
			goto tr41
		case 58:
			goto st10
		case 95:
			goto tr41
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st39
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr41
			}
		default:
			goto tr41
		}
		goto tr50
	st39:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch  lex.data[( lex.p)] {
		case 46:
			goto st35
		case 47:
			goto tr41
		case 95:
			goto tr41
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st39
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr41
			}
		default:
			goto tr41
		}
		goto tr50
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st11
		}
		goto tr12
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st12
		}
		goto tr12
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		if  lex.data[( lex.p)] == 58 {
			goto st13
		}
		goto tr12
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st14
		}
		goto tr12
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st15
		}
		goto tr12
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		if  lex.data[( lex.p)] == 46 {
			goto st16
		}
		goto tr12
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st17
		}
		goto tr12
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st18
		}
		goto tr12
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr21
		}
		goto tr12
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		if  lex.data[( lex.p)] == 70 {
			goto st20
		}
		goto st0
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		if  lex.data[( lex.p)] == 114 {
			goto st21
		}
		goto st0
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		if  lex.data[( lex.p)] == 105 {
			goto st22
		}
		goto st0
	st22:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		if  lex.data[( lex.p)] == 101 {
			goto st23
		}
		goto st0
	st23:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		if  lex.data[( lex.p)] == 110 {
			goto st24
		}
		goto st0
	st24:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		if  lex.data[( lex.p)] == 100 {
			goto st25
		}
		goto st0
	st25:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		if  lex.data[( lex.p)] == 108 {
			goto st26
		}
		goto st0
	st26:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		if  lex.data[( lex.p)] == 121 {
			goto st27
		}
		goto st0
	st27:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		if  lex.data[( lex.p)] == 70 {
			goto st28
		}
		goto st0
	st28:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		if  lex.data[( lex.p)] == 105 {
			goto st29
		}
		goto st0
	st29:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		if  lex.data[( lex.p)] == 114 {
			goto st30
		}
		goto st0
	st30:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		if  lex.data[( lex.p)] == 101 {
			goto st31
		}
		goto st0
	st31:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		if  lex.data[( lex.p)] == 62 {
			goto tr34
		}
		goto st0
	st40:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr41
		case 97:
			goto st41
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr41
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr41
			}
		default:
			goto tr41
		}
		goto tr57
	st41:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr41
		case 109:
			goto st42
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr41
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr41
			}
		default:
			goto tr41
		}
		goto tr57
	st42:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr41
		case 97:
			goto st43
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr41
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr41
			}
		default:
			goto tr41
		}
		goto tr57
	st43:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr41
		case 103:
			goto st44
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr41
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr41
			}
		default:
			goto tr41
		}
		goto tr57
	st44:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr41
		case 101:
			goto tr62
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr41
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr41
			}
		default:
			goto tr41
		}
		goto tr57
	st45:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch  lex.data[( lex.p)] {
		case 58:
			goto tr63
		case 95:
			goto tr41
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr41
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr41
			}
		default:
			goto tr41
		}
		goto tr57
	st46:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch  lex.data[( lex.p)] {
		case 58:
			goto tr64
		case 95:
			goto tr41
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr41
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr41
			}
		default:
			goto tr41
		}
		goto tr57
	st_out:
	_test_eof32:  lex.cs = 32; goto _test_eof
	_test_eof33:  lex.cs = 33; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof34:  lex.cs = 34; goto _test_eof
	_test_eof35:  lex.cs = 35; goto _test_eof
	_test_eof36:  lex.cs = 36; goto _test_eof
	_test_eof37:  lex.cs = 37; goto _test_eof
	_test_eof38:  lex.cs = 38; goto _test_eof
	_test_eof39:  lex.cs = 39; goto _test_eof
	_test_eof10:  lex.cs = 10; goto _test_eof
	_test_eof11:  lex.cs = 11; goto _test_eof
	_test_eof12:  lex.cs = 12; goto _test_eof
	_test_eof13:  lex.cs = 13; goto _test_eof
	_test_eof14:  lex.cs = 14; goto _test_eof
	_test_eof15:  lex.cs = 15; goto _test_eof
	_test_eof16:  lex.cs = 16; goto _test_eof
	_test_eof17:  lex.cs = 17; goto _test_eof
	_test_eof18:  lex.cs = 18; goto _test_eof
	_test_eof19:  lex.cs = 19; goto _test_eof
	_test_eof20:  lex.cs = 20; goto _test_eof
	_test_eof21:  lex.cs = 21; goto _test_eof
	_test_eof22:  lex.cs = 22; goto _test_eof
	_test_eof23:  lex.cs = 23; goto _test_eof
	_test_eof24:  lex.cs = 24; goto _test_eof
	_test_eof25:  lex.cs = 25; goto _test_eof
	_test_eof26:  lex.cs = 26; goto _test_eof
	_test_eof27:  lex.cs = 27; goto _test_eof
	_test_eof28:  lex.cs = 28; goto _test_eof
	_test_eof29:  lex.cs = 29; goto _test_eof
	_test_eof30:  lex.cs = 30; goto _test_eof
	_test_eof31:  lex.cs = 31; goto _test_eof
	_test_eof40:  lex.cs = 40; goto _test_eof
	_test_eof41:  lex.cs = 41; goto _test_eof
	_test_eof42:  lex.cs = 42; goto _test_eof
	_test_eof43:  lex.cs = 43; goto _test_eof
	_test_eof44:  lex.cs = 44; goto _test_eof
	_test_eof45:  lex.cs = 45; goto _test_eof
	_test_eof46:  lex.cs = 46; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 33:
			goto tr48
		case 1:
			goto tr0
		case 2:
			goto tr0
		case 3:
			goto tr0
		case 4:
			goto tr0
		case 5:
			goto tr0
		case 6:
			goto tr0
		case 7:
			goto tr0
		case 8:
			goto tr0
		case 34:
			goto tr50
		case 35:
			goto tr52
		case 36:
			goto tr53
		case 37:
			goto tr50
		case 38:
			goto tr50
		case 39:
			goto tr50
		case 10:
			goto tr12
		case 11:
			goto tr12
		case 12:
			goto tr12
		case 13:
			goto tr12
		case 14:
			goto tr12
		case 15:
			goto tr12
		case 16:
			goto tr12
		case 17:
			goto tr12
		case 18:
			goto tr12
		case 40:
			goto tr57
		case 41:
			goto tr57
		case 42:
			goto tr57
		case 43:
			goto tr57
		case 44:
			goto tr57
		case 45:
			goto tr57
		case 46:
			goto tr57
		}
	}

	_out: {}
	}

//line ./lexer.go.rl:125


    return tok;
}

func (lex *lexer) Error(e string) {
    fmt.Println("error:", e)
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