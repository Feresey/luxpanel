package game

import "time"

type LogLine interface {
	gameLine()
	setTime(Time)
	GetTime() time.Time
}

type Time time.Time

func (Time) gameLine()         {}
func (c *Time) setTime(t Time) { *c = t }

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

func (l *ClientAddPlayer) GetTime() time.Time {
	if l == nil {
		return time.Time{}
	}
	return time.Time(l.Time)
}

type ClientPlayerLeave struct {
	Time

	InGamePlayerID int
}

func (l *ClientPlayerLeave) GetTime() time.Time {
	if l == nil {
		return time.Time{}
	}
	return time.Time(l.Time)
}

type ClientConnected struct {
	Time

	ServerAddr string
	MTU        int
}

func (l *ClientConnected) GetTime() time.Time {
	if l == nil {
		return time.Time{}
	}
	return time.Time(l.Time)
}

type ClientConnectionClosed struct {
	Time

	Reason String
}

func (l *ClientConnectionClosed) GetTime() time.Time {
	if l == nil {
		return time.Time{}
	}
	return time.Time(l.Time)
}
