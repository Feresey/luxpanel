// DO NOT EDIT. This file was auto-generated

package combat

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var reKill = regexp.MustCompile(`(?s)^(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+CMBT\s+\|\s+Killed\s+(((?P<RecipientName>[a-zA-Z0-9_/-]+)\s+(?P<RecipientShip>[a-zA-Z0-9_/-]+))|((?P<RecipientObject>[a-zA-Z0-9_/-]+)|(?P<RecipientObjectName>[a-zA-Z0-9_/-]+)\((?P<RecipientObjectOwner>[a-zA-Z0-9_/-]+)\)))\|(?P<RecipientID>-?\d+);\s+killer\s+(?P<Initiator>[a-zA-Z0-9_/-]+)\|(?P<InitiatorID>-?\d+)(\s+(?P<ActionSource>\(?[a-zA-Z0-9_/-]+\)?))?(\s+(?P<FriendlyFire><FriendlyFire>))?\s*$`)

type Kill struct {
	LogTime              time.Time
	RecipientName        string
	RecipientShip        string
	RecipientObject      string
	RecipientObjectName  string
	RecipientObjectOwner string
	RecipientID          int
	Initiator            string
	InitiatorID          int
	ActionSource         string
	FriendlyFire         bool
}

func (c *Kill) Unmarshal(src string, now time.Time) (err error) {
	res := reKill.FindStringSubmatch(src)
	if len(res) != 17 {
		return fmt.Errorf("%w: %d", errWrongLineFormat, len(res))
	}

	c.LogTime, err = parseField(res[1], "LogTime", parseTime(now))
	if err != nil {
		return err
	}
	c.RecipientName = res[4]
	c.RecipientShip = res[5]
	c.RecipientObject = res[7]
	c.RecipientObjectName = res[8]
	c.RecipientObjectOwner = res[9]
	c.RecipientID, err = parseField(res[10], "RecipientID", strconv.Atoi)
	if err != nil {
		return err
	}
	c.Initiator = res[11]
	c.InitiatorID, err = parseField(res[12], "InitiatorID", strconv.Atoi)
	if err != nil {
		return err
	}
	c.ActionSource = res[14]
	c.FriendlyFire, err = parseField(res[16], "FriendlyFire", parseBool)
	if err != nil {
		return err
	}

	return nil
}

func (c *Kill) Type() CombatLineType {
	return KillLineType
}

func (c *Kill) Time() time.Time {
	return c.LogTime
}
