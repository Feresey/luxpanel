package combat

import (
	"fmt"
	"regexp"
	"time"
	"strconv"
)

var reHeal = regexp.MustCompile(`(?s)^(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+CMBT\s+\|\s+Heal\s+(?P<Initiator>[a-zA-Z0-9_/-]+)\|(?P<InitiatorID>-?\d+)\s+->\s+((?P<Recipient>[a-zA-Z0-9_/-]+)|(?P<ObjectName>[a-zA-Z0-9_/-]+)\((?P<ObjectOwner>[a-zA-Z0-9_/-]+)\))\|(?P<RecipientID>-?\d+)\s+(?P<Heal>-?\d+\.\d+)\s+(?P<ActionSource>\(?[a-zA-Z0-9_/-]+\)?)?\s*$`)

type Heal struct {
	LogTime time.Time
	Initiator string
	InitiatorID int
	Recipient string
	ObjectName string
	ObjectOwner string
	RecipientID int
	Heal float32
	ActionSource string
}

func (c *Heal) Unmarhsal(src string, now time.Time) (err error) {
	res := reHeal.FindStringSubmatch(src)
	if len(res) != 11 {
		return fmt.Errorf("%w: %d", ErrWrongLineFormat, len(res))
	}

	c.LogTime, err = parseField(res[1], "LogTime", parseTime(now))
	if err != nil {
		return err
	}
	c.Initiator = res[2]
	c.InitiatorID, err = parseField(res[3], "InitiatorID", strconv.Atoi)
	if err != nil {
		return err
	}
	c.Recipient = res[5]
	c.ObjectName = res[6]
	c.ObjectOwner = res[7]
	c.RecipientID, err = parseField(res[8], "RecipientID", strconv.Atoi)
	if err != nil {
		return err
	}
	c.Heal, err = parseField(res[9], "Heal", parseFloat)
	if err != nil {
		return err
	}
	c.ActionSource = res[10]

	return nil
}

func (c *Heal) Type() CombatLineType {
	return HealLineType
}

func (c *Heal) Time() time.Time {
	return c.LogTime
}
