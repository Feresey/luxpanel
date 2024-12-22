package game

import (
	"time"

	"github.com/Feresey/luxpanel/internal/parser/common"
)

type LogLine interface {
	gameLine()
	setTime(string)
	GetTime(time.Time) time.Time
}

type Time struct{ Time string }

func (Time) gameLine()           {}
func (t *Time) setTime(s string) { t.Time = s }

type ClientAddPlayer struct {
	Time

	InGamePlayerID int
	Name           string
	ClanTag        String
	PlayerID       int
	Status         int
	TeamID         int
	GroupID        int
}

func (l *ClientAddPlayer) GetTime(logTime time.Time) time.Time {
	if l == nil {
		return time.Time{}
	}
	return common.ParseTime(logTime, l.Time.Time)
}

type ClientPlayerLeave struct {
	Time

	InGamePlayerID int
}

func (l *ClientPlayerLeave) GetTime(logTime time.Time) time.Time {
	if l == nil {
		return time.Time{}
	}
	return common.ParseTime(logTime, l.Time.Time)
}

type ClientConnected struct {
	Time

	ServerAddr string
	MTU        int
}

func (l *ClientConnected) GetTime(logTime time.Time) time.Time {
	if l == nil {
		return time.Time{}
	}
	return common.ParseTime(logTime, l.Time.Time)
}

type ClientConnectionClosed struct {
	Time

	Reason String
}

func (l *ClientConnectionClosed) GetTime(logTime time.Time) time.Time {
	if l == nil {
		return time.Time{}
	}
	return common.ParseTime(logTime, l.Time.Time)
}
