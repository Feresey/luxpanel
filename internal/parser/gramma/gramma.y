%{
package gramma

import (
	"time"
	"strings"
)

type (
	Int = int
	String = string
	Strings = []string
	Float = float32
	Bool = bool
	Time = time.Time
)

%}

%union {
	// Common
	String;
	Strings;
	Int;
	Float;
	Bool;
	Time;
	// Combat log
	Damage;
	DamageModifiersMap;
	Object;
	Heal;
	Kill;
	Participant;
	BuffDebuff;
	ConnectToGameSession;
	Start;
	Finished;
	Reward;
	CombatLine;
	// Game log
	ClientAddPlayer;
	ClientPlayerLeave;
	ClientConnected;
	ClientConnectionClosed;
	GameLine;
	// Result
	LogLine;
}

// MAIN TOKENS

%left COMBAT
%left GAME
%right
%token ARROW

// BASIC TYPES
%token <Float> FLOAT
%token <Int> INT
%token <String> STRING
%token <String> SOURCE
%token FRIENDLY_FIRE
%left <Time> TIME

// TYPES

// Common
%type <Strings> strings

// COMBAT

%type <Object> object
%type <Object> player_or_object

// Damage
%left DAMAGE
%token<String> DAMAGE_MODIFIER
%token ROCKET

%type <Damage> damage
%type <Damage> damage_line
%type <DamageModifiersMap> damage_modifiers
%type <Bool> friendly_fire
%type <Int> rocket

// Heal
%left HEAL

%type <Heal> heal
%type <Heal> heal_line

// Kill
%left KILL

%type <Kill> kill
%type <Kill> kill_line

// Participation
%left PARTICIPANT
%token BUFF DEBUFF

%type <Participant> participation
%type <Participant> participation_line
%type <Participant> participation_damage
%type <BuffDebuff> buff_debuff

// Connect to game session
%left CONNECT_TO_GAME_SESSION_PREFIX
%token EQ_DELIM LOCAL_CLIENT_TEAM

%type <ConnectToGameSession> connect_to_game_session
%type <ConnectToGameSession> connect_to_game_session_line
%type <Int> local_client_team

// Start
%left START

%type <Start> start
%type <Start> start_line

// Finish
%left GAMEPLAY_FINISHED
%token ACTUAL_GAME_TIME FINISH_REASON

%type <Finished> finished
%type <Finished> finished_line

// Reward
%left REWARD

%type <Reward> reward
%type <Reward> reward_line


// GAME

// Add player
%token CLIENT_ADD_PLAYER

%type <Int> group
%type <ClientAddPlayer> add_player_line
%type <ClientAddPlayer> add_player

// Remove player
%token CLIENT_PLAYER_LEAVE

%type <ClientPlayerLeave> player_leave_line
%type <ClientPlayerLeave> player_leave

// Client connected
%token CLIENT_CONNECTED

%type <ClientConnected> client_connected_line
%type <ClientConnected> client_connected

// Client connection closed
%token CLIENT_CONNECTION_CLOSED

%type <ClientConnectionClosed> client_connection_closed_line
%type <ClientConnectionClosed> client_connection_closed

// RESULT

%type <LogLine> action

%type <CombatLine> combat_line
%type <GameLine> game_line

%%

main: action {
	yylex.(*lexer).res = $1
}


action: combat_line {
		$$.Combat = $1
	} |
	game_line {
		$$.Game = $1
	}

combat_line:
	damage_line {
		$$ = &$1
	} |
	heal_line {
		$$ = &$1
	} |
	kill_line {
		$$ = &$1
	} |
	participation_line {
		$$ = &$1
	} |
	start_line {
		$$ = &$1
	} |
	finished_line {
		$$ = &$1
	} |
	reward_line {
		$$ = &$1
	} |
	connect_to_game_session_line {
		$$ = &$1
	}

game_line:
	add_player_line {
		$$ = &$1
	} |
	player_leave_line {
		$$ = &$1
	} |
	client_connected_line {
		$$ = &$1
	} |
	client_connection_closed_line {
		$$ = &$1
	} | TIME GAME $unk {}

// LINES

// COMBAT

damage_line: TIME COMBAT DAMAGE damage {
	$$ = $4
	$$.Time = $1
}

heal_line: TIME COMBAT HEAL heal {
	$$ = $4
	$$.Time = $1
}

kill_line: TIME COMBAT KILL kill {
	$$ = $4
	$$.Time = $1
}

participation_line: TIME COMBAT PARTICIPANT participation {
	$$ = $4
	$$.Time = $1
}

connect_to_game_session_line: TIME COMBAT connect_to_game_session {
	$$ = $3
	$$.Time = $1
}

start_line: TIME COMBAT START start {
	$$ = $4
	$$.Time = $1
}

finished_line: TIME COMBAT GAMEPLAY_FINISHED finished {
	$$ = $4
	$$.Time = $1
}

reward_line: TIME COMBAT REWARD reward {

}

// GAME

add_player_line: TIME GAME add_player {
    $$ = $3
    $$.Time = $1
}

player_leave_line: TIME GAME player_leave {
    $$ = $3
    $$.Time = $1
}

client_connected_line: TIME GAME client_connected {
    $$ = $3
    $$.Time = $1
}

client_connection_closed_line: TIME GAME client_connection_closed {
    $$ = $3
    $$.Time = $1
}

// CONTENTS

// common

strings: STRING { $$ = append($$, $1) } | strings STRING { $$ = append($1, $2) }

object: STRING '|' INT {
	$$ = Object{
		Name: $1,
		ObjectID: $3,
	}
}
| STRING '(' STRING ')' '|' INT {
	$$ = Object{
		PlayerObject: PlayerObject{
			ObjectName: $1,
			ObjectOwner: $3,
		},
		ObjectID: $6,
	}
}

friendly_fire: FRIENDLY_FIRE { $$ = true } | /* empty */ {};

rocket: ROCKET INT { $$ = $2 } | /* empty */ {};

// Damage

// 21:41:15.644  CMBT   | Damage         aSpectro|0000002757 ->          Nafalar|0000002677 117.55 (h:0.00 s:117.55) Weapon_KineticStkBomb_T5_Epic KINETIC|PRIMARY_WEAPON|CRIT  
damage:
	object object
	FLOAT FLOAT FLOAT
	SOURCE
	damage_modifiers
	friendly_fire
	rocket {
	$$ = Damage{
		Initiator: $1,
		Recipient: $2,
		DamageFull: $3,
		DamageHull: $4,
		DamageShield: $5,
		Source: $6,
		DamageModifiers: $7,
		FriendlyFire: $8,
		Rocket: $9,
	}
}

damage_modifiers: DAMAGE_MODIFIER {
	$$ = DamageModifiersMap{
		DamageModifier($1): {},
	}
} | damage_modifiers '|' DAMAGE_MODIFIER {
	$$[DamageModifier($3)] = struct{}{}
}

// Heal

// 19:33:24.732  CMBT   | Heal            Feresey|0000000204 ->          Feresey|0000000204 244.00 Module_Lynx2Shield_T4_Epic
heal: object ARROW object FLOAT STRING {
	$$ = Heal{
		Initiator: $1,
		Recipient: $3,
		Heal: $4,
		Source: $5,
	}
}

// Kill

// 23:04:35.283  CMBT   | Killed Ship_Bot_ClanShipDroneBig_Empire|0000001860;	 killer Feresey|0000002076 Weapon_Plasmagun_Heavy_T5_Pirate 
// 19:33:59.527  CMBT   | Killed Py6Jl	 Ship_Race3_M_T2_Pirate|0000000248;	 killer Feresey|0000000204 Weapon_Plasmagun_Heavy_T5_Pirate 
// 19:44:55.746  CMBT   | Killed SwarmPack2(georgeatg)|0000001044;	 killer georgeatg|0000001044 (suicide) <FriendlyFire>
kill: player_or_object ';' '\t' STRING object SOURCE friendly_fire {
	if $4 != "killer" {
		yylex.Error("not a killer")
	}
	$$ = Kill{
		Killed: $1,
		Killer: $5,
		Source: $6,
		FriendlyFire: $7,
	}
}

player_or_object: STRING '\t' object {
	$$ = Object{
		Name: $1,
		PlayerObject: PlayerObject{
			ObjectName: $3.Name,
		},
		ObjectID: $3.ObjectID,
	}
} | object {
	$$ = $1
}

// Participation
buff_debuff: BUFF { $$.IsBuff = true } | DEBUFF { $$.IsDebuff = true }

// 23:04:35.283  CMBT   |    Participant            Py6Jl	 Ship_Race1_M_T5_Dlc             	 totalDamage 18029.51; mostDamageWith 'Weapon_ThorHammer_T5_Epic'; <debuff>
// 23:04:35.283  CMBT   |    Participant        Gladiator	 Ship_Race1_L_T5_CraftSpec       	 totalDamage 4711.33; mostDamageWith 'Weapon_ShieldHeal_T5_Epic'; <buff>
// 23:04:35.283  CMBT   |    Participant          Feresey	 Ship_Race5_M_ATTACK_Rank15      	 totalDamage 4597.25; mostDamageWith 'Weapon_Plasmagun_Heavy_T5_Pirate';
// 23:04:35.283  CMBT   |    Participant             OSN1	 Ship_Race2_M_T5_CraftUniq       	 <buff>
// 23:04:35.283  CMBT   |    Participant   MadmenRoverTit	 Ship_Race2_M_T5_CraftUniq       	 <buff>
participation: STRING '\t' STRING '\t' participation_damage buff_debuff friendly_fire  {
	$$ = Participant{
		Name: $1,
		Ship: $3,
		Damage: $5.Damage,
		MostDamageWith: $5.MostDamageWith,
		BuffDebuff: $6,
		FriendlyFire: $7,
	}
}

participation_damage: STRING FLOAT ';' STRING '\'' STRING '\'' ';' {
	$$.Damage = $2
	$$.MostDamageWith = $6
} | {}

// Start gameplay

local_client_team: ',' LOCAL_CLIENT_TEAM INT { $$ = $3 } | /*empty*/ {}

// 19:33:00.763  CMBT   | ======= Start gameplay 'BombTheBase' map 's1340_thar_aliendebris13', local client team 2 =======
// 19:42:14.670  CMBT   | ======= Start PVE mission 'pve_raid_waterharvest_t5' map 'pve_raid_waterharvest' =======
start: EQ_DELIM START strings '\'' STRING '\'' STRING '\'' STRING '\'' local_client_team EQ_DELIM {
	$$ = Start{
		What: strings.Join($3, " "),
		GameMode: $5,
		MapName: $9,
		LocalClientTeamID: $11,
	}
}

// Connect to game client

// 19:32:58.666  CMBT   | ======= Connect to game session 50419619 =======
connect_to_game_session: EQ_DELIM CONNECT_TO_GAME_SESSION_PREFIX INT EQ_DELIM {
	$$.SessionID = $3
}


// Finished

// 19:47:09.448  CMBT   | Gameplay finished. Winner team: 1(PVE_MISSION_COMPLETE_ALT_2). Finish reason: 'Mission complete'. Actual game time 275.9 sec
finished: INT '(' STRING ')' FINISH_REASON '\'' strings '\'' ACTUAL_GAME_TIME FLOAT STRING {
	$$ = Finished{
		WinnerTeamID: $1,
		WinReason: $3,
		FinishReason: strings.Join($7, " "),
		GameTime: $10,
	}
}

// Reward

// 19:33:35.728  CMBT   | Reward          Feresey Ship_Race3_M_T2_Pirate          	 136 experience                for damage Py6Jl
reward: STRING STRING '\t' INT STRING strings STRING {

}

// Add player

group: INT { $$ = $1 } | /* empy */ {}

// 19:42:11.984         | client: ADD_PLAYER 11 (idanceandkillyou [], 435146) status 6 team 1 group 5212334
add_player: CLIENT_ADD_PLAYER INT STRING STRING INT INT INT group {
	$$ = ClientAddPlayer{
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
player_leave: CLIENT_PLAYER_LEAVE INT {
	$$ = ClientPlayerLeave{
		InGamePlayerID: $2,
	}
}

// 21:46:40.862         | client: connected to 23.111.211.207|35020, MTU 1492. setting up session...
client_connected: CLIENT_CONNECTED STRING INT {
	$$ = ClientConnected{
		ServerAddr: $2,
		MTU: $3,
	}
}

// 21:51:50.272         | client: connection closed. DR_CLIENT_GAME_FINISHED
client_connection_closed: CLIENT_CONNECTION_CLOSED STRING {
	$$ = ClientConnectionClosed{
		Reason: $2,
	}
}

%%