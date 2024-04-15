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

var emptyAddPlayerLogTime time.Time

func (c *AddPlayer) GetLogTime() time.Time {
	if c == nil || c.LogTime == emptyAddPlayerLogTime {
		return emptyAddPlayerLogTime
	}
	return c.LogTime

}

var emptyAddPlayerSessionPlayerID int

func (c *AddPlayer) GetSessionPlayerID() int {
	if c == nil || c.SessionPlayerID == emptyAddPlayerSessionPlayerID {
		return emptyAddPlayerSessionPlayerID
	}
	return c.SessionPlayerID

}

var emptyAddPlayerPlayerName string

func (c *AddPlayer) GetPlayerName() string {
	if c == nil || c.PlayerName == emptyAddPlayerPlayerName {
		return emptyAddPlayerPlayerName
	}
	return c.PlayerName

}

var emptyAddPlayerPlayerCorpTag string

func (c *AddPlayer) GetPlayerCorpTag() string {
	if c == nil || c.PlayerCorpTag == emptyAddPlayerPlayerCorpTag {
		return emptyAddPlayerPlayerCorpTag
	}
	return c.PlayerCorpTag

}

var emptyAddPlayerPlayerID int

func (c *AddPlayer) GetPlayerID() int {
	if c == nil || c.PlayerID == emptyAddPlayerPlayerID {
		return emptyAddPlayerPlayerID
	}
	return c.PlayerID

}

var emptyAddPlayerConnectionStatus int

func (c *AddPlayer) GetConnectionStatus() int {
	if c == nil || c.ConnectionStatus == emptyAddPlayerConnectionStatus {
		return emptyAddPlayerConnectionStatus
	}
	return c.ConnectionStatus

}

var emptyAddPlayerTeamID int

func (c *AddPlayer) GetTeamID() int {
	if c == nil || c.TeamID == emptyAddPlayerTeamID {
		return emptyAddPlayerTeamID
	}
	return c.TeamID

}

var emptyAddPlayerGroupID *int

func (c *AddPlayer) GetGroupID() *int {
	if c == nil || c.GroupID == emptyAddPlayerGroupID {
		return emptyAddPlayerGroupID
	}
	return c.GroupID

}
