package gramma

import "time"

type CombatLine struct {
	*Damage
	*Heal
	*Kill
	*Participant
	*ConnectToGameSession
	*Start
	*Finished
	*Reward
}

type ConnectToGameSession struct {
	Time      time.Time
	SessionID int
}

type Start struct {
	Time              time.Time
	What              string
	GameMode          string
	MapName           string
	LocalClientTeamID int
}

type Finished struct {
	Time time.Time

	WinnerTeamID int
	WinReason    string
	FinishReason string
	GameTime     float32
}

type Reward struct {
	Time time.Time

	Recipient  string
	Ship       string
	Reward     int
	RewardType string
	Reason     string
}

type Damage struct {
	Time      time.Time
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
	Time      time.Time
	Initiator Object
	Recipient Object
	Source    string

	Heal float32
}

type Kill struct {
	Time         time.Time
	Killed       Object
	Killer       Object
	Source       string
	FriendlyFire bool
}

type Participant struct {
	Time           time.Time
	Name           string
	Ship           string
	Damage         float32
	MostDamageWith string
	BuffDebuff
	FriendlyFire bool
}

type BuffDebuff struct {
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

type GameLine struct {
	*ClientAddPlayer
	*ClientPlayerLeave
	*ClientConnected
	*ClientConnectionClosed
}

type ClientAddPlayer struct {
	Time           time.Time
	InGamePlayerID int
	Name           string
	ClanTag        String
	PlayerID       int
	Status         int
	TeamID         int
	GroupID        int
}

type ClientPlayerLeave struct {
	Time           time.Time
	InGamePlayerID int
}

type ClientConnected struct {
	Time       time.Time
	ServerAddr string
	MTU        int
}

type ClientConnectionClosed struct {
	Time   time.Time
	Reason String
}
