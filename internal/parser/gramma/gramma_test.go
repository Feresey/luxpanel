package gramma

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

// const data = `
// 01:25:10.331  CMBT   |    Participant           SKDnex	 Ship_Race1_S_T5                 	 totalDamage 244.46; mostDamageWith 'Module_ExhaustZone_T5_Epic';
// 01:25:15.658  CMBT   |    Participant        evergreen	 Ship_Race3_S_T4_Premium
// 01:25:16.845  CMBT   |    Participant          Feresey	 SwarmPack3                      	 totalDamage 28600.00; mostDamageWith '(suicide)'; <FriendlyFire>
// 23:07:15.679  CMBT   |    Participant        Gladiator	 Ship_Race1_L_T5_CraftSpec       	 <heal> <buff>
// `

const data = `
19:43:02.566  CMBT   | Damage           KANDIS|0000000609 -> Alien_Destroyer_SmallHearth_T5|0000001159 56083.77 (h:56083.77 s:0.00) Weapon_PlasmaWebLaser_T5_Epic EMP|PRIMARY_WEAPON
19:43:02.566  CMBT   | Damage           KANDIS|0000000609 -> Alien_Destroyer_SmallHearth_T5|0000001159 56083.77 (h:56083.77 s:0.00)  EMP|PRIMARY_WEAPON
19:47:09.448  CMBT   | Gameplay finished. Winner team: 1(PVE_MISSION_COMPLETE_ALT_2). Finish reason: 'Mission complete'. Actual game time 275.9 sec

19:33:59.527  CMBT   | Killed Py6Jl	 Ship_Race3_M_T2_Pirate|0000000248;	 killer Feresey|0000000204 Weapon_Plasmagun_Heavy_T5_Pirate
19:43:01.146  CMBT   | Killed Alien_Destroyer_Life_02_T5|0000001154;	 killer Feresey|0000000766 Weapon_PlasmaWebLaser_T5_Epic
19:43:19.332  CMBT   | Killed Module_CookCannon_T5_Epic(Jigolee)|0000000494;	 killer Turret_CruiserAlien_T5R|0000001179 Turret_CruiserAlien_T5R
19:43:19.332  CMBT   | Killed Module_CookCannon_T5_Epic(Jigolee)|0000000494;	 killer Turret_CruiserAlien_T5R|0000001179 Turret_CruiserAlien_T5R <FriendlyFire>
19:44:55.747  CMBT   | Killed SwarmPack2(georgeatg)|0000001043;	 killer georgeatg|0000001043 (suicide) <FriendlyFire>
`

func Test(t *testing.T) {
	// p := yyNewParser()
	raw, err := os.ReadFile("../test/parser/combat/testdata/finished.txt")
	_ = raw
	_ = err
	raw = []byte(data)
	rd := bufio.NewReader(bytes.NewReader(raw))
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("MAIN READ: %q\n", err)
			break
		}
		fmt.Println("======")
		fmt.Printf("MAIN LINE: %q\n", line)

		var t Tokenizer
		t.Parse(time.Now(), line)
		fmt.Println(t.errors, len(t.tokens))
		for _, tok := range t.tokens {
			fmt.Printf("%s\n", ShowTok(tok))
		}

		// _ = p.Parse(l)
		// if err := l.Err(); err != nil {
		// 	fmt.Printf("MAIN LEX: %q\n", err.Error())
		// 	break
		// }

		// res, err := l.Parse(line)
		// println(err)

		// enc := json.NewEncoder(os.Stdout)
		// enc.SetIndent("", "  ")
		// enc.Encode(toks)
	}
}
