package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const data = `19:33:59.527  CMBT   | Killed Py6Jl	 Ship_Race3_M_T2_Pirate|0000000248;	 killer Feresey|0000000204 Weapon_Plasmagun_Heavy_T5_Pirate 
19:43:01.146  CMBT   | Killed Alien_Destroyer_Life_02_T5|0000001154;	 killer Feresey|0000000766 Weapon_PlasmaWebLaser_T5_Epic 
19:43:19.332  CMBT   | Killed Module_CookCannon_T5_Epic(Jigolee)|0000000494;	 killer Turret_CruiserAlien_T5R|0000001179 Turret_CruiserAlien_T5R 
19:43:19.332  CMBT   | Killed Module_CookCannon_T5_Epic(Jigolee)|0000000494;	 killer Turret_CruiserAlien_T5R|0000001179 Turret_CruiserAlien_T5R <FriendlyFire>
19:44:55.747  CMBT   | Killed SwarmPack2(georgeatg)|0000001043;	 killer georgeatg|0000001043 (suicide) <FriendlyFire>
`

func main() {
	// yyDebug = 3
	p := yyNewParser()
	raw, err := os.ReadFile("/Users/pavelmilko/Desktop/my/sclogparser/test/parser/combat/testdata/kill.txt")
	_ = raw
	_ = err
	// raw = []byte(data)
	l := newLexer(string(raw), time.Now())

	_ = p.Parse(l)

	fmt.Println(len(l.res))

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	// enc.Encode(l.res)
}
