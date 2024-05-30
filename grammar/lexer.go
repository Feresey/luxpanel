
//line lexer.go.rl:1
// Generated code. DO NOT EDIT

package main

import (
    "fmt"
    "strconv"
    "time"
)


//line lexer.go:15
const logparser_start int = 119
const logparser_first_final int = 119
const logparser_error int = 0

const logparser_en_main int = 119


//line lexer.go.rl:22


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
    
//line lexer.go:68
	{
	 lex.cs = logparser_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line lexer.go.rl:66
    return lex
}

func (lex *lexer) Lex(out *yySymType) int {
    eof := lex.pe
    var tok int
    
//line lexer.go:84
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
	case 119:
		goto st_case_119
	case 0:
		goto st_case_0
	case 120:
		goto st_case_120
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
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
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
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
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
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
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
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	case 141:
		goto st_case_141
	case 142:
		goto st_case_142
	case 143:
		goto st_case_143
	case 144:
		goto st_case_144
	case 145:
		goto st_case_145
	case 146:
		goto st_case_146
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
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 147:
		goto st_case_147
	case 148:
		goto st_case_148
	case 149:
		goto st_case_149
	case 150:
		goto st_case_150
	case 151:
		goto st_case_151
	case 152:
		goto st_case_152
	case 153:
		goto st_case_153
	case 154:
		goto st_case_154
	case 155:
		goto st_case_155
	case 156:
		goto st_case_156
	case 157:
		goto st_case_157
	case 158:
		goto st_case_158
	case 159:
		goto st_case_159
	case 160:
		goto st_case_160
	case 161:
		goto st_case_161
	case 162:
		goto st_case_162
	case 163:
		goto st_case_163
	case 164:
		goto st_case_164
	case 165:
		goto st_case_165
	case 166:
		goto st_case_166
	case 167:
		goto st_case_167
	case 168:
		goto st_case_168
	case 169:
		goto st_case_169
	case 170:
		goto st_case_170
	case 171:
		goto st_case_171
	case 172:
		goto st_case_172
	case 173:
		goto st_case_173
	case 174:
		goto st_case_174
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 175:
		goto st_case_175
	}
	goto st_out
tr0:
//line lexer.go.rl:137
( lex.p) = ( lex.te) - 1

	goto st119
tr8:
//line lexer.go.rl:74
 lex.te = ( lex.p)+1
{tok = COMBAT; {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr11:
//line lexer.go.rl:93
 lex.te = ( lex.p)+1
{tok = ARROW; {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr30:
//line lexer.go.rl:86
 lex.te = ( lex.p)+1
{tok = ACTUAL_GAME_TIME; {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr44:
//line lexer.go.rl:85
 lex.te = ( lex.p)+1
{tok = FINISH_REASON; {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr45:
//line lexer.go.rl:108
( lex.p) = ( lex.te) - 1
{
                // fmt.Printf("int: %q\n", lex.data[lex.ts:lex.te])
                res, err := strconv.Atoi(lex.data[lex.ts:lex.te])
                if err != nil {
                    lex.Error(err.Error())
                    {( lex.p)++;  lex.cs = 119; goto _out }
                }
                out.Int = res
                tok = INT;
                {( lex.p)++;  lex.cs = 119; goto _out }
            }
	goto st119
tr54:
//line lexer.go.rl:97
 lex.te = ( lex.p)+1
{
                tok = TIME;
                // fmt.Printf("time:%q\n", lex.data[lex.ts:lex.te])
                res, err := lex.parseTime(lex.data[lex.ts:lex.te])
                if err != nil {
                    lex.Error(err.Error())
                    {( lex.p)++;  lex.cs = 119; goto _out }
                }
                out.Time = res
                {( lex.p)++;  lex.cs = 119; goto _out }
            }
	goto st119
tr67:
//line lexer.go.rl:96
 lex.te = ( lex.p)+1
{tok = FRIENDLY_FIRE; out.Bool = true; {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr73:
//line lexer.go.rl:75
 lex.te = ( lex.p)+1
{tok = EQ_DELIM; {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr74:
//line lexer.go.rl:130
( lex.p) = ( lex.te) - 1
{
                tok = STRING;
                // fmt.Printf("name: %q\n", lex.data[lex.ts:lex.te])
                out.String = lex.data[lex.ts:lex.te]
                {( lex.p)++;  lex.cs = 119; goto _out }
            }
	goto st119
tr89:
//line lexer.go.rl:83
 lex.te = ( lex.p)+1
{tok = CONNECT_TO_GAME_SESSION_PREFIX; {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr112:
//line lexer.go.rl:84
 lex.te = ( lex.p)+1
{tok = GAMEPLAY_FINISHED; {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr123:
//line lexer.go.rl:87
 lex.te = ( lex.p)+1
{tok = LOCAL_CLIENT_TEAM; {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr124:
//line lexer.go.rl:79
 lex.te = ( lex.p)+1
{tok = int('\t'); {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr125:
//line lexer.go.rl:136
 lex.te = ( lex.p)+1
{tok = EOL; {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr126:
//line lexer.go.rl:137
 lex.te = ( lex.p)+1

	goto st119
tr128:
//line lexer.go.rl:76
 lex.te = ( lex.p)+1
{tok = int('\''); {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr129:
//line lexer.go.rl:80
 lex.te = ( lex.p)+1
{tok = int('('); {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr130:
//line lexer.go.rl:82
 lex.te = ( lex.p)+1
{tok = int(')'); {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr131:
//line lexer.go.rl:77
 lex.te = ( lex.p)+1
{tok = int(','); {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr136:
//line lexer.go.rl:78
 lex.te = ( lex.p)+1
{tok = int(';'); {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr149:
//line lexer.go.rl:81
 lex.te = ( lex.p)+1
{tok = int('|'); {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr150:
//line lexer.go.rl:137
 lex.te = ( lex.p)
( lex.p)--

	goto st119
tr152:
//line lexer.go.rl:108
 lex.te = ( lex.p)
( lex.p)--
{
                // fmt.Printf("int: %q\n", lex.data[lex.ts:lex.te])
                res, err := strconv.Atoi(lex.data[lex.ts:lex.te])
                if err != nil {
                    lex.Error(err.Error())
                    {( lex.p)++;  lex.cs = 119; goto _out }
                }
                out.Int = res
                tok = INT;
                {( lex.p)++;  lex.cs = 119; goto _out }
            }
	goto st119
tr154:
//line lexer.go.rl:119
 lex.te = ( lex.p)
( lex.p)--
{
                // fmt.Printf("float: %q\n", lex.data[lex.ts:lex.te])
                res, err := strconv.ParseFloat(lex.data[lex.ts:lex.te], 32)
                if err != nil {
                    lex.Error(err.Error())
                    {( lex.p)++;  lex.cs = 119; goto _out }
                }
                out.Float = float32(res)
                tok = FLOAT;
                {( lex.p)++;  lex.cs = 119; goto _out }
            }
	goto st119
tr155:
//line NONE:1
	switch  lex.act {
	case 15:
	{( lex.p) = ( lex.te) - 1
tok = START; {( lex.p)++;  lex.cs = 119; goto _out }}
	case 16:
	{( lex.p) = ( lex.te) - 1
tok = DAMAGE; {( lex.p)++;  lex.cs = 119; goto _out }}
	case 17:
	{( lex.p) = ( lex.te) - 1
tok = KILL; {( lex.p)++;  lex.cs = 119; goto _out }}
	case 18:
	{( lex.p) = ( lex.te) - 1
tok = PARTICIPANT; {( lex.p)++;  lex.cs = 119; goto _out }}
	case 19:
	{( lex.p) = ( lex.te) - 1
tok = HEAL; {( lex.p)++;  lex.cs = 119; goto _out }}
	case 27:
	{( lex.p) = ( lex.te) - 1

                tok = STRING;
                // fmt.Printf("name: %q\n", lex.data[lex.ts:lex.te])
                out.String = lex.data[lex.ts:lex.te]
                {( lex.p)++;  lex.cs = 119; goto _out }
            }
	}
	
	goto st119
tr159:
//line lexer.go.rl:130
 lex.te = ( lex.p)
( lex.p)--
{
                tok = STRING;
                // fmt.Printf("name: %q\n", lex.data[lex.ts:lex.te])
                out.String = lex.data[lex.ts:lex.te]
                {( lex.p)++;  lex.cs = 119; goto _out }
            }
	goto st119
tr202:
//line lexer.go.rl:94
 lex.te = ( lex.p)+1
{tok = DAMAGE_HULL_START; {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
tr208:
//line lexer.go.rl:95
 lex.te = ( lex.p)+1
{tok = DAMAGE_SHIELD_START; {( lex.p)++;  lex.cs = 119; goto _out }}
	goto st119
	st119:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof119
		}
	st_case_119:
//line NONE:1
 lex.ts = ( lex.p)

//line lexer.go:678
		switch  lex.data[( lex.p)] {
		case 9:
			goto tr124
		case 10:
			goto tr125
		case 32:
			goto tr127
		case 39:
			goto tr128
		case 40:
			goto tr129
		case 41:
			goto tr130
		case 44:
			goto tr131
		case 45:
			goto st9
		case 46:
			goto st10
		case 47:
			goto tr134
		case 59:
			goto tr136
		case 60:
			goto st51
		case 61:
			goto st64
		case 67:
			goto st127
		case 68:
			goto st134
		case 71:
			goto st139
		case 72:
			goto st147
		case 75:
			goto st150
		case 80:
			goto st155
		case 83:
			goto st165
		case 95:
			goto tr134
		case 104:
			goto st169
		case 108:
			goto st170
		case 115:
			goto st175
		case 124:
			goto tr149
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 11 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr126
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr134
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr134
			}
		default:
			goto st124
		}
		goto st0
st_case_0:
	st0:
		 lex.cs = 0
		goto _out
tr127:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st120
	st120:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof120
		}
	st_case_120:
//line lexer.go:763
		if  lex.data[( lex.p)] == 67 {
			goto st1
		}
		goto tr150
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
			goto st121
		}
		goto st0
	st121:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof121
		}
	st_case_121:
		if  lex.data[( lex.p)] == 46 {
			goto st122
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st121
		}
		goto tr152
	st122:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof122
		}
	st_case_122:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st122
		}
		goto tr154
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		if  lex.data[( lex.p)] == 32 {
			goto st11
		}
		goto st0
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch  lex.data[( lex.p)] {
		case 65:
			goto st12
		case 70:
			goto st28
		}
		goto st0
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		if  lex.data[( lex.p)] == 99 {
			goto st13
		}
		goto st0
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		if  lex.data[( lex.p)] == 116 {
			goto st14
		}
		goto st0
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		if  lex.data[( lex.p)] == 117 {
			goto st15
		}
		goto st0
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		if  lex.data[( lex.p)] == 97 {
			goto st16
		}
		goto st0
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		if  lex.data[( lex.p)] == 108 {
			goto st17
		}
		goto st0
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		if  lex.data[( lex.p)] == 32 {
			goto st18
		}
		goto st0
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		if  lex.data[( lex.p)] == 103 {
			goto st19
		}
		goto st0
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		if  lex.data[( lex.p)] == 97 {
			goto st20
		}
		goto st0
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		if  lex.data[( lex.p)] == 109 {
			goto st21
		}
		goto st0
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		if  lex.data[( lex.p)] == 101 {
			goto st22
		}
		goto st0
	st22:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		if  lex.data[( lex.p)] == 32 {
			goto st23
		}
		goto st0
	st23:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		if  lex.data[( lex.p)] == 116 {
			goto st24
		}
		goto st0
	st24:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		if  lex.data[( lex.p)] == 105 {
			goto st25
		}
		goto st0
	st25:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		if  lex.data[( lex.p)] == 109 {
			goto st26
		}
		goto st0
	st26:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		if  lex.data[( lex.p)] == 101 {
			goto st27
		}
		goto st0
	st27:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		if  lex.data[( lex.p)] == 32 {
			goto tr30
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
		if  lex.data[( lex.p)] == 110 {
			goto st30
		}
		goto st0
	st30:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		if  lex.data[( lex.p)] == 105 {
			goto st31
		}
		goto st0
	st31:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		if  lex.data[( lex.p)] == 115 {
			goto st32
		}
		goto st0
	st32:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		if  lex.data[( lex.p)] == 104 {
			goto st33
		}
		goto st0
	st33:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		if  lex.data[( lex.p)] == 32 {
			goto st34
		}
		goto st0
	st34:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		if  lex.data[( lex.p)] == 114 {
			goto st35
		}
		goto st0
	st35:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		if  lex.data[( lex.p)] == 101 {
			goto st36
		}
		goto st0
	st36:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		if  lex.data[( lex.p)] == 97 {
			goto st37
		}
		goto st0
	st37:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		if  lex.data[( lex.p)] == 115 {
			goto st38
		}
		goto st0
	st38:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		if  lex.data[( lex.p)] == 111 {
			goto st39
		}
		goto st0
	st39:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		if  lex.data[( lex.p)] == 110 {
			goto st40
		}
		goto st0
	st40:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		if  lex.data[( lex.p)] == 58 {
			goto st41
		}
		goto st0
	st41:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		if  lex.data[( lex.p)] == 32 {
			goto tr44
		}
		goto st0
tr134:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.go.rl:130
 lex.act = 27;
	goto st123
tr171:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.go.rl:89
 lex.act = 16;
	goto st123
tr182:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.go.rl:92
 lex.act = 19;
	goto st123
tr187:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.go.rl:90
 lex.act = 17;
	goto st123
tr197:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.go.rl:91
 lex.act = 18;
	goto st123
tr201:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.go.rl:88
 lex.act = 15;
	goto st123
	st123:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof123
		}
	st_case_123:
//line lexer.go:1211
		if  lex.data[( lex.p)] == 95 {
			goto tr134
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr155
	st124:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof124
		}
	st_case_124:
		switch  lex.data[( lex.p)] {
		case 46:
			goto st122
		case 47:
			goto tr134
		case 95:
			goto tr134
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr156
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr152
tr156:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st125
	st125:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof125
		}
	st_case_125:
//line lexer.go:1264
		switch  lex.data[( lex.p)] {
		case 46:
			goto st122
		case 47:
			goto tr134
		case 58:
			goto st42
		case 95:
			goto tr134
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st126
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr152
	st126:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof126
		}
	st_case_126:
		switch  lex.data[( lex.p)] {
		case 46:
			goto st122
		case 47:
			goto tr134
		case 95:
			goto tr134
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st126
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr152
	st42:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st43
		}
		goto tr45
	st43:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st44
		}
		goto tr45
	st44:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		if  lex.data[( lex.p)] == 58 {
			goto st45
		}
		goto tr45
	st45:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st46
		}
		goto tr45
	st46:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st47
		}
		goto tr45
	st47:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof47
		}
	st_case_47:
		if  lex.data[( lex.p)] == 46 {
			goto st48
		}
		goto tr45
	st48:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof48
		}
	st_case_48:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st49
		}
		goto tr45
	st49:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof49
		}
	st_case_49:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st50
		}
		goto tr45
	st50:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof50
		}
	st_case_50:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr54
		}
		goto tr45
	st51:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof51
		}
	st_case_51:
		if  lex.data[( lex.p)] == 70 {
			goto st52
		}
		goto st0
	st52:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof52
		}
	st_case_52:
		if  lex.data[( lex.p)] == 114 {
			goto st53
		}
		goto st0
	st53:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof53
		}
	st_case_53:
		if  lex.data[( lex.p)] == 105 {
			goto st54
		}
		goto st0
	st54:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof54
		}
	st_case_54:
		if  lex.data[( lex.p)] == 101 {
			goto st55
		}
		goto st0
	st55:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof55
		}
	st_case_55:
		if  lex.data[( lex.p)] == 110 {
			goto st56
		}
		goto st0
	st56:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof56
		}
	st_case_56:
		if  lex.data[( lex.p)] == 100 {
			goto st57
		}
		goto st0
	st57:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof57
		}
	st_case_57:
		if  lex.data[( lex.p)] == 108 {
			goto st58
		}
		goto st0
	st58:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof58
		}
	st_case_58:
		if  lex.data[( lex.p)] == 121 {
			goto st59
		}
		goto st0
	st59:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof59
		}
	st_case_59:
		if  lex.data[( lex.p)] == 70 {
			goto st60
		}
		goto st0
	st60:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof60
		}
	st_case_60:
		if  lex.data[( lex.p)] == 105 {
			goto st61
		}
		goto st0
	st61:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof61
		}
	st_case_61:
		if  lex.data[( lex.p)] == 114 {
			goto st62
		}
		goto st0
	st62:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof62
		}
	st_case_62:
		if  lex.data[( lex.p)] == 101 {
			goto st63
		}
		goto st0
	st63:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof63
		}
	st_case_63:
		if  lex.data[( lex.p)] == 62 {
			goto tr67
		}
		goto st0
	st64:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof64
		}
	st_case_64:
		if  lex.data[( lex.p)] == 61 {
			goto st65
		}
		goto st0
	st65:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof65
		}
	st_case_65:
		if  lex.data[( lex.p)] == 61 {
			goto st66
		}
		goto st0
	st66:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof66
		}
	st_case_66:
		if  lex.data[( lex.p)] == 61 {
			goto st67
		}
		goto st0
	st67:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof67
		}
	st_case_67:
		if  lex.data[( lex.p)] == 61 {
			goto st68
		}
		goto st0
	st68:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof68
		}
	st_case_68:
		if  lex.data[( lex.p)] == 61 {
			goto st69
		}
		goto st0
	st69:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof69
		}
	st_case_69:
		if  lex.data[( lex.p)] == 61 {
			goto tr73
		}
		goto st0
	st127:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof127
		}
	st_case_127:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 111:
			goto st128
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st128:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof128
		}
	st_case_128:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 110:
			goto st129
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st129:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof129
		}
	st_case_129:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 110:
			goto st130
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st130:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof130
		}
	st_case_130:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 101:
			goto st131
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st131:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof131
		}
	st_case_131:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 99:
			goto st132
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st132:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof132
		}
	st_case_132:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 116:
			goto tr165
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
tr165:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st133
	st133:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof133
		}
	st_case_133:
//line lexer.go:1720
		switch  lex.data[( lex.p)] {
		case 32:
			goto st70
		case 95:
			goto tr134
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st70:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof70
		}
	st_case_70:
		if  lex.data[( lex.p)] == 116 {
			goto st71
		}
		goto tr74
	st71:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof71
		}
	st_case_71:
		if  lex.data[( lex.p)] == 111 {
			goto st72
		}
		goto tr74
	st72:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof72
		}
	st_case_72:
		if  lex.data[( lex.p)] == 32 {
			goto st73
		}
		goto tr74
	st73:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof73
		}
	st_case_73:
		if  lex.data[( lex.p)] == 103 {
			goto st74
		}
		goto tr74
	st74:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof74
		}
	st_case_74:
		if  lex.data[( lex.p)] == 97 {
			goto st75
		}
		goto tr74
	st75:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof75
		}
	st_case_75:
		if  lex.data[( lex.p)] == 109 {
			goto st76
		}
		goto tr74
	st76:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof76
		}
	st_case_76:
		if  lex.data[( lex.p)] == 101 {
			goto st77
		}
		goto tr74
	st77:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof77
		}
	st_case_77:
		if  lex.data[( lex.p)] == 32 {
			goto st78
		}
		goto tr74
	st78:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof78
		}
	st_case_78:
		if  lex.data[( lex.p)] == 115 {
			goto st79
		}
		goto tr74
	st79:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof79
		}
	st_case_79:
		if  lex.data[( lex.p)] == 101 {
			goto st80
		}
		goto tr74
	st80:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof80
		}
	st_case_80:
		if  lex.data[( lex.p)] == 115 {
			goto st81
		}
		goto tr74
	st81:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof81
		}
	st_case_81:
		if  lex.data[( lex.p)] == 115 {
			goto st82
		}
		goto tr74
	st82:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof82
		}
	st_case_82:
		if  lex.data[( lex.p)] == 105 {
			goto st83
		}
		goto tr74
	st83:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof83
		}
	st_case_83:
		if  lex.data[( lex.p)] == 111 {
			goto st84
		}
		goto tr74
	st84:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof84
		}
	st_case_84:
		if  lex.data[( lex.p)] == 110 {
			goto tr89
		}
		goto tr74
	st134:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof134
		}
	st_case_134:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 97:
			goto st135
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st135:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof135
		}
	st_case_135:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 109:
			goto st136
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st136:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof136
		}
	st_case_136:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 97:
			goto st137
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st137:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof137
		}
	st_case_137:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 103:
			goto st138
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st138:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof138
		}
	st_case_138:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 101:
			goto tr171
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st139:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof139
		}
	st_case_139:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 97:
			goto st140
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st140:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof140
		}
	st_case_140:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 109:
			goto st141
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st141:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof141
		}
	st_case_141:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 101:
			goto st142
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st142:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof142
		}
	st_case_142:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 112:
			goto st143
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st143:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof143
		}
	st_case_143:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 108:
			goto st144
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st144:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof144
		}
	st_case_144:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 97:
			goto st145
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st145:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof145
		}
	st_case_145:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 121:
			goto tr178
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
tr178:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st146
	st146:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof146
		}
	st_case_146:
//line lexer.go:2173
		switch  lex.data[( lex.p)] {
		case 32:
			goto st85
		case 95:
			goto tr134
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st85:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof85
		}
	st_case_85:
		if  lex.data[( lex.p)] == 102 {
			goto st86
		}
		goto tr74
	st86:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof86
		}
	st_case_86:
		if  lex.data[( lex.p)] == 105 {
			goto st87
		}
		goto tr74
	st87:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof87
		}
	st_case_87:
		if  lex.data[( lex.p)] == 110 {
			goto st88
		}
		goto tr74
	st88:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof88
		}
	st_case_88:
		if  lex.data[( lex.p)] == 105 {
			goto st89
		}
		goto tr74
	st89:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof89
		}
	st_case_89:
		if  lex.data[( lex.p)] == 115 {
			goto st90
		}
		goto tr74
	st90:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof90
		}
	st_case_90:
		if  lex.data[( lex.p)] == 104 {
			goto st91
		}
		goto tr74
	st91:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof91
		}
	st_case_91:
		if  lex.data[( lex.p)] == 101 {
			goto st92
		}
		goto tr74
	st92:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof92
		}
	st_case_92:
		if  lex.data[( lex.p)] == 100 {
			goto st93
		}
		goto tr74
	st93:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof93
		}
	st_case_93:
		if  lex.data[( lex.p)] == 46 {
			goto st94
		}
		goto tr74
	st94:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof94
		}
	st_case_94:
		if  lex.data[( lex.p)] == 32 {
			goto st95
		}
		goto tr74
	st95:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof95
		}
	st_case_95:
		if  lex.data[( lex.p)] == 87 {
			goto st96
		}
		goto tr74
	st96:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof96
		}
	st_case_96:
		if  lex.data[( lex.p)] == 105 {
			goto st97
		}
		goto tr74
	st97:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof97
		}
	st_case_97:
		if  lex.data[( lex.p)] == 110 {
			goto st98
		}
		goto tr74
	st98:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof98
		}
	st_case_98:
		if  lex.data[( lex.p)] == 110 {
			goto st99
		}
		goto tr74
	st99:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof99
		}
	st_case_99:
		if  lex.data[( lex.p)] == 101 {
			goto st100
		}
		goto tr74
	st100:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof100
		}
	st_case_100:
		if  lex.data[( lex.p)] == 114 {
			goto st101
		}
		goto tr74
	st101:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof101
		}
	st_case_101:
		if  lex.data[( lex.p)] == 32 {
			goto st102
		}
		goto tr74
	st102:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof102
		}
	st_case_102:
		if  lex.data[( lex.p)] == 116 {
			goto st103
		}
		goto tr74
	st103:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof103
		}
	st_case_103:
		if  lex.data[( lex.p)] == 101 {
			goto st104
		}
		goto tr74
	st104:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof104
		}
	st_case_104:
		if  lex.data[( lex.p)] == 97 {
			goto st105
		}
		goto tr74
	st105:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof105
		}
	st_case_105:
		if  lex.data[( lex.p)] == 109 {
			goto st106
		}
		goto tr74
	st106:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof106
		}
	st_case_106:
		if  lex.data[( lex.p)] == 58 {
			goto st107
		}
		goto tr74
	st107:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof107
		}
	st_case_107:
		if  lex.data[( lex.p)] == 32 {
			goto tr112
		}
		goto tr74
	st147:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof147
		}
	st_case_147:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 101:
			goto st148
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st148:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof148
		}
	st_case_148:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 97:
			goto st149
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st149:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof149
		}
	st_case_149:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 108:
			goto tr182
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st150:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof150
		}
	st_case_150:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 105:
			goto st151
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st151:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof151
		}
	st_case_151:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 108:
			goto st152
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st152:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof152
		}
	st_case_152:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 108:
			goto st153
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st153:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof153
		}
	st_case_153:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 101:
			goto st154
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st154:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof154
		}
	st_case_154:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 100:
			goto tr187
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st155:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof155
		}
	st_case_155:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 97:
			goto st156
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st156:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof156
		}
	st_case_156:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 114:
			goto st157
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st157:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof157
		}
	st_case_157:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 116:
			goto st158
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st158:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof158
		}
	st_case_158:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 105:
			goto st159
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st159:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof159
		}
	st_case_159:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 99:
			goto st160
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st160:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof160
		}
	st_case_160:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 105:
			goto st161
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st161:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof161
		}
	st_case_161:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 112:
			goto st162
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st162:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof162
		}
	st_case_162:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 97:
			goto st163
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st163:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof163
		}
	st_case_163:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 110:
			goto st164
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st164:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof164
		}
	st_case_164:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 116:
			goto tr197
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st165:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof165
		}
	st_case_165:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 116:
			goto st166
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st166:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof166
		}
	st_case_166:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 97:
			goto st167
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st167:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof167
		}
	st_case_167:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 114:
			goto st168
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st168:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof168
		}
	st_case_168:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 116:
			goto tr201
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st169:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof169
		}
	st_case_169:
		switch  lex.data[( lex.p)] {
		case 58:
			goto tr202
		case 95:
			goto tr134
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st170:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof170
		}
	st_case_170:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 111:
			goto st171
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st171:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof171
		}
	st_case_171:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 99:
			goto st172
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st172:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof172
		}
	st_case_172:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 97:
			goto st173
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st173:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof173
		}
	st_case_173:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr134
		case 108:
			goto tr206
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
tr206:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st174
	st174:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof174
		}
	st_case_174:
//line lexer.go:3058
		switch  lex.data[( lex.p)] {
		case 32:
			goto st108
		case 95:
			goto tr134
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st108:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof108
		}
	st_case_108:
		if  lex.data[( lex.p)] == 99 {
			goto st109
		}
		goto tr74
	st109:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof109
		}
	st_case_109:
		if  lex.data[( lex.p)] == 108 {
			goto st110
		}
		goto tr74
	st110:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof110
		}
	st_case_110:
		if  lex.data[( lex.p)] == 105 {
			goto st111
		}
		goto tr74
	st111:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof111
		}
	st_case_111:
		if  lex.data[( lex.p)] == 101 {
			goto st112
		}
		goto tr74
	st112:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof112
		}
	st_case_112:
		if  lex.data[( lex.p)] == 110 {
			goto st113
		}
		goto tr74
	st113:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof113
		}
	st_case_113:
		if  lex.data[( lex.p)] == 116 {
			goto st114
		}
		goto tr74
	st114:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof114
		}
	st_case_114:
		if  lex.data[( lex.p)] == 32 {
			goto st115
		}
		goto tr74
	st115:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof115
		}
	st_case_115:
		if  lex.data[( lex.p)] == 116 {
			goto st116
		}
		goto tr74
	st116:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof116
		}
	st_case_116:
		if  lex.data[( lex.p)] == 101 {
			goto st117
		}
		goto tr74
	st117:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof117
		}
	st_case_117:
		if  lex.data[( lex.p)] == 97 {
			goto st118
		}
		goto tr74
	st118:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof118
		}
	st_case_118:
		if  lex.data[( lex.p)] == 109 {
			goto tr123
		}
		goto tr74
	st175:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof175
		}
	st_case_175:
		switch  lex.data[( lex.p)] {
		case 58:
			goto tr208
		case 95:
			goto tr134
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr134
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr134
			}
		default:
			goto tr134
		}
		goto tr159
	st_out:
	_test_eof119:  lex.cs = 119; goto _test_eof
	_test_eof120:  lex.cs = 120; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof121:  lex.cs = 121; goto _test_eof
	_test_eof122:  lex.cs = 122; goto _test_eof
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
	_test_eof38:  lex.cs = 38; goto _test_eof
	_test_eof39:  lex.cs = 39; goto _test_eof
	_test_eof40:  lex.cs = 40; goto _test_eof
	_test_eof41:  lex.cs = 41; goto _test_eof
	_test_eof123:  lex.cs = 123; goto _test_eof
	_test_eof124:  lex.cs = 124; goto _test_eof
	_test_eof125:  lex.cs = 125; goto _test_eof
	_test_eof126:  lex.cs = 126; goto _test_eof
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
	_test_eof64:  lex.cs = 64; goto _test_eof
	_test_eof65:  lex.cs = 65; goto _test_eof
	_test_eof66:  lex.cs = 66; goto _test_eof
	_test_eof67:  lex.cs = 67; goto _test_eof
	_test_eof68:  lex.cs = 68; goto _test_eof
	_test_eof69:  lex.cs = 69; goto _test_eof
	_test_eof127:  lex.cs = 127; goto _test_eof
	_test_eof128:  lex.cs = 128; goto _test_eof
	_test_eof129:  lex.cs = 129; goto _test_eof
	_test_eof130:  lex.cs = 130; goto _test_eof
	_test_eof131:  lex.cs = 131; goto _test_eof
	_test_eof132:  lex.cs = 132; goto _test_eof
	_test_eof133:  lex.cs = 133; goto _test_eof
	_test_eof70:  lex.cs = 70; goto _test_eof
	_test_eof71:  lex.cs = 71; goto _test_eof
	_test_eof72:  lex.cs = 72; goto _test_eof
	_test_eof73:  lex.cs = 73; goto _test_eof
	_test_eof74:  lex.cs = 74; goto _test_eof
	_test_eof75:  lex.cs = 75; goto _test_eof
	_test_eof76:  lex.cs = 76; goto _test_eof
	_test_eof77:  lex.cs = 77; goto _test_eof
	_test_eof78:  lex.cs = 78; goto _test_eof
	_test_eof79:  lex.cs = 79; goto _test_eof
	_test_eof80:  lex.cs = 80; goto _test_eof
	_test_eof81:  lex.cs = 81; goto _test_eof
	_test_eof82:  lex.cs = 82; goto _test_eof
	_test_eof83:  lex.cs = 83; goto _test_eof
	_test_eof84:  lex.cs = 84; goto _test_eof
	_test_eof134:  lex.cs = 134; goto _test_eof
	_test_eof135:  lex.cs = 135; goto _test_eof
	_test_eof136:  lex.cs = 136; goto _test_eof
	_test_eof137:  lex.cs = 137; goto _test_eof
	_test_eof138:  lex.cs = 138; goto _test_eof
	_test_eof139:  lex.cs = 139; goto _test_eof
	_test_eof140:  lex.cs = 140; goto _test_eof
	_test_eof141:  lex.cs = 141; goto _test_eof
	_test_eof142:  lex.cs = 142; goto _test_eof
	_test_eof143:  lex.cs = 143; goto _test_eof
	_test_eof144:  lex.cs = 144; goto _test_eof
	_test_eof145:  lex.cs = 145; goto _test_eof
	_test_eof146:  lex.cs = 146; goto _test_eof
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
	_test_eof97:  lex.cs = 97; goto _test_eof
	_test_eof98:  lex.cs = 98; goto _test_eof
	_test_eof99:  lex.cs = 99; goto _test_eof
	_test_eof100:  lex.cs = 100; goto _test_eof
	_test_eof101:  lex.cs = 101; goto _test_eof
	_test_eof102:  lex.cs = 102; goto _test_eof
	_test_eof103:  lex.cs = 103; goto _test_eof
	_test_eof104:  lex.cs = 104; goto _test_eof
	_test_eof105:  lex.cs = 105; goto _test_eof
	_test_eof106:  lex.cs = 106; goto _test_eof
	_test_eof107:  lex.cs = 107; goto _test_eof
	_test_eof147:  lex.cs = 147; goto _test_eof
	_test_eof148:  lex.cs = 148; goto _test_eof
	_test_eof149:  lex.cs = 149; goto _test_eof
	_test_eof150:  lex.cs = 150; goto _test_eof
	_test_eof151:  lex.cs = 151; goto _test_eof
	_test_eof152:  lex.cs = 152; goto _test_eof
	_test_eof153:  lex.cs = 153; goto _test_eof
	_test_eof154:  lex.cs = 154; goto _test_eof
	_test_eof155:  lex.cs = 155; goto _test_eof
	_test_eof156:  lex.cs = 156; goto _test_eof
	_test_eof157:  lex.cs = 157; goto _test_eof
	_test_eof158:  lex.cs = 158; goto _test_eof
	_test_eof159:  lex.cs = 159; goto _test_eof
	_test_eof160:  lex.cs = 160; goto _test_eof
	_test_eof161:  lex.cs = 161; goto _test_eof
	_test_eof162:  lex.cs = 162; goto _test_eof
	_test_eof163:  lex.cs = 163; goto _test_eof
	_test_eof164:  lex.cs = 164; goto _test_eof
	_test_eof165:  lex.cs = 165; goto _test_eof
	_test_eof166:  lex.cs = 166; goto _test_eof
	_test_eof167:  lex.cs = 167; goto _test_eof
	_test_eof168:  lex.cs = 168; goto _test_eof
	_test_eof169:  lex.cs = 169; goto _test_eof
	_test_eof170:  lex.cs = 170; goto _test_eof
	_test_eof171:  lex.cs = 171; goto _test_eof
	_test_eof172:  lex.cs = 172; goto _test_eof
	_test_eof173:  lex.cs = 173; goto _test_eof
	_test_eof174:  lex.cs = 174; goto _test_eof
	_test_eof108:  lex.cs = 108; goto _test_eof
	_test_eof109:  lex.cs = 109; goto _test_eof
	_test_eof110:  lex.cs = 110; goto _test_eof
	_test_eof111:  lex.cs = 111; goto _test_eof
	_test_eof112:  lex.cs = 112; goto _test_eof
	_test_eof113:  lex.cs = 113; goto _test_eof
	_test_eof114:  lex.cs = 114; goto _test_eof
	_test_eof115:  lex.cs = 115; goto _test_eof
	_test_eof116:  lex.cs = 116; goto _test_eof
	_test_eof117:  lex.cs = 117; goto _test_eof
	_test_eof118:  lex.cs = 118; goto _test_eof
	_test_eof175:  lex.cs = 175; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 120:
			goto tr150
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
		case 121:
			goto tr152
		case 122:
			goto tr154
		case 123:
			goto tr155
		case 124:
			goto tr152
		case 125:
			goto tr152
		case 126:
			goto tr152
		case 42:
			goto tr45
		case 43:
			goto tr45
		case 44:
			goto tr45
		case 45:
			goto tr45
		case 46:
			goto tr45
		case 47:
			goto tr45
		case 48:
			goto tr45
		case 49:
			goto tr45
		case 50:
			goto tr45
		case 127:
			goto tr159
		case 128:
			goto tr159
		case 129:
			goto tr159
		case 130:
			goto tr159
		case 131:
			goto tr159
		case 132:
			goto tr159
		case 133:
			goto tr159
		case 70:
			goto tr74
		case 71:
			goto tr74
		case 72:
			goto tr74
		case 73:
			goto tr74
		case 74:
			goto tr74
		case 75:
			goto tr74
		case 76:
			goto tr74
		case 77:
			goto tr74
		case 78:
			goto tr74
		case 79:
			goto tr74
		case 80:
			goto tr74
		case 81:
			goto tr74
		case 82:
			goto tr74
		case 83:
			goto tr74
		case 84:
			goto tr74
		case 134:
			goto tr159
		case 135:
			goto tr159
		case 136:
			goto tr159
		case 137:
			goto tr159
		case 138:
			goto tr159
		case 139:
			goto tr159
		case 140:
			goto tr159
		case 141:
			goto tr159
		case 142:
			goto tr159
		case 143:
			goto tr159
		case 144:
			goto tr159
		case 145:
			goto tr159
		case 146:
			goto tr159
		case 85:
			goto tr74
		case 86:
			goto tr74
		case 87:
			goto tr74
		case 88:
			goto tr74
		case 89:
			goto tr74
		case 90:
			goto tr74
		case 91:
			goto tr74
		case 92:
			goto tr74
		case 93:
			goto tr74
		case 94:
			goto tr74
		case 95:
			goto tr74
		case 96:
			goto tr74
		case 97:
			goto tr74
		case 98:
			goto tr74
		case 99:
			goto tr74
		case 100:
			goto tr74
		case 101:
			goto tr74
		case 102:
			goto tr74
		case 103:
			goto tr74
		case 104:
			goto tr74
		case 105:
			goto tr74
		case 106:
			goto tr74
		case 107:
			goto tr74
		case 147:
			goto tr159
		case 148:
			goto tr159
		case 149:
			goto tr159
		case 150:
			goto tr159
		case 151:
			goto tr159
		case 152:
			goto tr159
		case 153:
			goto tr159
		case 154:
			goto tr159
		case 155:
			goto tr159
		case 156:
			goto tr159
		case 157:
			goto tr159
		case 158:
			goto tr159
		case 159:
			goto tr159
		case 160:
			goto tr159
		case 161:
			goto tr159
		case 162:
			goto tr159
		case 163:
			goto tr159
		case 164:
			goto tr159
		case 165:
			goto tr159
		case 166:
			goto tr159
		case 167:
			goto tr159
		case 168:
			goto tr159
		case 169:
			goto tr159
		case 170:
			goto tr159
		case 171:
			goto tr159
		case 172:
			goto tr159
		case 173:
			goto tr159
		case 174:
			goto tr159
		case 108:
			goto tr74
		case 109:
			goto tr74
		case 110:
			goto tr74
		case 111:
			goto tr74
		case 112:
			goto tr74
		case 113:
			goto tr74
		case 114:
			goto tr74
		case 115:
			goto tr74
		case 116:
			goto tr74
		case 117:
			goto tr74
		case 118:
			goto tr74
		case 175:
			goto tr159
		}
	}

	_out: {}
	}

//line lexer.go.rl:141


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