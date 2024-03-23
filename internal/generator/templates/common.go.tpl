package {{ (index . 0).PackageName}}

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	ErrUndefinedLineType = errors.New("undefined line type")
	ErrWrongLineFormat   = errors.New("wrong format")
)

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

func parseBool(s string) (bool, error) {
	return s != "", nil
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

type {{ (index . 0).PackageName | camelcase}}LineType string

const (
	{{- range $i, $type := .}}
	{{$type.TypeName}}LineType = "{{$type.TypeName}}"
	{{- end}}
)