package combat

import (
	"time"

	"github.com/Feresey/luxpanel/internal/parser/common"
)

type Time struct{ Time string }

func (Time) combatLine()         {}
func (t *Time) setTime(s string) { t.Time = s }

type LogLine interface {
	combatLine()
	setTime(string)
	GetTime(time.Time) time.Time
}

type ConnectToGameSession struct {
	Time

	SessionID int
}

func (l *ConnectToGameSession) IsEmpty() bool {
	return l.SessionID == 0
}

func (l *ConnectToGameSession) GetTime(logTime time.Time) time.Time {
	if l == nil {
		return time.Time{}
	}
	return common.ParseTime(logTime, l.Time.Time)
}

type Start struct {
	Time

	What              string
	GameMode          string
	MapName           string
	LocalClientTeamID int
}

func (l *Start) GetTime(logTime time.Time) time.Time {
	if l == nil {
		return time.Time{}
	}
	return common.ParseTime(logTime, l.Time.Time)
}

func (l *Start) IsEmpty() bool {
	if l.What == "" || l.MapName == "" || l.LocalClientTeamID == 0 {
		return true
	}
	return false
}

type Finished struct {
	Time

	WinnerTeamID int
	WinReason    string
	FinishReason string
	GameTime     float32
}

func (l *Finished) GetTime(logTime time.Time) time.Time {
	if l == nil {
		return time.Time{}
	}
	return common.ParseTime(logTime, l.Time.Time)
}

func (l *Finished) IsEmpty() bool {
	if l.WinnerTeamID == 0 || l.WinReason == "" || l.FinishReason == "" || l.GameTime < 1 {
		return true
	}
	return false
}

type Reward struct {
	Time

	Recipient  string
	Ship       string
	Reward     int
	RewardType string
	Reason     string
}

func (l *Reward) GetTime(logTime time.Time) time.Time {
	if l == nil {
		return time.Time{}
	}
	return common.ParseTime(logTime, l.Time.Time)
}

func (l *Reward) IsEmpty() bool {
	if l.Recipient == "" {
		return true
	}
	return false
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

func (l *Damage) GetTime(logTime time.Time) time.Time {
	if l == nil {
		return time.Time{}
	}
	return common.ParseTime(logTime, l.Time.Time)
}

func (l *Damage) IsEmpty() bool {
	if l.DamageFull < 0.001 {
		return true
	}
	return false
}

type Heal struct {
	Time

	Initiator Object
	Recipient Object
	Source    string

	Heal float32
}

func (l *Heal) GetTime(logTime time.Time) time.Time {
	if l == nil {
		return time.Time{}
	}
	return common.ParseTime(logTime, l.Time.Time)
}

func (l *Heal) IsEmpty() bool {
	if l.Heal < 0.001 {
		return true
	}
	return false
}

type Kill struct {
	Time

	Killed       Object
	Killer       Object
	Source       string
	FriendlyFire bool
}

func (l *Kill) GetTime(logTime time.Time) time.Time {
	if l == nil {
		return time.Time{}
	}
	return common.ParseTime(logTime, l.Time.Time)
}

func (l *Kill) IsEmpty() bool {
	if l.Killed.ObjectID == 0 {
		return true
	}
	return false
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

func (l *Participant) GetTime(logTime time.Time) time.Time {
	if l == nil {
		return time.Time{}
	}
	return common.ParseTime(logTime, l.Time.Time)
}

func (l *Participant) IsEmpty() bool {
	if l.Name == "" {
		return true
	}
	return false
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
