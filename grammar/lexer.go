
//line lexer.go.rl:1
// Generated code. DO NOT EDIT

package main

import (
    "fmt"
    "strconv"
    "time"
)


//line lexer.go:15
const logparser_start int = 64
const logparser_first_final int = 64
const logparser_error int = 0

const logparser_en_main int = 64


//line lexer.go.rl:22


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

//line lexer.go.rl:64
    return lex
}

func (lex *lexer) Lex(out *yySymType) int {
    eof := lex.pe
    var tok int
    
//line lexer.go:82
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
	case 64:
		goto st_case_64
	case 0:
		goto st_case_0
	case 65:
		goto st_case_65
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
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
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
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
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
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 97:
		goto st_case_97
	}
	goto st_out
tr0:
//line lexer.go.rl:130
( lex.p) = ( lex.te) - 1

	goto st64
tr8:
//line lexer.go.rl:72
 lex.te = ( lex.p)+1
{tok = COMBAT; {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr11:
//line lexer.go.rl:83
 lex.te = ( lex.p)+1
{tok = ARROW; {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr12:
//line lexer.go.rl:101
( lex.p) = ( lex.te) - 1
{
                // fmt.Printf("int: %q\n", lex.data[lex.ts:lex.te])
                res, err := strconv.Atoi(lex.data[lex.ts:lex.te])
                if err != nil {
                    lex.Error(err.Error())
                    {( lex.p)++;  lex.cs = 64; goto _out }
                }
                out.Int = res
                tok = INT;
                {( lex.p)++;  lex.cs = 64; goto _out }
            }
	goto st64
tr21:
//line lexer.go.rl:90
 lex.te = ( lex.p)+1
{
                tok = TIME;
                // fmt.Printf("time:%q\n", lex.data[lex.ts:lex.te])
                res, err := lex.parseTime(lex.data[lex.ts:lex.te])
                if err != nil {
                    lex.Error(err.Error())
                    {( lex.p)++;  lex.cs = 64; goto _out }
                }
                out.Time = res
                {( lex.p)++;  lex.cs = 64; goto _out }
            }
	goto st64
tr34:
//line lexer.go.rl:89
 lex.te = ( lex.p)+1
{tok = FRIENDLY_FIRE; out.Bool = true; {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr40:
//line lexer.go.rl:73
 lex.te = ( lex.p)+1
{tok = EQ_DELIM; {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr41:
//line lexer.go.rl:123
( lex.p) = ( lex.te) - 1
{
                tok = STRING;
                // fmt.Printf("name: %q\n", lex.data[lex.ts:lex.te])
                out.String = lex.data[lex.ts:lex.te]
                {( lex.p)++;  lex.cs = 64; goto _out }
            }
	goto st64
tr56:
//line lexer.go.rl:78
 lex.te = ( lex.p)+1
{tok = CONNECT_TO_GAME_SESSION_PREFIX; {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr67:
//line lexer.go.rl:79
 lex.te = ( lex.p)+1
{tok = LOCAL_CLIENT_TEAM; {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr68:
//line lexer.go.rl:77
 lex.te = ( lex.p)+1
{tok = int('\t'); {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr69:
//line lexer.go.rl:129
 lex.te = ( lex.p)+1
{tok = EOL; {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr70:
//line lexer.go.rl:130
 lex.te = ( lex.p)+1

	goto st64
tr72:
//line lexer.go.rl:74
 lex.te = ( lex.p)+1
{tok = QUOTE; {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr73:
//line lexer.go.rl:84
 lex.te = ( lex.p)+1
{tok = LBRACE; {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr74:
//line lexer.go.rl:86
 lex.te = ( lex.p)+1
{tok = RBRACE; {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr75:
//line lexer.go.rl:75
 lex.te = ( lex.p)+1
{tok = COMMA; {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr79:
//line lexer.go.rl:76
 lex.te = ( lex.p)+1
{tok = int(';'); {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr89:
//line lexer.go.rl:85
 lex.te = ( lex.p)+1
{tok = VSLASH; {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr90:
//line lexer.go.rl:130
 lex.te = ( lex.p)
( lex.p)--

	goto st64
tr92:
//line lexer.go.rl:101
 lex.te = ( lex.p)
( lex.p)--
{
                // fmt.Printf("int: %q\n", lex.data[lex.ts:lex.te])
                res, err := strconv.Atoi(lex.data[lex.ts:lex.te])
                if err != nil {
                    lex.Error(err.Error())
                    {( lex.p)++;  lex.cs = 64; goto _out }
                }
                out.Int = res
                tok = INT;
                {( lex.p)++;  lex.cs = 64; goto _out }
            }
	goto st64
tr94:
//line lexer.go.rl:112
 lex.te = ( lex.p)
( lex.p)--
{
                // fmt.Printf("float: %q\n", lex.data[lex.ts:lex.te])
                res, err := strconv.ParseFloat(lex.data[lex.ts:lex.te], 32)
                if err != nil {
                    lex.Error(err.Error())
                    {( lex.p)++;  lex.cs = 64; goto _out }
                }
                out.Float = float32(res)
                tok = FLOAT;
                {( lex.p)++;  lex.cs = 64; goto _out }
            }
	goto st64
tr95:
//line NONE:1
	switch  lex.act {
	case 9:
	{( lex.p) = ( lex.te) - 1
tok = START; {( lex.p)++;  lex.cs = 64; goto _out }}
	case 10:
	{( lex.p) = ( lex.te) - 1
tok = DAMAGE; {( lex.p)++;  lex.cs = 64; goto _out }}
	case 11:
	{( lex.p) = ( lex.te) - 1
tok = HEAL; {( lex.p)++;  lex.cs = 64; goto _out }}
	case 22:
	{( lex.p) = ( lex.te) - 1

                tok = STRING;
                // fmt.Printf("name: %q\n", lex.data[lex.ts:lex.te])
                out.String = lex.data[lex.ts:lex.te]
                {( lex.p)++;  lex.cs = 64; goto _out }
            }
	}
	
	goto st64
tr99:
//line lexer.go.rl:123
 lex.te = ( lex.p)
( lex.p)--
{
                tok = STRING;
                // fmt.Printf("name: %q\n", lex.data[lex.ts:lex.te])
                out.String = lex.data[lex.ts:lex.te]
                {( lex.p)++;  lex.cs = 64; goto _out }
            }
	goto st64
tr119:
//line lexer.go.rl:87
 lex.te = ( lex.p)+1
{tok = DAMAGE_HULL_START; {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
tr125:
//line lexer.go.rl:88
 lex.te = ( lex.p)+1
{tok = DAMAGE_SHIELD_START; {( lex.p)++;  lex.cs = 64; goto _out }}
	goto st64
	st64:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof64
		}
	st_case_64:
//line NONE:1
 lex.ts = ( lex.p)

//line lexer.go:499
		switch  lex.data[( lex.p)] {
		case 9:
			goto tr68
		case 10:
			goto tr69
		case 32:
			goto tr71
		case 39:
			goto tr72
		case 40:
			goto tr73
		case 41:
			goto tr74
		case 44:
			goto tr75
		case 45:
			goto st9
		case 47:
			goto tr77
		case 59:
			goto tr79
		case 60:
			goto st19
		case 61:
			goto st32
		case 67:
			goto st72
		case 68:
			goto st79
		case 72:
			goto st84
		case 83:
			goto st87
		case 95:
			goto tr77
		case 104:
			goto st91
		case 108:
			goto st92
		case 115:
			goto st97
		case 124:
			goto tr89
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 11 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr70
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr77
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr77
			}
		default:
			goto st69
		}
		goto st0
st_case_0:
	st0:
		 lex.cs = 0
		goto _out
tr71:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st65
	st65:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof65
		}
	st_case_65:
//line lexer.go:576
		if  lex.data[( lex.p)] == 67 {
			goto st1
		}
		goto tr90
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
			goto st66
		}
		goto st0
	st66:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof66
		}
	st_case_66:
		if  lex.data[( lex.p)] == 46 {
			goto st67
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st66
		}
		goto tr92
	st67:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof67
		}
	st_case_67:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st67
		}
		goto tr94
tr77:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.go.rl:123
 lex.act = 22;
	goto st68
tr111:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.go.rl:81
 lex.act = 10;
	goto st68
tr114:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.go.rl:82
 lex.act = 11;
	goto st68
tr118:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.go.rl:80
 lex.act = 9;
	goto st68
	st68:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof68
		}
	st_case_68:
//line lexer.go:719
		if  lex.data[( lex.p)] == 95 {
			goto tr77
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr95
	st69:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof69
		}
	st_case_69:
		switch  lex.data[( lex.p)] {
		case 46:
			goto st67
		case 47:
			goto tr77
		case 95:
			goto tr77
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr96
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr92
tr96:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st70
	st70:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof70
		}
	st_case_70:
//line lexer.go:772
		switch  lex.data[( lex.p)] {
		case 46:
			goto st67
		case 47:
			goto tr77
		case 58:
			goto st10
		case 95:
			goto tr77
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st71
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr92
	st71:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof71
		}
	st_case_71:
		switch  lex.data[( lex.p)] {
		case 46:
			goto st67
		case 47:
			goto tr77
		case 95:
			goto tr77
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st71
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr92
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
	st32:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		if  lex.data[( lex.p)] == 61 {
			goto st33
		}
		goto st0
	st33:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		if  lex.data[( lex.p)] == 61 {
			goto st34
		}
		goto st0
	st34:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		if  lex.data[( lex.p)] == 61 {
			goto st35
		}
		goto st0
	st35:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		if  lex.data[( lex.p)] == 61 {
			goto st36
		}
		goto st0
	st36:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		if  lex.data[( lex.p)] == 61 {
			goto st37
		}
		goto st0
	st37:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		if  lex.data[( lex.p)] == 61 {
			goto tr40
		}
		goto st0
	st72:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof72
		}
	st_case_72:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 111:
			goto st73
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st73:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof73
		}
	st_case_73:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 110:
			goto st74
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st74:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof74
		}
	st_case_74:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 110:
			goto st75
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st75:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof75
		}
	st_case_75:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 101:
			goto st76
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st76:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof76
		}
	st_case_76:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 99:
			goto st77
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st77:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof77
		}
	st_case_77:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 116:
			goto tr105
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
tr105:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st78
	st78:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof78
		}
	st_case_78:
//line lexer.go:1228
		switch  lex.data[( lex.p)] {
		case 32:
			goto st38
		case 95:
			goto tr77
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st38:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		if  lex.data[( lex.p)] == 116 {
			goto st39
		}
		goto tr41
	st39:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		if  lex.data[( lex.p)] == 111 {
			goto st40
		}
		goto tr41
	st40:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		if  lex.data[( lex.p)] == 32 {
			goto st41
		}
		goto tr41
	st41:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		if  lex.data[( lex.p)] == 103 {
			goto st42
		}
		goto tr41
	st42:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		if  lex.data[( lex.p)] == 97 {
			goto st43
		}
		goto tr41
	st43:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		if  lex.data[( lex.p)] == 109 {
			goto st44
		}
		goto tr41
	st44:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		if  lex.data[( lex.p)] == 101 {
			goto st45
		}
		goto tr41
	st45:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		if  lex.data[( lex.p)] == 32 {
			goto st46
		}
		goto tr41
	st46:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		if  lex.data[( lex.p)] == 115 {
			goto st47
		}
		goto tr41
	st47:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof47
		}
	st_case_47:
		if  lex.data[( lex.p)] == 101 {
			goto st48
		}
		goto tr41
	st48:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof48
		}
	st_case_48:
		if  lex.data[( lex.p)] == 115 {
			goto st49
		}
		goto tr41
	st49:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof49
		}
	st_case_49:
		if  lex.data[( lex.p)] == 115 {
			goto st50
		}
		goto tr41
	st50:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof50
		}
	st_case_50:
		if  lex.data[( lex.p)] == 105 {
			goto st51
		}
		goto tr41
	st51:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof51
		}
	st_case_51:
		if  lex.data[( lex.p)] == 111 {
			goto st52
		}
		goto tr41
	st52:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof52
		}
	st_case_52:
		if  lex.data[( lex.p)] == 110 {
			goto tr56
		}
		goto tr41
	st79:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof79
		}
	st_case_79:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 97:
			goto st80
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st80:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof80
		}
	st_case_80:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 109:
			goto st81
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st81:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof81
		}
	st_case_81:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 97:
			goto st82
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st82:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof82
		}
	st_case_82:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 103:
			goto st83
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st83:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof83
		}
	st_case_83:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 101:
			goto tr111
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st84:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof84
		}
	st_case_84:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 101:
			goto st85
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st85:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof85
		}
	st_case_85:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 97:
			goto st86
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st86:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof86
		}
	st_case_86:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 108:
			goto tr114
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st87:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof87
		}
	st_case_87:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 116:
			goto st88
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st88:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof88
		}
	st_case_88:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 97:
			goto st89
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st89:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof89
		}
	st_case_89:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 114:
			goto st90
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st90:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof90
		}
	st_case_90:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 116:
			goto tr118
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st91:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof91
		}
	st_case_91:
		switch  lex.data[( lex.p)] {
		case 58:
			goto tr119
		case 95:
			goto tr77
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st92:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof92
		}
	st_case_92:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 111:
			goto st93
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st93:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof93
		}
	st_case_93:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 99:
			goto st94
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st94:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof94
		}
	st_case_94:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 97:
			goto st95
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st95:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof95
		}
	st_case_95:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr77
		case 108:
			goto tr123
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
tr123:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st96
	st96:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof96
		}
	st_case_96:
//line lexer.go:1801
		switch  lex.data[( lex.p)] {
		case 32:
			goto st53
		case 95:
			goto tr77
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st53:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof53
		}
	st_case_53:
		if  lex.data[( lex.p)] == 99 {
			goto st54
		}
		goto tr41
	st54:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof54
		}
	st_case_54:
		if  lex.data[( lex.p)] == 108 {
			goto st55
		}
		goto tr41
	st55:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof55
		}
	st_case_55:
		if  lex.data[( lex.p)] == 105 {
			goto st56
		}
		goto tr41
	st56:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof56
		}
	st_case_56:
		if  lex.data[( lex.p)] == 101 {
			goto st57
		}
		goto tr41
	st57:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof57
		}
	st_case_57:
		if  lex.data[( lex.p)] == 110 {
			goto st58
		}
		goto tr41
	st58:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof58
		}
	st_case_58:
		if  lex.data[( lex.p)] == 116 {
			goto st59
		}
		goto tr41
	st59:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof59
		}
	st_case_59:
		if  lex.data[( lex.p)] == 32 {
			goto st60
		}
		goto tr41
	st60:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof60
		}
	st_case_60:
		if  lex.data[( lex.p)] == 116 {
			goto st61
		}
		goto tr41
	st61:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof61
		}
	st_case_61:
		if  lex.data[( lex.p)] == 101 {
			goto st62
		}
		goto tr41
	st62:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof62
		}
	st_case_62:
		if  lex.data[( lex.p)] == 97 {
			goto st63
		}
		goto tr41
	st63:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof63
		}
	st_case_63:
		if  lex.data[( lex.p)] == 109 {
			goto tr67
		}
		goto tr41
	st97:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof97
		}
	st_case_97:
		switch  lex.data[( lex.p)] {
		case 58:
			goto tr125
		case 95:
			goto tr77
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr77
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto tr99
	st_out:
	_test_eof64:  lex.cs = 64; goto _test_eof
	_test_eof65:  lex.cs = 65; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof66:  lex.cs = 66; goto _test_eof
	_test_eof67:  lex.cs = 67; goto _test_eof
	_test_eof68:  lex.cs = 68; goto _test_eof
	_test_eof69:  lex.cs = 69; goto _test_eof
	_test_eof70:  lex.cs = 70; goto _test_eof
	_test_eof71:  lex.cs = 71; goto _test_eof
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
	_test_eof32:  lex.cs = 32; goto _test_eof
	_test_eof33:  lex.cs = 33; goto _test_eof
	_test_eof34:  lex.cs = 34; goto _test_eof
	_test_eof35:  lex.cs = 35; goto _test_eof
	_test_eof36:  lex.cs = 36; goto _test_eof
	_test_eof37:  lex.cs = 37; goto _test_eof
	_test_eof72:  lex.cs = 72; goto _test_eof
	_test_eof73:  lex.cs = 73; goto _test_eof
	_test_eof74:  lex.cs = 74; goto _test_eof
	_test_eof75:  lex.cs = 75; goto _test_eof
	_test_eof76:  lex.cs = 76; goto _test_eof
	_test_eof77:  lex.cs = 77; goto _test_eof
	_test_eof78:  lex.cs = 78; goto _test_eof
	_test_eof38:  lex.cs = 38; goto _test_eof
	_test_eof39:  lex.cs = 39; goto _test_eof
	_test_eof40:  lex.cs = 40; goto _test_eof
	_test_eof41:  lex.cs = 41; goto _test_eof
	_test_eof42:  lex.cs = 42; goto _test_eof
	_test_eof43:  lex.cs = 43; goto _test_eof
	_test_eof44:  lex.cs = 44; goto _test_eof
	_test_eof45:  lex.cs = 45; goto _test_eof
	_test_eof46:  lex.cs = 46; goto _test_eof
	_test_eof47:  lex.cs = 47; goto _test_eof
	_test_eof48:  lex.cs = 48; goto _test_eof
	_test_eof49:  lex.cs = 49; goto _test_eof
	_test_eof50:  lex.cs = 50; goto _test_eof
	_test_eof51:  lex.cs = 51; goto _test_eof
	_test_eof52:  lex.cs = 52; goto _test_eof
	_test_eof79:  lex.cs = 79; goto _test_eof
	_test_eof80:  lex.cs = 80; goto _test_eof
	_test_eof81:  lex.cs = 81; goto _test_eof
	_test_eof82:  lex.cs = 82; goto _test_eof
	_test_eof83:  lex.cs = 83; goto _test_eof
	_test_eof84:  lex.cs = 84; goto _test_eof
	_test_eof85:  lex.cs = 85; goto _test_eof
	_test_eof86:  lex.cs = 86; goto _test_eof
	_test_eof87:  lex.cs = 87; goto _test_eof
	_test_eof88:  lex.cs = 88; goto _test_eof
	_test_eof89:  lex.cs = 89; goto _test_eof
	_test_eof90:  lex.cs = 90; goto _test_eof
	_test_eof91:  lex.cs = 91; goto _test_eof
	_test_eof92:  lex.cs = 92; goto _test_eof
	_test_eof93:  lex.cs = 93; goto _test_eof
	_test_eof94:  lex.cs = 94; goto _test_eof
	_test_eof95:  lex.cs = 95; goto _test_eof
	_test_eof96:  lex.cs = 96; goto _test_eof
	_test_eof53:  lex.cs = 53; goto _test_eof
	_test_eof54:  lex.cs = 54; goto _test_eof
	_test_eof55:  lex.cs = 55; goto _test_eof
	_test_eof56:  lex.cs = 56; goto _test_eof
	_test_eof57:  lex.cs = 57; goto _test_eof
	_test_eof58:  lex.cs = 58; goto _test_eof
	_test_eof59:  lex.cs = 59; goto _test_eof
	_test_eof60:  lex.cs = 60; goto _test_eof
	_test_eof61:  lex.cs = 61; goto _test_eof
	_test_eof62:  lex.cs = 62; goto _test_eof
	_test_eof63:  lex.cs = 63; goto _test_eof
	_test_eof97:  lex.cs = 97; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 65:
			goto tr90
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
		case 66:
			goto tr92
		case 67:
			goto tr94
		case 68:
			goto tr95
		case 69:
			goto tr92
		case 70:
			goto tr92
		case 71:
			goto tr92
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
		case 72:
			goto tr99
		case 73:
			goto tr99
		case 74:
			goto tr99
		case 75:
			goto tr99
		case 76:
			goto tr99
		case 77:
			goto tr99
		case 78:
			goto tr99
		case 38:
			goto tr41
		case 39:
			goto tr41
		case 40:
			goto tr41
		case 41:
			goto tr41
		case 42:
			goto tr41
		case 43:
			goto tr41
		case 44:
			goto tr41
		case 45:
			goto tr41
		case 46:
			goto tr41
		case 47:
			goto tr41
		case 48:
			goto tr41
		case 49:
			goto tr41
		case 50:
			goto tr41
		case 51:
			goto tr41
		case 52:
			goto tr41
		case 79:
			goto tr99
		case 80:
			goto tr99
		case 81:
			goto tr99
		case 82:
			goto tr99
		case 83:
			goto tr99
		case 84:
			goto tr99
		case 85:
			goto tr99
		case 86:
			goto tr99
		case 87:
			goto tr99
		case 88:
			goto tr99
		case 89:
			goto tr99
		case 90:
			goto tr99
		case 91:
			goto tr99
		case 92:
			goto tr99
		case 93:
			goto tr99
		case 94:
			goto tr99
		case 95:
			goto tr99
		case 96:
			goto tr99
		case 53:
			goto tr41
		case 54:
			goto tr41
		case 55:
			goto tr41
		case 56:
			goto tr41
		case 57:
			goto tr41
		case 58:
			goto tr41
		case 59:
			goto tr41
		case 60:
			goto tr41
		case 61:
			goto tr41
		case 62:
			goto tr41
		case 63:
			goto tr41
		case 97:
			goto tr99
		}
	}

	_out: {}
	}

//line lexer.go.rl:134


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