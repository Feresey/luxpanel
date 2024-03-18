package splitter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/maps"

	"github.com/Feresey/sclogparser/internal/logger"
	"github.com/Feresey/sclogparser/internal/parser"
)

var ErrLogsCorrupted = errors.New("logs corrupted")

type Splitter struct {
	lg logger.Factory
	tr trace.Tracer
}

func NewSplitter(lg logger.Factory, tr trace.TracerProvider) *Splitter {
	return &Splitter{lg: lg, tr: tr.Tracer("splitter")}
}

func (f *Splitter) SplitLevels(ctx context.Context, gameLog []parser.GameLogLine, combatLog []parser.CombatLogLine) ([]*Level, error) {
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

	playerTeamsMap := make(map[int]map[string]Player)
	for _, logPlayer := range gameLevel.AddPlayer {
		player := Player{
			ID:      logPlayer.PlayerID,
			Name:    logPlayer.PlayerName,
			CorpTag: logPlayer.PlayerCorp,
			TeamID:  logPlayer.TeamID,
		}
		teamMap := playerTeamsMap[logPlayer.TeamID]
		if teamMap == nil {
			teamMap = make(map[string]Player)
			playerTeamsMap[logPlayer.TeamID] = teamMap
		}

		teamMap[player.String()] = player
	}

	f.lg.For(ctx).Debugw("detected teams", "team_ids", maps.Keys(playerTeamsMap))
	lvl.Teams = make(map[int][]Player, len(playerTeamsMap))
	for teamID, team := range playerTeamsMap {
		players := maps.Values(team)
		f.lg.For(ctx).Debugw("detected team players", "team_id", teamID, "players", players)
		lvl.Teams[teamID] = players
	}

	if len(lvl.Teams) == 0 {
		return nil, fmt.Errorf("no teams found")
	}

	return lvl, nil
}

type Player struct {
	ID      int
	Name    string
	CorpTag string
	TeamID  int
}

func (p Player) String() string {
	return fmt.Sprintf("%s [%s] (id: %d) (in_game_id: %d)", p.Name, p.CorpTag, p.ID, p.TeamID)
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
	Lines []parser.GameLogLine

	StartGameplay  *parser.GameLogLineConnected
	AddPlayer      []*parser.GameLogLineAddPlayer
	LeavePlayer    []*parser.GameLogLinePlayerLeave
	FinishGameplay *parser.GameLogLineConnectionClosed
}

type CombatLogLevel struct {
	Lines []parser.CombatLogLine

	Connect  *parser.CombatLogLineConnectToGameSession
	Start    *parser.CombatLogLineStartGameplay
	Damage   []*parser.CombatLogLineDamage
	Heal     []*parser.CombatLogLineHeal
	Kill     []*parser.CombatLogLineKill
	Finished *parser.CombatLogLineGameFinished
}

func (f *Splitter) GetGameLogLevels(ctx context.Context, lines []parser.GameLogLine) (res []*GameLogLevel) {
	ctx, span := f.tr.Start(ctx, "GetGameLogLevels")
	defer span.End()

	currLevel := new(GameLogLevel)

	for _, line := range lines {
		currLevel.Lines = append(currLevel.Lines, line)
		switch line := line.(type) {
		case *parser.GameLogLineConnected:
			if currLevel.StartGameplay != nil {
				f.lg.For(ctx).Warnw("start gameplay twice", "prev", currLevel.StartGameplay, "next", line,
					"start", currLevel.StartGameplay, "end", currLevel.FinishGameplay)
				res = append(res, currLevel)
				currLevel = new(GameLogLevel)
			}
			currLevel.StartGameplay = line
		case *parser.GameLogLineAddPlayer:
			currLevel.AddPlayer = append(currLevel.AddPlayer, line)
		case *parser.GameLogLineConnectionClosed:
			currLevel.FinishGameplay = line
			f.lg.For(ctx).Infow("detected could not connect log", "line", line)
			if line.Reason == parser.ConnectionClosedReasonClientCouldNotConnect {
				currLevel = new(GameLogLevel)
				continue
			}
			res = append(res, currLevel)
			currLevel = new(GameLogLevel)
		case *parser.GameLogLinePlayerLeave:
			currLevel.LeavePlayer = append(currLevel.LeavePlayer, line)
		case *parser.GameLogLineNetStat: // Я хз что с этим делать
		}
	}

	f.lg.For(ctx).Infow("got levels", "count", len(res))
	return res
}

func (f *Splitter) GetCombatLogLevels(ctx context.Context, lines []parser.CombatLogLine) (res []*CombatLogLevel) {
	ctx, span := f.tr.Start(ctx, "GetCombatLogLevels")
	defer span.End()

	currLevel := new(CombatLogLevel)

	for _, line := range lines {
		currLevel.Lines = append(currLevel.Lines, line)
		switch line := line.(type) {
		case *parser.CombatLogLineConnectToGameSession:
			if currLevel.Connect != nil && currLevel.Connect.SessionID != line.SessionID {
				f.lg.For(ctx).Warnw("connect to game session twice",
					"prev", currLevel.Connect, "next", line,
					"connect", currLevel.Connect, "start", currLevel.Start, "end", currLevel.Finished)
				res = append(res, currLevel)
				currLevel = new(CombatLogLevel)
			}
			currLevel.Connect = line
		case *parser.CombatLogLineStartGameplay:
			if currLevel.Start != nil {
				f.lg.For(ctx).Warnw("start gameplay twice", "prev", currLevel.Start, "next", line,
					"connect", currLevel.Connect, "start", currLevel.Start, "end", currLevel.Finished)
				res = append(res, currLevel)
				currLevel = new(CombatLogLevel)
			}
			currLevel.Start = line
		case *parser.CombatLogLineDamage:
			currLevel.Damage = append(currLevel.Damage, line)
		case *parser.CombatLogLineHeal:
			currLevel.Heal = append(currLevel.Heal, line)
		case *parser.CombatLogLineKill:
			currLevel.Kill = append(currLevel.Kill, line)
		case *parser.CombatLogLineGameFinished:
			currLevel.Finished = line
			f.lg.For(ctx).Debugw("finished level", "connect", currLevel.Connect, "start", currLevel.Start, "end", currLevel.Finished)
			res = append(res, currLevel)
			currLevel = new(CombatLogLevel)
		}
	}

	f.lg.For(ctx).Infow("got levels", "count", len(res))
	return res
}
