package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

//go:generate stringer -type=GameLogLineType
type GameLogLineType int

const (
	GameLogLineTypeConnected GameLogLineType = iota
	GameLogLineTypeAddPlayer
	GameLogLineTypePlayerLeave
	GameLogLineTypeNetStat
	GameLogLineTypeGameFinished
)

type GameLogLine interface {
	Type() GameLogLineType
	Unmarshal(raw []byte, now time.Time) error
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
		line = &GameLogLineConnectionClosed{}
	case gameRe.netStat.Match(raw):
		line = &GameLogLineConnected{}
	default:
		return nil, fmt.Errorf("%w: %s", ErrUndefinedLineType, raw)
	}
	if err := line.Unmarshal(raw, now); err != nil {
		return nil, fmt.Errorf("line matched to regex(%s), but failed to parse: %s: %s", line.Type().String(), raw, err.Error())
	}

	return line, nil
}

//nolint:gochecknoglobals // this is const
var gameRe = struct {
	connected,
	addPlayer,
	playerLeave,
	netStat,
	finished,
	_ *regexp.Regexp
}{
	connected:   regexp.MustCompile(gameLineConnected),
	addPlayer:   regexp.MustCompile(gameLineAddPlayer),
	playerLeave: regexp.MustCompile(gameLinePlayerLeave),
	netStat:     regexp.MustCompile(gameLineNetStat),
	finished:    regexp.MustCompile(gameLineConnectionClosed),
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
	playerNameAndState = `\(([[:word:]]+)(\s+\[([[:word:]]*)\])?, (\d+)\)`
)

const (
	// 20:21:01.615         | client: ADD_PLAYER 0 (Feresey [LuX], 2516405) status 2 team 1 group 5212349
	gameLineAddPlayer = gameLinePrefix + `client: ADD_PLAYER (\d+)` + space +
		playerNameAndState + `\s+status\s+(\d+) ` +
		`team\s+(\d+)` +
		`(\s+group\s+(\d+))?\s*$`
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
	gameLinePlayerLeave = gameLinePrefix + `client: player (\d+) leave game\s*$`
)
const (
	gameLinePlayerLeaveTime = iota + 1
	gameLinePlayerLeaveSessionPlayerID
	gameLinePlayerLeaveTotal
)

const (
	// 20:30:17.253         | client: avgPing 2.5/1.4; avgPacketLoss 0.0/0.0%; avgSnapshotLoss 0.0/0.2%, MTU 1492
	gameLineNetStat = gameLinePrefix + `client:` + space +
		`avgPing (\d+\.\d+)/(\d+\.\d+);` + space +
		`avgPackageLoss (\d+\.\d+)/(\d+\.\d+)%;` + space +
		`avgSnapshotLoss (\d+\.\d+)/(\d+\.\d+)%,` + space +
		`MTU (\d+)`
	gameLineNetStatTime = iota
	gameLineNetStatPing1
	gameLineNetStatPing2
	gameLineNetStatPacketLoss1
	gameLineNetStatPacketLoss2
	gameLineNetStatSnapshotLoss1
	gameLineNetStatSnapshotLoss2
	gameLineNetStatMTU
	gameLineNetStatTotal
)

const gameLineConnectionClosed = gameLinePrefix + `client: connection closed\. ([A-Z_]+)\s*$`

const (
	gameLineConnectionClosedTime = iota + 1
	gameLineConnectionClosedReason
	gameLineConnectionClosedTotal
)

type GameLogLineConnected struct {
	LogTime time.Time
}

func (g GameLogLineConnected) Type() GameLogLineType {
	return GameLogLineTypeConnected
}

func (g *GameLogLineConnected) Unmarshal(raw []byte, now time.Time) (err error) {
	res := gameRe.connected.FindStringSubmatch(string(raw))
	if len(res) != gameLineConnectedTotal {
		return ErrWrongLineFormat
	}

	g.LogTime, err = ParseField(res[gameLineConnectedTime], "time", parseLogTime(now))
	if err != nil {
		return err
	}
	return nil
}

type GameLogLineAddPlayer struct {
	LogTime         time.Time
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
		return ErrWrongLineFormat
	}

	g.PlayerName = res[gameLineAddPlayerPlayerName]
	g.PlayerCorp = res[gameLineAddPlayerCorp]

	g.LogTime, err = ParseField(res[gameLineAddPlayerTime], "time", parseLogTime(now))
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

type ConnectionClosedReason string

const (
	ConnectionClosedReasonGameFinished          ConnectionClosedReason = "DR_CLIENT_GAME_FINISHED"
	ConnectionClosedReasonClientCouldNotConnect ConnectionClosedReason = "DR_CLIENT_COULD_NOT_CONNECT"
	ConnectionClosedReasonClientQuit            ConnectionClosedReason = "DR_CLIENT_QUIT"
)

func (c ConnectionClosedReason) Validate() error {
	switch c {
	case ConnectionClosedReasonGameFinished:
	case ConnectionClosedReasonClientCouldNotConnect:
	case ConnectionClosedReasonClientQuit:
	default:
		return fmt.Errorf("undefined connection closed reason: %q", c)
	}
	return nil
}

func parseConnectionClosedReason(s string) (res ConnectionClosedReason, err error) {
	res = ConnectionClosedReason(s)
	if err := res.Validate(); err != nil {
		return res, err
	}
	return res, nil
}

type GameLogLineConnectionClosed struct {
	LogTime time.Time
	Reason  ConnectionClosedReason
}

func (g GameLogLineConnectionClosed) Type() GameLogLineType {
	return GameLogLineTypeGameFinished
}

func (g *GameLogLineConnectionClosed) Unmarshal(raw []byte, now time.Time) (err error) {
	res := gameRe.finished.FindStringSubmatch(string(raw))
	if len(res) != gameLineConnectionClosedTotal {
		return ErrWrongLineFormat
	}

	g.LogTime, err = ParseField(res[gameLineConnectionClosedTime], "time", parseLogTime(now))
	if err != nil {
		return err
	}

	g.Reason, err = ParseField(res[gameLineConnectionClosedReason], "reason", parseConnectionClosedReason)
	if err != nil {
		return err
	}
	return nil
}

type GameLogLineNetStat struct {
	LogTime         time.Time
	SessionPlayerID int
	PlayerName      string
	PlayerCorp      string
	Status          int
	TeamID          int
	GroupID         int
}

func (g GameLogLineNetStat) Type() GameLogLineType {
	return GameLogLineTypeNetStat
}

func (g *GameLogLineNetStat) Unmarshal(raw []byte, now time.Time) (err error) {
	res := gameRe.netStat.FindStringSubmatch(string(raw))
	if len(res) != gameLineNetStatTotal {
		return ErrWrongLineFormat
	}

	g.LogTime, err = ParseField(res[gameLineNetStatTime], "time", parseLogTime(now))
	if err != nil {
		return err
	}
	return nil
}

type GameLogLinePlayerLeave struct {
	LogTime  time.Time
	PlayerID int
}

func (g GameLogLinePlayerLeave) Type() GameLogLineType {
	return GameLogLineTypePlayerLeave
}

func (g *GameLogLinePlayerLeave) Unmarshal(raw []byte, now time.Time) (err error) {
	res := gameRe.playerLeave.FindStringSubmatch(string(raw))
	if len(res) != gameLinePlayerLeaveTotal {
		return ErrWrongLineFormat
	}

	g.LogTime, err = ParseField(res[gameLinePlayerLeaveTime], "time", parseLogTime(now))
	if err != nil {
		return err
	}

	g.PlayerID, err = ParseField(res[gameLinePlayerLeaveSessionPlayerID], "player_id", strconv.Atoi)
	if err != nil {
		return err
	}
	return nil
}
