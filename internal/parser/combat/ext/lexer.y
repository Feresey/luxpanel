%{
package combat

import (
	"fmt"
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
	any;
	string;
	strArr;
	int;
	float32;
	bool;
	ParticipationModifiers;
	Spell;
}

// MAIN TOKENS

// BASIC TYPES
%token <int> INT
%token <string> STRING
%left <string> TIME

// TYPES

%left COMBAT
%token ARROW
%token <float32> FLOAT
%token <string> SOURCE
%token FRIENDLY_FIRE

// COMBAT

%type <any> object
%type <any> player_or_object

// Damage
%left DAMAGE
%token<string> DAMAGE_MODIFIER
%token ROCKET

%type <string> source
%type <any> damage
%type <any> damage_modifiers
%type <bool> friendly_fire
%type <int> rocket

// Heal
%left HEAL

%type <any> heal

// Kill
%left KILL

%type <any> kill

// Participation
%left PARTICIPANT
%token <string> PARTICIPATION_MODIFIER
%right PARTICIPATION_MODIFIERS_END

%type <string> string
%type <any> participation
%type <any> participation_damage
%type <ParticipationModifiers> participation_modifiers

// Connect to game session
%left CONNECT_TO_GAME_SESSION_PREFIX

%type <any> connect_to_game_session
%type <int> local_client_team

// Spawn

%left SPAWN_PREFIX

%type <any> spawn

// Spell

%left SPELL_PREFIX
%left SPELL_NORMALIZER
%left SPELL_GODGIFT

%type <any> spell
%type <Spell> spell_type
%type <strArr> spell_targets

// Start
%left START

%type <any> start

// Finish
%left GAMEPLAY_FINISHED

%type <any> finished

// Reward
%left REWARD

%type <any> reward
%type <string> ship

// RESULT

%type <any> action

%right EOL

%%

main: TIME COMBAT action EOL {
	ll := $3.(LogLine)
	ll.setTime($1)
	Yacclex.(*Lexer).line = ll
}

action:
	damage {
		$$ = $1
	} |
	heal {
		$$ = $1
	} |
	kill {
		$$ = $1
	} |
	participation {
		$$ = $1
	} |
	start {
		$$ = $1
	} |
	finished {
		$$ = $1
	} |
	reward {
		$$ = $1
	} |
	connect_to_game_session {
		$$ = $1
	} |
	spawn {
		$$ = $1
	} |
	spell {
		$$ = $1
	}

// CONTENTS

// common

object:
	STRING '|' INT {
		$$ = &Object{
			Name: $1,
			ObjectID: $3,
		}
	} |
	STRING '(' STRING ')' '|' INT {
		$$ = &Object{
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
	$$ = &Damage{
		Initiator: *$2.(*Object),
		Recipient: *$4.(*Object),
		DamageFull: $5,
		DamageHull: $6,
		DamageShield: $7,
		Source: $8,
		DamageModifiers: $9.([]DamageModifier),
		FriendlyFire: $10,
		Rocket: $11,
	}
}

damage_modifiers: DAMAGE_MODIFIER {
	$$ = []DamageModifier{DamageModifier($1)}
} | damage_modifiers '|' DAMAGE_MODIFIER {
	$$ = append($$.([]DamageModifier), DamageModifier($3))
}

// Heal

// 19:33:24.732  CMBT   | Heal            Feresey|0000000204 ->          Feresey|0000000204 244.00 Module_Lynx2Shield_T4_Epic
heal: HEAL object ARROW object FLOAT SOURCE {
	$$ = &Heal{
		Initiator: *$2.(*Object),
		Recipient: *$4.(*Object),
		Heal: $5,
		Source: $6,
	}
}

// Kill

// 23:04:35.283  CMBT   | Killed Ship_Bot_ClanShipDroneBig_Empire|0000001860;	 killer Feresey|0000002076 Weapon_Plasmagun_Heavy_T5_Pirate 
// 19:33:59.527  CMBT   | Killed Py6Jl	 Ship_Race3_M_T2_Pirate|0000000248;	 killer Feresey|0000000204 Weapon_Plasmagun_Heavy_T5_Pirate 
// 19:44:55.746  CMBT   | Killed SwarmPack2(georgeatg)|0000001044;	 killer georgeatg|0000001044 (suicide) <FriendlyFire>
kill: KILL player_or_object ';' object source friendly_fire {
	$$ = &Kill{
		Killed: *$2.(*Object),
		Killer: *$4.(*Object),
		Source: $5,
		FriendlyFire: $6,
	}
}

player_or_object: STRING '\t' object {
	$$ = &Object{
		Name: $1,
		PlayerObject: PlayerObject{
			ObjectName: $3.(*Object).Name,
		},
		ObjectID: $3.(*Object).ObjectID,
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
	$$ = &Participant{
		Name: $2,
		Ship: $3,
		Damage: $4.([]any)[0].(float32),
		MostDamageWith: $4.([]any)[1].(string),
		Modifiers: $5,
		FriendlyFire: $6,
	}
}

participation_damage: FLOAT SOURCE {
	$$ = []any{$1, $2}
} | {
	$$ = []any{float32(0), ""}
}

// Start gameplay

local_client_team: INT { $$ = $1 } | /*empty*/ { $$ = 0 }

// 19:33:00.763  CMBT   | ======= Start gameplay 'BombTheBase' map 's1340_thar_aliendebris13', local client team 2 =======
// 19:42:14.670  CMBT   | ======= Start PVE mission 'pve_raid_waterharvest_t5' map 'pve_raid_waterharvest' =======
start: START STRING STRING STRING local_client_team {
	$$ = &Start{
		What: $2,
		GameMode: $3,
		MapName: $4,
		LocalClientTeamID: $5,
	}
}

// Connect to game client

// 19:32:58.666  CMBT   | ======= Connect to game session 50419619 =======
connect_to_game_session: CONNECT_TO_GAME_SESSION_PREFIX INT {
	$$ = &ConnectToGameSession{
		SessionID: $2,
	}
}

// Spawn ship

// 19:46:04.732  CMBT   | Spawn SpaceShip for player5 (KANDIS, #3996749). 'Ship_Race3_H_T5_Uniq'
spawn: SPAWN_PREFIX STRING INT STRING {
	$$ = &Spawn{
		Name: $2,
		ID: $3,
		Ship: $4,
	}
}

// Spell casted

// 19:45:51.379  CMBT   | Spell 'TurretThermo_T5_Mk3' by georgeatg(Module_TurretThermo_T5_Mk3) targets(1): Turret_CruiserAlien_T5R
// 19:45:55.009  CMBT   | Spell 'Spell_AG_DamageDec_Autogen_6859' by Therm0Nuclear(Weapon_PlasmaBlade_T5_Mk3) targets(0):
spell: SPELL_PREFIX spell_type spell_targets {
	$$ = &Spell{
		Normalizer: $2.Normalizer,
		GodGift: $2.GodGift,
		Targets: $3,
	}
} | SPELL_PREFIX spell_type {
	$$ = &Spell{}
}

spell_type: SPELL_NORMALIZER INT { $$.Normalizer = $2 } |
	SPELL_GODGIFT { $$.GodGift = true } | {}

spell_targets: spell_targets STRING { $$ = append($$, $2)} | {}

// Finished

// 19:47:09.448  CMBT   | Gameplay finished. Winner team: 1(PVE_MISSION_COMPLETE_ALT_2). Finish reason: 'Mission complete'. Actual game time 275.9 sec
finished: GAMEPLAY_FINISHED INT STRING STRING FLOAT {
	$$ = &Finished{
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
	$$ = &Reward{
		Recipient: $2,
		Ship: $3,
		Reward: $4,
		RewardType: $5,
		Reason: $6,
	}
}

%%