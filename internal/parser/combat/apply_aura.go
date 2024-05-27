// DO NOT EDIT. This file was auto-generated

package combat

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	reApplyAura      = regexp.MustCompile(`(?s)^(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+CMBT\s+\|\s+Apply aura\s+'(?P<AuraName>[a-zA-Z][a-zA-Z0-9_/-]*)'\s+id\s+(?P<AuraID>\d+)\s+type\s+(?P<AuraType>[A-Z0-9_]+)\s+to\s+'((?P<Recipient>[a-zA-Z][a-zA-Z0-9_/-]*)|(?P<ObjectName>[a-zA-Z][a-zA-Z0-9_/-]*)\((?P<ObjectOwner>[a-zA-Z][a-zA-Z0-9_/-]*)\))'\s*$`)
	shortReApplyAura = regexp.MustCompile(`^(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+CMBT\s+\|\s+Apply aura\s+`)
)

type ApplyAura struct {
	LogTime     time.Time
	AuraName    string
	AuraID      int
	AuraType    string
	Recipient   string
	ObjectName  string
	ObjectOwner string
}

func (c *ApplyAura) Unmarshal(src string, now time.Time) (err error) {
	res := reApplyAura.FindStringSubmatch(src)
	if len(res) != 9 {
		return fmt.Errorf("%w: %d", errWrongLineFormat, len(res))
	}

	c.LogTime, err = parseField(res[1], "LogTime", parseTime(now))
	if err != nil {
		return err
	}
	c.AuraName = res[2]
	c.AuraID, err = parseField(res[3], "AuraID", strconv.Atoi)
	if err != nil {
		return err
	}
	c.AuraType = res[4]
	c.Recipient = res[6]
	c.ObjectName = res[7]
	c.ObjectOwner = res[8]

	return nil
}

func (c *ApplyAura) Type() CombatLineType {
	return ApplyAuraLineType
}

var emptyApplyAuraLogTime time.Time

func (c *ApplyAura) GetLogTime() time.Time {
	if c == nil || c.LogTime == emptyApplyAuraLogTime {
		return emptyApplyAuraLogTime
	}
	return c.LogTime

}

var emptyApplyAuraAuraName string

func (c *ApplyAura) GetAuraName() string {
	if c == nil || c.AuraName == emptyApplyAuraAuraName {
		return emptyApplyAuraAuraName
	}
	return c.AuraName

}

var emptyApplyAuraAuraID int

func (c *ApplyAura) GetAuraID() int {
	if c == nil || c.AuraID == emptyApplyAuraAuraID {
		return emptyApplyAuraAuraID
	}
	return c.AuraID

}

var emptyApplyAuraAuraType string

func (c *ApplyAura) GetAuraType() string {
	if c == nil || c.AuraType == emptyApplyAuraAuraType {
		return emptyApplyAuraAuraType
	}
	return c.AuraType

}

var emptyApplyAuraRecipient string

func (c *ApplyAura) GetRecipient() string {
	if c == nil || c.Recipient == emptyApplyAuraRecipient {
		return emptyApplyAuraRecipient
	}
	return c.Recipient

}

var emptyApplyAuraObjectName string

func (c *ApplyAura) GetObjectName() string {
	if c == nil || c.ObjectName == emptyApplyAuraObjectName {
		return emptyApplyAuraObjectName
	}
	return c.ObjectName

}

var emptyApplyAuraObjectOwner string

func (c *ApplyAura) GetObjectOwner() string {
	if c == nil || c.ObjectOwner == emptyApplyAuraObjectOwner {
		return emptyApplyAuraObjectOwner
	}
	return c.ObjectOwner

}
