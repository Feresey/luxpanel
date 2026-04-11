package splitter

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"slices"
	"strings"
	"time"

	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/maps"

	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/Feresey/luxpanel/internal/parser/common"
	"github.com/Feresey/luxpanel/internal/parser/game"
)

var ErrLogsCorrupted = errors.New("logs corrupted")

type Splitter struct {
	lg     logger.Factory
	tr     trace.Tracer
	parser *parser.Parser
}

func NewSplitter(lg logger.Factory, tr trace.TracerProvider, parser *parser.Parser) *Splitter {
	return &Splitter{lg: lg, tr: tr.Tracer("splitter"), parser: parser}
}

// LevelSpan is a half-open byte range [Start, End) into the original game.log / combat.log strings.
type LevelSpan struct {
	GameStart, GameEnd     int
	CombatStart, CombatEnd int
}

// DiscoverResult is the outcome of the lazy first pass: per-match byte spans and preview levels for UI meta.
type DiscoverResult struct {
	LogTime          time.Time
	GameLog          string
	CombatLog        string
	GameHeaderEnd    int
	CombatHeaderEnd  int
	Spans            []LevelSpan
	PreviewLevels    []*Level
}

func (s *Splitter) SplitLevels(ctx context.Context, fsys fs.FS) ([]*Level, error) {
	ctx, span := s.tr.Start(ctx, "SplitLevels")
	defer span.End()

	gameBytes, err := fs.ReadFile(fsys, "game.log")
	if err != nil {
		return nil, fmt.Errorf("fs.ReadFile(game.log): %w", err)
	}
	combatBytes, err := fs.ReadFile(fsys, "combat.log")
	if err != nil {
		return nil, fmt.Errorf("fs.ReadFile(combat.log): %w", err)
	}

	disc, err := s.DiscoverLevels(ctx, string(gameBytes), string(combatBytes))
	if err != nil {
		return nil, err
	}

	levels := make([]*Level, len(disc.Spans))
	for i := range disc.Spans {
		lvl, err := s.HydrateLevel(ctx, disc, i)
		if err != nil {
			return nil, fmt.Errorf("HydrateLevel %d: %w", i, err)
		}
		levels[i] = lvl
	}
	return levels, nil
}

// DiscoverLevels runs lazy parsers over full logs, pairs levels, records byte spans, and builds preview levels (roster meta only).
func (s *Splitter) DiscoverLevels(ctx context.Context, gameLog, combatLog string) (*DiscoverResult, error) {
	ctx, span := s.tr.Start(ctx, "DiscoverLevels")
	defer span.End()

	logTime, gameHeaderEnd, err := parser.ParseLogHeader(gameLog)
	if err != nil {
		return nil, fmt.Errorf("game.log header: %w", err)
	}
	_, combatHeaderEnd, err := parser.ParseLogHeader(combatLog)
	if err != nil {
		return nil, fmt.Errorf("combat.log header: %w", err)
	}

	var gameLines []parser.LogLine[game.LogLine]
	if _, err := s.parser.WalkGameLogLazyString(ctx, gameLog, func(ll parser.LogLine[game.LogLine]) error {
		gameLines = append(gameLines, ll)
		return nil
	}); err != nil {
		return nil, fmt.Errorf("WalkGameLogLazyString: %w", err)
	}

	var combatLines []parser.LogLine[combat.LogLine]
	if _, err := s.parser.WalkCombatLogLazyString(ctx, combatLog, func(ll parser.LogLine[combat.LogLine]) error {
		combatLines = append(combatLines, ll)
		return nil
	}); err != nil {
		return nil, fmt.Errorf("WalkCombatLogLazyString: %w", err)
	}

	gameLevels, gameSpans := s.collectGameLogLevels(ctx, gameLines)
	combatLevels, combatSpans, combatErrs := s.collectCombatLogLevels(ctx, logTime, combatLines)
	_ = combatErrs // same as before: non-fatal

	if len(gameLevels) != len(combatLevels) {
		mx := max(len(gameLevels), len(combatLevels))
		s.lg.For(ctx).Infow("levels count mismatch", "combat", len(combatLevels), "game", len(gameLevels))
		for i := range mx {
			if i < len(gameLevels) {
				gm := gameLevels[i]
				var gs, ge time.Time
				if gm != nil {
					if gm.StartGameplay != nil {
						gs = gm.StartGameplay.GetTime(logTime)
					}
					if gm.FinishGameplay != nil {
						ge = gm.FinishGameplay.GetTime(logTime)
					}
				}
				s.lg.For(ctx).Infow("game log time", "number", i, "start", gs, "end", ge, "gameplay", gm.StartGameplay, "finish", gm.FinishGameplay)
			}
			if i < len(combatLevels) {
				cm := combatLevels[i]
				s.lg.For(ctx).Infow("combat log time", "number", i, "start", cm.Start.GetTime(logTime), "end", cm.Finished.GetTime(logTime), "cmbt_start", cm.Start, "cmbnt_finish", cm.Finished, "connect", cm.Connect)
			}
		}

		for i, cm := range combatLevels {
			s.lg.For(ctx).Infow("combat log time", "number", i, "start", cm.Start.GetTime(logTime), "end", cm.Finished.GetTime(logTime), "cmbt_start", cm.Start, "cmbnt_finish", cm.Finished, "connect", cm.Connect)
		}
		return nil, fmt.Errorf("%w: levels count mismatch: game logs: %d, combat logs: %d", ErrLogsCorrupted, len(gameLevels), len(combatLevels))
	}

	if len(gameSpans) != len(gameLevels) || len(combatSpans) != len(combatLevels) {
		return nil, fmt.Errorf("internal span mismatch")
	}

	spans := make([]LevelSpan, len(gameLevels))
	for i := range gameLevels {
		spans[i] = LevelSpan{
			GameStart:   gameSpans[i].Start,
			GameEnd:     gameSpans[i].End,
			CombatStart: combatSpans[i].Start,
			CombatEnd:   combatSpans[i].End,
		}
	}

	preview := make([]*Level, len(gameLevels))
	for i := range gameLevels {
		lvl, err := s.makeLevel(ctx, logTime, gameLevels[i], combatLevels[i])
		if err != nil {
			return nil, fmt.Errorf("makeLevel preview %d: %w", i, err)
		}
		s.LogMatchParseStats(ctx, i, len(gameLevels), logTime, lvl)
		preview[i] = lvl
	}

	s.lg.For(ctx).Debugw("discovered games", "count", len(preview))
	return &DiscoverResult{
		LogTime:         logTime,
		GameLog:         gameLog,
		CombatLog:       combatLog,
		GameHeaderEnd:   gameHeaderEnd,
		CombatHeaderEnd: combatHeaderEnd,
		Spans:           spans,
		PreviewLevels:   preview,
	}, nil
}

// lineSpan is a half-open byte range [Start, End) into one log file.
type lineSpan struct {
	Start, End int
}

// HydrateLevel runs full parsers only on the byte ranges for match idx.
func (s *Splitter) HydrateLevel(ctx context.Context, disc *DiscoverResult, idx int) (*Level, error) {
	if disc == nil || idx < 0 || idx >= len(disc.Spans) {
		return nil, fmt.Errorf("invalid discover result or index")
	}
	ctx, span := s.tr.Start(ctx, "HydrateLevel")
	defer span.End()

	sp := disc.Spans[idx]
	if sp.GameStart < 0 || sp.GameEnd < sp.GameStart || sp.GameEnd > len(disc.GameLog) {
		return nil, fmt.Errorf("game span out of range")
	}
	if sp.CombatStart < 0 || sp.CombatEnd < sp.CombatStart || sp.CombatEnd > len(disc.CombatLog) {
		return nil, fmt.Errorf("combat span out of range")
	}

	gameR := io.MultiReader(
		strings.NewReader(disc.GameLog[:disc.GameHeaderEnd]),
		strings.NewReader(disc.GameLog[sp.GameStart:sp.GameEnd]),
	)
	var gameLines []parser.LogLine[game.LogLine]
	if _, err := s.parser.WalkGameLogFull(ctx, gameR, func(ll parser.LogLine[game.LogLine]) error {
		gameLines = append(gameLines, ll)
		return nil
	}); err != nil {
		return nil, fmt.Errorf("WalkGameLogFull: %w", err)
	}

	combatR := io.MultiReader(
		strings.NewReader(disc.CombatLog[:disc.CombatHeaderEnd]),
		strings.NewReader(disc.CombatLog[sp.CombatStart:sp.CombatEnd]),
	)
	var combatLines []parser.LogLine[combat.LogLine]
	if _, err := s.parser.WalkCombatLogFull(ctx, combatR, func(ll parser.LogLine[combat.LogLine]) error {
		combatLines = append(combatLines, ll)
		return nil
	}); err != nil {
		return nil, fmt.Errorf("WalkCombatLogFull: %w", err)
	}

	gLevels, _ := s.collectGameLogLevels(ctx, gameLines)
	cLevels, _, _ := s.collectCombatLogLevels(ctx, disc.LogTime, combatLines)
	if len(gLevels) != 1 || len(cLevels) != 1 {
		return nil, fmt.Errorf("hydrated slice expected exactly one game and one combat level, got game=%d combat=%d", len(gLevels), len(cLevels))
	}

	lvl, err := s.makeLevel(ctx, disc.LogTime, gLevels[0], cLevels[0])
	if err != nil {
		return nil, fmt.Errorf("makeLevel: %w", err)
	}
	s.LogMatchParseStats(ctx, idx, len(disc.Spans), disc.LogTime, lvl)
	return lvl, nil
}

func (s *Splitter) makeLevel(ctx context.Context, logTime time.Time, gameLevel *GameLogLevel, combatLevel *CombatLogLevel) (*Level, error) {
	ctx, span := s.tr.Start(ctx, "makeLevel")
	defer span.End()

	lvl := &Level{
		GameLog:   gameLevel,
		CombatLog: combatLevel,
	}

	if gameLevel != nil && gameLevel.StartGameplay != nil {
		lvl.StartLevelTime = gameLevel.StartGameplay.GetTime(logTime)
	}
	if combatLevel != nil {
		for _, t := range []time.Time{
			combatLevel.Connect.GetTime(logTime),
			combatLevel.Start.GetTime(logTime),
		} {
			if t.IsZero() {
				continue
			}
			if lvl.StartLevelTime.IsZero() || t.Before(lvl.StartLevelTime) {
				lvl.StartLevelTime = t
			}
		}
		if lvl.StartLevelTime.IsZero() {
			if et := earliestCombatEventTime(combatLevel, logTime); !et.IsZero() {
				lvl.StartLevelTime = et
			}
		}
	}
	// День из заголовка лога — лучше, чем нулевое время (Unix -62135596800000 в JSON).
	if lvl.StartLevelTime.IsZero() && !logTime.IsZero() {
		lvl.StartLevelTime = time.Date(
			logTime.Year(), logTime.Month(), logTime.Day(),
			0, 0, 0, 0, logTime.Location(),
		)
	}
	lvl.EndLevelTime = gameLevel.FinishGameplay.GetTime(logTime)
	if cmbt := combatLevel.Finished; lvl.EndLevelTime.Before(cmbt.GetTime(logTime)) {
		lvl.EndLevelTime = cmbt.GetTime(logTime)
	}
	// Незавершённый матч: нет finish в game/combat — иначе EndLevelTime=0, span=0 и
	// timeInRangeFromStart(..., lo=0, hi=0) отбрасывает весь урон/отхил/киллы.
	// Не подмешиваем latestCombatEventTime, если конец уже известен: после combat.Finished
	// в тот же уровень ещё долго сыплются Spell/Heal/Damage до следующего Start — иначе
	// таймлайн и графики растягиваются на минуты «пустого» хвоста с последним счётом живых.
	if lvl.EndLevelTime.IsZero() {
		if lt := latestCombatEventTime(combatLevel, logTime); !lt.IsZero() {
			lvl.EndLevelTime = lt
		}
	}
	if !lvl.StartLevelTime.IsZero() && (lvl.EndLevelTime.IsZero() || lvl.EndLevelTime.Before(lvl.StartLevelTime)) {
		lvl.EndLevelTime = lvl.StartLevelTime
	}

	playerTeamsMap := make(map[int]map[int]Player)
	for _, logPlayer := range gameLevel.AddPlayer {
		player := Player{
			PlayerID:        logPlayer.PlayerID,
			SessionPlayerID: logPlayer.InGamePlayerID,
			Name:            logPlayer.Name,
			CorpTag:         logPlayer.ClanTag,
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
	s.lg.For(ctx).Debugw("level", "level", lvl.CombatLog.Start, "team_ids", teams)

	lvl.Teams = make(map[int][]Player, len(playerTeamsMap))
	for teamID, team := range playerTeamsMap {
		players := maps.Values(team)
		playerNames := make([]string, 0, len(players))
		for _, p := range players {
			playerNames = append(playerNames, p.String())
		}
		slices.Sort(playerNames)
		if os.Getenv("SHOW_WATCHERS") != "" && teamID == 0 {
			s.lg.For(ctx).Debugw("team players", "team_id", teamID, "players", playerNames)
		}
		lvl.Teams[teamID] = players
	}

	if len(lvl.Teams) == 0 {
		return nil, fmt.Errorf("no teams found")
	}

	return lvl, nil
}

func earliestCombatEventTime(level *CombatLogLevel, logTime time.Time) time.Time {
	if level == nil {
		return time.Time{}
	}
	best := time.Time{}
	setMin := func(t time.Time) {
		if t.IsZero() {
			return
		}
		if best.IsZero() || t.Before(best) {
			best = t
		}
	}
	for _, x := range level.Damage {
		if x != nil {
			setMin(x.GetTime(logTime))
		}
	}
	for _, x := range level.Heal {
		if x != nil {
			setMin(x.GetTime(logTime))
		}
	}
	for _, x := range level.Kill {
		if x != nil {
			setMin(x.GetTime(logTime))
		}
	}
	for _, x := range level.Spawn {
		if x != nil {
			setMin(x.GetTime(logTime))
		}
	}
	for _, x := range level.Reward {
		if x != nil {
			setMin(x.GetTime(logTime))
		}
	}
	for _, x := range level.Spell {
		if x != nil {
			setMin(x.GetTime(logTime))
		}
	}
	return best
}

func latestCombatEventTime(level *CombatLogLevel, logTime time.Time) time.Time {
	if level == nil {
		return time.Time{}
	}
	best := time.Time{}
	setMax := func(t time.Time) {
		if t.IsZero() {
			return
		}
		if best.IsZero() || t.After(best) {
			best = t
		}
	}
	for _, x := range level.Damage {
		if x != nil {
			setMax(x.GetTime(logTime))
		}
	}
	for _, x := range level.Heal {
		if x != nil {
			setMax(x.GetTime(logTime))
		}
	}
	for _, x := range level.Kill {
		if x != nil {
			setMax(x.GetTime(logTime))
		}
	}
	for _, x := range level.Spawn {
		if x != nil {
			setMax(x.GetTime(logTime))
		}
	}
	for _, x := range level.Reward {
		if x != nil {
			setMax(x.GetTime(logTime))
		}
	}
	for _, x := range level.Spell {
		if x != nil {
			setMax(x.GetTime(logTime))
		}
	}
	return best
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

	StartGameplay  *game.ClientConnected
	AddPlayer      []*game.ClientAddPlayer
	LeavePlayer    []*game.ClientPlayerLeave
	FinishGameplay *game.ClientConnectionClosed
}

func (g *GameLogLevel) IsEmpty() bool {
	return g == nil || (g.StartGameplay == nil && g.FinishGameplay == nil && len(g.AddPlayer) == 0 && len(g.LeavePlayer) == 0)
}

type CombatLogLevel struct {
	logTime time.Time
	// LogLines — распарсенные записи combat.log в порядке файла (только успешно распознанные строки).
	LogLines []combat.LogLine

	Connect  combat.ConnectToGameSession
	Start    combat.Start
	Damage   []*combat.Damage
	Heal     []*combat.Heal
	Kill     []*combat.Kill
	Spawn    []*combat.Spawn
	Reward   []*combat.Reward
	Spell    []*combat.Spell
	Finished combat.Finished
}

func (g *CombatLogLevel) String() string {
	if len(g.LogLines) == 0 {
		return "empty game level"
	}

	var sb strings.Builder

	sb.WriteString("game level: ")
	fmt.Fprintf(&sb, "time: %s ", g.LogLines[0].GetTime(g.logTime))
	fmt.Fprintf(&sb, "lines: %d ", len(g.LogLines))
	fmt.Fprintf(&sb, "connect(session_id): %d ", g.Connect.SessionID)
	fmt.Fprintf(&sb, "game_mode: %s map_name: %s start_time: %s", g.Start.GameMode, g.Start.MapName, g.Start.GetTime(g.logTime))
	fmt.Fprintf(&sb, "finished: %s reason: %s game_time: %f", g.Finished.FinishReason, g.Finished.WinReason, g.Finished.GameTime)

	return sb.String()
}

func (g *CombatLogLevel) IsEmpty() bool {
	return g == nil || (g.Connect.IsEmpty() && g.Start.IsEmpty() && g.Finished.IsEmpty())
}

// LogAnchorTime is the calendar date from the combat log header; used with line clocks in GetTime.
func (g *CombatLogLevel) LogAnchorTime() time.Time {
	if g == nil {
		return time.Time{}
	}
	return g.logTime
}

const (
	ConnectionClosedReasonGameFinished          = "DR_CLIENT_GAME_FINISHED"
	ConnectionClosedReasonClientCouldNotConnect = "DR_CLIENT_COULD_NOT_CONNECT"
	ConnectionClosedReasonQuit                  = "DR_CLIENT_QUIT"
	ConnectionClosedReasonServerTransfer        = "DR_CLIENT_SERVER_TRANSFER"
	ConnectionClosedReasonDockSpaceStation      = "DR_CLIENT_DOCK_SPACE_STATION"
	ConnectionClosedReasonReturnSpaceStation    = "DR_CLIENT_RETURN_SPACE_STATION"
)

func (s *Splitter) collectGameLogLevels(ctx context.Context, lines []parser.LogLine[game.LogLine]) ([]*GameLogLevel, []lineSpan) {
	ctx, span := s.tr.Start(ctx, "collectGameLogLevels")
	defer span.End()

	var errs []error
	var out []*GameLogLevel
	var outSpans []lineSpan
	currLevel := new(GameLogLevel)
	spanStart, spanEnd := -1, -1

	touch := func(ll parser.LogLine[game.LogLine]) {
		if ll.Data == nil {
			return
		}
		if spanStart < 0 {
			spanStart = ll.ByteOffset
		}
		spanEnd = ll.ByteEnd
	}

	pushLevel := func() {
		var sp lineSpan
		if spanStart >= 0 {
			sp = lineSpan{Start: spanStart, End: spanEnd}
		}
		out = append(out, currLevel)
		outSpans = append(outSpans, sp)
		currLevel = new(GameLogLevel)
		errs = nil
		spanStart, spanEnd = -1, -1
	}

	for _, logLine := range lines {
		var cutErr *common.LineIsNotFinishedError
		if logLine.Err != nil && !errors.As(logLine.Err, &cutErr) {
			errs = append(errs, logLine.Err)
		}
		if logLine.Data == nil {
			continue
		}

		switch line := logLine.Data.(type) {
		case *game.ClientConnected:
			if currLevel.StartGameplay != nil {
				s.lg.For(ctx).Warnw("start gameplay twice", "prev", currLevel.StartGameplay, "next", line,
					"start", currLevel.StartGameplay, "end", currLevel.FinishGameplay)
				pushLevel()
			}
			touch(logLine)
			currLevel.Lines = append(currLevel.Lines, logLine.Data)
			currLevel.StartGameplay = line
		case *game.ClientAddPlayer:
			touch(logLine)
			currLevel.Lines = append(currLevel.Lines, logLine.Data)
			currLevel.AddPlayer = append(currLevel.AddPlayer, line)
		case *game.ClientConnectionClosed:
			touch(logLine)
			currLevel.Lines = append(currLevel.Lines, logLine.Data)
			currLevel.FinishGameplay = line
			if line.Reason == ConnectionClosedReasonClientCouldNotConnect {
				s.lg.For(ctx).Infow("detected could not connect log", "line", line)
				currLevel = new(GameLogLevel)
				spanStart, spanEnd = -1, -1
				continue
			}
			pushLevel()
		case *game.ClientPlayerLeave:
			touch(logLine)
			currLevel.Lines = append(currLevel.Lines, logLine.Data)
			currLevel.LeavePlayer = append(currLevel.LeavePlayer, line)
		}
	}

	if !currLevel.IsEmpty() {
		pushLevel()
	}
	return out, outSpans
}

func (s *Splitter) collectCombatLogLevels(ctx context.Context, logTime time.Time, lines []parser.LogLine[combat.LogLine]) (res []*CombatLogLevel, spans []lineSpan, errs []error) {
	ctx, span := s.tr.Start(ctx, "collectCombatLogLevels")
	defer span.End()

	newLevel := func() *CombatLogLevel {
		l := new(CombatLogLevel)
		l.logTime = logTime
		return l
	}
	currLevel := newLevel()
	spanStart, spanEnd := -1, -1

	touch := func(ll parser.LogLine[combat.LogLine]) {
		if ll.Data == nil {
			return
		}
		if spanStart < 0 {
			spanStart = ll.ByteOffset
		}
		spanEnd = ll.ByteEnd
	}

	seal := func() {
		var sp lineSpan
		if spanStart >= 0 {
			sp = lineSpan{Start: spanStart, End: spanEnd}
		}
		res = append(res, currLevel)
		spans = append(spans, sp)
		currLevel = newLevel()
		spanStart, spanEnd = -1, -1
	}

	for _, logLine := range lines {
		var cutErr *common.LineIsNotFinishedError
		if logLine.Err != nil && !errors.As(logLine.Err, &cutErr) {
			errs = append(errs, logLine.Err)
		}
		if logLine.Data == nil {
			continue
		}
		switch line := logLine.Data.(type) {
		case *combat.ConnectToGameSession:
			if !currLevel.Connect.IsEmpty() && currLevel.Connect.SessionID != line.SessionID || !currLevel.Start.IsEmpty() {
				seal()
			}
			touch(logLine)
			currLevel.LogLines = append(currLevel.LogLines, logLine.Data)
			currLevel.Connect = *line
		case *combat.Start:
			if !currLevel.Start.IsEmpty() {
				seal()
			}
			touch(logLine)
			currLevel.LogLines = append(currLevel.LogLines, logLine.Data)
			currLevel.Start = *line
		case *combat.Finished:
			s.lg.For(ctx).Debugw("finished", "line", line, "start", currLevel.Start, "connect", currLevel.Connect)
			if !currLevel.Finished.IsEmpty() {
				seal()
			}
			touch(logLine)
			currLevel.LogLines = append(currLevel.LogLines, logLine.Data)
			currLevel.Finished = *line
		case *combat.Damage:
			touch(logLine)
			currLevel.LogLines = append(currLevel.LogLines, logLine.Data)
			currLevel.Damage = append(currLevel.Damage, line)
		case *combat.Heal:
			touch(logLine)
			currLevel.LogLines = append(currLevel.LogLines, logLine.Data)
			currLevel.Heal = append(currLevel.Heal, line)
		case *combat.Kill:
			touch(logLine)
			currLevel.LogLines = append(currLevel.LogLines, logLine.Data)
			currLevel.Kill = append(currLevel.Kill, line)
		case *combat.Spawn:
			touch(logLine)
			currLevel.LogLines = append(currLevel.LogLines, logLine.Data)
			currLevel.Spawn = append(currLevel.Spawn, line)
		case *combat.Reward:
			touch(logLine)
			currLevel.LogLines = append(currLevel.LogLines, logLine.Data)
			currLevel.Reward = append(currLevel.Reward, line)
		case *combat.Spell:
			touch(logLine)
			currLevel.LogLines = append(currLevel.LogLines, logLine.Data)
			currLevel.Spell = append(currLevel.Spell, line)
		}
	}

	if !currLevel.IsEmpty() {
		s.lg.For(ctx).Debugw("level", "level", currLevel.String())
		var sp lineSpan
		if spanStart >= 0 {
			sp = lineSpan{Start: spanStart, End: spanEnd}
		}
		res = append(res, currLevel)
		spans = append(spans, sp)
	}

	s.lg.For(ctx).Infow("got combat log levels", "count", len(res))
	return res, spans, errs
}
