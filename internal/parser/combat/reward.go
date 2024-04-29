// DO NOT EDIT. This file was auto-generated

package combat

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var reReward = regexp.MustCompile(`(?s)^(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+CMBT\s+\|\s+Reward\s+(?P<Recipient>[a-zA-Z][a-zA-Z0-9_/-]*)(\s(?P<Ship>[a-zA-Z][a-zA-Z0-9_/-]*))?\s*\t\s*(?P<Reward>\d+)\s+(?P<RewardType>((experience)|(credits)|(effective points)))\s+for\s(?P<Reason>.+)\s*$`)

type Reward struct {
	LogTime    time.Time
	Recipient  string
	Ship       string
	Reward     int
	RewardType string
	Reason     string
}

func (c *Reward) Unmarshal(src string, now time.Time) (err error) {
	res := reReward.FindStringSubmatch(src)
	if len(res) != 12 {
		return fmt.Errorf("%w: %d", errWrongLineFormat, len(res))
	}

	c.LogTime, err = parseField(res[1], "LogTime", parseTime(now))
	if err != nil {
		return err
	}
	c.Recipient = res[2]
	c.Ship = res[4]
	c.Reward, err = parseField(res[5], "Reward", strconv.Atoi)
	if err != nil {
		return err
	}
	c.RewardType = res[6]
	c.Reason = res[11]

	return nil
}

func (c *Reward) Type() CombatLineType {
	return RewardLineType
}

var emptyRewardLogTime time.Time

func (c *Reward) GetLogTime() time.Time {
	if c == nil || c.LogTime == emptyRewardLogTime {
		return emptyRewardLogTime
	}
	return c.LogTime

}

var emptyRewardRecipient string

func (c *Reward) GetRecipient() string {
	if c == nil || c.Recipient == emptyRewardRecipient {
		return emptyRewardRecipient
	}
	return c.Recipient

}

var emptyRewardShip string

func (c *Reward) GetShip() string {
	if c == nil || c.Ship == emptyRewardShip {
		return emptyRewardShip
	}
	return c.Ship

}

var emptyRewardReward int

func (c *Reward) GetReward() int {
	if c == nil || c.Reward == emptyRewardReward {
		return emptyRewardReward
	}
	return c.Reward

}

var emptyRewardRewardType string

func (c *Reward) GetRewardType() string {
	if c == nil || c.RewardType == emptyRewardRewardType {
		return emptyRewardRewardType
	}
	return c.RewardType

}

var emptyRewardReason string

func (c *Reward) GetReason() string {
	if c == nil || c.Reason == emptyRewardReason {
		return emptyRewardReason
	}
	return c.Reason

}
