// DO NOT EDIT. This file was auto-generated

package combat

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var reDamage = regexp.MustCompile(`(?s)^(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+CMBT\s+\|\s+Damage\s+(?P<Initiator>[a-zA-Z0-9_/-]+)\|(?P<InitiatorID>-?\d+)\s+->\s+((?P<Recipient>[a-zA-Z0-9_/-]+)|(?P<ObjectName>[a-zA-Z0-9_/-]+)\((?P<ObjectOwner>[a-zA-Z0-9_/-]+)\))\|(?P<RecipientID>-?\d+)\s+(?P<DamageTotal>-?\d+\.\d+)\s+\((h:(?P<DamageHull>-?\d+\.\d+))\s+(s:(?P<DamageShield>-?\d+\.\d+))\)\s+((?P<ActionSource>\(?[a-zA-Z0-9_/-]+\)?))?\s+(?P<DamageModifiers>[A-Z\|_]+)(\s+(?P<FriendlyFire><FriendlyFire>))?\s*$`)

type Damage struct {
	LogTime         time.Time
	Initiator       string
	InitiatorID     int
	Recipient       string
	ObjectName      string
	ObjectOwner     string
	RecipientID     int
	DamageTotal     float32
	DamageHull      float32
	DamageShield    float32
	ActionSource    string
	DamageModifiers DamageModifiersMap
	FriendlyFire    bool
}

func (c *Damage) Unmarhsal(src string, now time.Time) (err error) {
	res := reDamage.FindStringSubmatch(src)
	if len(res) != 19 {
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
	c.DamageTotal, err = parseField(res[9], "DamageTotal", parseFloat)
	if err != nil {
		return err
	}
	c.DamageHull, err = parseField(res[11], "DamageHull", parseFloat)
	if err != nil {
		return err
	}
	c.DamageShield, err = parseField(res[13], "DamageShield", parseFloat)
	if err != nil {
		return err
	}
	c.ActionSource = res[15]
	c.DamageModifiers, err = parseField(res[16], "DamageModifiers", parseDamageModifiers)
	if err != nil {
		return err
	}
	c.FriendlyFire, err = parseField(res[18], "FriendlyFire", parseBool)
	if err != nil {
		return err
	}

	return nil
}

func (c *Damage) Type() CombatLineType {
	return DamageLineType
}

func (c *Damage) Time() time.Time {
	return c.LogTime
}
