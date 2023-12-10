package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type CombatLogLineType int

const (
	CombatLogLineTypeConnectToGameSession CombatLogLineType = iota
	CombatLogLineTypeStartGameplay
	CombatLogLineTypeDamage
	CombatLogLineTypeHeal
	CombatLogLineTypeKill
	CombatLogLineTypeGameplayFinished
)

type CombatLogLine interface {
	CombatLogLineType() CombatLogLineType
	Unmarshal(raw []byte, now time.Time) error
}

const timeFormat = "15:04:05.000"

func parseLogTime(nowTime time.Time) func(string) (time.Time, error) {
	return func(s string) (time.Time, error) {
		t, err := time.Parse(timeFormat, s)
		if err != nil {
			return t, err
		}
		return time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local), nil
	}
}

func parseFloat(s string) (float32, error) {
	res, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
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

func ParseField[T any](raw string, fieldName string, convert func(raw string) (res T, err error)) (res T, err error) {
	if res, err = convert(raw); err != nil {
		return res, ParseFieldError{
			Raw:       raw,
			FieldName: fieldName,
			Err:       err,
		}
	}
	return res, nil
}

const (
	cmbtLinePrefix = `(?s)^(\d{2}:\d{2}:\d{2}\.\d{3})  CMBT   \| `
	playerIDLine   = `([a-zA-Z0-9()_/-]+)\|(-?\d+)`
	floatLine      = `(-?\d+\.\d+)`
	actionReason   = `([a-zA-Z0-9()_/-]+)?`
	friendlyFire   = `(\s+(<FriendlyFire>))?`
	cmbtLineSuffix = `\s*$`
)

const (
	// 19:32:58.666  CMBT   | ======= Connect to game session 50419619 =======
	connectToGameSessionLine = cmbtLinePrefix + `======= Connect to game session (\d+) =======` + cmbtLineSuffix
)

const (
	connectToGameSessionLineTime = iota + 1
	connectToGameSessionLineSessionID
	connectToGameSessionLineTotal
)

const (
	// 19:33:00.763  CMBT   | ======= Start gameplay 'BombTheBase' map 's1340_thar_aliendebris13', local client team 2 =======
	startGameplayLine = cmbtLinePrefix + `======= Start .* '(.+)' map '(.+)'(, local client team \d+)?\s*=======` + cmbtLineSuffix
)

const (
	startGameplayLineTime = iota + 1
	startGameplayLineGameMode
	startGameplayLineMap
	_
	startGameplayTotal
)

const (
	damageDetailed  = floatLine + `\s+\(h:` + floatLine + `\s+s:` + floatLine + `\)`
	damageModifiers = `([A-Z\|_]+)`
	// 19:33:24.165  CMBT   | Damage              n/a|-000000001 ->          Feresey|0000000204 558.90 (h:0.00 s:558.90) (crash) TRUE_DAMAGE|COLLISION
	// 19:42:53.450  CMBT   | Damage            Py6Jl|0000000395 ->            Py6Jl|0000000395   0.00 (h:0.00 s:0.00) Weapon_OrbGun_T5_Epic EMP|PRIMARY_WEAPON|EXPLOSION <FriendlyFire>
	// 19:44:04.074  CMBT   | Damage Megabomb_RW_BlackHole|0000000155 ->            tuman|0000000824   0.00 (h:0.00 s:0.00)  KINETIC
	damageLine      = cmbtLinePrefix + `Damage\s+` + playerIDLine + `\s+->\s+` + playerIDLine + `\s+` + damageDetailed + `\s` + actionReason + `\s` + damageModifiers + friendlyFire + cmbtLineSuffix
	damageLineShort = cmbtLinePrefix + `Damage`
)

const (
	damageLineTime = iota + 1
	damageLineInitiator
	damageLineInitiatorID
	damageLineRecipient
	damageLineRecipientID
	damageLineDamage
	damageLineHullDamage
	damageLineShieldDamage
	damageLineWeapon
	damageLineWeaponModifiers
	_
	damageLineFriendlyFire
	damageLineTotal
)

const (
	// 19:33:24.732  CMBT   | Heal            Feresey|0000000204 ->          Feresey|0000000204 244.00 Module_Lynx2Shield_T4_Epic
	healLine      = cmbtLinePrefix + `Heal\s+` + playerIDLine + `\s+->\s+` + playerIDLine + `\s+` + floatLine + `\s+` + actionReason + cmbtLineSuffix
	healLineShort = cmbtLinePrefix + `Heal`
)

const (
	healLineTime = iota + 1
	healLineInitiator
	healLineInitiatorID
	healLineRecipient
	healLineRecipientID
	healLineHeal
	healLineReason
	healLineTotal
)

const (
	killedPlayer = `(([a-zA-Z0-9()_/-]+)\s+)?([a-zA-Z0-9()_/-]+)\|(-?\d+)`
	// 19:33:59.527  CMBT   | Killed Py6Jl      Ship_Race3_M_T2_Pirate|0000000248;      killer Feresey|0000000204 Weapon_Plasmagun_Heavy_T5_Pirate
	// 19:43:01.146  CMBT   | Killed Alien_Destroyer_Life_02_T5|0000001154;     killer Feresey|0000000766 Weapon_PlasmaWebLaser_T5_Epic
	killLine      = cmbtLinePrefix + `Killed\s+` + killedPlayer + `;\s+killer\s+` + playerIDLine + `\s+` + actionReason + friendlyFire + cmbtLineSuffix
	killLineShort = cmbtLinePrefix + `Killed`
)

const (
	killLineTime = iota + 1
	_
	killLineRecipient
	killLineKilledShip
	killLineRecipientID
	killLineInitiator
	killLineInitiatorID
	killLineWeapon
	_
	killLineFriendlyFire
	killLineTotal
)

const (
	winnerTeam     = `(\d+)\((.*)\)\.`
	finishReason   = `Finish reason: '(.*)'\.`
	actualGameTime = `Actual game time ` + floatLine + ` sec`
	// 19:47:09.448  CMBT   | Gameplay finished. Winner team: 1(PVE_MISSION_COMPLETE_ALT_2). Finish reason: 'Mission complete'. Actual game time 275.9 sec
	// 20:18:37.406  CMBT   | Gameplay finished. Winner team: 1(ALL_ENEMY_VITALPOINTS_DEAD). Finish reason: 'All beacons captured'. Actual game time 591.1 sec
	// 20:30:08.030  CMBT   | Gameplay finished. Winner team: 2(ALL_ENEMY_SHIPS_KILLED). Finish reason: 'All SpaceShips destroyed'. Actual game time 521.4 sec
	// 20:45:59.862  CMBT   | Gameplay finished. Winner team: 1(MORE_ALIVE_VITALPOINTS_LEFT). Finish reason: 'Timeout'. Actual game time 720.0 sec
	gameFinishedLine      = cmbtLinePrefix + `Gameplay finished\. Winner team: ` + winnerTeam + `\s+` + finishReason + `\s+` + actualGameTime + cmbtLineSuffix
	gameFinishedLineShort = cmbtLinePrefix + `Gameplay finished`
)

const (
	gameFinishedLineTime = iota + 1
	gameFinishedLineWinnerTeamID
	gameFinishedLineWinnerTeamReason
	gameFinishedLineFinishReason
	gameFinishedLineActualGameTime
	gameFinishedLineTotal
)

var combatRe = struct {
	connectToGameSession,
	startGame,
	damage, damageShort,
	heal, healShort,
	kill, killShort,
	gameFinished, gameFinishedShort,
	_ *regexp.Regexp
}{
	connectToGameSession: regexp.MustCompile(connectToGameSessionLine),
	startGame:            regexp.MustCompile(startGameplayLine),
	damage:               regexp.MustCompile(damageLine),
	damageShort:          regexp.MustCompile(damageLineShort),
	heal:                 regexp.MustCompile(healLine),
	healShort:            regexp.MustCompile(healLineShort),
	kill:                 regexp.MustCompile(killLine),
	killShort:            regexp.MustCompile(killLineShort),
	gameFinished:         regexp.MustCompile(gameFinishedLine),
	gameFinishedShort:    regexp.MustCompile(gameFinishedLineShort),
}

type CombatLogLineConnectToGameSession struct {
	Time      time.Time
	SessionID int
}

func (c *CombatLogLineConnectToGameSession) CombatLogLineType() CombatLogLineType {
	return CombatLogLineTypeConnectToGameSession
}

func (c *CombatLogLineConnectToGameSession) Unmarshal(raw []byte, now time.Time) (err error) {
	res := combatRe.connectToGameSession.FindStringSubmatch(string(raw))
	if len(res) != connectToGameSessionLineTotal {
		return fmt.Errorf("wrong format: %s", raw)
	}

	c.Time, err = ParseField(res[connectToGameSessionLineTime], "time", parseLogTime(now))
	if err != nil {
		return err
	}
	c.SessionID, err = ParseField(res[connectToGameSessionLineSessionID], "session_id", strconv.Atoi)
	if err != nil {
		return err
	}
	return nil
}

type CombatLogLineStartGameplay struct {
	Time     time.Time
	GameMode string
	MapName  string
}

func (c *CombatLogLineStartGameplay) CombatLogLineType() CombatLogLineType {
	return CombatLogLineTypeStartGameplay
}

func (c *CombatLogLineStartGameplay) Unmarshal(raw []byte, now time.Time) (err error) {
	res := combatRe.startGame.FindStringSubmatch(string(raw))
	if len(res) != startGameplayTotal {
		return fmt.Errorf("wrong format: %s", raw)
	}
	c.GameMode = res[startGameplayLineGameMode]
	c.MapName = res[startGameplayLineMap]

	c.Time, err = ParseField(res[startGameplayLineTime], "time", parseLogTime(now))
	if err != nil {
		return err
	}
	return nil
}

type CombatPlayers struct {
	Initiator   string
	InitiatorID int
	Recipient   string
	RecipientID int
}

type DamageModifier string

const (
	DamageTypeEMP         = "EMP"
	DamageTypeKinetic     = "KINETIC"
	DamageTypeThermal     = "THERMAL"
	DamageTypeTrueDamage  = "TRUE_DAMAGE"
	DamageUnintention     = "UNINTENTION"
	DamageCrit            = "CRIT"
	DamageExplosion       = "EXPLOSION"
	DamageCollision       = "COLLISION"
	DamageWeaponPrimary   = "PRIMARY_WEAPON"
	DamageWeaponSecondary = "SECONDARY_WEAPON"
	DamageIgnoreScale     = "IGNORE_DAMAGE_SCALE"
	DamageIgoreShield     = "IGNORE_SHIELD"
	DamageModule          = "MODULE"
)

type CombatLogLineDamage struct {
	Time            time.Time
	Players         CombatPlayers
	DamageTotal     float32
	DamageHull      float32
	DamageShield    float32
	Weapon          string
	DamageModifiers []DamageModifier
	IsFriendlyFire  bool
}

func (c CombatLogLineDamage) CombatLogLineType() CombatLogLineType {
	return CombatLogLineTypeDamage
}

func (c *CombatLogLineDamage) Unmarshal(raw []byte, now time.Time) (err error) {
	res := combatRe.damage.FindStringSubmatch(string(raw))
	if len(res) != damageLineTotal {
		return fmt.Errorf("wrong format: %s", raw)
	}
	c.Players.Initiator = res[damageLineInitiator]
	c.Players.Recipient = res[damageLineRecipient]
	c.Weapon = res[damageLineWeapon]

	c.Time, err = ParseField(res[damageLineTime], "time", parseLogTime(now))
	if err != nil {
		return err
	}
	c.Players.InitiatorID, err = ParseField(res[damageLineInitiatorID], "initiator_id", strconv.Atoi)
	if err != nil {
		return err
	}
	c.Players.RecipientID, err = ParseField(res[damageLineRecipientID], "recipient_id", strconv.Atoi)
	if err != nil {
		return err
	}
	c.DamageTotal, err = ParseField(res[damageLineDamage], "damage", parseFloat)
	if err != nil {
		return err
	}
	c.DamageHull, err = ParseField(res[damageLineHullDamage], "hull_damage", parseFloat)
	if err != nil {
		return err
	}
	c.DamageShield, err = ParseField(res[damageLineShieldDamage], "shield_damage", parseFloat)
	if err != nil {
		return err
	}
	c.DamageModifiers, err = ParseField(res[damageLineWeaponModifiers], "damage_modifiers", func(raw string) (res []DamageModifier, err error) {
		parts := strings.Split(raw, "|")
		if len(parts) == 0 {
			return res, fmt.Errorf("wrong parts number")
		}
		res = make([]DamageModifier, 0, len(parts))
		for _, p := range parts {
			res = append(res, DamageModifier(p))
		}
		return res, nil
	})
	if err != nil {
		return err
	}
	if len(res[damageLineFriendlyFire]) != 0 {
		c.IsFriendlyFire = true
	}

	return nil
}

type CombatLogLineHeal struct {
	Time    time.Time
	Players CombatPlayers
	Heal    float32
	Reason  string
}

func (c CombatLogLineHeal) CombatLogLineType() CombatLogLineType {
	return CombatLogLineTypeHeal
}

func (c *CombatLogLineHeal) Unmarshal(raw []byte, now time.Time) (err error) {
	res := combatRe.heal.FindStringSubmatch(string(raw))
	if len(res) != healLineTotal {
		return fmt.Errorf("wrong format: %s", raw)
	}
	c.Players.Initiator = res[healLineInitiator]
	c.Players.Recipient = res[healLineRecipient]
	c.Reason = res[healLineReason]

	c.Time, err = ParseField(res[healLineTime], "time", parseLogTime(now))
	if err != nil {
		return err
	}
	c.Players.InitiatorID, err = ParseField(res[healLineInitiatorID], "initiator_id", strconv.Atoi)
	if err != nil {
		return err
	}
	c.Players.RecipientID, err = ParseField(res[healLineRecipientID], "recipient_id", strconv.Atoi)
	if err != nil {
		return err
	}
	c.Heal, err = ParseField(res[healLineHeal], "heal", parseFloat)
	if err != nil {
		return err
	}
	return nil
}

type CombatLogLineKill struct {
	Time           time.Time
	Players        CombatPlayers
	KilledShip     string
	IsFriendlyFire bool
	Weapon         string
}

func (c CombatLogLineKill) CombatLogLineType() CombatLogLineType {
	return CombatLogLineTypeKill
}

func (c *CombatLogLineKill) Unmarshal(raw []byte, now time.Time) (err error) {
	res := combatRe.kill.FindStringSubmatch(string(raw))
	if len(res) != killLineTotal {
		return fmt.Errorf("wrong format: %s", raw)
	}
	c.Players.Initiator = res[killLineInitiator]
	c.Players.Recipient = res[killLineRecipient]
	c.KilledShip = res[killLineKilledShip]
	c.Weapon = res[killLineWeapon]

	c.Time, err = ParseField(res[killLineTime], "time", parseLogTime(now))
	if err != nil {
		return err
	}
	c.Players.InitiatorID, err = ParseField(res[killLineInitiatorID], "initiator_id", strconv.Atoi)
	if err != nil {
		return err
	}
	c.Players.RecipientID, err = ParseField(res[killLineRecipientID], "recipient_id", strconv.Atoi)
	if err != nil {
		return err
	}
	if len(res[killLineFriendlyFire]) != 0 {
		c.IsFriendlyFire = true
	}

	return nil
}

type CombatLogLineGameFinished struct {
	Time             time.Time
	WinnerTeamID     int
	WinnerTeamReason string
	FinishReason     string
	GameDuration     time.Duration
}

func (c CombatLogLineGameFinished) CombatLogLineType() CombatLogLineType {
	return CombatLogLineTypeGameplayFinished
}

func (c *CombatLogLineGameFinished) Unmarshal(raw []byte, now time.Time) (err error) {
	res := combatRe.gameFinished.FindStringSubmatch(string(raw))
	if len(res) != gameFinishedLineTotal {
		return fmt.Errorf("wrong format: %s", raw)
	}
	c.FinishReason = res[gameFinishedLineFinishReason]
	c.WinnerTeamReason = res[gameFinishedLineWinnerTeamReason]

	c.Time, err = ParseField(res[gameFinishedLineTime], "time", parseLogTime(now))
	if err != nil {
		return err
	}
	c.WinnerTeamID, err = ParseField(res[gameFinishedLineWinnerTeamID], "winner_team_id", strconv.Atoi)
	if err != nil {
		return err
	}
	c.GameDuration, err = ParseField(res[gameFinishedLineActualGameTime]+"s", "game_duration", time.ParseDuration)
	if err != nil {
		return err
	}

	return nil
}

var ErrUndefinedLineType = errors.New("undefined line type")

func ParseCombatLogLine(raw []byte, now time.Time) (line CombatLogLine, err error) {
	switch {
	case combatRe.connectToGameSession.Match(raw):
		line = &CombatLogLineConnectToGameSession{}
	case combatRe.startGame.Match(raw):
		line = &CombatLogLineStartGameplay{}
	case combatRe.damageShort.Match(raw):
		line = &CombatLogLineDamage{}
	case combatRe.healShort.Match(raw):
		line = &CombatLogLineHeal{}
	case combatRe.killShort.Match(raw):
		line = &CombatLogLineKill{}
	case combatRe.gameFinishedShort.Match(raw):
		line = &CombatLogLineGameFinished{}
	default:
		return nil, ErrUndefinedLineType
	}

	if err := line.Unmarshal(raw, now); err != nil {
		return nil, err
	}

	return line, nil
}
