%{
package main

import "time"

type LogLine[T any] struct {
    LogTime time.Time
    Line T
}

type CombatLine struct {
    Damage *Damage
}

type Damage struct {
	Initiator Object
	Recipient Object
	Reason    string

	DamageFull      float32
	DamageHull      float32
	DamageShield    float32
	DamageModifiers DamageModifiersMap

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

%}

%union {
    String string
    Int int
    Float float32
    Bool bool

    Time time.Time

    CombatLine *LogLine[*CombatLine]
    CombatLines []*LogLine[*CombatLine]
    Combat *CombatLine
    Damage *Damage
    DamageSucks damageSucks
    DamageModifiersMap DamageModifiersMap
    Object Object
}

%token ARROW DAMAGE COMBAT EOL

%token LBRACE RBRACE VSLASH
%token DAMAGE_HULL_START DAMAGE_SHIELD_START

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
%type <Bool> friendly_fire
%type <DamageModifiersMap> damage_modifiers
%type <String> source
%type <DamageSucks> damage_sucks

%%

start: lines {
    yylex.(*lexer).res = $1
}

lines: lines line { $$ = append($1, $2)} | line { $$ = append($$, $1)} | EOL {};

line: TIME combat_line EOL {
    // println($1.String())
    $$ = &LogLine[*CombatLine]{
        LogTime: $1,
        Line: $2,
    }
}

combat_line: COMBAT DAMAGE damage {
    $$ = &CombatLine{
        Damage: $3,
    }
}

// 21:41:15.644  CMBT   | Damage         aSpectro|0000002757 ->          Nafalar|0000002677 117.55 (h:0.00 s:117.55) Weapon_KineticStkBomb_T5_Epic KINETIC|PRIMARY_WEAPON|CRIT  
damage: object ARROW object FLOAT LBRACE DAMAGE_HULL_START FLOAT DAMAGE_SHIELD_START FLOAT RBRACE damage_sucks friendly_fire {
    $$ = &Damage{
        Initiator: $1,
        Recipient: $3,
        DamageFull: $4,
        DamageHull: $7,
        DamageShield: $9,
        Reason: $11.Source,
        DamageModifiers: $11.Modifiers,
        FriendlyFire: $12,
    }
}

object: STRING VSLASH INT {
    $$ = Object{
        Name: $1,
        ObjectID: $3,
    }
}
| STRING LBRACE STRING RBRACE VSLASH INT {
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

source: STRING { $$ = $1} | LBRACE STRING RBRACE { $$ = $2}

damage_modifiers: STRING {
    $$ = DamageModifiersMap{
        DamageModifier($1): {},
    }
} | damage_modifiers VSLASH STRING {
    $$[DamageModifier($3)] = struct{}{}
}

friendly_fire: FRIENDLY_FIRE { $$ = true } | /* empty */ {};

%%