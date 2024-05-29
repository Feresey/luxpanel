%{
package main

import (
	"time"
	"strings"
)

type LogLine[T any] struct {
	LogTime time.Time
	Line T
}

type CombatLine struct {
	*Damage
	*Heal
	*ConnectToGameSession
	*Start
	*Kill
}

type ConnectToGameSession struct {
	SessionID int
}

type Start struct {
	What              string
	GameMode          string
	MapName           string
	LocalClientTeamID int
}

type Damage struct {
	Initiator Object
	Recipient Object
	Source    string

	DamageFull      float32
	DamageHull      float32
	DamageShield    float32
	DamageModifiers DamageModifiersMap

	FriendlyFire bool
}

type Heal struct {
	Initiator Object
	Recipient Object
	Source    string

	Heal float32
}

type Kill struct {
	Killed Object
	Killer Object
	Source string
	FriendlyFire bool
}

type PlayerObject struct {
	ObjectName  string
	ObjectOwner string
}

type Object struct {
	Name string
	PlayerObject
	ObjectID int
}

type DamageModifier string
type DamageModifiersMap map[DamageModifier]struct{}

// TODO string to enum 1 << iota
const (
	DamageTypeEMP         DamageModifier = "EMP"
	DamageTypeKinetic     DamageModifier = "KINETIC"
	DamageTypeThermal     DamageModifier = "THERMAL"
	DamageTypeTrueDamage  DamageModifier = "TRUE_DAMAGE"
	DamageUnintention     DamageModifier = "UNINTENTION"
	DamageCrit            DamageModifier = "CRIT"
	DamageExplosion       DamageModifier = "EXPLOSION"
	DamageCollision       DamageModifier = "COLLISION"
	DamageWeaponPrimary   DamageModifier = "PRIMARY_WEAPON"
	DamageWeaponSecondary DamageModifier = "SECONDARY_WEAPON"
	DamageIgnoreScale     DamageModifier = "IGNORE_DAMAGE_SCALE"
	DamageIgnoreShield    DamageModifier = "IGNORE_SHIELD"
	DamageModule          DamageModifier = "MODULE"
)

type damageSucks struct {
	Source string
	Modifiers DamageModifiersMap
}

type killSucks struct {
	Source string
	FriendlyFire bool
}

%}

%union {
	String  string
	Strings []string
	Int     int
	Float   float32
	Bool    bool

	Time time.Time

	CombatLine  *LogLine[*CombatLine]
	CombatLines []*LogLine[*CombatLine]
	Combat      *CombatLine

	Damage          *Damage
	DamageSucks     damageSucks
	DamageModifiers DamageModifiersMap
	Object          Object

	Heal *Heal
	Kill *Kill
	KillSucks killSucks
}

%token COMBAT GAME EOL
%token ARROW

%token DAMAGE DAMAGE_HULL_START DAMAGE_SHIELD_START
%token HEAL
%token KILL

%token CONNECT_TO_GAME_SESSION_PREFIX EQ_DELIM LOCAL_CLIENT_TEAM START

%token <Float> FLOAT
%token <Int> INT
%token <String> STRING
%token <Bool> FRIENDLY_FIRE

%token <Time> TIME

%type <CombatLines> lines
%type <CombatLine> line
%type <Combat> combat_line


%type <Damage> damage
%type <Object> object
%type <Object> player_or_object
%type <Bool> friendly_fire
%type <DamageModifiers> damage_modifiers
%type <String> source
%type <DamageSucks> damage_sucks
%type <KillSucks> kill_sucks
%type <Int> local_client_team
%type <Strings> strings

%type <Heal> heal
%type <Kill> kill

%%

start: lines {
	yylex.(*lexer).res = $1
}

lines: lines line { $$ = append($1, $2)} | line { $$ = append($$, $1)} | EOL {};

line: TIME COMBAT combat_line EOL {
	// println($1.String())
	$$ = &LogLine[*CombatLine]{
		LogTime: $1,
		Line: $3,
	}
}

// 19:32:58.666  CMBT   | ======= Connect to game session 50419619 =======
// 19:33:00.763  CMBT   | ======= Start gameplay 'BombTheBase' map 's1340_thar_aliendebris13', local client team 2 =======
// 19:42:14.670  CMBT   | ======= Start PVE mission 'pve_raid_waterharvest_t5' map 'pve_raid_waterharvest' =======
combat_line: DAMAGE damage {
	$$ = &CombatLine{
		Damage: $2,
	}
} | HEAL heal {
	$$ = &CombatLine{
		Heal: $2,
	}
} | KILL kill {
	$$ = &CombatLine{
		Kill: $2,
	}
} | EQ_DELIM CONNECT_TO_GAME_SESSION_PREFIX INT EQ_DELIM {
	$$ = &CombatLine{
		ConnectToGameSession: &ConnectToGameSession{SessionID: $3},
	}
} | EQ_DELIM START strings '\'' STRING '\'' STRING '\'' STRING '\'' local_client_team EQ_DELIM {
	$$ = &CombatLine{
		Start: &Start{
			What: strings.Join($3, " "),
			GameMode: $5,
			MapName: $9,
			LocalClientTeamID: $11,
		},
	}
}

strings: STRING { $$ = append($$, $1) } | strings STRING { $$ = append($1, $2) }

local_client_team: ',' LOCAL_CLIENT_TEAM INT { $$ = $3 } | /*empty*/ {}

// 21:41:15.644  CMBT   | Damage         aSpectro|0000002757 ->          Nafalar|0000002677 117.55 (h:0.00 s:117.55) Weapon_KineticStkBomb_T5_Epic KINETIC|PRIMARY_WEAPON|CRIT  
damage: object ARROW object FLOAT '(' DAMAGE_HULL_START FLOAT DAMAGE_SHIELD_START FLOAT ')' damage_sucks friendly_fire {
	$$ = &Damage{
		Initiator: $1,
		Recipient: $3,
		DamageFull: $4,
		DamageHull: $7,
		DamageShield: $9,
		Source: $11.Source,
		DamageModifiers: $11.Modifiers,
		FriendlyFire: $12,
	}
}

// 19:33:24.732  CMBT   | Heal            Feresey|0000000204 ->          Feresey|0000000204 244.00 Module_Lynx2Shield_T4_Epic
heal: object ARROW object FLOAT STRING {
	$$ = &Heal{
		Initiator: $1,
		Recipient: $3,
		Heal: $4,
		Source: $5,
	}
}

// 23:04:35.283  CMBT   | Killed Ship_Bot_ClanShipDroneBig_Empire|0000001860;	 killer Feresey|0000002076 Weapon_Plasmagun_Heavy_T5_Pirate 
// 19:33:59.527  CMBT   | Killed Py6Jl	 Ship_Race3_M_T2_Pirate|0000000248;	 killer Feresey|0000000204 Weapon_Plasmagun_Heavy_T5_Pirate 
// 19:44:55.746  CMBT   | Killed SwarmPack2(georgeatg)|0000001044;	 killer georgeatg|0000001044 (suicide) <FriendlyFire>
kill: player_or_object ';' '\t' STRING object kill_sucks {
	if $4 != "killer" {
		yylex.Error("not a killer")
	}
	$$ = &Kill{
		Killed: $1,
		Killer: $5,
		Source: $6.Source,
		FriendlyFire: $6.FriendlyFire,
	}
}

kill_sucks: source friendly_fire {
	$$.Source =  $1
	$$.FriendlyFire = $2
} | friendly_fire { $$.FriendlyFire = $1}

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

damage_sucks: source damage_modifiers {
	$$.Source = $1
	$$.Modifiers = $2
} | damage_modifiers {
	$$.Modifiers = $1
}

source: STRING { $$ = $1} | '(' STRING ')' { $$ = $2}

damage_modifiers: STRING {
	$$ = DamageModifiersMap{
		DamageModifier($1): {},
	}
} | damage_modifiers '|' STRING {
	$$[DamageModifier($3)] = struct{}{}
}

friendly_fire: FRIENDLY_FIRE { $$ = true } | /* empty */ {};

%%