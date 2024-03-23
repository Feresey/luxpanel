// DO NOT EDIT. This file was auto-generated

package combat

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var reConnectToGameSession = regexp.MustCompile(`(?s)(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+CMBT\s+\| ======= Connect to game session (?P<SessionID>\d+) =======\s*$`)

type ConnectToGameSession struct {
	LogTime   time.Time
	SessionID int
}

func (c *ConnectToGameSession) Unmarshal(src string, now time.Time) (err error) {
	res := reConnectToGameSession.FindStringSubmatch(src)
	if len(res) != 3 {
		return fmt.Errorf("%w: %d", errWrongLineFormat, len(res))
	}

	c.LogTime, err = parseField(res[1], "LogTime", parseTime(now))
	if err != nil {
		return err
	}
	c.SessionID, err = parseField(res[2], "SessionID", strconv.Atoi)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConnectToGameSession) Type() CombatLineType {
	return ConnectToGameSessionLineType
}

func (c *ConnectToGameSession) Time() time.Time {
	return c.LogTime
}
