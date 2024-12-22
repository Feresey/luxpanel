%{
package game

import "fmt"

type (
	Int = int
	String = string
	Float = float32
	Bool = bool
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

%}

%union {
	// Common
	String;
	Int;
	Float;
	Bool;
	// Game log
	*ClientAddPlayer;
	*ClientPlayerLeave;
	*ClientConnected;
	*ClientConnectionClosed;
	LogLine;
}

// MAIN TOKENS

%left GAME

// BASIC TYPES
%token <Int> INT
%token <String> STRING
%left <String> TIME

// GAME

// Add player
%token CLIENT_ADD_PLAYER

%type <Int> group
%type <ClientAddPlayer> add_player
%type <String> player_clan_tag

// Remove player
%token CLIENT_PLAYER_LEAVE
%token LEAVE

%type <ClientPlayerLeave> player_leave

// Client connected
%token CLIENT_CONNECTED

%type <ClientConnected> client_connected

// Client connection closed
%token CLIENT_CONNECTION_CLOSED

%type <ClientConnectionClosed> client_connection_closed

// RESULT

%type <LogLine> line

%right EOL

%%

main: TIME GAME line EOL {
	$3.setTime($1)
	Yacclex.(*Lexer).line = $3
}

line:
	add_player {
		$$ = $1
	} |
	player_leave {
		$$ = $1
	} |
	client_connected {
		$$ = $1
	} |
	client_connection_closed {
		$$ = $1
	}

// CONTENTS

// Add player

group: INT { $$ = $1 } | /* empy */ {}

player_clan_tag: STRING { $$ = $1 } | /* empy */ {}

// 19:42:11.984         | client: ADD_PLAYER 11 (idanceandkillyou [], 435146) status 6 team 1 group 5212334
add_player: CLIENT_ADD_PLAYER INT STRING player_clan_tag INT INT INT group {
	$$ = &ClientAddPlayer{
		InGamePlayerID: $2,
		Name: $3,
		ClanTag: $4,
		PlayerID: $5,
		Status: $6,
		TeamID: $7,
		GroupID: $8,
	}
}

// 21:44:21.601         | client: player 3 leave game
player_leave: CLIENT_PLAYER_LEAVE INT LEAVE {
	$$ = &ClientPlayerLeave{
		InGamePlayerID: $2,
	}
}

// 21:46:40.862         | client: connected to 23.111.211.207|35020, MTU 1492. setting up session...
client_connected: CLIENT_CONNECTED STRING INT {
	$$ = &ClientConnected{
		ServerAddr: $2,
		MTU: $3,
	}
}

// 21:51:50.272         | client: connection closed. DR_CLIENT_GAME_FINISHED
client_connection_closed: CLIENT_CONNECTION_CLOSED STRING {
	$$ = &ClientConnectionClosed{
		Reason: $2,
	}
}

%%