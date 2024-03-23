// DO NOT EDIT. This file was auto-generated

package game

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var reAddPlayer = regexp.MustCompile(`(?s)(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+\|\s+client: ADD_PLAYER (?P<SessionPlayerID>\d+)\s+\((?P<PlayerName>[a-zA-Z0-9_/-]+)(\s+\[(?P<PlayerCorpTag>[a-zA-Z0-9_/-]*)\])?,\s+(?P<PlayerID>\d+)\)\s+status\s+(?P<ConnectionStatus>\d+)\s+team\s+(?P<TeamID>\d+)(\s+group\s+(?P<GroupID>\d+))?\s*$`)

type AddPlayer struct {
	LogTime          time.Time
	SessionPlayerID  int
	PlayerName       string
	PlayerCorpTag    string
	PlayerID         int
	ConnectionStatus int
	TeamID           int
	GroupID          *int
}

func (c *AddPlayer) Unmarshal(src string, now time.Time) (err error) {
	res := reAddPlayer.FindStringSubmatch(src)
	if len(res) != 11 {
		return fmt.Errorf("%w: %d", errWrongLineFormat, len(res))
	}

	c.LogTime, err = parseField(res[1], "LogTime", parseTime(now))
	if err != nil {
		return err
	}
	c.SessionPlayerID, err = parseField(res[2], "SessionPlayerID", strconv.Atoi)
	if err != nil {
		return err
	}
	c.PlayerName = res[3]
	c.PlayerCorpTag = res[5]
	c.PlayerID, err = parseField(res[6], "PlayerID", strconv.Atoi)
	if err != nil {
		return err
	}
	c.ConnectionStatus, err = parseField(res[7], "ConnectionStatus", strconv.Atoi)
	if err != nil {
		return err
	}
	c.TeamID, err = parseField(res[8], "TeamID", strconv.Atoi)
	if err != nil {
		return err
	}
	c.GroupID, err = parseField(res[10], "GroupID", parseOptional(strconv.Atoi))
	if err != nil {
		return err
	}

	return nil
}

func (c *AddPlayer) Type() GameLineType {
	return AddPlayerLineType
}

func (c *AddPlayer) Time() time.Time {
	return c.LogTime
}
