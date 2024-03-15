package parser_test

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/Feresey/sclogparser/internal/parser"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed testdata/combat/connect.txt
	connectRaw string
	//go:embed testdata/combat/start.txt
	startRaw string
	//go:embed testdata/combat/damage.txt
	damageRaw string
	//go:embed testdata/combat/heal.txt
	healRaw string
	//go:embed testdata/combat/kill.txt
	killRaw string
	//go:embed testdata/combat/finished.txt
	finishedRaw string
)

func TestCombatConnectUnmarshal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name      string
		raw       string
		want      *parser.CombatLogLineConnectToGameSession
		wantError bool
	}{
		{
			name: "ok",
			raw:  "19:32:58.666  CMBT   | ======= Connect to game session 50419619 =======",
			want: &parser.CombatLogLineConnectToGameSession{
				LogTime:   time.Date(2023, 1, 0, 19, 32, 58, 666000000, time.Local),
				SessionID: 50419619,
			},
		},
		{
			name:      "cutted",
			raw:       "19:32:58.666  CMBT   | ======= Connect to game session 50419619",
			wantError: true,
		},
		{
			name:      "empty",
			raw:       "",
			wantError: true,
		},
		{
			name: "double",
			raw: "19:32:58.666  CMBT   | ======= Connect to game session 50419619 =======\n" +
				"19:32:58.666  CMBT   | ======= Connect to game session 50419619 =======",
			wantError: true,
		},
		{
			name:      "wrong time",
			raw:       "19:32:58.66  CMBT   | ======= Connect to game session 50419619 =======",
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			val, err := parser.ParseCombatLogLine([]byte(tt.raw), now)
			if tt.wantError {
				r.Error(err)
				return
			} else {
				r.NoError(err)
			}

			r.Equal(tt.want, val)
		})
	}

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(connectRaw, "\n")
		for _, line := range lines {
			if line == "" {
				return
			}
			_, err := parser.ParseCombatLogLine([]byte(line), now)
			r.NoError(err)
		}
	})
}

func TestCombatStartGameplayUnmarshal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name      string
		raw       string
		want      *parser.CombatLogLineStartGameplay
		wantError bool
	}{
		{
			name: "pve",
			raw:  `19:42:14.670  CMBT   | ======= Start PVE mission 'pve_raid_waterharvest_t5' map 'pve_raid_waterharvest' =======`,
			want: &parser.CombatLogLineStartGameplay{
				LogTime:  time.Date(2023, 1, 0, 19, 42, 14, 670000000, time.Local),
				GameMode: "pve_raid_waterharvest_t5",
				MapName:  "pve_raid_waterharvest",
			},
		},
		{
			name: "pvp",
			raw:  `20:21:02.744  CMBT   | ======= Start gameplay 'CaptureTheBase' map 's1420_ceres3_asteroidcity', local client team 1 =======`,
			want: &parser.CombatLogLineStartGameplay{
				LogTime:  time.Date(2023, 1, 0, 20, 21, 0o2, 744000000, time.Local),
				GameMode: "CaptureTheBase",
				MapName:  "s1420_ceres3_asteroidcity",
			},
		},
		{
			name:      "cutted",
			raw:       "19:42:14.670  CMBT   | ======= Start PVE mission 'pve_raid_waterharvest_t5' map 'pve_raid_waterharvest'",
			wantError: true,
		},
		{
			name:      "empty",
			raw:       "",
			wantError: true,
		},
		{
			name:      "wrong time",
			raw:       `19:42:14.60  CMBT   | ======= Start PVE mission 'pve_raid_waterharvest_t5' map 'pve_raid_waterharvest' =======`,
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			val, err := parser.ParseCombatLogLine([]byte(tt.raw), now)
			if tt.wantError {
				r.Error(err)
				return
			} else {
				r.NoError(err)
			}

			r.Equal(tt.want, val)
		})
	}

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(startRaw, "\n")
		for _, line := range lines {
			if line == "" {
				return
			}
			_, err := parser.ParseCombatLogLine([]byte(line), now)
			r.NoError(err)
		}
	})
}

func TestCombatDamageUnmarshal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name      string
		raw       string
		want      *parser.CombatLogLineDamage
		wantError bool
	}{
		{
			name: "ok",
			raw:  `21:17:13.938  CMBT   | Damage        Gladiator|0000003117 ->           YanFei|0000167786  73.78 (h:0.00 s:73.78) Weapon_PlasmaBursts_T5_Rel EMP`,
			want: &parser.CombatLogLineDamage{
				LogTime: time.Date(2023, 1, 0, 21, 17, 13, 938000000, time.Local),
				Players: parser.CombatPlayers{
					Initiator:   "Gladiator",
					InitiatorID: 3117,
					Recipient:   "YanFei",
					RecipientID: 167786,
				},
				DamageTotal:  73.78,
				DamageHull:   0,
				DamageShield: 73.78,
				Weapon:       "Weapon_PlasmaBursts_T5_Rel",
				DamageModifiers: []parser.DamageModifier{
					"EMP",
				},
			},
		},
		{
			name: "no action",
			raw:  `19:44:04.074  CMBT   | Damage Megabomb_RW_BlackHole|0000000155 ->            tuman|0000000824   0.00 (h:0.00 s:0.00)  KINETIC`,
			want: &parser.CombatLogLineDamage{
				LogTime: time.Date(2023, 1, 0, 19, 44, 4, 74000000, time.Local),
				Players: parser.CombatPlayers{
					Initiator:   "Megabomb_RW_BlackHole",
					InitiatorID: 155,
					Recipient:   "tuman",
					RecipientID: 824,
				},
				DamageTotal:  0,
				DamageHull:   0,
				DamageShield: 0,
				Weapon:       "",
				DamageModifiers: []parser.DamageModifier{
					"KINETIC",
				},
			},
		},
		{
			name: "crash",
			raw:  `19:42:53.450  CMBT   | Damage            Py6Jl|0000000395 ->            Py6Jl|0000000395   0.00 (h:0.00 s:0.00) Weapon_OrbGun_T5_Epic EMP|PRIMARY_WEAPON|EXPLOSION <FriendlyFire>`,
			want: &parser.CombatLogLineDamage{
				LogTime: time.Date(2023, 1, 0, 19, 42, 53, 450000000, time.Local),
				Players: parser.CombatPlayers{
					Initiator:   "Py6Jl",
					InitiatorID: 395,
					Recipient:   "Py6Jl",
					RecipientID: 395,
				},
				DamageTotal:    0,
				DamageHull:     0,
				DamageShield:   0,
				Weapon:         "Weapon_OrbGun_T5_Epic",
				IsFriendlyFire: true,
				DamageModifiers: []parser.DamageModifier{
					"EMP", "PRIMARY_WEAPON", "EXPLOSION",
				},
			},
		},

		{
			name:      "cutted",
			raw:       `21:17:13.938  CMBT   | Damage        Gladiator|0000003117 ->           YanFei|0000167786  73.78 (h:0.00 s:73.78) Weapon_PlasmaBursts_T5_Rel`,
			wantError: true,
		},
		{
			name:      "empty",
			raw:       "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			val, err := parser.ParseCombatLogLine([]byte(tt.raw), now)
			if tt.wantError {
				r.Error(err)
				return
			} else {
				r.NoError(err)
			}

			r.Equal(tt.want, val)
		})
	}

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(damageRaw, "\n")

		modifiers := make(map[parser.DamageModifier]int)
		for _, line := range lines {
			if line == "" {
				break
			}
			val, err := parser.ParseCombatLogLine([]byte(line), now)
			r.NoError(err)

			for _, m := range val.(*parser.CombatLogLineDamage).DamageModifiers {
				modifiers[m]++
			}
		}

		fmt.Printf("%+#v\n", modifiers)
	})
}

func TestCombatHealUnmarshal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name      string
		raw       string
		want      *parser.CombatLogLineHeal
		wantError bool
	}{
		{
			name: "ok",
			raw:  `19:33:24.732  CMBT   | Heal            Feresey|0000000204 ->          Feresey|0000000204 244.00 Module_Lynx2Shield_T4_Epic`,
			want: &parser.CombatLogLineHeal{
				LogTime: time.Date(2023, 1, 0, 19, 33, 24, 732000000, time.Local),
				Players: parser.CombatPlayers{
					Initiator:   "Feresey",
					InitiatorID: 204,
					Recipient:   "Feresey",
					RecipientID: 204,
				},
				Heal:   244,
				Reason: "Module_Lynx2Shield_T4_Epic",
			},
		},

		{
			name:      "cutted",
			raw:       `19:33:24.732  CMBT   | Heal            Feresey|0000000204 ->          Feresey|0000000204 244.00`,
			wantError: true,
		},
		{
			name:      "empty",
			raw:       "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			val, err := parser.ParseCombatLogLine([]byte(tt.raw), now)
			if tt.wantError {
				r.Error(err)
				return
			} else {
				r.NoError(err)
			}

			r.Equal(tt.want, val)
		})
	}

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(healRaw, "\n")

		for _, line := range lines {
			if line == "" {
				return
			}
			_, err := parser.ParseCombatLogLine([]byte(line), now)
			r.NoError(err)
		}
	})
}

func TestCombatKillUnmarshal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name      string
		raw       string
		want      *parser.CombatLogLineKill
		wantError bool
	}{
		{
			name: "player",
			raw:  `19:33:59.527  CMBT   | Killed Py6Jl      Ship_Race3_M_T2_Pirate|0000000248;      killer Feresey|0000000204 Weapon_Plasmagun_Heavy_T5_Pirate`,
			want: &parser.CombatLogLineKill{
				LogTime: time.Date(2023, 1, 0, 19, 33, 59, 527000000, time.Local),
				Players: parser.CombatPlayers{
					Initiator:   "Feresey",
					InitiatorID: 204,
					Recipient:   "Py6Jl",
					RecipientID: 248,
				},
				KilledShip: "Ship_Race3_M_T2_Pirate",
				Weapon:     "Weapon_Plasmagun_Heavy_T5_Pirate",
			},
		},
		{
			name: "not player",
			raw:  `19:43:01.146  CMBT   | Killed Alien_Destroyer_Life_02_T5|0000001154;     killer Feresey|0000000766 Weapon_PlasmaWebLaser_T5_Epic`,
			want: &parser.CombatLogLineKill{
				LogTime: time.Date(2023, 1, 0, 19, 43, 1, 146000000, time.Local),
				Players: parser.CombatPlayers{
					Initiator:   "Feresey",
					InitiatorID: 766,
					Recipient:   "",
					RecipientID: 1154,
				},
				KilledShip: "Alien_Destroyer_Life_02_T5",
				Weapon:     "Weapon_PlasmaWebLaser_T5_Epic",
			},
		},
		{
			name: "friendly fire",
			raw:  `19:46:16.971  CMBT   | Killed HealBot_Armor(Therm0Nuclear)|0000039068;	 killer Therm0Nuclear|0000039068 (suicide) <FriendlyFire>`,
			want: &parser.CombatLogLineKill{
				LogTime: time.Date(2023, 1, 0, 19, 46, 16, 971000000, time.Local),
				Players: parser.CombatPlayers{
					Initiator:   "Therm0Nuclear",
					InitiatorID: 39068,
					Recipient:   "",
					RecipientID: 39068,
				},
				KilledShip:     "HealBot_Armor(Therm0Nuclear)",
				Weapon:         "(suicide)",
				IsFriendlyFire: true,
			},
		},
		{
			name:      "cutted",
			raw:       `19:33:59.527  CMBT   | Killed Py6Jl      Ship_Race3_M_T2_Pirate|0000000248;      killer Feresey|`,
			wantError: true,
		},
		{
			name:      "empty",
			raw:       "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			val, err := parser.ParseCombatLogLine([]byte(tt.raw), now)
			if tt.wantError {
				r.Error(err)
				return
			} else {
				r.NoError(err)
			}

			r.Equal(tt.want, val)
		})
	}

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(killRaw, "\n")

		for _, line := range lines {
			if line == "" {
				return
			}
			_, err := parser.ParseCombatLogLine([]byte(line), now)
			r.NoError(err)
		}
	})
}

func TestCombatGameFinishedUnmarshal(t *testing.T) {
	now := time.Date(2023, time.January, 0, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name      string
		raw       string
		want      *parser.CombatLogLineGameFinished
		wantError bool
	}{
		{
			name: "ok",
			raw:  `19:47:09.448  CMBT   | Gameplay finished. Winner team: 1(PVE_MISSION_COMPLETE_ALT_2). Finish reason: 'Mission complete'. Actual game time 275.9 sec`,
			want: &parser.CombatLogLineGameFinished{
				LogTime:          time.Date(2023, 1, 0, 19, 47, 9, 448000000, time.Local),
				WinnerTeamID:     1,
				WinnerTeamReason: "PVE_MISSION_COMPLETE_ALT_2",
				FinishReason:     "Mission complete",
				GameDuration:     (275 * time.Second) + 900*time.Millisecond,
			},
		},
		{
			name:      "cutted",
			raw:       `19:47:09.448  CMBT   | Gameplay finished. Winner team: 1(PVE_MISSION_COMPLETE_ALT_2). Finish reason: 'Mission complete'. Actual game time`,
			wantError: true,
		},
		{
			name:      "empty",
			raw:       "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			val, err := parser.ParseCombatLogLine([]byte(tt.raw), now)
			if tt.wantError {
				r.Error(err)
				return
			} else {
				r.NoError(err)
			}

			r.Equal(tt.want, val)
		})
	}

	t.Run("raw", func(t *testing.T) {
		r := require.New(t)
		lines := strings.Split(finishedRaw, "\n")

		for _, line := range lines {
			if line == "" {
				return
			}
			_, err := parser.ParseCombatLogLine([]byte(line), now)
			r.NoError(err)
		}
	})
}
