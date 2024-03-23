// DO NOT EDIT. This file was auto-generated

package combat

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var reFinishedGameplay = regexp.MustCompile(`(?s)(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+CMBT\s+\|\s+Gameplay finished\. Winner team:\s+(?P<WinnerTeamID>\d+)\((?P<WinReason>.*)\)\.\s+Finish reason:\s+'(?P<FinishReason>.*)'\.\s+Actual game time\s+(?P<GameTime>-?\d+\.\d+)\s+sec\s*$`)

type FinishedGameplay struct {
	LogTime      time.Time
	WinnerTeamID int
	WinReason    string
	FinishReason string
	GameTime     time.Duration
}

func (c *FinishedGameplay) Unmarshal(src string, now time.Time) (err error) {
	res := reFinishedGameplay.FindStringSubmatch(src)
	if len(res) != 6 {
		return fmt.Errorf("%w: %d", errWrongLineFormat, len(res))
	}

	c.LogTime, err = parseField(res[1], "LogTime", parseTime(now))
	if err != nil {
		return err
	}
	c.WinnerTeamID, err = parseField(res[2], "WinnerTeamID", strconv.Atoi)
	if err != nil {
		return err
	}
	c.WinReason = res[3]
	c.FinishReason = res[4]
	c.GameTime, err = parseField(res[5], "GameTime", parseSeconds)
	if err != nil {
		return err
	}

	return nil
}

func (c *FinishedGameplay) Type() CombatLineType {
	return FinishedGameplayLineType
}

func (c *FinishedGameplay) Time() time.Time {
	return c.LogTime
}
