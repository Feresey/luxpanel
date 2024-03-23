// DO NOT EDIT. This file was auto-generated

package game

import (
	"fmt"
	"regexp"
	"time"
)

var reConnectionClosed = regexp.MustCompile(`(?s)(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+\|\s+client: connection closed\. (?P<CloseReason>[A-Z_]+)\s*$`)

type ConnectionClosed struct {
	LogTime     time.Time
	CloseReason ConnectionClosedReason
}

func (c *ConnectionClosed) Unmarhsal(src string, now time.Time) (err error) {
	res := reConnectionClosed.FindStringSubmatch(src)
	if len(res) != 3 {
		return fmt.Errorf("%w: %d", ErrWrongLineFormat, len(res))
	}

	c.LogTime, err = parseField(res[1], "LogTime", parseTime(now))
	if err != nil {
		return err
	}
	c.CloseReason, err = parseField(res[2], "CloseReason", parseConnectionClosedReason)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConnectionClosed) Type() GameLineType {
	return ConnectionClosedLineType
}

func (c *ConnectionClosed) Time() time.Time {
	return c.LogTime
}
