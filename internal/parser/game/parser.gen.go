
//line internal/parser/game/ext/parser.rl:1
// Generated code. DO NOT EDIT

package game

import (
	"fmt"
	"strconv"
	"errors"

	"github.com/Feresey/luxpanel/internal/parser/common"
)


//line internal/parser/game/parser.gen.go:15
const logparser_start int = 1
const logparser_first_final int = 175
const logparser_error int = 0

const logparser_en_main int = 1


//line internal/parser/game/ext/parser.rl:21


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
	tokens []Token
	errors []error
	state
}


//line internal/parser/game/ext/parser.rl:134


func (t *Tokenizer) Parse(data string) ([]Token, error) {
	var parseErr error
	var temp YaccSymType

	t.state = state{}
	
//line internal/parser/game/parser.gen.go:78
	{
	 t.cs = logparser_start
	}

//line internal/parser/game/ext/parser.rl:142

	t.prev = 0
	t.pe = len(data)
	t.data = data

	t.tokens = t.tokens[:0]
	t.errors = nil

	
//line internal/parser/game/parser.gen.go:91
	{
	if ( t.p) == ( t.pe) {
		goto _test_eof
	}
	switch  t.cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
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
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 175:
		goto st_case_175
	case 176:
		goto st_case_176
	case 177:
		goto st_case_177
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
	case 178:
		goto st_case_178
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
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
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
	case 179:
		goto st_case_179
	case 180:
		goto st_case_180
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
	}
	goto st_out
	st_case_1:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr0
		}
		goto st0
st_case_0:
	st0:
		 t.cs = 0
		goto _out
tr0:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st2
	st2:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof2
		}
	st_case_2:
//line internal/parser/game/parser.gen.go:484
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st3
		}
		goto st0
	st3:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof3
		}
	st_case_3:
		if ( t.data)[( t.p)] == 58 {
			goto st4
		}
		goto st0
	st4:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof4
		}
	st_case_4:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st5
		}
		goto st0
	st5:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof5
		}
	st_case_5:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st6
		}
		goto st0
	st6:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof6
		}
	st_case_6:
		if ( t.data)[( t.p)] == 58 {
			goto st7
		}
		goto st0
	st7:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof7
		}
	st_case_7:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st8
		}
		goto st0
	st8:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof8
		}
	st_case_8:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st9
		}
		goto st0
	st9:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof9
		}
	st_case_9:
		if ( t.data)[( t.p)] == 46 {
			goto st10
		}
		goto st0
	st10:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st11
		}
		goto st0
	st11:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof11
		}
	st_case_11:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st12
		}
		goto st0
	st12:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof12
		}
	st_case_12:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st13
		}
		goto st0
	st13:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof13
		}
	st_case_13:
		if ( t.data)[( t.p)] == 32 {
			goto tr13
		}
		goto st0
tr13:
//line internal/parser/game/ext/parser.rl:110
t.tokval(newAnyVal(TIME, strTok(t.data[t.prev:t.p])))
	goto st14
	st14:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof14
		}
	st_case_14:
//line internal/parser/game/parser.gen.go:597
		if ( t.data)[( t.p)] == 32 {
			goto tr15
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto tr14
		}
		goto st0
tr14:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st15
	st15:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof15
		}
	st_case_15:
//line internal/parser/game/parser.gen.go:619
		switch ( t.data)[( t.p)] {
		case 32:
			goto st15
		case 124:
			goto tr17
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st15
		}
		goto st0
tr17:
//line internal/parser/game/ext/parser.rl:111
 t.tok(GAME) 
	goto st16
	st16:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof16
		}
	st_case_16:
//line internal/parser/game/parser.gen.go:639
		if ( t.data)[( t.p)] == 32 {
			goto st17
		}
		goto st0
	st17:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch ( t.data)[( t.p)] {
		case 32:
			goto st17
		case 99:
			goto tr19
		}
		goto st0
tr19:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st18
	st18:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof18
		}
	st_case_18:
//line internal/parser/game/parser.gen.go:670
		if ( t.data)[( t.p)] == 108 {
			goto st19
		}
		goto st0
	st19:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof19
		}
	st_case_19:
		if ( t.data)[( t.p)] == 105 {
			goto st20
		}
		goto st0
	st20:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof20
		}
	st_case_20:
		if ( t.data)[( t.p)] == 101 {
			goto st21
		}
		goto st0
	st21:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof21
		}
	st_case_21:
		if ( t.data)[( t.p)] == 110 {
			goto st22
		}
		goto st0
	st22:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof22
		}
	st_case_22:
		if ( t.data)[( t.p)] == 116 {
			goto st23
		}
		goto st0
	st23:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof23
		}
	st_case_23:
		if ( t.data)[( t.p)] == 58 {
			goto st24
		}
		goto st0
	st24:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof24
		}
	st_case_24:
		if ( t.data)[( t.p)] == 32 {
			goto st25
		}
		goto st0
	st25:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch ( t.data)[( t.p)] {
		case 65:
			goto st26
		case 99:
			goto st81
		case 112:
			goto st155
		}
		goto st0
	st26:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof26
		}
	st_case_26:
		if ( t.data)[( t.p)] == 68 {
			goto st27
		}
		goto st0
	st27:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof27
		}
	st_case_27:
		if ( t.data)[( t.p)] == 68 {
			goto st28
		}
		goto st0
	st28:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof28
		}
	st_case_28:
		if ( t.data)[( t.p)] == 95 {
			goto st29
		}
		goto st0
	st29:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof29
		}
	st_case_29:
		if ( t.data)[( t.p)] == 80 {
			goto st30
		}
		goto st0
	st30:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof30
		}
	st_case_30:
		if ( t.data)[( t.p)] == 76 {
			goto st31
		}
		goto st0
	st31:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof31
		}
	st_case_31:
		if ( t.data)[( t.p)] == 65 {
			goto st32
		}
		goto st0
	st32:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof32
		}
	st_case_32:
		if ( t.data)[( t.p)] == 89 {
			goto st33
		}
		goto st0
	st33:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof33
		}
	st_case_33:
		if ( t.data)[( t.p)] == 69 {
			goto st34
		}
		goto st0
	st34:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof34
		}
	st_case_34:
		if ( t.data)[( t.p)] == 82 {
			goto st35
		}
		goto st0
	st35:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof35
		}
	st_case_35:
		if ( t.data)[( t.p)] == 32 {
			goto st36
		}
		goto st0
	st36:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof36
		}
	st_case_36:
		if ( t.data)[( t.p)] == 45 {
			goto tr40
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr41
		}
		goto st0
tr40:
//line internal/parser/game/ext/parser.rl:119
t.tok(CLIENT_ADD_PLAYER)
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st37
	st37:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof37
		}
	st_case_37:
//line internal/parser/game/parser.gen.go:861
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st38
		}
		goto st0
tr41:
//line internal/parser/game/ext/parser.rl:119
t.tok(CLIENT_ADD_PLAYER)
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st38
	st38:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof38
		}
	st_case_38:
//line internal/parser/game/parser.gen.go:882
		if ( t.data)[( t.p)] == 32 {
			goto tr43
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st38
		}
		goto st0
tr43:
//line internal/parser/game/ext/parser.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(common.ParseTokenError{TokType: "INT", Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 39; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st39
	st39:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof39
		}
	st_case_39:
//line internal/parser/game/parser.gen.go:906
		if ( t.data)[( t.p)] == 40 {
			goto st40
		}
		goto st0
	st40:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr45
		case 95:
			goto tr46
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr46
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr46
			}
		default:
			goto tr46
		}
		goto st0
tr45:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st41
	st41:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof41
		}
	st_case_41:
//line internal/parser/game/parser.gen.go:949
		if ( t.data)[( t.p)] == 95 {
			goto st42
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st42
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		goto st0
	st42:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st43
		case 95:
			goto st42
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st42
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		goto st0
	st43:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr49
		case 44:
			goto tr50
		}
		goto st0
tr49:
//line internal/parser/game/ext/parser.rl:76

		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st44
	st44:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof44
		}
	st_case_44:
//line internal/parser/game/parser.gen.go:1013
		if ( t.data)[( t.p)] == 91 {
			goto st45
		}
		goto st0
	st45:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr52
		case 93:
			goto st49
		case 95:
			goto tr53
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr53
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr53
			}
		default:
			goto tr53
		}
		goto st0
tr52:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st46
	st46:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof46
		}
	st_case_46:
//line internal/parser/game/parser.gen.go:1058
		if ( t.data)[( t.p)] == 95 {
			goto st47
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st47
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st47:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st48
		case 95:
			goto st47
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st47
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st48:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof48
		}
	st_case_48:
		if ( t.data)[( t.p)] == 93 {
			goto tr57
		}
		goto st0
tr57:
//line internal/parser/game/ext/parser.rl:76

		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st49
	st49:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof49
		}
	st_case_49:
//line internal/parser/game/parser.gen.go:1119
		if ( t.data)[( t.p)] == 44 {
			goto st50
		}
		goto st0
tr50:
//line internal/parser/game/ext/parser.rl:76

		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st50
	st50:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof50
		}
	st_case_50:
//line internal/parser/game/parser.gen.go:1135
		if ( t.data)[( t.p)] == 32 {
			goto st51
		}
		goto st0
	st51:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof51
		}
	st_case_51:
		if ( t.data)[( t.p)] == 45 {
			goto tr60
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr61
		}
		goto st0
tr60:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st52
	st52:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof52
		}
	st_case_52:
//line internal/parser/game/parser.gen.go:1166
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st53
		}
		goto st0
tr61:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st53
	st53:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof53
		}
	st_case_53:
//line internal/parser/game/parser.gen.go:1185
		if ( t.data)[( t.p)] == 41 {
			goto tr63
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st53
		}
		goto st0
tr63:
//line internal/parser/game/ext/parser.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(common.ParseTokenError{TokType: "INT", Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 54; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st54
	st54:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof54
		}
	st_case_54:
//line internal/parser/game/parser.gen.go:1209
		if ( t.data)[( t.p)] == 32 {
			goto st55
		}
		goto st0
	st55:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof55
		}
	st_case_55:
		if ( t.data)[( t.p)] == 115 {
			goto st56
		}
		goto st0
	st56:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof56
		}
	st_case_56:
		if ( t.data)[( t.p)] == 116 {
			goto st57
		}
		goto st0
	st57:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof57
		}
	st_case_57:
		if ( t.data)[( t.p)] == 97 {
			goto st58
		}
		goto st0
	st58:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof58
		}
	st_case_58:
		if ( t.data)[( t.p)] == 116 {
			goto st59
		}
		goto st0
	st59:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof59
		}
	st_case_59:
		if ( t.data)[( t.p)] == 117 {
			goto st60
		}
		goto st0
	st60:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof60
		}
	st_case_60:
		if ( t.data)[( t.p)] == 115 {
			goto st61
		}
		goto st0
	st61:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof61
		}
	st_case_61:
		if ( t.data)[( t.p)] == 32 {
			goto st62
		}
		goto st0
	st62:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof62
		}
	st_case_62:
		if ( t.data)[( t.p)] == 45 {
			goto tr72
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr73
		}
		goto st0
tr72:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st63
	st63:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof63
		}
	st_case_63:
//line internal/parser/game/parser.gen.go:1303
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st64
		}
		goto st0
tr73:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st64
	st64:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof64
		}
	st_case_64:
//line internal/parser/game/parser.gen.go:1322
		if ( t.data)[( t.p)] == 32 {
			goto tr75
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st64
		}
		goto st0
tr75:
//line internal/parser/game/ext/parser.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(common.ParseTokenError{TokType: "INT", Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 65; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st65
	st65:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof65
		}
	st_case_65:
//line internal/parser/game/parser.gen.go:1346
		if ( t.data)[( t.p)] == 116 {
			goto st66
		}
		goto st0
	st66:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof66
		}
	st_case_66:
		if ( t.data)[( t.p)] == 101 {
			goto st67
		}
		goto st0
	st67:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof67
		}
	st_case_67:
		if ( t.data)[( t.p)] == 97 {
			goto st68
		}
		goto st0
	st68:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof68
		}
	st_case_68:
		if ( t.data)[( t.p)] == 109 {
			goto st69
		}
		goto st0
	st69:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof69
		}
	st_case_69:
		if ( t.data)[( t.p)] == 32 {
			goto st70
		}
		goto st0
	st70:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof70
		}
	st_case_70:
		if ( t.data)[( t.p)] == 45 {
			goto tr81
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr82
		}
		goto st0
tr81:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st71
	st71:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof71
		}
	st_case_71:
//line internal/parser/game/parser.gen.go:1413
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st175
		}
		goto st0
tr82:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st175
	st175:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof175
		}
	st_case_175:
//line internal/parser/game/parser.gen.go:1432
		if ( t.data)[( t.p)] == 32 {
			goto tr192
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st175
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr191
		}
		goto st0
tr191:
//line internal/parser/game/ext/parser.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(common.ParseTokenError{TokType: "INT", Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 176; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st176
tr194:
//line internal/parser/game/ext/parser.rl:76

		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st176
tr190:
//line internal/parser/game/ext/parser.rl:121
t.tok(LEAVE)
	goto st176
	st176:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof176
		}
	st_case_176:
//line internal/parser/game/parser.gen.go:1471
		if ( t.data)[( t.p)] == 32 {
			goto st176
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st176
		}
		goto st0
tr192:
//line internal/parser/game/ext/parser.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(common.ParseTokenError{TokType: "INT", Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 177; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st177
	st177:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof177
		}
	st_case_177:
//line internal/parser/game/parser.gen.go:1495
		switch ( t.data)[( t.p)] {
		case 32:
			goto st176
		case 103:
			goto st72
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto st176
		}
		goto st0
	st72:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof72
		}
	st_case_72:
		if ( t.data)[( t.p)] == 114 {
			goto st73
		}
		goto st0
	st73:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof73
		}
	st_case_73:
		if ( t.data)[( t.p)] == 111 {
			goto st74
		}
		goto st0
	st74:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof74
		}
	st_case_74:
		if ( t.data)[( t.p)] == 117 {
			goto st75
		}
		goto st0
	st75:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof75
		}
	st_case_75:
		if ( t.data)[( t.p)] == 112 {
			goto st76
		}
		goto st0
	st76:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof76
		}
	st_case_76:
		if ( t.data)[( t.p)] == 32 {
			goto st77
		}
		goto st0
	st77:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof77
		}
	st_case_77:
		if ( t.data)[( t.p)] == 45 {
			goto tr89
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr90
		}
		goto st0
tr89:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st78
	st78:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof78
		}
	st_case_78:
//line internal/parser/game/parser.gen.go:1577
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st178
		}
		goto st0
tr90:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st178
	st178:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof178
		}
	st_case_178:
//line internal/parser/game/parser.gen.go:1596
		if ( t.data)[( t.p)] == 32 {
			goto tr191
		}
		switch {
		case ( t.data)[( t.p)] > 13:
			if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st178
			}
		case ( t.data)[( t.p)] >= 9:
			goto tr191
		}
		goto st0
tr53:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st79
	st79:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof79
		}
	st_case_79:
//line internal/parser/game/parser.gen.go:1623
		switch ( t.data)[( t.p)] {
		case 93:
			goto tr57
		case 95:
			goto st79
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st79
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st79
			}
		default:
			goto st79
		}
		goto st0
tr46:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st80
	st80:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof80
		}
	st_case_80:
//line internal/parser/game/parser.gen.go:1657
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr49
		case 44:
			goto tr50
		case 95:
			goto st80
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st80
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st80
			}
		default:
			goto st80
		}
		goto st0
	st81:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof81
		}
	st_case_81:
		if ( t.data)[( t.p)] == 111 {
			goto st82
		}
		goto st0
	st82:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof82
		}
	st_case_82:
		if ( t.data)[( t.p)] == 110 {
			goto st83
		}
		goto st0
	st83:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof83
		}
	st_case_83:
		if ( t.data)[( t.p)] == 110 {
			goto st84
		}
		goto st0
	st84:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof84
		}
	st_case_84:
		if ( t.data)[( t.p)] == 101 {
			goto st85
		}
		goto st0
	st85:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof85
		}
	st_case_85:
		if ( t.data)[( t.p)] == 99 {
			goto st86
		}
		goto st0
	st86:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof86
		}
	st_case_86:
		if ( t.data)[( t.p)] == 116 {
			goto st87
		}
		goto st0
	st87:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof87
		}
	st_case_87:
		switch ( t.data)[( t.p)] {
		case 101:
			goto st88
		case 105:
			goto st141
		}
		goto st0
	st88:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof88
		}
	st_case_88:
		if ( t.data)[( t.p)] == 100 {
			goto st89
		}
		goto st0
	st89:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof89
		}
	st_case_89:
		if ( t.data)[( t.p)] == 32 {
			goto st90
		}
		goto st0
	st90:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof90
		}
	st_case_90:
		if ( t.data)[( t.p)] == 116 {
			goto st91
		}
		goto st0
	st91:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof91
		}
	st_case_91:
		if ( t.data)[( t.p)] == 111 {
			goto st92
		}
		goto st0
	st92:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof92
		}
	st_case_92:
		if ( t.data)[( t.p)] == 32 {
			goto st93
		}
		goto st0
	st93:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof93
		}
	st_case_93:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr107
		}
		goto st0
tr107:
//line internal/parser/game/ext/parser.rl:122
t.tok(CLIENT_CONNECTED)
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st94
	st94:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof94
		}
	st_case_94:
//line internal/parser/game/parser.gen.go:1815
		if ( t.data)[( t.p)] == 46 {
			goto st95
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st139
		}
		goto st0
	st95:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof95
		}
	st_case_95:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st96
		}
		goto st0
	st96:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof96
		}
	st_case_96:
		if ( t.data)[( t.p)] == 46 {
			goto st97
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st137
		}
		goto st0
	st97:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof97
		}
	st_case_97:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st98
		}
		goto st0
	st98:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof98
		}
	st_case_98:
		if ( t.data)[( t.p)] == 46 {
			goto st99
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st135
		}
		goto st0
	st99:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof99
		}
	st_case_99:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st100
		}
		goto st0
	st100:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof100
		}
	st_case_100:
		if ( t.data)[( t.p)] == 124 {
			goto st103
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st101
		}
		goto st0
	st101:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof101
		}
	st_case_101:
		if ( t.data)[( t.p)] == 124 {
			goto st103
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st102
		}
		goto st0
	st102:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof102
		}
	st_case_102:
		if ( t.data)[( t.p)] == 124 {
			goto st103
		}
		goto st0
	st103:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof103
		}
	st_case_103:
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st104
		}
		goto st0
	st104:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof104
		}
	st_case_104:
		if ( t.data)[( t.p)] == 44 {
			goto tr121
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st104
		}
		goto st0
tr121:
//line internal/parser/game/ext/parser.rl:76

		t.tokval(strTok(t.data[t.prev:t.p]))
	
	goto st105
	st105:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof105
		}
	st_case_105:
//line internal/parser/game/parser.gen.go:1939
		if ( t.data)[( t.p)] == 32 {
			goto st106
		}
		goto st0
	st106:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof106
		}
	st_case_106:
		if ( t.data)[( t.p)] == 77 {
			goto st107
		}
		goto st0
	st107:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof107
		}
	st_case_107:
		if ( t.data)[( t.p)] == 84 {
			goto st108
		}
		goto st0
	st108:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof108
		}
	st_case_108:
		if ( t.data)[( t.p)] == 85 {
			goto st109
		}
		goto st0
	st109:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof109
		}
	st_case_109:
		if ( t.data)[( t.p)] == 32 {
			goto st110
		}
		goto st0
	st110:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof110
		}
	st_case_110:
		if ( t.data)[( t.p)] == 45 {
			goto tr127
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr128
		}
		goto st0
tr127:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st111
	st111:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof111
		}
	st_case_111:
//line internal/parser/game/parser.gen.go:2006
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st112
		}
		goto st0
tr128:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st112
	st112:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof112
		}
	st_case_112:
//line internal/parser/game/parser.gen.go:2025
		if ( t.data)[( t.p)] == 46 {
			goto tr130
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st112
		}
		goto st0
tr130:
//line internal/parser/game/ext/parser.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(common.ParseTokenError{TokType: "INT", Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 113; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st113
	st113:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof113
		}
	st_case_113:
//line internal/parser/game/parser.gen.go:2049
		if ( t.data)[( t.p)] == 32 {
			goto st114
		}
		goto st0
	st114:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof114
		}
	st_case_114:
		if ( t.data)[( t.p)] == 115 {
			goto st115
		}
		goto st0
	st115:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof115
		}
	st_case_115:
		if ( t.data)[( t.p)] == 101 {
			goto st116
		}
		goto st0
	st116:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof116
		}
	st_case_116:
		if ( t.data)[( t.p)] == 116 {
			goto st117
		}
		goto st0
	st117:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof117
		}
	st_case_117:
		if ( t.data)[( t.p)] == 116 {
			goto st118
		}
		goto st0
	st118:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof118
		}
	st_case_118:
		if ( t.data)[( t.p)] == 105 {
			goto st119
		}
		goto st0
	st119:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof119
		}
	st_case_119:
		if ( t.data)[( t.p)] == 110 {
			goto st120
		}
		goto st0
	st120:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof120
		}
	st_case_120:
		if ( t.data)[( t.p)] == 103 {
			goto st121
		}
		goto st0
	st121:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof121
		}
	st_case_121:
		if ( t.data)[( t.p)] == 32 {
			goto st122
		}
		goto st0
	st122:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof122
		}
	st_case_122:
		if ( t.data)[( t.p)] == 117 {
			goto st123
		}
		goto st0
	st123:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof123
		}
	st_case_123:
		if ( t.data)[( t.p)] == 112 {
			goto st124
		}
		goto st0
	st124:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof124
		}
	st_case_124:
		if ( t.data)[( t.p)] == 32 {
			goto st125
		}
		goto st0
	st125:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof125
		}
	st_case_125:
		if ( t.data)[( t.p)] == 115 {
			goto st126
		}
		goto st0
	st126:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof126
		}
	st_case_126:
		if ( t.data)[( t.p)] == 101 {
			goto st127
		}
		goto st0
	st127:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof127
		}
	st_case_127:
		if ( t.data)[( t.p)] == 115 {
			goto st128
		}
		goto st0
	st128:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof128
		}
	st_case_128:
		if ( t.data)[( t.p)] == 115 {
			goto st129
		}
		goto st0
	st129:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof129
		}
	st_case_129:
		if ( t.data)[( t.p)] == 105 {
			goto st130
		}
		goto st0
	st130:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof130
		}
	st_case_130:
		if ( t.data)[( t.p)] == 111 {
			goto st131
		}
		goto st0
	st131:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof131
		}
	st_case_131:
		if ( t.data)[( t.p)] == 110 {
			goto st132
		}
		goto st0
	st132:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof132
		}
	st_case_132:
		if ( t.data)[( t.p)] == 46 {
			goto st133
		}
		goto st0
	st133:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof133
		}
	st_case_133:
		if ( t.data)[( t.p)] == 46 {
			goto st134
		}
		goto st0
	st134:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof134
		}
	st_case_134:
		if ( t.data)[( t.p)] == 46 {
			goto st176
		}
		goto st0
	st135:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof135
		}
	st_case_135:
		if ( t.data)[( t.p)] == 46 {
			goto st99
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st136
		}
		goto st0
	st136:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof136
		}
	st_case_136:
		if ( t.data)[( t.p)] == 46 {
			goto st99
		}
		goto st0
	st137:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof137
		}
	st_case_137:
		if ( t.data)[( t.p)] == 46 {
			goto st97
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st138
		}
		goto st0
	st138:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof138
		}
	st_case_138:
		if ( t.data)[( t.p)] == 46 {
			goto st97
		}
		goto st0
	st139:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof139
		}
	st_case_139:
		if ( t.data)[( t.p)] == 46 {
			goto st95
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st140
		}
		goto st0
	st140:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof140
		}
	st_case_140:
		if ( t.data)[( t.p)] == 46 {
			goto st95
		}
		goto st0
	st141:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof141
		}
	st_case_141:
		if ( t.data)[( t.p)] == 111 {
			goto st142
		}
		goto st0
	st142:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof142
		}
	st_case_142:
		if ( t.data)[( t.p)] == 110 {
			goto st143
		}
		goto st0
	st143:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof143
		}
	st_case_143:
		if ( t.data)[( t.p)] == 32 {
			goto st144
		}
		goto st0
	st144:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof144
		}
	st_case_144:
		if ( t.data)[( t.p)] == 99 {
			goto st145
		}
		goto st0
	st145:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof145
		}
	st_case_145:
		if ( t.data)[( t.p)] == 108 {
			goto st146
		}
		goto st0
	st146:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof146
		}
	st_case_146:
		if ( t.data)[( t.p)] == 111 {
			goto st147
		}
		goto st0
	st147:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof147
		}
	st_case_147:
		if ( t.data)[( t.p)] == 115 {
			goto st148
		}
		goto st0
	st148:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof148
		}
	st_case_148:
		if ( t.data)[( t.p)] == 101 {
			goto st149
		}
		goto st0
	st149:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof149
		}
	st_case_149:
		if ( t.data)[( t.p)] == 100 {
			goto st150
		}
		goto st0
	st150:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof150
		}
	st_case_150:
		if ( t.data)[( t.p)] == 46 {
			goto st151
		}
		goto st0
	st151:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof151
		}
	st_case_151:
		if ( t.data)[( t.p)] == 32 {
			goto st152
		}
		goto st0
	st152:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof152
		}
	st_case_152:
		switch ( t.data)[( t.p)] {
		case 40:
			goto tr167
		case 95:
			goto tr168
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto tr168
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto tr168
			}
		default:
			goto tr168
		}
		goto st0
tr167:
//line internal/parser/game/ext/parser.rl:125
t.tok(CLIENT_CONNECTION_CLOSED)
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st153
	st153:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof153
		}
	st_case_153:
//line internal/parser/game/parser.gen.go:2445
		if ( t.data)[( t.p)] == 95 {
			goto st154
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st154
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st154
			}
		default:
			goto st154
		}
		goto st0
	st154:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof154
		}
	st_case_154:
		switch ( t.data)[( t.p)] {
		case 41:
			goto st179
		case 95:
			goto st154
		}
		switch {
		case ( t.data)[( t.p)] < 65:
			if 47 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
				goto st154
			}
		case ( t.data)[( t.p)] > 90:
			if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
				goto st154
			}
		default:
			goto st154
		}
		goto st0
	st179:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof179
		}
	st_case_179:
		if ( t.data)[( t.p)] == 32 {
			goto tr194
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto tr194
		}
		goto st0
tr168:
//line internal/parser/game/ext/parser.rl:125
t.tok(CLIENT_CONNECTION_CLOSED)
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st180
	st180:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof180
		}
	st_case_180:
//line internal/parser/game/parser.gen.go:2514
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr194
		case 95:
			goto st180
		}
		switch {
		case ( t.data)[( t.p)] < 47:
			if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
				goto tr194
			}
		case ( t.data)[( t.p)] > 57:
			switch {
			case ( t.data)[( t.p)] > 90:
				if 97 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 122 {
					goto st180
				}
			case ( t.data)[( t.p)] >= 65:
				goto st180
			}
		default:
			goto st180
		}
		goto st0
	st155:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof155
		}
	st_case_155:
		if ( t.data)[( t.p)] == 108 {
			goto st156
		}
		goto st0
	st156:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof156
		}
	st_case_156:
		if ( t.data)[( t.p)] == 97 {
			goto st157
		}
		goto st0
	st157:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof157
		}
	st_case_157:
		if ( t.data)[( t.p)] == 121 {
			goto st158
		}
		goto st0
	st158:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof158
		}
	st_case_158:
		if ( t.data)[( t.p)] == 101 {
			goto st159
		}
		goto st0
	st159:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof159
		}
	st_case_159:
		if ( t.data)[( t.p)] == 114 {
			goto st160
		}
		goto st0
	st160:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof160
		}
	st_case_160:
		if ( t.data)[( t.p)] == 32 {
			goto st161
		}
		goto st0
	st161:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof161
		}
	st_case_161:
		if ( t.data)[( t.p)] == 45 {
			goto tr177
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto tr178
		}
		goto st0
tr177:
//line internal/parser/game/ext/parser.rl:121
t.tok(CLIENT_PLAYER_LEAVE)
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st162
	st162:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof162
		}
	st_case_162:
//line internal/parser/game/parser.gen.go:2621
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st163
		}
		goto st0
tr178:
//line internal/parser/game/ext/parser.rl:121
t.tok(CLIENT_PLAYER_LEAVE)
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st163
	st163:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof163
		}
	st_case_163:
//line internal/parser/game/parser.gen.go:2642
		if ( t.data)[( t.p)] == 32 {
			goto tr180
		}
		if 48 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 57 {
			goto st163
		}
		goto st0
tr180:
//line internal/parser/game/ext/parser.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(common.ParseTokenError{TokType: "INT", Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 164; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
	goto st164
	st164:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof164
		}
	st_case_164:
//line internal/parser/game/parser.gen.go:2666
		if ( t.data)[( t.p)] == 108 {
			goto st165
		}
		goto st0
	st165:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof165
		}
	st_case_165:
		if ( t.data)[( t.p)] == 101 {
			goto st166
		}
		goto st0
	st166:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof166
		}
	st_case_166:
		if ( t.data)[( t.p)] == 97 {
			goto st167
		}
		goto st0
	st167:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof167
		}
	st_case_167:
		if ( t.data)[( t.p)] == 118 {
			goto st168
		}
		goto st0
	st168:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof168
		}
	st_case_168:
		if ( t.data)[( t.p)] == 101 {
			goto st169
		}
		goto st0
	st169:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof169
		}
	st_case_169:
		if ( t.data)[( t.p)] == 32 {
			goto st170
		}
		goto st0
	st170:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof170
		}
	st_case_170:
		if ( t.data)[( t.p)] == 103 {
			goto st171
		}
		goto st0
	st171:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof171
		}
	st_case_171:
		if ( t.data)[( t.p)] == 97 {
			goto st172
		}
		goto st0
	st172:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof172
		}
	st_case_172:
		if ( t.data)[( t.p)] == 109 {
			goto st173
		}
		goto st0
	st173:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof173
		}
	st_case_173:
		if ( t.data)[( t.p)] == 101 {
			goto tr190
		}
		goto st0
tr15:
//line internal/parser/game/ext/parser.rl:95

		if debugTokenizer {
			fmt.Printf("start: %s\n", t.state)
		}
		t.prev = t.p
	
	goto st174
	st174:
		if ( t.p)++; ( t.p) == ( t.pe) {
			goto _test_eof174
		}
	st_case_174:
//line internal/parser/game/parser.gen.go:2766
		switch ( t.data)[( t.p)] {
		case 32:
			goto tr15
		case 124:
			goto tr17
		}
		if 9 <= ( t.data)[( t.p)] && ( t.data)[( t.p)] <= 13 {
			goto tr14
		}
		goto st0
	st_out:
	_test_eof2:  t.cs = 2; goto _test_eof
	_test_eof3:  t.cs = 3; goto _test_eof
	_test_eof4:  t.cs = 4; goto _test_eof
	_test_eof5:  t.cs = 5; goto _test_eof
	_test_eof6:  t.cs = 6; goto _test_eof
	_test_eof7:  t.cs = 7; goto _test_eof
	_test_eof8:  t.cs = 8; goto _test_eof
	_test_eof9:  t.cs = 9; goto _test_eof
	_test_eof10:  t.cs = 10; goto _test_eof
	_test_eof11:  t.cs = 11; goto _test_eof
	_test_eof12:  t.cs = 12; goto _test_eof
	_test_eof13:  t.cs = 13; goto _test_eof
	_test_eof14:  t.cs = 14; goto _test_eof
	_test_eof15:  t.cs = 15; goto _test_eof
	_test_eof16:  t.cs = 16; goto _test_eof
	_test_eof17:  t.cs = 17; goto _test_eof
	_test_eof18:  t.cs = 18; goto _test_eof
	_test_eof19:  t.cs = 19; goto _test_eof
	_test_eof20:  t.cs = 20; goto _test_eof
	_test_eof21:  t.cs = 21; goto _test_eof
	_test_eof22:  t.cs = 22; goto _test_eof
	_test_eof23:  t.cs = 23; goto _test_eof
	_test_eof24:  t.cs = 24; goto _test_eof
	_test_eof25:  t.cs = 25; goto _test_eof
	_test_eof26:  t.cs = 26; goto _test_eof
	_test_eof27:  t.cs = 27; goto _test_eof
	_test_eof28:  t.cs = 28; goto _test_eof
	_test_eof29:  t.cs = 29; goto _test_eof
	_test_eof30:  t.cs = 30; goto _test_eof
	_test_eof31:  t.cs = 31; goto _test_eof
	_test_eof32:  t.cs = 32; goto _test_eof
	_test_eof33:  t.cs = 33; goto _test_eof
	_test_eof34:  t.cs = 34; goto _test_eof
	_test_eof35:  t.cs = 35; goto _test_eof
	_test_eof36:  t.cs = 36; goto _test_eof
	_test_eof37:  t.cs = 37; goto _test_eof
	_test_eof38:  t.cs = 38; goto _test_eof
	_test_eof39:  t.cs = 39; goto _test_eof
	_test_eof40:  t.cs = 40; goto _test_eof
	_test_eof41:  t.cs = 41; goto _test_eof
	_test_eof42:  t.cs = 42; goto _test_eof
	_test_eof43:  t.cs = 43; goto _test_eof
	_test_eof44:  t.cs = 44; goto _test_eof
	_test_eof45:  t.cs = 45; goto _test_eof
	_test_eof46:  t.cs = 46; goto _test_eof
	_test_eof47:  t.cs = 47; goto _test_eof
	_test_eof48:  t.cs = 48; goto _test_eof
	_test_eof49:  t.cs = 49; goto _test_eof
	_test_eof50:  t.cs = 50; goto _test_eof
	_test_eof51:  t.cs = 51; goto _test_eof
	_test_eof52:  t.cs = 52; goto _test_eof
	_test_eof53:  t.cs = 53; goto _test_eof
	_test_eof54:  t.cs = 54; goto _test_eof
	_test_eof55:  t.cs = 55; goto _test_eof
	_test_eof56:  t.cs = 56; goto _test_eof
	_test_eof57:  t.cs = 57; goto _test_eof
	_test_eof58:  t.cs = 58; goto _test_eof
	_test_eof59:  t.cs = 59; goto _test_eof
	_test_eof60:  t.cs = 60; goto _test_eof
	_test_eof61:  t.cs = 61; goto _test_eof
	_test_eof62:  t.cs = 62; goto _test_eof
	_test_eof63:  t.cs = 63; goto _test_eof
	_test_eof64:  t.cs = 64; goto _test_eof
	_test_eof65:  t.cs = 65; goto _test_eof
	_test_eof66:  t.cs = 66; goto _test_eof
	_test_eof67:  t.cs = 67; goto _test_eof
	_test_eof68:  t.cs = 68; goto _test_eof
	_test_eof69:  t.cs = 69; goto _test_eof
	_test_eof70:  t.cs = 70; goto _test_eof
	_test_eof71:  t.cs = 71; goto _test_eof
	_test_eof175:  t.cs = 175; goto _test_eof
	_test_eof176:  t.cs = 176; goto _test_eof
	_test_eof177:  t.cs = 177; goto _test_eof
	_test_eof72:  t.cs = 72; goto _test_eof
	_test_eof73:  t.cs = 73; goto _test_eof
	_test_eof74:  t.cs = 74; goto _test_eof
	_test_eof75:  t.cs = 75; goto _test_eof
	_test_eof76:  t.cs = 76; goto _test_eof
	_test_eof77:  t.cs = 77; goto _test_eof
	_test_eof78:  t.cs = 78; goto _test_eof
	_test_eof178:  t.cs = 178; goto _test_eof
	_test_eof79:  t.cs = 79; goto _test_eof
	_test_eof80:  t.cs = 80; goto _test_eof
	_test_eof81:  t.cs = 81; goto _test_eof
	_test_eof82:  t.cs = 82; goto _test_eof
	_test_eof83:  t.cs = 83; goto _test_eof
	_test_eof84:  t.cs = 84; goto _test_eof
	_test_eof85:  t.cs = 85; goto _test_eof
	_test_eof86:  t.cs = 86; goto _test_eof
	_test_eof87:  t.cs = 87; goto _test_eof
	_test_eof88:  t.cs = 88; goto _test_eof
	_test_eof89:  t.cs = 89; goto _test_eof
	_test_eof90:  t.cs = 90; goto _test_eof
	_test_eof91:  t.cs = 91; goto _test_eof
	_test_eof92:  t.cs = 92; goto _test_eof
	_test_eof93:  t.cs = 93; goto _test_eof
	_test_eof94:  t.cs = 94; goto _test_eof
	_test_eof95:  t.cs = 95; goto _test_eof
	_test_eof96:  t.cs = 96; goto _test_eof
	_test_eof97:  t.cs = 97; goto _test_eof
	_test_eof98:  t.cs = 98; goto _test_eof
	_test_eof99:  t.cs = 99; goto _test_eof
	_test_eof100:  t.cs = 100; goto _test_eof
	_test_eof101:  t.cs = 101; goto _test_eof
	_test_eof102:  t.cs = 102; goto _test_eof
	_test_eof103:  t.cs = 103; goto _test_eof
	_test_eof104:  t.cs = 104; goto _test_eof
	_test_eof105:  t.cs = 105; goto _test_eof
	_test_eof106:  t.cs = 106; goto _test_eof
	_test_eof107:  t.cs = 107; goto _test_eof
	_test_eof108:  t.cs = 108; goto _test_eof
	_test_eof109:  t.cs = 109; goto _test_eof
	_test_eof110:  t.cs = 110; goto _test_eof
	_test_eof111:  t.cs = 111; goto _test_eof
	_test_eof112:  t.cs = 112; goto _test_eof
	_test_eof113:  t.cs = 113; goto _test_eof
	_test_eof114:  t.cs = 114; goto _test_eof
	_test_eof115:  t.cs = 115; goto _test_eof
	_test_eof116:  t.cs = 116; goto _test_eof
	_test_eof117:  t.cs = 117; goto _test_eof
	_test_eof118:  t.cs = 118; goto _test_eof
	_test_eof119:  t.cs = 119; goto _test_eof
	_test_eof120:  t.cs = 120; goto _test_eof
	_test_eof121:  t.cs = 121; goto _test_eof
	_test_eof122:  t.cs = 122; goto _test_eof
	_test_eof123:  t.cs = 123; goto _test_eof
	_test_eof124:  t.cs = 124; goto _test_eof
	_test_eof125:  t.cs = 125; goto _test_eof
	_test_eof126:  t.cs = 126; goto _test_eof
	_test_eof127:  t.cs = 127; goto _test_eof
	_test_eof128:  t.cs = 128; goto _test_eof
	_test_eof129:  t.cs = 129; goto _test_eof
	_test_eof130:  t.cs = 130; goto _test_eof
	_test_eof131:  t.cs = 131; goto _test_eof
	_test_eof132:  t.cs = 132; goto _test_eof
	_test_eof133:  t.cs = 133; goto _test_eof
	_test_eof134:  t.cs = 134; goto _test_eof
	_test_eof135:  t.cs = 135; goto _test_eof
	_test_eof136:  t.cs = 136; goto _test_eof
	_test_eof137:  t.cs = 137; goto _test_eof
	_test_eof138:  t.cs = 138; goto _test_eof
	_test_eof139:  t.cs = 139; goto _test_eof
	_test_eof140:  t.cs = 140; goto _test_eof
	_test_eof141:  t.cs = 141; goto _test_eof
	_test_eof142:  t.cs = 142; goto _test_eof
	_test_eof143:  t.cs = 143; goto _test_eof
	_test_eof144:  t.cs = 144; goto _test_eof
	_test_eof145:  t.cs = 145; goto _test_eof
	_test_eof146:  t.cs = 146; goto _test_eof
	_test_eof147:  t.cs = 147; goto _test_eof
	_test_eof148:  t.cs = 148; goto _test_eof
	_test_eof149:  t.cs = 149; goto _test_eof
	_test_eof150:  t.cs = 150; goto _test_eof
	_test_eof151:  t.cs = 151; goto _test_eof
	_test_eof152:  t.cs = 152; goto _test_eof
	_test_eof153:  t.cs = 153; goto _test_eof
	_test_eof154:  t.cs = 154; goto _test_eof
	_test_eof179:  t.cs = 179; goto _test_eof
	_test_eof180:  t.cs = 180; goto _test_eof
	_test_eof155:  t.cs = 155; goto _test_eof
	_test_eof156:  t.cs = 156; goto _test_eof
	_test_eof157:  t.cs = 157; goto _test_eof
	_test_eof158:  t.cs = 158; goto _test_eof
	_test_eof159:  t.cs = 159; goto _test_eof
	_test_eof160:  t.cs = 160; goto _test_eof
	_test_eof161:  t.cs = 161; goto _test_eof
	_test_eof162:  t.cs = 162; goto _test_eof
	_test_eof163:  t.cs = 163; goto _test_eof
	_test_eof164:  t.cs = 164; goto _test_eof
	_test_eof165:  t.cs = 165; goto _test_eof
	_test_eof166:  t.cs = 166; goto _test_eof
	_test_eof167:  t.cs = 167; goto _test_eof
	_test_eof168:  t.cs = 168; goto _test_eof
	_test_eof169:  t.cs = 169; goto _test_eof
	_test_eof170:  t.cs = 170; goto _test_eof
	_test_eof171:  t.cs = 171; goto _test_eof
	_test_eof172:  t.cs = 172; goto _test_eof
	_test_eof173:  t.cs = 173; goto _test_eof
	_test_eof174:  t.cs = 174; goto _test_eof

	_test_eof: {}
	if ( t.p) == ( t.pe) {
		switch  t.cs {
		case 176, 177:
//line internal/parser/game/ext/parser.rl:108
 t.tok(EOL) 
		case 179, 180:
//line internal/parser/game/ext/parser.rl:76

		t.tokval(strTok(t.data[t.prev:t.p]))
	
//line internal/parser/game/ext/parser.rl:108
 t.tok(EOL) 
		case 175, 178:
//line internal/parser/game/ext/parser.rl:79

		temp.Int, parseErr = strconv.Atoi(t.data[t.prev:t.p])
		if parseErr != nil {
			t.err(common.ParseTokenError{TokType: "INT", Raw: t.data[t.prev:t.p], Err: fmt.Errorf("strconv.Atoi: %w", parseErr)})
			{( t.p)++;  t.cs = 0; goto _out }
		}
		t.tokval(intTok(temp.Int))
	
//line internal/parser/game/ext/parser.rl:108
 t.tok(EOL) 
//line internal/parser/game/parser.gen.go:2983
		}
	}

	_out: {}
	}

//line internal/parser/game/ext/parser.rl:151

	if debugTokenizer {
		fmt.Printf("exited: %s\n", t.state)
	}
	if t.p != t.pe || (len(t.tokens) > 0 && t.tokens[len(t.tokens)-1] != voidTok(EOL)) {
		t.err(common.LineIsNotFinishedError{Parsed: t.data[:t.p], Rest: t.data[t.p:t.pe]})
	}
	return t.tokens, errors.Join(t.errors...)
}

func (t *Tokenizer) tokval(token Token) {
	t.tokens = append(t.tokens, token)
}

func (t *Tokenizer) tok(tok int) {
	t.tokens = append(t.tokens, voidTok(tok))
}

func (t *Tokenizer) err(err error) {
	t.errors = append(t.errors, err)
}