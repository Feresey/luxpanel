// DO NOT EDIT. This file was auto-generated

package game

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	rePlayerLeave      = regexp.MustCompile(`(?s)(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+\|\s+client: player (?P<PlayerID>\d+) leave game\s*$`)
	shortRePlayerLeave = regexp.MustCompile(`(?s)(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+\|\s+client: player (?P<PlayerID>\d+) leave game\s*$`)
)

type PlayerLeave struct {
	LogTime  time.Time
	PlayerID int
}

func (c *PlayerLeave) Unmarshal(src string, now time.Time) (err error) {
	res := rePlayerLeave.FindStringSubmatch(src)
	if len(res) != 3 {
		return fmt.Errorf("%w: %d", errWrongLineFormat, len(res))
	}

	c.LogTime, err = parseField(res[1], "LogTime", parseTime(now))
	if err != nil {
		return err
	}
	c.PlayerID, err = parseField(res[2], "PlayerID", strconv.Atoi)
	if err != nil {
		return err
	}

	return nil
}

func (c *PlayerLeave) Type() GameLineType {
	return PlayerLeaveLineType
}

var emptyPlayerLeaveLogTime time.Time

func (c *PlayerLeave) GetLogTime() time.Time {
	if c == nil || c.LogTime == emptyPlayerLeaveLogTime {
		return emptyPlayerLeaveLogTime
	}
	return c.LogTime

}

var emptyPlayerLeavePlayerID int

func (c *PlayerLeave) GetPlayerID() int {
	if c == nil || c.PlayerID == emptyPlayerLeavePlayerID {
		return emptyPlayerLeavePlayerID
	}
	return c.PlayerID

}
