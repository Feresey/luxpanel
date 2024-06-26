// DO NOT EDIT. This file was auto-generated

package game

import (
	"fmt"
	"regexp"
	"time"
)

var (
	reConnected      = regexp.MustCompile(`(?s)(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+\|\s+client: connected to`)
	shortReConnected = regexp.MustCompile(`(?s)(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+\|\s+client: connected to`)
)

type Connected struct {
	LogTime time.Time
}

func (c *Connected) Unmarshal(src string, now time.Time) (err error) {
	res := reConnected.FindStringSubmatch(src)
	if len(res) != 2 {
		return fmt.Errorf("%w: %d", errWrongLineFormat, len(res))
	}

	c.LogTime, err = parseField(res[1], "LogTime", parseTime(now))
	if err != nil {
		return err
	}

	return nil
}

func (c *Connected) Type() GameLineType {
	return ConnectedLineType
}

var emptyConnectedLogTime time.Time

func (c *Connected) GetLogTime() time.Time {
	if c == nil || c.LogTime == emptyConnectedLogTime {
		return emptyConnectedLogTime
	}
	return c.LogTime

}
