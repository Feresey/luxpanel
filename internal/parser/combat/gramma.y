%{
package combat

type (
	Int = int
	String = string
	Strings = []string
	Float = float32
	Bool = bool
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
	DamageModifiers;
	Object;
	Heal;
	Kill;
	Participant;
	ParticipationModifiers;
	ConnectToGameSession;
	Start;
	Finished;
	Reward;
	LogLine;
}

// MAIN TOKENS

// BASIC TYPES
%token <Int> INT
%token <String> STRING
%left <Time> TIME

// TYPES

%left COMBAT
%token ARROW
%token <Float> FLOAT
%token <String> SOURCE
%token FRIENDLY_FIRE

// COMBAT

%type <Object> object
%type <Object> player_or_object

// Damage
%left DAMAGE
%token<String> DAMAGE_MODIFIER
%token ROCKET

%type <String> source
%type <Damage> damage
%type <DamageModifiers> damage_modifiers
%type <Bool> friendly_fire
%type <Int> rocket

// Heal
%left HEAL

%type <Heal> heal

// Kill
%left KILL

%type <Kill> kill

// Participation
%left PARTICIPANT
%token <String> PARTICIPATION_MODIFIER
%right PARTICIPATION_MODIFIERS_END

%type <String> string
%type <Participant> participation
%type <Participant> participation_damage
%type <ParticipationModifiers> participation_modifiers

// Connect to game session
%left CONNECT_TO_GAME_SESSION_PREFIX

%type <ConnectToGameSession> connect_to_game_session
%type <Int> local_client_team

// Start
%left START

%type <Start> start

// Finish
%left GAMEPLAY_FINISHED

%type <Finished> finished

// Reward
%left REWARD

%type <Reward> reward
%type <String> ship

// RESULT

%type <LogLine> action

%right EOL

%%

main: TIME COMBAT action EOL {
	$3.setTime($1)
	yylex.(*lexer).line = $3
}

action:
	damage {
		$$ = &$1
	} |
	heal {
		$$ = &$1
	} |
	kill {
		$$ = &$1
	} |
	participation {
		$$ = &$1
	} |
	start {
		$$ = &$1
	} |
	finished {
		$$ = &$1
	} |
	reward {
		$$ = &$1
	} |
	connect_to_game_session {
		$$ = &$1
	}

// CONTENTS

// common

object:
	STRING '|' INT {
		$$ = Object{
			Name: $1,
			ObjectID: $3,
		}
	} |
	STRING '(' STRING ')' '|' INT {
		$$ = Object{
			PlayerObject: PlayerObject{
				ObjectName: $1,
				ObjectOwner: $3,
			},
			ObjectID: $6,
		}
	}

friendly_fire: FRIENDLY_FIRE { $$ = true } | /* empty */ { $$ = false };

rocket: ROCKET INT { $$ = $2 } | /* empty */ { $$ = 0 };

// Damage

source: SOURCE { $$ = $1} | { $$ = ""}

// 21:41:15.644  CMBT   | Damage         aSpectro|0000002757 ->          Nafalar|0000002677 117.55 (h:0.00 s:117.55) Weapon_KineticStkBomb_T5_Epic KINETIC|PRIMARY_WEAPON|CRIT  
damage:
	DAMAGE
	object ARROW object
	FLOAT FLOAT FLOAT
	source
	damage_modifiers
	friendly_fire
	rocket {
	$$ = Damage{
		Initiator: $2,
		Recipient: $4,
		DamageFull: $5,
		DamageHull: $6,
		DamageShield: $7,
		Source: $8,
		DamageModifiers: $9,
		FriendlyFire: $10,
		Rocket: $11,
	}
}

damage_modifiers: DAMAGE_MODIFIER {
	$$ = []DamageModifier{DamageModifier($1)}
} | damage_modifiers '|' DAMAGE_MODIFIER {
	$$ = append($$, DamageModifier($3))
}

// Heal

// 19:33:24.732  CMBT   | Heal            Feresey|0000000204 ->          Feresey|0000000204 244.00 Module_Lynx2Shield_T4_Epic
heal: HEAL object ARROW object FLOAT SOURCE {
	$$ = Heal{
		Initiator: $2,
		Recipient: $4,
		Heal: $5,
		Source: $6,
	}
}

// Kill

// 23:04:35.283  CMBT   | Killed Ship_Bot_ClanShipDroneBig_Empire|0000001860;	 killer Feresey|0000002076 Weapon_Plasmagun_Heavy_T5_Pirate 
// 19:33:59.527  CMBT   | Killed Py6Jl	 Ship_Race3_M_T2_Pirate|0000000248;	 killer Feresey|0000000204 Weapon_Plasmagun_Heavy_T5_Pirate 
// 19:44:55.746  CMBT   | Killed SwarmPack2(georgeatg)|0000001044;	 killer georgeatg|0000001044 (suicide) <FriendlyFire>
kill: KILL player_or_object ';' object source friendly_fire {
	$$ = Kill{
		Killed: $2,
		Killer: $4,
		Source: $5,
		FriendlyFire: $6,
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

participation_modifiers: participation_modifiers PARTICIPATION_MODIFIER {
	$$ = append($1, ParticipationModifier($2))
} | participation_modifiers PARTICIPATION_MODIFIERS_END {
	$$ = $1
} | {
	$$ = nil
}

string: STRING { $$ = $1 } | { $$ = "" }

// 23:04:35.283  CMBT   |    Participant            Py6Jl	 Ship_Race1_M_T5_Dlc             	 totalDamage 18029.51; mostDamageWith 'Weapon_ThorHammer_T5_Epic'; <debuff>
// 23:04:35.283  CMBT   |    Participant        Gladiator	 Ship_Race1_L_T5_CraftSpec       	 totalDamage 4711.33; mostDamageWith 'Weapon_ShieldHeal_T5_Epic'; <buff>
// 23:04:35.283  CMBT   |    Participant          Feresey	 Ship_Race5_M_ATTACK_Rank15      	 totalDamage 4597.25; mostDamageWith 'Weapon_Plasmagun_Heavy_T5_Pirate';
// 23:04:35.283  CMBT   |    Participant             OSN1	 Ship_Race2_M_T5_CraftUniq       	 <buff>
// 23:04:35.283  CMBT   |    Participant   MadmenRoverTit	 Ship_Race2_M_T5_CraftUniq       	 <buff>
participation: PARTICIPANT STRING string participation_damage participation_modifiers friendly_fire {
	$$ = Participant{
		Name: $2,
		Ship: $3,
		Damage: $4.Damage,
		MostDamageWith: $4.MostDamageWith,
		Modifiers: $5,
		FriendlyFire: $6,
	}
}

participation_damage: FLOAT SOURCE {
	$$.Damage = $1
	$$.MostDamageWith = $2
} | {
	$$.Damage = 0
	$$.MostDamageWith = ""
}

// Start gameplay

local_client_team: INT { $$ = $1 } | /*empty*/ { $$ = 0 }

// 19:33:00.763  CMBT   | ======= Start gameplay 'BombTheBase' map 's1340_thar_aliendebris13', local client team 2 =======
// 19:42:14.670  CMBT   | ======= Start PVE mission 'pve_raid_waterharvest_t5' map 'pve_raid_waterharvest' =======
start: START STRING STRING STRING local_client_team {
	$$ = Start{
		What: $2,
		GameMode: $3,
		MapName: $4,
		LocalClientTeamID: $5,
	}
}

// Connect to game client

// 19:32:58.666  CMBT   | ======= Connect to game session 50419619 =======
connect_to_game_session: CONNECT_TO_GAME_SESSION_PREFIX INT {
	$$.SessionID = $2
}


// Finished

// 19:47:09.448  CMBT   | Gameplay finished. Winner team: 1(PVE_MISSION_COMPLETE_ALT_2). Finish reason: 'Mission complete'. Actual game time 275.9 sec
finished: GAMEPLAY_FINISHED INT STRING STRING FLOAT {
	$$ = Finished{
		WinnerTeamID: $2,
		WinReason: $3,
		FinishReason: $4,
		GameTime: $5,
	}
}

// Reward

ship: STRING { $$ = $1 } | { $$ = ""}

// 19:33:35.728  CMBT   | Reward          Feresey Ship_Race3_M_T2_Pirate          	 136 experience                for damage Py6Jl
reward: REWARD STRING ship INT STRING STRING {
	$$ = Reward{
		Recipient: $2,
		Ship: $3,
		Reward: $4,
		RewardType: $5,
		Reason: $6,
	}
}

%%