package gramma

import "time"

type combatLineImpl struct {
	Time time.Time
}

func (combatLineImpl) combatLine() {}

type CombatLine interface {
	combatLine()
}

type ConnectToGameSession struct {
	combatLineImpl

	SessionID int
}

type Start struct {
	combatLineImpl

	What              string
	GameMode          string
	MapName           string
	LocalClientTeamID int
}

type Finished struct {
	combatLineImpl

	WinnerTeamID int
	WinReason    string
	FinishReason string
	GameTime     float32
}

type Reward struct {
	combatLineImpl

	Recipient  string
	Ship       string
	Reward     int
	RewardType string
	Reason     string
}

type Damage struct {
	combatLineImpl

	Initiator Object
	Recipient Object
	Source    string

	DamageFull      float32
	DamageHull      float32
	DamageShield    float32
	DamageModifiers DamageModifiersMap

	FriendlyFire bool
	Rocket       int
}

type Heal struct {
	combatLineImpl

	Initiator Object
	Recipient Object
	Source    string

	Heal float32
}

type Kill struct {
	combatLineImpl

	Killed       Object
	Killer       Object
	Source       string
	FriendlyFire bool
}

type Participant struct {
	combatLineImpl

	Name           string
	Ship           string
	Damage         float32
	MostDamageWith string
	BuffDebuff
	FriendlyFire bool
}

type BuffDebuff struct {
	combatLineImpl

	IsBuff   bool
	IsDebuff bool
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

type gameLineImpl struct {
	Time time.Time
}

func (gameLineImpl) gameLine() {}

type GameLine interface {
	gameLine()
}

type ClientAddPlayer struct {
	gameLineImpl

	InGamePlayerID int
	Name           string
	ClanTag        String
	PlayerID       int
	Status         int
	TeamID         int
	GroupID        int
}

type ClientPlayerLeave struct {
	gameLineImpl

	InGamePlayerID int
}

type ClientConnected struct {
	gameLineImpl

	ServerAddr string
	MTU        int
}

type ClientConnectionClosed struct {
	gameLineImpl

	Reason String
}

type LogLine struct {
	Combat CombatLine
	Game   GameLine
}
