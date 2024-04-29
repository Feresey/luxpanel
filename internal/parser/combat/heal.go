// DO NOT EDIT. This file was auto-generated

package combat

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var reHeal = regexp.MustCompile(`(?s)^(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+CMBT\s+\|\s+Heal\s+(?P<Initiator>[a-zA-Z][a-zA-Z0-9_/-]*)\|(?P<InitiatorID>-?\d+)\s+->\s+((?P<Recipient>[a-zA-Z][a-zA-Z0-9_/-]*)|(?P<ObjectName>[a-zA-Z][a-zA-Z0-9_/-]*)\((?P<ObjectOwner>[a-zA-Z][a-zA-Z0-9_/-]*)\))\|(?P<RecipientID>-?\d+)\s+(?P<Heal>-?\d+\.\d+)\s+(?P<ActionSource>\(?[a-zA-Z0-9_/-]+\)?)?\s*$`)

type Heal struct {
	LogTime      time.Time
	Initiator    string
	InitiatorID  int
	Recipient    string
	ObjectName   string
	ObjectOwner  string
	RecipientID  int
	Heal         float32
	ActionSource string
}

func (c *Heal) Unmarshal(src string, now time.Time) (err error) {
	res := reHeal.FindStringSubmatch(src)
	if len(res) != 11 {
		return fmt.Errorf("%w: %d", errWrongLineFormat, len(res))
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

var emptyHealLogTime time.Time

func (c *Heal) GetLogTime() time.Time {
	if c == nil || c.LogTime == emptyHealLogTime {
		return emptyHealLogTime
	}
	return c.LogTime

}

var emptyHealInitiator string

func (c *Heal) GetInitiator() string {
	if c == nil || c.Initiator == emptyHealInitiator {
		return emptyHealInitiator
	}
	return c.Initiator

}

var emptyHealInitiatorID int

func (c *Heal) GetInitiatorID() int {
	if c == nil || c.InitiatorID == emptyHealInitiatorID {
		return emptyHealInitiatorID
	}
	return c.InitiatorID

}

var emptyHealRecipient string

func (c *Heal) GetRecipient() string {
	if c == nil || c.Recipient == emptyHealRecipient {
		return emptyHealRecipient
	}
	return c.Recipient

}

var emptyHealObjectName string

func (c *Heal) GetObjectName() string {
	if c == nil || c.ObjectName == emptyHealObjectName {
		return emptyHealObjectName
	}
	return c.ObjectName

}

var emptyHealObjectOwner string

func (c *Heal) GetObjectOwner() string {
	if c == nil || c.ObjectOwner == emptyHealObjectOwner {
		return emptyHealObjectOwner
	}
	return c.ObjectOwner

}

var emptyHealRecipientID int

func (c *Heal) GetRecipientID() int {
	if c == nil || c.RecipientID == emptyHealRecipientID {
		return emptyHealRecipientID
	}
	return c.RecipientID

}

var emptyHealHeal float32

func (c *Heal) GetHeal() float32 {
	if c == nil || c.Heal == emptyHealHeal {
		return emptyHealHeal
	}
	return c.Heal

}

var emptyHealActionSource string

func (c *Heal) GetActionSource() string {
	if c == nil || c.ActionSource == emptyHealActionSource {
		return emptyHealActionSource
	}
	return c.ActionSource

}
