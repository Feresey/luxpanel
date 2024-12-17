package combat

import (
	"time"
)

type Time time.Time

func (Time) combatLine()       {}
func (c *Time) setTime(t Time) { *c = t }

type LogLine interface {
	combatLine()
	setTime(Time)
	GetTime() time.Time
}

type ConnectToGameSession struct {
	Time

	SessionID int
}

func (l *ConnectToGameSession) GetTime() time.Time {
	if l == nil {
		return time.Time{}
	}
	return time.Time(l.Time)
}

type Start struct {
	Time

	What              string
	GameMode          string
	MapName           string
	LocalClientTeamID int
}

func (l *Start) GetTime() time.Time {
	if l == nil {
		return time.Time{}
	}
	return time.Time(l.Time)
}

type Finished struct {
	Time

	WinnerTeamID int
	WinReason    string
	FinishReason string
	GameTime     float32
}

func (l *Finished) GetTime() time.Time {
	if l == nil {
		return time.Time{}
	}
	return time.Time(l.Time)
}

type Reward struct {
	Time

	Recipient  string
	Ship       string
	Reward     int
	RewardType string
	Reason     string
}

func (l *Reward) GetTime() time.Time {
	if l == nil {
		return time.Time{}
	}
	return time.Time(l.Time)
}

type Damage struct {
	Time

	Initiator Object
	Recipient Object
	Source    string

	DamageFull      float32
	DamageHull      float32
	DamageShield    float32
	DamageModifiers DamageModifiers

	FriendlyFire bool
	Rocket       int
}

func (l *Damage) GetTime() time.Time {
	if l == nil {
		return time.Time{}
	}
	return time.Time(l.Time)
}

type Heal struct {
	Time

	Initiator Object
	Recipient Object
	Source    string

	Heal float32
}

func (l *Heal) GetTime() time.Time {
	if l == nil {
		return time.Time{}
	}
	return time.Time(l.Time)
}

type Kill struct {
	Time

	Killed       Object
	Killer       Object
	Source       string
	FriendlyFire bool
}

func (l *Kill) GetTime() time.Time {
	if l == nil {
		return time.Time{}
	}
	return time.Time(l.Time)
}

type Participant struct {
	Time

	Name           string
	Ship           string
	Damage         float32
	MostDamageWith string
	FriendlyFire   bool

	Modifiers ParticipationModifiers
}

func (l *Participant) GetTime() time.Time {
	if l == nil {
		return time.Time{}
	}
	return time.Time(l.Time)
}

type ParticipationModifier string
type ParticipationModifiers []ParticipationModifier

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
type DamageModifiers []DamageModifier
