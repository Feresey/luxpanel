package main

import (
	"fmt"
	"strings"
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
