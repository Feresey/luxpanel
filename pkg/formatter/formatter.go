package formatter

import (
	"fmt"
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

	startTime = f.level.Game.Connected.LogTime
	if startTime.After(f.level.Combat.Connect.LogTime) {
		startTime = f.level.Combat.Connect.LogTime
	}

	finishTime = f.level.Game.Connected.LogTime
	if finishTime.Before(f.level.Game.Finished.LogTime) {
		finishTime = f.level.Game.Finished.LogTime
	}

	sheetName := formatSheetName(f.level.Combat.Start.GameMode, startTime, finishTime)
	sheet, err := out.AddSheet(sheetName)
	if err != nil {
		return fmt.Errorf("xlxs.AddSheet(%s): %w", sheetName, err)
	}

	f.addHeaderRow(sheet)
	f.addDamageLogs(sheet)
	f.addKillLogs(sheet)

	for idx := range f.nameByIndex {
		sheet.SetColAutoWidth(idx+1, xlsx.DefaultAutoWidth)
	}

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
		cells := map[int]CellSetter{
			f.indexByName[columnActionType]:   StringCell("damage"),
			f.indexByName[columnTime]:         TimeCell(dmg.LogTime),
			f.indexByName[columnCollision]:    BoolCell(isCollision),
			f.indexByName[columnCrit]:         BoolCell(isCrit),
			f.indexByName[columnExplosion]:    BoolCell(isExplosion),
			f.indexByName[columnDamage]:       FloatCell(dmg.DamageTotal),
			f.indexByName[columnDamageType]:   StringCell(damageType),
			f.indexByName[columnDamageHull]:   FloatCell(dmg.DamageHull),
			f.indexByName[columnDamageShield]: FloatCell(dmg.DamageShield),
			f.indexByName[columnFirendlyFire]: BoolCell(dmg.IsFriendlyFire),
			f.indexByName[columnInitiator]:    StringCell(dmg.Players.Initiator),
			f.indexByName[columnRecipient]:    StringCell(dmg.Players.Recipient),
			f.indexByName[columnReason]:       StringCell(dmg.Weapon),
		}

		f.addRow(sheet, cells)
	}
}

func (f *Formatter) addHealLogs(sheet *xlsx.Sheet) {
	for _, heal := range f.level.Combat.Heal {
		cells := map[int]CellSetter{
			f.indexByName[columnActionType]: StringCell("heal"),
			f.indexByName[columnTime]:       TimeCell(heal.LogTime),
			f.indexByName[columnHeal]:       FloatCell(heal.Heal),
			f.indexByName[columnInitiator]:  StringCell(heal.Players.Initiator),
			f.indexByName[columnRecipient]:  StringCell(heal.Players.Recipient),
			f.indexByName[columnReason]:     StringCell(heal.Reason),
		}

		f.addRow(sheet, cells)
	}
}

func (f *Formatter) addKillLogs(sheet *xlsx.Sheet) {
	for _, kill := range f.level.Combat.Kill {
		cells := map[int]CellSetter{
			f.indexByName[columnActionType]:   StringCell("kill"),
			f.indexByName[columnTime]:         TimeCell(kill.LogTime),
			f.indexByName[columnKilledShip]:   StringCell(kill.KilledShip),
			f.indexByName[columnInitiator]:    StringCell(kill.Players.Initiator),
			f.indexByName[columnRecipient]:    StringCell(kill.Players.Recipient),
			f.indexByName[columnReason]:       StringCell(kill.Weapon),
			f.indexByName[columnFirendlyFire]: BoolCell(kill.IsFriendlyFire),
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

type CellSetter interface {
	SetCell(cell *xlsx.Cell)
}

type BoolCell bool

func (c BoolCell) SetCell(cell *xlsx.Cell) {
	cell.SetBool(bool(c))
}

type StringCell string

func (c StringCell) SetCell(cell *xlsx.Cell) {
	cell.SetString(string(c))
}

type FloatCell float64

func (c FloatCell) SetCell(cell *xlsx.Cell) {
	cell.SetFloat(float64(c))
}

type TimeCell time.Time

func (c TimeCell) SetCell(cell *xlsx.Cell) {
	cell.SetDateTime(time.Time(c))
}

func (f *Formatter) addRow(sheet *xlsx.Sheet, data map[int]CellSetter) {
	row := sheet.AddRow()

	for idx, val := range data {
		cell := row.GetCell(idx)
		val.SetCell(cell)
		cell.GetStyle().Font = *xlsx.NewFont(12, xlsx.TimesNewRoman)
	}
}

const (
	timeOnly = "15-04-05"
)

func formatSheetName(gameMode string, start, end time.Time) string {
	return fmt.Sprintf("%s %s", gameMode, start.Format(timeOnly))
}
