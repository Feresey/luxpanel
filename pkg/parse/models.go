package parse

import (
	"time"
)

type Player struct {
	Name    string
	ID      uint64
	ClanTag string
	GroupID int
}

type GameLogLevel struct {
	MapName  string
	GameMode string
	YourTeam int
	// Players is map[team_id]Player
	Players map[int][]Player

	LevelEnd time.Time
}
