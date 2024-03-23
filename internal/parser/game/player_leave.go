// DO NOT EDIT. This file was auto-generated

package game

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var rePlayerLeave = regexp.MustCompile(`(?s)(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+\|\s+client: player (?P<PlayerID>\d+) leave game\s*$`)

type PlayerLeave struct {
	LogTime  time.Time
	PlayerID int
}

func (c *PlayerLeave) Unmarhsal(src string, now time.Time) (err error) {
	res := rePlayerLeave.FindStringSubmatch(src)
	if len(res) != 3 {
		return fmt.Errorf("%w: %d", ErrWrongLineFormat, len(res))
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

func (c *PlayerLeave) Time() time.Time {
	return c.LogTime
}
