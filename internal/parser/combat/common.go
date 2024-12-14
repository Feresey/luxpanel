// DO NOT EDIT. This file was auto-generated

package combat

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type CombatLineType string

func (s CombatLineType) String() string { return string(s) }

const (
	ConnectToGameSessionLineType = "ConnectToGameSession"
	DamageLineType               = "Damage"
	FinishedGameplayLineType     = "FinishedGameplay"
	HealLineType                 = "Heal"
	KillLineType                 = "Kill"
	StartGameplayLineType        = "StartGameplay"
)

type LogLine interface {
	Type() CombatLineType
	GetLogTime() time.Time
	Unmarshal(raw string, now time.Time) error
}

var errWrongLineFormat = errors.New("Combat: wrong format")

func ParseLogLine(raw string, now time.Time) (line LogLine, matchedToRegexp bool, err error) {
	switch {
	case shortReConnectToGameSession.MatchString(raw):
		line = &ConnectToGameSession{}
	case shortReDamage.MatchString(raw):
		line = &Damage{}
	case shortReFinishedGameplay.MatchString(raw):
		line = &FinishedGameplay{}
	case shortReHeal.MatchString(raw):
		line = &Heal{}
	case shortReKill.MatchString(raw):
		line = &Kill{}
	case shortReStartGameplay.MatchString(raw):
		line = &StartGameplay{}
	default:
		return nil, false, nil
	}

	if err := line.Unmarshal(raw, now); err != nil {
		return nil, true, fmt.Errorf("line matched to regex(%s), but failed to parse: %s: %s", line.Type().String(), raw, err.Error())
	}

	return line, true, nil
}

const timeFormat = "15:04:05.000"

func parseTime(nowTime time.Time) func(string) (time.Time, error) {
	return func(s string) (time.Time, error) {
		t, err := time.Parse(timeFormat, s)
		if err != nil {
			return t, fmt.Errorf("time.Parse(%s): %w", timeFormat, err)
		}
		res := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
		if res.Before(nowTime) {
			return res.AddDate(0, 0, 1), nil
		}
		return res, nil
	}
}

func parseSeconds(s string) (time.Duration, error) {
	return time.ParseDuration(s + "s")
}

func parseBool(s string) (bool, error) {
	return s != "", nil
}

func parseOptional[T any](parseFunc func(string) (T, error)) func(string) (*T, error) {
	return func(s string) (res *T, err error) {
		if s == "" {
			return nil, nil
		}
		r, err := parseFunc(s)
		if err != nil {
			return nil, err
		}
		return &r, nil
	}
}

func parseFloat(s string) (float32, error) {
	res, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, fmt.Errorf("strconv.ParseFloat: %w", err)
	}
	return float32(res), nil
}

type ParseFieldError struct {
	Raw       string
	FieldName string
	Err       error
}

func (e ParseFieldError) Unwrap() error {
	return e.Err
}

func (p ParseFieldError) Error() string {
	return fmt.Sprintf("parse %s: (%s) %v", p.FieldName, p.Raw, p.Err)
}

func parseField[T any](raw string, fieldName string, convert func(raw string) (res T, err error)) (res T, err error) {
	if res, err = convert(raw); err != nil {
		return res, ParseFieldError{
			Raw:       raw,
			FieldName: fieldName,
			Err:       err,
		}
	}
	return res, nil
}

type DamageModifier string
type DamageModifiersMap map[DamageModifier]struct{}

const (
	DamageTypeEMP         DamageModifier = "EMP"
	DamageTypeKinetic     DamageModifier = "KINETIC"
	DamageTypeThermal     DamageModifier = "THERMAL"
	DamageTypeTrueDamage  DamageModifier = "TRUE_DAMAGE"
	DamageUnintention     DamageModifier = "UNINTENTION"
	DamageCrit            DamageModifier = "CRIT"
	DamageExplosion       DamageModifier = "EXPLOSION"
	DamageCollision       DamageModifier = "COLLISION"
	DamageWeaponPrimary   DamageModifier = "PRIMARY_WEAPON"
	DamageWeaponSecondary DamageModifier = "SECONDARY_WEAPON"
	DamageIgnoreScale     DamageModifier = "IGNORE_DAMAGE_SCALE"
	DamageIgoreShield     DamageModifier = "IGNORE_SHIELD"
	DamageModule          DamageModifier = "MODULE"
)

func parseDamageModifiers(raw string) (res DamageModifiersMap, err error) {
	parts := strings.Split(raw, "|")
	if len(parts) == 0 {
		return res, fmt.Errorf("wrong parts number")
	}
	res = make(DamageModifiersMap, len(parts))
	for _, p := range parts {
		res[DamageModifier(p)] = struct{}{}
	}
	return res, nil
}
