package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type GameLogLineType int

const (
	GameLogLineTypeConnected GameLogLineType = iota
	GameLogLineTypeAddPlayer
	GameLogLineTypePlayerLeave
	// GameLogLineTypeNetStat
	GameLogLineTypeGameFinished
)

type GameLogLine interface {
	Type() GameLogLineType
	Unmarshal(raw []byte, now time.Time) error
}

const (
	gameLinePrefix = `(?s)^(\d{2}:\d{2}:\d{2}\.\d{3})\s+\| `
)

const (
	gameLineConnected     = gameLinePrefix + `client: connected to`
	gameLineConnectedTime = iota
	gameLineConnectedTotal
)

const (
	// 20:21:01.615         | client: ADD_PLAYER 0 (Feresey [LuX], 2516405) status 2 team 1 group 5212349
	gameLineAddPlayer     = gameLinePrefix + `client: ADD_PLAYER (\d+) \(([[:word:]]+)(\s+\[([[:word:]]*)\])?, (\d+)\) status (\d+) team (\d+)(\s+group (\d+))?\s*$`
	gameLineAddPlayerTime = iota
	gameLineAddPlayerSessionPlayerID
	gameLineAddPlayerPlayerName
	_
	gameLineAddPlayerCorp
	gameLineAddPlayerID
	gameLineAddPlayerStatus
	gameLineAddPlayerTeam
	_
	gameLineAddPlayerGroup
	gameLineAddPlayerTotal
)

const (
	gameLinePlayerLeave     = gameLinePrefix + `client: player (\d+) leave game\s*$`
	gameLinePlayerLeaveTime = iota
	gameLinePlayerLeaveSessionPlayerID
	gameLinePlayerLeaveTotal
)

// const (
// 	// 20:30:17.253         | client: avgPing 2.5/1.4; avgPacketLoss 0.0/0.0%; avgSnapshotLoss 0.0/0.2%, MTU 1492
// 	gameLineNetStat     = gameLinePrefix + `client: avgPing (\d+\.\d+)/(\d+\.\d+); avgPackageLoss (\d+\.\d+)/(\d+\.\d+)%; avgSnapshotLoss (\d+\.\d+)/(\d+\.\d+)%, MTU (\d+)`
// 	gameLineNetStatTime = iota
// 	gameLineNetStatPing1
// 	gameLineNetStatPing2
// 	gameLineNetStatPacketLoss1
// 	gameLineNetStatPacketLoss2
// 	gameLineNetStatSnapshotLoss1
// 	gameLineNetStatSnapshotLoss2
// 	gameLineNetStatMTU
// 	gameLineNetStatTotal
// )

const (
	gameLineFinished     = gameLinePrefix + `client: connection closed\. DR_CLIENT_GAME_FINISHED\s*$`
	gameLineFinishedTime = iota
	gameLineFinishedTotal
)

var gameRe = struct {
	connected,
	addPlayer,
	playerLeave,
	// netStat,
	finished,
	_ *regexp.Regexp
}{
	connected:   regexp.MustCompile(gameLineConnected),
	addPlayer:   regexp.MustCompile(gameLineAddPlayer),
	playerLeave: regexp.MustCompile(gameLinePlayerLeave),
	// netStat:     regexp.MustCompile(gameLineNetStat),
	finished: regexp.MustCompile(gameLineFinished),
}

type GameLogLineConnected struct {
	Time time.Time
}

func (g GameLogLineConnected) Type() GameLogLineType {
	return GameLogLineTypeConnected
}

func (g *GameLogLineConnected) Unmarshal(raw []byte, now time.Time) (err error) {
	res := gameRe.connected.FindStringSubmatch(string(raw))
	if len(res) != gameLineConnectedTotal {
		return fmt.Errorf("wrong format: %s", raw)
	}

	g.Time, err = ParseField(res[gameLineConnectedTime], "time", parseLogTime(now))
	if err != nil {
		return err
	}
	return nil
}

type GameLogLineAddPlayer struct {
	Time            time.Time
	SessionPlayerID int
	PlayerName      string
	PlayerCorp      string
	PlayerID        int
	Status          int
	TeamID          int
	GroupID         int
}

func (g GameLogLineAddPlayer) Type() GameLogLineType {
	return GameLogLineTypeAddPlayer
}

func (g *GameLogLineAddPlayer) Unmarshal(raw []byte, now time.Time) (err error) {
	res := gameRe.addPlayer.FindStringSubmatch(string(raw))
	if len(res) != gameLineAddPlayerTotal {
		return fmt.Errorf("wrong format: %s", raw)
	}

	g.PlayerName = res[gameLineAddPlayerPlayerName]
	g.PlayerCorp = res[gameLineAddPlayerCorp]

	g.Time, err = ParseField(res[gameLineAddPlayerTime], "time", parseLogTime(now))
	if err != nil {
		return err
	}
	g.SessionPlayerID, err = ParseField(res[gameLineAddPlayerSessionPlayerID], "session_player_id", strconv.Atoi)
	if err != nil {
		return err
	}
	g.PlayerID, err = ParseField(res[gameLineAddPlayerID], "player_id", strconv.Atoi)
	if err != nil {
		return err
	}
	g.Status, err = ParseField(res[gameLineAddPlayerStatus], "status", strconv.Atoi)
	if err != nil {
		return err
	}
	g.TeamID, err = ParseField(res[gameLineAddPlayerTeam], "team_id", strconv.Atoi)
	if err != nil {
		return err
	}
	if len(res[gameLineAddPlayerGroup]) != 0 {
		g.GroupID, err = ParseField(res[gameLineAddPlayerGroup], "group_id", strconv.Atoi)
		if err != nil {
			return err
		}
	}
	return nil
}

type GameLogLineFinished struct {
	Time time.Time
}

func (g GameLogLineFinished) Type() GameLogLineType {
	return GameLogLineTypeGameFinished
}

func (g *GameLogLineFinished) Unmarshal(raw []byte, now time.Time) (err error) {
	res := gameRe.finished.FindStringSubmatch(string(raw))
	if len(res) != gameLineFinishedTotal {
		return fmt.Errorf("wrong format: %s", raw)
	}

	g.Time, err = ParseField(res[gameLineFinishedTime], "time", parseLogTime(now))
	if err != nil {
		return err
	}
	return nil
}

// type GameLogLineNetStat struct {
// 	Time            time.Time
// 	SessionPlayerID int
// 	PlayerName      string
// 	PlayerCorp      string
// 	Status          int
// 	TeamID          int
// 	GroupID         int
// }

// func (g GameLogLineNetStat) GameLogLineType() GameLogLineType {
// 	return GameLogLineTypeNetStat
// }

// func (g *GameLogLineNetStat) Unmarshal(raw []byte, now time.Time) (err error) {
// 	res := gameRe.netStat.FindStringSubmatch(string(raw))
// 	if len(res) != gameLineNetStatTotal {
// 		return fmt.Errorf("wrong format: %s", raw)
// 	}

// 	g.Time, err = ParseField(res[gameLineNetStatTime], "time", parseLogTime(now))
// 	if err != nil {
// 		return err
// 	}
// 	g.Time, err = ParseField(res[gameLineNetStatTime], "time", parseLogTime(now))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

type GameLogLinePlayerLeave struct {
	Time     time.Time
	PlayerID int
}

func (g GameLogLinePlayerLeave) Type() GameLogLineType {
	return GameLogLineTypePlayerLeave
}

func (g *GameLogLinePlayerLeave) Unmarshal(raw []byte, now time.Time) (err error) {
	res := gameRe.playerLeave.FindStringSubmatch(string(raw))
	if len(res) != gameLinePlayerLeaveTotal {
		return fmt.Errorf("wrong format: %s", raw)
	}

	g.Time, err = ParseField(res[gameLinePlayerLeaveTime], "time", parseLogTime(now))
	if err != nil {
		return err
	}
	return nil
}

func ParseGameLogLine(raw []byte, now time.Time) (line GameLogLine, err error) {
	switch {
	case gameRe.connected.Match(raw):
		line = &GameLogLineConnected{}
	case gameRe.addPlayer.Match(raw):
		line = &GameLogLineAddPlayer{}
	case gameRe.playerLeave.Match(raw):
		line = &GameLogLinePlayerLeave{}
	case gameRe.finished.Match(raw):
		line = &GameLogLineFinished{}
	// case gameRe.netStat.Match(raw):
	// 	line = &GameLogLineConnected{}
	default:
		return nil, ErrUndefinedLineType
	}
	if err := line.Unmarshal(raw, now); err != nil {
		return nil, err
	}

	return line, nil
}
