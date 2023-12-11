package formatter

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Feresey/sclogparser/pkg/parser"
	"github.com/tealeg/xlsx/v3"
)

type Player struct {
	ID   int
	Name string
	Corp string
}

type Formatter struct {
	level       *parser.Level
	nameByIndex map[int]string
	indexByName map[string]int
	players     map[string]Player
}

const (
	columnActionType   = "action type"
	columnTime         = "time"
	columnInitiator    = "initiator"
	columnRecipient    = "recipient"
	columnDamage       = "damage"
	columnDamageHull   = "damage hull"
	columnDamageShield = "damage shield"
	columnHeal         = "heal"
	columnKilledShip   = "killed_ship"
	columnReason       = "reason"
	columnDamageType   = "damage type"
	columnExplosion    = "explosion"
	columnCrit         = "crit"
	columnCollision    = "collision"
	columnFirendlyFire = "firendly fire"
)

func New(level *parser.Level) *Formatter {
	f := &Formatter{
		level: level,
		nameByIndex: map[int]string{
			0:  columnActionType, // damage, kill, heal
			1:  columnTime,
			2:  columnInitiator,
			3:  columnRecipient,
			4:  columnDamage,
			5:  columnDamageHull,
			6:  columnDamageShield,
			7:  columnHeal,
			8:  columnKilledShip,
			9:  columnReason,
			10: columnDamageType,
			11: columnExplosion,
			12: columnCrit,
			13: columnCollision,
			14: columnFirendlyFire,
		},
		players: make(map[string]Player),
	}
	f.processPlayers()
	f.indexByName = make(map[string]int, len(f.nameByIndex))
	for idx, name := range f.nameByIndex {
		f.indexByName[name] = idx
	}
	return f
}

func (f *Formatter) processPlayers() {
	for _, p := range f.level.Game.AddedPlayers {
		if p.PlayerID == 0 || p.PlayerName == "" {
			continue
		}
		f.players[p.PlayerName] = Player{
			ID:   p.PlayerID,
			Name: p.PlayerName,
			Corp: p.PlayerCorp,
		}
	}
}

func (f Formatter) Format(out *xlsx.File) error {
	var (
		startTime  time.Time
		finishTime time.Time
	)

	startTime = f.level.Game.Connected.Time
	if startTime.After(f.level.Combat.Connect.Time) {
		startTime = f.level.Combat.Connect.Time
	}

	finishTime = f.level.Game.Connected.Time
	if finishTime.Before(f.level.Game.Finished.Time) {
		finishTime = f.level.Game.Finished.Time
	}

	sheetName := formatSheetName(f.level.Combat.Start.GameMode, startTime, finishTime)
	sheet, err := out.AddSheet(sheetName)
	if err != nil {
		return fmt.Errorf("xlxs.AddSheet(%s): %w", sheetName, err)
	}

	f.addHeaderRow(sheet)
	f.addDamageLogs(sheet)
	f.addKillLogs(sheet)

	return nil
}

func (f *Formatter) addHeaderRow(sheet *xlsx.Sheet) {
	header := sheet.AddRow()
	sortedNames := make([]string, len(f.nameByIndex))

	for idx, name := range f.nameByIndex {
		sortedNames[idx] = name
	}
	for _, name := range sortedNames {
		cell := header.AddCell()
		cell.Value = name
	}
}

func (f *Formatter) addDamageLogs(sheet *xlsx.Sheet) {
	for _, dmg := range f.level.Combat.Damage {
		var (
			isCollision bool
			isCrit      bool
			isExplosion bool
			damageType  string
		)
		for _, mod := range dmg.DamageModifiers {
			switch mod {
			case parser.DamageCollision:
				isCollision = true
			case parser.DamageCrit:
				isCrit = true
			case parser.DamageExplosion:
				isExplosion = true
			case parser.DamageTypeEMP,
				parser.DamageTypeThermal,
				parser.DamageTypeKinetic,
				parser.DamageTypeTrueDamage:
				damageType = string(mod)
			}
		}
		cells := []CellWithIndex{
			{
				Index: f.indexByName[columnActionType],
				Value: "damage",
			},
			{
				Index: f.indexByName[columnTime],
				Value: timeVal(dmg.Time),
			},
			{
				Index: f.indexByName[columnCollision],
				Value: boolVal(isCollision),
			},
			{
				Index: f.indexByName[columnCrit],
				Value: boolVal(isCrit),
			},
			{
				Index: f.indexByName[columnExplosion],
				Value: boolVal(isExplosion),
			},
			{
				Index: f.indexByName[columnDamage],
				Value: floatVal(dmg.DamageTotal),
			},
			{
				Index: f.indexByName[columnDamageType],
				Value: damageType,
			},
			{
				Index: f.indexByName[columnDamageHull],
				Value: floatVal(dmg.DamageHull),
			},
			{
				Index: f.indexByName[columnDamageShield],
				Value: floatVal(dmg.DamageShield),
			},
			{
				Index: f.indexByName[columnFirendlyFire],
				Value: boolVal(dmg.IsFriendlyFire),
			},
			{
				Index: f.indexByName[columnInitiator],
				Value: dmg.Players.Initiator,
			},
			{
				Index: f.indexByName[columnRecipient],
				Value: dmg.Players.Recipient,
			},
			{
				Index: f.indexByName[columnReason],
				Value: dmg.Weapon,
			},
		}

		f.addRow(sheet, cells)
	}
}

func (f *Formatter) addHealLogs(sheet *xlsx.Sheet) {
	for _, heal := range f.level.Combat.Heal {
		cells := []CellWithIndex{
			{
				Index: f.indexByName[columnActionType],
				Value: "heal",
			},
			{
				Index: f.indexByName[columnTime],
				Value: timeVal(heal.Time),
			},
			{
				Index: f.indexByName[columnHeal],
				Value: floatVal(heal.Heal),
			},
			{
				Index: f.indexByName[columnInitiator],
				Value: heal.Players.Initiator,
			},
			{
				Index: f.indexByName[columnRecipient],
				Value: heal.Players.Recipient,
			},
			{
				Index: f.indexByName[columnReason],
				Value: heal.Reason,
			},
		}

		f.addRow(sheet, cells)
	}
}

func (f *Formatter) addKillLogs(sheet *xlsx.Sheet) {
	for _, kill := range f.level.Combat.Kill {
		cells := []CellWithIndex{
			{
				Index: f.indexByName[columnActionType],
				Value: "kill",
			},
			{
				Index: f.indexByName[columnTime],
				Value: timeVal(kill.Time),
			},
			{
				Index: f.indexByName[columnKilledShip],
				Value: kill.KilledShip,
			},
			{
				Index: f.indexByName[columnInitiator],
				Value: kill.Players.Initiator,
			},
			{
				Index: f.indexByName[columnRecipient],
				Value: kill.Players.Recipient,
			},
			{
				Index: f.indexByName[columnReason],
				Value: kill.Weapon,
			},
			{
				Index: f.indexByName[columnFirendlyFire],
				Value: boolVal(kill.IsFriendlyFire),
			},
		}

		f.addRow(sheet, cells)
	}
}

func boolVal(val bool) string {
	if val {
		return "1"
	}
	return "0"
}

func floatVal(val float32) string {
	return strconv.FormatFloat(float64(val), 'f', 10, 64)
}

func intVal(val int) string {
	return strconv.Itoa(val)
}

func timeVal(val time.Time) string {
	return floatVal(float32(xlsx.TimeToExcelTime(val, false)))
}

type CellWithIndex struct {
	Index int
	Value string
}

func (f *Formatter) addRow(sheet *xlsx.Sheet, data []CellWithIndex) {
	row := sheet.AddRow()

	for _, one := range data {
		cell := row.GetCell(one.Index)
		cell.Value = one.Value
	}
}

const (
	timeOnly = "15-04-05"
)

func formatSheetName(gameMode string, start, end time.Time) string {
	return fmt.Sprintf("%s %s", gameMode, start.Format(timeOnly))
}
