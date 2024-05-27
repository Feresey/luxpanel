// DO NOT EDIT. This file was auto-generated

package combat

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	reKill      = regexp.MustCompile(`(?s)^(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+CMBT\s+\|\s+Killed\s+(((?P<RecipientName>[a-zA-Z][a-zA-Z0-9_/-]*)\s+(?P<RecipientShip>[a-zA-Z][a-zA-Z0-9_/-]*))|((?P<RecipientObject>[a-zA-Z][a-zA-Z0-9_/-]*)|(?P<RecipientObjectName>[a-zA-Z][a-zA-Z0-9_/-]*)\((?P<RecipientObjectOwner>[a-zA-Z][a-zA-Z0-9_/-]*)\)))\|(?P<RecipientID>-?\d+);\s+killer\s+(?P<Initiator>[a-zA-Z][a-zA-Z0-9_/-]*)\|(?P<InitiatorID>-?\d+)(\s+(?P<ActionSource>\(?[a-zA-Z0-9_/-]+\)?))?(\s+(?P<FriendlyFire><FriendlyFire>))?\s*$`)
	shortReKill = regexp.MustCompile(`^(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})\s+CMBT\s+\|\s+Killed\s+`)
)

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

var emptyKillLogTime time.Time

func (c *Kill) GetLogTime() time.Time {
	if c == nil || c.LogTime == emptyKillLogTime {
		return emptyKillLogTime
	}
	return c.LogTime

}

var emptyKillRecipientName string

func (c *Kill) GetRecipientName() string {
	if c == nil || c.RecipientName == emptyKillRecipientName {
		return emptyKillRecipientName
	}
	return c.RecipientName

}

var emptyKillRecipientShip string

func (c *Kill) GetRecipientShip() string {
	if c == nil || c.RecipientShip == emptyKillRecipientShip {
		return emptyKillRecipientShip
	}
	return c.RecipientShip

}

var emptyKillRecipientObject string

func (c *Kill) GetRecipientObject() string {
	if c == nil || c.RecipientObject == emptyKillRecipientObject {
		return emptyKillRecipientObject
	}
	return c.RecipientObject

}

var emptyKillRecipientObjectName string

func (c *Kill) GetRecipientObjectName() string {
	if c == nil || c.RecipientObjectName == emptyKillRecipientObjectName {
		return emptyKillRecipientObjectName
	}
	return c.RecipientObjectName

}

var emptyKillRecipientObjectOwner string

func (c *Kill) GetRecipientObjectOwner() string {
	if c == nil || c.RecipientObjectOwner == emptyKillRecipientObjectOwner {
		return emptyKillRecipientObjectOwner
	}
	return c.RecipientObjectOwner

}

var emptyKillRecipientID int

func (c *Kill) GetRecipientID() int {
	if c == nil || c.RecipientID == emptyKillRecipientID {
		return emptyKillRecipientID
	}
	return c.RecipientID

}

var emptyKillInitiator string

func (c *Kill) GetInitiator() string {
	if c == nil || c.Initiator == emptyKillInitiator {
		return emptyKillInitiator
	}
	return c.Initiator

}

var emptyKillInitiatorID int

func (c *Kill) GetInitiatorID() int {
	if c == nil || c.InitiatorID == emptyKillInitiatorID {
		return emptyKillInitiatorID
	}
	return c.InitiatorID

}

var emptyKillActionSource string

func (c *Kill) GetActionSource() string {
	if c == nil || c.ActionSource == emptyKillActionSource {
		return emptyKillActionSource
	}
	return c.ActionSource

}

var emptyKillFriendlyFire bool

func (c *Kill) GetFriendlyFire() bool {
	if c == nil || c.FriendlyFire == emptyKillFriendlyFire {
		return emptyKillFriendlyFire
	}
	return c.FriendlyFire

}
