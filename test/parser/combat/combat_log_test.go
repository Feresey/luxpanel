package parser_test

import (
	_ "embed"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed testdata/connect.txt
	connectRaw string
	//go:embed testdata/start.txt
	startRaw string
	//go:embed testdata/damage.txt
	damageRaw string
	//go:embed testdata/heal.txt
	healRaw string
	//go:embed testdata/kill.txt
	killRaw string
	//go:embed testdata/finished.txt
	finishedRaw string
	//go:embed testdata/reward.txt
	rewardRaw string
	//go:embed testdata/apply_aura.txt
	applyAuraRaw string
)

type testData[Out any] struct {
	name      string
	input     string
	want      Out
	wantError bool
}

func runTests[Out any](t *testing.T, tests []testData[Out], rawTests string, runFunc func(*testing.T, string) (Out, error)) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			r := require.New(t)

			got, err := runFunc(t, tt.input)
			if tt.wantError {
				raw, _ := json.Marshal(got)
				r.Error(err, "res: %s, err: %v", raw, err)
				return
			}
			r.NoError(err)
			r.Equal(tt.want, got)
		})
	}

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(rawTests, "\n")

		for _, line := range lines {
			if line == "" {
				return
			}
			res, err := runFunc(t, line)
			r.NoError(err)
			r.NotEmpty(res)
		}
	})
}

func TestCombatConnect(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.UTC)
	tests := []testData[combat.LogLine]{
		{
			name:  "ok",
			input: "19:32:58.666  CMBT   | ======= Connect to game session 50419619 =======",
			want: &combat.ConnectToGameSession{
				Time:      combat.Time(time.Date(2023, 1, 0, 19, 32, 58, 666000000, time.UTC)),
				SessionID: 50419619,
			},
		},
		{
			name:      "cutted",
			input:     "19:32:58.666  CMBT   | ======= Connect to game session 50419619",
			wantError: true,
		},
		{
			name:      "empty",
			input:     "",
			wantError: false,
		},
		{
			name:      "wrong time",
			input:     "19:32:58.66  CMBT   | ======= Connect to game session 50419619 =======",
			wantError: true,
		},
	}

	p := combat.NewParser()
	runTests(t, tests, connectRaw, func(t *testing.T, raw string) (combat.LogLine, error) {
		return p.Parse(now, raw)
	})
}

func New[T any](val T) *T {
	return &val
}

func TestCombatStartGameplay(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.UTC)
	tests := []testData[combat.LogLine]{
		{
			name:  "pve",
			input: `19:42:14.670  CMBT   | ======= Start PVE mission 'pve_raid_waterharvest_t5' map 'pve_raid_waterharvest' =======`,
			want: &combat.Start{
				Time:              combat.Time(time.Date(2023, 1, 0, 19, 42, 14, 670000000, time.UTC)),
				What:              "PVE mission",
				GameMode:          "pve_raid_waterharvest_t5",
				MapName:           "pve_raid_waterharvest",
				LocalClientTeamID: 0,
			},
		},
		{
			name:  "pvp",
			input: `20:21:02.744  CMBT   | ======= Start gameplay 'CaptureTheBase' map 's1420_ceres3_asteroidcity', local client team 1 =======`,
			want: &combat.Start{
				Time:              combat.Time(time.Date(2023, 1, 0, 20, 21, 0o2, 744000000, time.UTC)),
				What:              "gameplay",
				GameMode:          "CaptureTheBase",
				MapName:           "s1420_ceres3_asteroidcity",
				LocalClientTeamID: 1,
			},
		},
		{
			name:      "cutted",
			input:     "19:42:14.670  CMBT   | ======= Start PVE mission 'pve_raid_waterharvest_t5' map 'pve_raid",
			wantError: true,
		},
		{
			name:      "empty",
			input:     "",
			wantError: false,
		},
		{
			name:      "wrong time",
			input:     `19:42:14.60  CMBT   | ======= Start PVE mission 'pve_raid_waterharvest_t5' map 'pve_raid_waterharvest' =======`,
			wantError: true,
		},
	}
	p := combat.NewParser()
	runTests(t, tests, startRaw, func(t *testing.T, raw string) (combat.LogLine, error) {
		return p.Parse(now, raw)
	})

}

func TestCombatDamage(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.UTC)
	tests := []testData[combat.LogLine]{
		{
			name:  "ok",
			input: `21:17:13.938  CMBT   | Damage        Gladiator|0000003117 ->           YanFei|0000167786  73.78 (h:0.00 s:73.78) Weapon_PlasmaBursts_T5_Rel EMP`,
			want: &combat.Damage{
				Time: combat.Time(time.Date(2023, 1, 0, 21, 17, 13, 938000000, time.UTC)),
				Initiator: combat.Object{
					Name:     "Gladiator",
					ObjectID: 3117,
				},
				Recipient: combat.Object{
					Name:     "YanFei",
					ObjectID: 167786,
				},
				DamageFull:   73.78,
				DamageHull:   0,
				DamageShield: 73.78,
				Source:       "Weapon_PlasmaBursts_T5_Rel",
				DamageModifiers: combat.DamageModifiers{
					"EMP",
				},
			},
		},
		{
			name:  "no action",
			input: `19:44:04.074  CMBT   | Damage Megabomb_RW_BlackHole|0000000155 ->            tuman|0000000824   0.00 (h:0.00 s:0.00)  KINETIC`,
			want: &combat.Damage{
				Time: combat.Time(time.Date(2023, 1, 0, 19, 44, 4, 74000000, time.UTC)),
				Initiator: combat.Object{
					Name:     "Megabomb_RW_BlackHole",
					ObjectID: 155,
				},
				Recipient: combat.Object{
					Name:     "tuman",
					ObjectID: 824,
				},
				DamageFull:   0,
				DamageHull:   0,
				DamageShield: 0,
				Source:       "",
				DamageModifiers: combat.DamageModifiers{
					"KINETIC",
				},
			},
		},
		{
			name:  "crash",
			input: `19:42:53.450  CMBT   | Damage            Py6Jl|0000000395 ->            Py6Jl|0000000395   0.00 (h:0.00 s:0.00) Weapon_OrbGun_T5_Epic EMP|PRIMARY_WEAPON|EXPLOSION <FriendlyFire>`,
			want: &combat.Damage{
				Time: combat.Time(time.Date(2023, 1, 0, 19, 42, 53, 450000000, time.UTC)),
				Initiator: combat.Object{
					Name:     "Py6Jl",
					ObjectID: 395,
				},
				Recipient: combat.Object{
					Name:     "Py6Jl",
					ObjectID: 395,
				},
				DamageFull:   0,
				DamageHull:   0,
				DamageShield: 0,
				Source:       "Weapon_OrbGun_T5_Epic",
				FriendlyFire: true,
				DamageModifiers: combat.DamageModifiers{
					"EMP", "PRIMARY_WEAPON", "EXPLOSION",
				},
			},
		},

		{
			name:      "cutted",
			input:     `21:17:13.938  CMBT   | Damage        Gladiator|0000003117 ->           YanFei|0000167786  73.78 (h:0.00 s:73.78) Weapon_PlasmaBursts_T5_Rel`,
			wantError: true,
		},
		{
			name:      "empty",
			input:     "",
			wantError: false,
		},
	}

	p := combat.NewParser()
	runTests(t, tests, damageRaw, func(t *testing.T, raw string) (combat.LogLine, error) {
		return p.Parse(now, raw)
	})
}

func TestCombatHeal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.UTC)
	tests := []testData[combat.LogLine]{
		{
			name:  "ok",
			input: `19:33:24.732  CMBT   | Heal            Feresey|0000000204 ->          Feresey|0000000204 244.00 Module_Lynx2Shield_T4_Epic`,
			want: &combat.Heal{
				Time: combat.Time(time.Date(2023, 1, 0, 19, 33, 24, 732000000, time.UTC)),
				Initiator: combat.Object{
					Name:     "Feresey",
					ObjectID: 204,
				},
				Recipient: combat.Object{
					Name:     "Feresey",
					ObjectID: 204,
				},
				Heal:   244,
				Source: "Module_Lynx2Shield_T4_Epic",
			},
		},

		{
			name:      "cutted",
			input:     `19:33:24.732  CMBT   | Heal            Feresey|0000000204 ->          Feresey|0000000204 244.00`,
			wantError: true,
		},
		{
			name:      "empty",
			input:     "",
			wantError: false,
		},
	}

	p := combat.NewParser()
	runTests(t, tests, healRaw, func(t *testing.T, raw string) (combat.LogLine, error) {
		return p.Parse(now, raw)
	})
}

func TestCombatKill(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.UTC)
	tests := []testData[combat.LogLine]{
		{
			name:  "player",
			input: `19:33:59.527  CMBT   | Killed Py6Jl	Ship_Race3_M_T2_Pirate|0000000248;	killer Feresey|0000000204 Weapon_Plasmagun_Heavy_T5_Pirate`,
			want: &combat.Kill{
				Time: combat.Time(time.Date(2023, 1, 0, 19, 33, 59, 527000000, time.UTC)),
				Killer: combat.Object{
					Name:     "Feresey",
					ObjectID: 204,
				},
				Killed: combat.Object{
					Name: "Py6Jl",
					PlayerObject: combat.PlayerObject{
						ObjectName: "Ship_Race3_M_T2_Pirate",
					},
					ObjectID: 248,
				},
				Source: "Weapon_Plasmagun_Heavy_T5_Pirate",
			},
		},
		{
			name:  "not player",
			input: `19:43:01.146  CMBT   | Killed Alien_Destroyer_Life_02_T5|0000001154;	killer Feresey|0000000766 Weapon_PlasmaWebLaser_T5_Epic`,
			want: &combat.Kill{
				Time: combat.Time(time.Date(2023, 1, 0, 19, 43, 1, 146000000, time.UTC)),
				Killer: combat.Object{
					Name:     "Feresey",
					ObjectID: 766,
				},
				Killed: combat.Object{
					ObjectID: 1154,
					Name:     "Alien_Destroyer_Life_02_T5",
				},
				Source: "Weapon_PlasmaWebLaser_T5_Epic",
			},
		},
		{
			name:  "friendly fire",
			input: `19:46:16.971  CMBT   | Killed HealBot_Armor(Therm0Nuclear)|0000039068;	 killer Therm0Nuclear|0000039068 (suicide) <FriendlyFire>`,
			want: &combat.Kill{
				Time: combat.Time(time.Date(2023, 1, 0, 19, 46, 16, 971000000, time.UTC)),
				Killer: combat.Object{
					Name:     "Therm0Nuclear",
					ObjectID: 39068,
				},
				Killed: combat.Object{
					PlayerObject: combat.PlayerObject{
						ObjectOwner: "Therm0Nuclear",
						ObjectName:  "HealBot_Armor",
					},
					ObjectID: 39068,
				},
				Source:       "(suicide)",
				FriendlyFire: true,
			},
		},
		{
			name:  "killed object",
			input: `15:55:08.879  CMBT   | Killed SwarmPack3(MADEinHEAVEN)|0000056377;	 killer MADEinHEAVEN|0000056377 (suicide) <FriendlyFire>`,
			want: &combat.Kill{
				Time: combat.Time(time.Date(2023, 1, 0, 15, 55, 8, 879000000, time.UTC)),
				Killer: combat.Object{
					Name:     "MADEinHEAVEN",
					ObjectID: 56377,
				},
				Killed: combat.Object{
					ObjectID: 56377,
					PlayerObject: combat.PlayerObject{
						ObjectName:  "SwarmPack3",
						ObjectOwner: "MADEinHEAVEN",
					},
				},
				Source:       "(suicide)",
				FriendlyFire: true,
			},
		},
		{
			name:      "cutted",
			input:     `19:33:59.527  CMBT   | Killed Py6Jl	Ship_Race3_M_T2_Pirate|0000000248;	killer Feresey|`,
			wantError: true,
		},
		{
			name:      "empty",
			input:     "",
			wantError: false,
		},
	}

	p := combat.NewParser()
	runTests(t, tests, killRaw, func(t *testing.T, raw string) (combat.LogLine, error) {
		return p.Parse(now, raw)
	})
}

func TestCombatGameFinished(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.UTC)
	tests := []testData[combat.LogLine]{
		{
			name:  "ok",
			input: `19:47:09.448  CMBT   | Gameplay finished. Winner team: 1(PVE_MISSION_COMPLETE_ALT_2). Finish reason: 'Mission complete'. Actual game time 275.9 sec`,
			want: &combat.Finished{
				Time:         combat.Time(time.Date(2023, 1, 0, 19, 47, 9, 448000000, time.UTC)),
				WinnerTeamID: 1,
				WinReason:    "PVE_MISSION_COMPLETE_ALT_2",
				FinishReason: "Mission complete",
				GameTime:     275.9,
			},
		},
		{
			name:      "cutted",
			input:     `19:47:09.448  CMBT   | Gameplay finished. Winner team: 1(PVE_MISSION_COMPLETE_ALT_2). Finish reason: 'Mission complete'. Actual game time`,
			wantError: true,
		},
		{
			name:      "empty",
			input:     "",
			wantError: false,
		},
	}

	p := combat.NewParser()
	runTests(t, tests, finishedRaw, func(t *testing.T, raw string) (combat.LogLine, error) {
		return p.Parse(now, raw)
	})
}

func TestCombatReward(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.UTC)
	tests := []testData[combat.LogLine]{
		{
			name:  "heal",
			input: `19:42:59.796  CMBT   | Reward idanceandkillyou Ship_Race3_L_T5_Faction1        	   1 experience                for heal Feresey`,
			want: &combat.Reward{
				Time:       combat.Time(time.Date(2023, 1, 0, 19, 42, 59, 796000000, time.UTC)),
				Recipient:  "idanceandkillyou",
				Ship:       "Ship_Race3_L_T5_Faction1",
				Reward:     1,
				RewardType: "experience",
				Reason:     "heal Feresey",
			},
		},
		{
			name:  "kill",
			input: `21:51:39.757  CMBT   | Reward        Gladiator Ship_Race5_H_OVERSEER_Rank15_13 	  46 effective points          for kill Khushal64n6`,
			want: &combat.Reward{
				Time:       combat.Time(time.Date(2023, 1, 0, 21, 51, 39, 757000000, time.UTC)),
				Recipient:  "Gladiator",
				Ship:       "Ship_Race5_H_OVERSEER_Rank15_13",
				Reward:     46,
				RewardType: "effective points",
				Reason:     "kill Khushal64n6",
			},
		},
		{
			name:  "victory",
			input: `21:51:41.115  CMBT   | Reward          Jigolee	722906 credits                   for victory`,
			want: &combat.Reward{
				Time:       combat.Time(time.Date(2023, 1, 0, 21, 51, 41, 115000000, time.UTC)),
				Recipient:  "Jigolee",
				Ship:       "",
				Reward:     722906,
				RewardType: "credits",
				Reason:     "victory",
			},
		},

		{
			name:  "strange kill",
			input: `21:51:31.180  CMBT   | Reward    PomogitePsixy Ship_Race2_M_T5_CraftUniq       	   9 effective points          for damage assist to kill `,
			want: &combat.Reward{
				Time:       combat.Time(time.Date(2023, 1, 0, 21, 51, 31, 180000000, time.UTC)),
				Recipient:  "PomogitePsixy",
				Ship:       "Ship_Race2_M_T5_CraftUniq",
				Reward:     9,
				RewardType: "effective points",
				Reason:     "damage assist to kill ",
			},
		},
		{
			name:  "cutted kill",
			input: `21:51:36.614  CMBT   | Reward      Khushal64n6 Ship_Race3_H_T5_Uniq            	 123 experience                for kill `,
			want: &combat.Reward{
				Time:       combat.Time(time.Date(2023, 1, 0, 21, 51, 36, 614000000, time.UTC)),
				Recipient:  "Khushal64n6",
				Ship:       "Ship_Race3_H_T5_Uniq",
				Reward:     123,
				RewardType: "experience",
				Reason:     "kill ",
			},
		},
	}

	p := combat.NewParser()
	runTests(t, tests, rewardRaw, func(t *testing.T, raw string) (combat.LogLine, error) {
		return p.Parse(now, raw)
	})
}

// func TestCombatApplyAura(t *testing.T) {
// 	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.UTC)
// 	tests := []testData[combat.LogLine]{
// 		{
// 			name: "apply aura",
// 			input:  `01:05:47.750  CMBT   | Apply aura 'TestKPMNormalizer_1' id 97 type AURA_RESIST_ALL to 'Taurefinne'`,
// 			want:combat.LogLine{
// 				Time: combat.Time( time.Date(2023, 1, 0, 1, 5, 47, 750000000, time.UTC)),
// Line: & combat.ApplyAura{
// 				AuraName:  "TestKPMNormalizer_1",
// 				AuraID:    97,
// 				AuraType:  "AURA_RESIST_ALL",
// 				Recipient: combat.Object{
// Name:"Taurefinne",
// 			ObjectID
// },
// 		},
// 		{
// 			name: "apply aura to object",
// 			input:  `23:04:41.679  CMBT   | Apply aura 'Spell_AdvancedHeal_T5_Epic' id 433 type AURA_HEALING_MOD to 'DestrMunition_SectorShield_T5_Mk3(Khushal64n6)'`,
// 			want:combat.LogLine{
// 				Time: combat.Time(   time.Date(2023, 1, 0, 23, 04, 41, 679000000, time.UTC)),
// Line: & combat.ApplyAura{
// 				AuraName:    "Spell_AdvancedHeal_T5_Epic",
// 				AuraID:      433,
// 				AuraType:    "AURA_HEALING_MOD",
// 				ObjectName:  "DestrMunition_SectorShield_T5_Mk3",
// 				ObjectOwner: "Khushal64n6",
// 			},
// 		},
// 		{
// 			name: "3aсtpeJIucb",
// 			input:  `00:39:57.865  CMBT   | Apply aura 'SpawnWarpInvulnerability' id 62 type AURA_INVULNERABILITY to '3acTPeJIuCb'`,
// 			want:combat.LogLine{
// 				Time: combat.Time( time.Date(2023, 1, 0, 0, 39, 57, 865000000, time.UTC)),
// Line: & combat.ApplyAura{
// 				AuraName:  "SpawnWarpInvulnerability",
// 				AuraID:    62,
// 				AuraType:  "AURA_INVULNERABILITY",
// 				Recipient: combat.Object{
// Name:"3acTPeJIuCb",
// 			ObjectID
// },
// 		},
// 	}

// }
