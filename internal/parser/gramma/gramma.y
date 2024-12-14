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
	String;
	Strings;
	Int;
	Float;
	Bool;
	Time;
	Damage;
	DamageModifiersMap;
	Object;
	Heal;
	Kill;
	Participant;
	BuffDebuff;
	ConnectToGameSession;
	Start;
	CombatLine;
	Finished;
	Reward;
}

// MAIN TOKENS

%left COMBAT
%left GAME
%right EOL
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
%type <Object> object
%type <Object> player_or_object
%type <Bool> friendly_fire

// Damage
%left DAMAGE
%token<String> DAMAGE_MODIFIER

%type <Damage> damage
%type <Damage> damage_line
%type <DamageModifiersMap> damage_modifiers

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

// RESULT

%type <CombatLine> action

// TODO
%token CLIENT_ADD_PLAYER CLIENT_PLAYER_LEAVE CLIENT_CONNECTED CLIENT_CONNECTION_CLOSED

%%

main: action {
	yylex.(*Lexer).res = $1
}

action:
	damage_line {
		$$.Damage = &$1
	} |
	heal_line {
		$$.Heal = &$1
	} |
	kill_line {
		$$.Kill = &$1
	} |
	participation_line {
		$$.Participant = &$1
	} |
	start_line {
		$$.Start = &$1
	} |
	finished_line {
		$$.Finished = &$1
	} |
	reward_line {
		$$.Reward = &$1
	} |
	connect_to_game_session_line {
		$$.ConnectToGameSession = &$1
	} | EOL {}

// LINES

damage_line: TIME COMBAT DAMAGE damage EOL {
	$$ = $4
	$$.Time = $1
}

heal_line: TIME COMBAT HEAL heal EOL {
	$$ = $4
	$$.Time = $1
}

kill_line: TIME COMBAT KILL kill EOL {
	$$ = $4
	$$.Time = $1
}

participation_line: TIME COMBAT PARTICIPANT participation EOL {
	$$ = $4
	$$.Time = $1
}

connect_to_game_session_line: TIME COMBAT connect_to_game_session EOL {
	$$ = $3
	$$.Time = $1
}

start_line: TIME COMBAT START start EOL {
	$$ = $4
	$$.Time = $1
}

finished_line: TIME COMBAT GAMEPLAY_FINISHED finished EOL {
	$$ = $4
	$$.Time = $1
}

reward_line: TIME COMBAT REWARD reward EOL {

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

// Damage

// 21:41:15.644  CMBT   | Damage         aSpectro|0000002757 ->          Nafalar|0000002677 117.55 (h:0.00 s:117.55) Weapon_KineticStkBomb_T5_Epic KINETIC|PRIMARY_WEAPON|CRIT  
damage:
	object object
	FLOAT FLOAT FLOAT
	SOURCE
	damage_modifiers
	friendly_fire {
	$$ = Damage{
		Initiator: $1,
		Recipient: $2,
		DamageFull: $3,
		DamageHull: $4,
		DamageShield: $5,
		Source: $6,
		DamageModifiers: $7,
		FriendlyFire: $8,
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

%%