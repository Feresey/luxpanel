package splitter

import (
	"context"
	"errors"
	"fmt"
	"os"
	"slices"
	"time"

	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/maps"

	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/Feresey/luxpanel/internal/parser/game"
)

var ErrLogsCorrupted = errors.New("logs corrupted")

type Splitter struct {
	lg logger.Factory
	tr trace.Tracer
}

func NewSplitter(lg logger.Factory, tr trace.TracerProvider) *Splitter {
	return &Splitter{lg: lg, tr: tr.Tracer("splitter")}
}

func (f *Splitter) SplitLevels(ctx context.Context, gameLog []game.LogLine, combatLog []combat.LogLine) ([]*Level, error) {
	ctx, span := f.tr.Start(ctx, "SplitLevels")
	defer span.End()

	gameLevels := f.GetGameLogLevels(ctx, gameLog)
	combatLevels := f.GetCombatLogLevels(ctx, combatLog)
	if len(gameLevels) != len(combatLevels) {
		mx := max(len(gameLevels), len(combatLevels))
		f.lg.For(ctx).Infow("levels count mismatch", "combat", len(combatLevels), "game", len(gameLevels))
		for i := range mx {
			if i < len(gameLevels) {
				gm := gameLevels[i]
				f.lg.For(ctx).Infow("game log time", "number", i, "start", gm.StartGameplay, "end", gm.FinishGameplay)
			}
			if i < len(combatLevels) {
				cm := combatLevels[i]
				f.lg.For(ctx).Infow("combat log time", "number", i, "start", cm.Start, "end", cm.Finished)
			}
		}
		return nil, fmt.Errorf("%w: levels count mismatch: game logs: %d, combat logs: %d", ErrLogsCorrupted, len(gameLevels), len(combatLevels))
	}

	levels := make([]*Level, 0, len(gameLevels))
	mx := max(len(gameLevels), len(combatLevels))
	for i := range mx {
		var gm *GameLogLevel
		var cm *CombatLogLevel
		if i < len(gameLevels) {
			gm = gameLevels[i]
		}
		if i < len(combatLevels) {
			cm = combatLevels[i]
		}

		lvl, err := f.makeLevel(ctx, gm, cm)
		if err != nil {
			return nil, fmt.Errorf("makeLevel: %w", err)
		}
		levels = append(levels, lvl)
	}

	f.lg.For(ctx).Debugw("got games", "count", len(levels))
	return levels, nil
}

func (f *Splitter) makeLevel(ctx context.Context, gameLevel *GameLogLevel, combatLevel *CombatLogLevel) (*Level, error) {
	ctx, span := f.tr.Start(ctx, "makeLevel")
	defer span.End()

	lvl := &Level{
		GameLog:   gameLevel,
		CombatLog: combatLevel,
	}

	if gameLevel != nil {
		lvl.StartLevelTime = gameLevel.StartGameplay.LogTime
	}
	if combatLevel != nil {
		if combatLevel.Start != nil && lvl.StartLevelTime.After(combatLevel.Start.LogTime) {
			lvl.StartLevelTime = combatLevel.Start.LogTime
		}
		if combatLevel.Connect != nil && lvl.StartLevelTime.After(combatLevel.Connect.LogTime) {
			lvl.StartLevelTime = combatLevel.Connect.LogTime
		}
	}
	if gameLevel != nil && gameLevel.FinishGameplay != nil {
		lvl.EndLevelTime = gameLevel.FinishGameplay.LogTime
	}
	if combatLevel != nil && combatLevel.Finished != nil && lvl.EndLevelTime.Before(combatLevel.Finished.LogTime) {
		lvl.EndLevelTime = combatLevel.Finished.LogTime
	}

	playerTeamsMap := make(map[int]map[int]Player)
	for _, logPlayer := range gameLevel.AddPlayer {
		player := Player{
			PlayerID:        logPlayer.PlayerID,
			SessionPlayerID: logPlayer.SessionPlayerID,
			Name:            logPlayer.PlayerName,
			CorpTag:         logPlayer.PlayerCorpTag,
			TeamID:          logPlayer.TeamID,
		}
		teamMap := playerTeamsMap[logPlayer.TeamID]
		if teamMap == nil {
			teamMap = make(map[int]Player)
			playerTeamsMap[logPlayer.TeamID] = teamMap
		}

		teamMap[player.SessionPlayerID] = player
	}

	teams := maps.Keys(playerTeamsMap)
	slices.Sort(teams)
	if os.Getenv("SHOW_WATCHERS") == "" {
		if teams[0] == 0 {
			teams = teams[1:]
		}
	}
	f.lg.For(ctx).Debugw("level", "level", lvl.CombatLog.Start, "team_ids", teams)

	lvl.Teams = make(map[int][]Player, len(playerTeamsMap))
	for teamID, team := range playerTeamsMap {
		players := maps.Values(team)
		playerNames := make([]string, 0, len(players))
		for _, p := range players {
			playerNames = append(playerNames, p.String())
		}
		slices.Sort(playerNames)
		if os.Getenv("SHOW_WATCHERS") != "" && teamID == 0 {
			f.lg.For(ctx).Debugw("team players", "team_id", teamID, "players", playerNames)
		}
		lvl.Teams[teamID] = players
	}

	if len(lvl.Teams) == 0 {
		return nil, fmt.Errorf("no teams found")
	}

	return lvl, nil
}

type Player struct {
	PlayerID        int
	SessionPlayerID int
	Name            string
	CorpTag         string
	TeamID          int
}

func (p Player) String() string {
	return fmt.Sprintf("%s [%s] (id: %d)", p.Name, p.CorpTag, p.PlayerID)
}

type Level struct {
	StartLevelTime time.Time
	EndLevelTime   time.Time
	Teams          map[int][]Player

	GameLog   *GameLogLevel
	CombatLog *CombatLogLevel
}

// TODO make timeline func for level

type GameLogLevel struct {
	Lines []game.LogLine

	StartGameplay  *game.Connected
	AddPlayer      []*game.AddPlayer
	LeavePlayer    []*game.PlayerLeave
	FinishGameplay *game.ConnectionClosed
}

type CombatLogLevel struct {
	Lines []combat.LogLine

	Connect  *combat.ConnectToGameSession
	Start    *combat.StartGameplay
	Damage   []*combat.Damage
	Heal     []*combat.Heal
	Kill     []*combat.Kill
	Finished *combat.FinishedGameplay
}

func (f *Splitter) GetGameLogLevels(ctx context.Context, lines []game.LogLine) (res []*GameLogLevel) {
	ctx, span := f.tr.Start(ctx, "GetGameLogLevels")
	defer span.End()

	currLevel := new(GameLogLevel)

	for _, line := range lines {
		currLevel.Lines = append(currLevel.Lines, line)
		switch line := line.(type) {
		case *game.Connected:
			if currLevel.StartGameplay != nil {
				f.lg.For(ctx).Warnw("start gameplay twice", "prev", currLevel.StartGameplay, "next", line,
					"start", currLevel.StartGameplay, "end", currLevel.FinishGameplay)
				res = append(res, currLevel)
				currLevel = new(GameLogLevel)
			}
			currLevel.StartGameplay = line
		case *game.AddPlayer:
			currLevel.AddPlayer = append(currLevel.AddPlayer, line)
		case *game.ConnectionClosed:
			currLevel.FinishGameplay = line
			if line.CloseReason == game.ConnectionClosedReasonClientCouldNotConnect {
				f.lg.For(ctx).Infow("detected could not connect log", "line", line)
				currLevel = new(GameLogLevel)
				continue
			}
			res = append(res, currLevel)
			currLevel = new(GameLogLevel)
		case *game.PlayerLeave:
			currLevel.LeavePlayer = append(currLevel.LeavePlayer, line)
		}
	}

	f.lg.For(ctx).Infow("got game log levels", "count", len(res))
	return res
}

func (f *Splitter) GetCombatLogLevels(ctx context.Context, lines []combat.LogLine) (res []*CombatLogLevel) {
	ctx, span := f.tr.Start(ctx, "GetCombatLogLevels")
	defer span.End()

	currLevel := new(CombatLogLevel)

	for _, line := range lines {
		currLevel.Lines = append(currLevel.Lines, line)
		switch line := line.(type) {
		case *combat.ConnectToGameSession:
			if currLevel.Connect != nil && currLevel.Connect.SessionID != line.SessionID {
				f.lg.For(ctx).Warnw("connect to game session twice",
					"prev", currLevel.Connect, "next", line,
					"connect", currLevel.Connect, "start", currLevel.Start, "end", currLevel.Finished)
				res = append(res, currLevel)
				currLevel = new(CombatLogLevel)
			}
			currLevel.Connect = line
		case *combat.StartGameplay:
			if currLevel.Start != nil {
				f.lg.For(ctx).Warnw("start gameplay twice", "prev", currLevel.Start, "next", line,
					"connect", currLevel.Connect, "start", currLevel.Start, "end", currLevel.Finished)
				res = append(res, currLevel)
				currLevel = new(CombatLogLevel)
			}
			currLevel.Start = line
		case *combat.Damage:
			currLevel.Damage = append(currLevel.Damage, line)
		case *combat.Heal:
			currLevel.Heal = append(currLevel.Heal, line)
		case *combat.Kill:
			currLevel.Kill = append(currLevel.Kill, line)
		case *combat.FinishedGameplay:
			currLevel.Finished = line
			// f.lg.For(ctx).Debugw("finished level", "connect", currLevel.Connect, "start", currLevel.Start, "end", currLevel.Finished)
			res = append(res, currLevel)
			currLevel = new(CombatLogLevel)
		}
	}

	f.lg.For(ctx).Infow("got combat log levels", "count", len(res))
	return res
}
