// Code generated by "enumer -type DamageType -trimprefix DamageType -text -transform snake"; DO NOT EDIT.

package service

import (
	"fmt"
)

const _DamageTypeName = "totalhullshield"

var _DamageTypeIndex = [...]uint8{0, 5, 9, 15}

func (i DamageType) String() string {
	if i < 0 || i >= DamageType(len(_DamageTypeIndex)-1) {
		return fmt.Sprintf("DamageType(%d)", i)
	}
	return _DamageTypeName[_DamageTypeIndex[i]:_DamageTypeIndex[i+1]]
}

var _DamageTypeValues = []DamageType{0, 1, 2}

var _DamageTypeNameToValueMap = map[string]DamageType{
	_DamageTypeName[0:5]:  0,
	_DamageTypeName[5:9]:  1,
	_DamageTypeName[9:15]: 2,
}

// DamageTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DamageTypeString(s string) (DamageType, error) {
	if val, ok := _DamageTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to DamageType values", s)
}

// DamageTypeValues returns all values of the enum
func DamageTypeValues() []DamageType {
	return _DamageTypeValues
}

// IsADamageType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i DamageType) IsADamageType() bool {
	for _, v := range _DamageTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for DamageType
func (i DamageType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for DamageType
func (i *DamageType) UnmarshalText(text []byte) error {
	var err error
	*i, err = DamageTypeString(string(text))
	return err
}
