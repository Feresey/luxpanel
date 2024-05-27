// DO NOT EDIT. This file was auto-generated

package combat

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	reStartGameplay      = regexp.MustCompile(`(?s)^(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+CMBT\s+\| ======= Start .* '(?P<GameMode>.+)' map '(?P<MapName>.+)'(, local client team (?P<ClientTeamID>\d+))?\s*=======\s*$`)
	shortReStartGameplay = regexp.MustCompile(`======= Start`)
)

type StartGameplay struct {
	LogTime      time.Time
	GameMode     string
	MapName      string
	ClientTeamID *int
}

func (c *StartGameplay) Unmarshal(src string, now time.Time) (err error) {
	res := reStartGameplay.FindStringSubmatch(src)
	if len(res) != 6 {
		return fmt.Errorf("%w: %d", errWrongLineFormat, len(res))
	}

	c.LogTime, err = parseField(res[1], "LogTime", parseTime(now))
	if err != nil {
		return err
	}
	c.GameMode = res[2]
	c.MapName = res[3]
	c.ClientTeamID, err = parseField(res[5], "ClientTeamID", parseOptional(strconv.Atoi))
	if err != nil {
		return err
	}

	return nil
}

func (c *StartGameplay) Type() CombatLineType {
	return StartGameplayLineType
}

var emptyStartGameplayLogTime time.Time

func (c *StartGameplay) GetLogTime() time.Time {
	if c == nil || c.LogTime == emptyStartGameplayLogTime {
		return emptyStartGameplayLogTime
	}
	return c.LogTime

}

var emptyStartGameplayGameMode string

func (c *StartGameplay) GetGameMode() string {
	if c == nil || c.GameMode == emptyStartGameplayGameMode {
		return emptyStartGameplayGameMode
	}
	return c.GameMode

}

var emptyStartGameplayMapName string

func (c *StartGameplay) GetMapName() string {
	if c == nil || c.MapName == emptyStartGameplayMapName {
		return emptyStartGameplayMapName
	}
	return c.MapName

}

var emptyStartGameplayClientTeamID *int

func (c *StartGameplay) GetClientTeamID() *int {
	if c == nil || c.ClientTeamID == emptyStartGameplayClientTeamID {
		return emptyStartGameplayClientTeamID
	}
	return c.ClientTeamID

}
