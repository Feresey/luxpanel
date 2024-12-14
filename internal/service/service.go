package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/maps"

	"github.com/Feresey/luxpanel/config"
	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/Feresey/luxpanel/internal/parser/combat"
	"github.com/Feresey/luxpanel/internal/splitter"
)

type Service struct {
	cfg      *config.ServiceConfig
	lg       logger.Factory
	tr       trace.Tracer
	splitter *splitter.Splitter
	parser   *parser.Parser
}

func NewService(
	cfg *config.ServiceConfig,
	lg logger.Factory,
	tr trace.TracerProvider,
	splitter *splitter.Splitter,
	parser *parser.Parser,
) *Service {
	return &Service{
		cfg:      cfg,
		lg:       lg,
		tr:       tr.Tracer("service"),
		splitter: splitter,
		parser:   parser,
	}
}

func (s *Service) Run(ctx context.Context) error {
	ctx, span := s.tr.Start(ctx, "Run")
	defer span.End()

	levels, err := s.splitter.SplitLevels(ctx, os.DirFS(s.cfg.InputDir))
	if err != nil {
		return fmt.Errorf("parseLevels: %w", err)
	}

	if s.cfg.OutputDir != "" {
		if err := s.saveToDir(ctx, levels); err != nil {
			return fmt.Errorf("saveToDir: %w", err)
		}
	}

	if s.cfg.TextOut != "" {
		if err := s.writeTextStatistics(ctx, levels); err != nil {
			return fmt.Errorf("writeTextStatistics: %w", err)
		}
	}

	return nil
}

const fileModePerm = 0600

func (s *Service) writeTextStatistics(ctx context.Context, levels []*splitter.Level) (err error) {
	ctx, span := s.tr.Start(ctx, "writeTextStatistics")
	defer span.End()
	outFile, err := os.OpenFile(s.cfg.TextOut, os.O_CREATE|os.O_TRUNC|os.O_RDWR, fileModePerm)
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}
	defer func() {
		err = errors.Join(err, outFile.Close())
	}()

	s.lg.For(ctx).Infow("write to file", "name", s.cfg.TextOut)

	for _, level := range levels {
		_, err := fmt.Fprintf(outFile, "\n===START  LEVEL=== time: %s, game_mode: %s, map: %s\n", level.StartLevelTime, level.CombatLog.Start.GameMode, level.CombatLog.Start.MapName)
		if err != nil {
			return fmt.Errorf("write: %w", err)
		}
		if err := s.writeTeams(ctx, level, outFile); err != nil {
			return fmt.Errorf("writeTeams: %w", err)
		}
		if err := s.writeLevelStatistics(ctx, level, outFile); err != nil {
			return fmt.Errorf("writeLevelStatistics: %w", err)
		}
		_, err = fmt.Fprintf(outFile, "\n===FINISH LEVEL=== time: %s, finish_reason: %s, win_reason: %s, winner_team_id: %d\n",
			level.EndLevelTime,
			level.CombatLog.Finished.GetFinishReason(),
			level.CombatLog.Finished.GetWinReason(),
			level.CombatLog.Finished.GetWinnerTeamID(),
		)

		if err != nil {
			return fmt.Errorf("write: %w", err)
		}
	}

	return nil
}

func (s *Service) writeTeams(ctx context.Context, lvl *splitter.Level, w io.Writer) error {
	ctx, span := s.tr.Start(ctx, "writeTeams")
	defer span.End()

	teamIDs := maps.Keys(lvl.Teams)
	slices.Sort(teamIDs)
	if len(teamIDs) == 0 {
		return nil
	}
	if watchers, ok := lvl.Teams[0]; ok {
		if _, err := fmt.Fprintln(w, strings.Repeat(" ", len(watchers))); err != nil {
			return err
		}
		if os.Getenv("SHOW_WATCHERS") != "" {
			fmt.Fprintln(w, watchers)
		}
		teamIDs = teamIDs[1:]
	} else {
		if _, err := fmt.Fprintln(w, ""); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(w, "teams: %v\n", teamIDs); err != nil {
		return err
	}

	for _, teamID := range teamIDs {
		players := lvl.Teams[teamID]
		if len(players) == 0 {
			continue
		}
		if _, err := fmt.Fprintf(w, "team %d\n", teamID); err != nil {
			return err
		}
		for _, player := range players {
			if _, err := fmt.Fprintf(w, "team %d: %v\n", teamID, player); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Service) writeLevelStatistics(ctx context.Context, lvl *splitter.Level, w io.Writer) error {
	ctx, span := s.tr.Start(ctx, "writeLevelStatistics")
	defer span.End()

	switch s.cfg.Show {
	case config.ShowDamage:
		return s.writeDamage(ctx, lvl, w)
	case config.ShowHeal:
		return s.writeHeal(ctx, lvl, w)
	case config.ShowKill:
		return s.writeKills(ctx, lvl, w)
	}

	return nil
}

func (s *Service) writeDamage(ctx context.Context, lvl *splitter.Level, w io.Writer) error {
	filters := s.getDamageFilters(ctx, lvl)
	for _, filter := range filters {
		res := ProcessArray(lvl.CombatLog.Damage, FilterPlayerDamage(filter), SummDamageBySource)
		if res == nil || len(res.BySource) == 0 {
			continue
		}

		if len(res.BySource) > 1 {
			var total float32
			keys := maps.Keys(res.BySource)
			slices.Sort(keys)
			for _, key := range keys {
				total += res.BySource[key].Value
			}
			_, err := fmt.Fprintf(w, "%s: %f\n", filter.String(), total)
			if err != nil {
				return fmt.Errorf("write: %w", err)
			}
		}
		for source, value := range res.BySource {
			_, err := fmt.Fprintf(w, "%s: %s: %d %f\n", filter.String(), source, value.Count, value.Value)
			if err != nil {
				return fmt.Errorf("write: %w", err)
			}
		}
	}

	return nil
}

func (s *Service) writeHeal(ctx context.Context, lvl *splitter.Level, w io.Writer) error {
	filters := s.getHealFilters(ctx, lvl)
	for _, filter := range filters {
		res := ProcessArray(lvl.CombatLog.Heal, FilterPlayerHeal(filter), SummHealBySource)
		if res == nil || len(res.BySource) == 0 {
			continue
		}

		if len(res.BySource) > 1 {
			var total float32
			keys := maps.Keys(res.BySource)
			slices.Sort(keys)
			for _, key := range keys {
				total += res.BySource[key].Value
			}
			_, err := fmt.Fprintf(w, "%s: %f\n", filter.String(), total)
			if err != nil {
				return fmt.Errorf("write: %w", err)
			}
		}
		for source, value := range res.BySource {
			_, err := fmt.Fprintf(w, "%s: %s: %d %f\n", filter.String(), source, value.Count, value.Value)
			if err != nil {
				return fmt.Errorf("write: %w", err)
			}
		}
	}

	return nil
}

func (s *Service) writeKills(ctx context.Context, lvl *splitter.Level, w io.Writer) error {
	filters := s.getKillsFilters(ctx, lvl)
	for _, filter := range filters {
		res := ProcessArray(lvl.CombatLog.Kill, FilterPlayerKills(filter), SummKillsBySource)
		if res == nil || len(res.BySource) == 0 {
			continue
		}

		if len(res.BySource) > 1 {
			var total int
			keys := maps.Keys(res.BySource)
			slices.Sort(keys)
			for _, key := range keys {
				total += res.BySource[key].Count
			}
			_, err := fmt.Fprintf(w, "%s: %d\n", filter.String(), total)
			if err != nil {
				return fmt.Errorf("write: %w", err)
			}
		}
		for source, value := range res.BySource {
			_, err := fmt.Fprintf(w, "%s: %s: %d\n", filter.String(), source, value.Count)
			if err != nil {
				return fmt.Errorf("write: %w", err)
			}
		}
	}

	return nil
}

func makeDamageFilters(filter *PlayerDamageFilterConfig) []*PlayerDamageFilterConfig {
	copyFilter := func(modifiers DamageModifiersMap) *PlayerDamageFilterConfig {
		return &PlayerDamageFilterConfig{
			InitiatorName:   filter.InitiatorName,
			RecipientName:   filter.RecipientName,
			DamageToObject:  filter.DamageToObject,
			DamageType:      filter.DamageType,
			DamageModifiers: modifiers,
			FriendlyFire:    filter.FriendlyFire,
		}
	}
	return []*PlayerDamageFilterConfig{
		copyFilter(DamageModifiersMap{}),
		copyFilter(DamageModifiersMap{
			combat.DamageCrit: true,
		}),
		copyFilter(DamageModifiersMap{
			combat.DamageExplosion: true,
		}),
		copyFilter(DamageModifiersMap{
			combat.DamageTypeEMP: true,
		}),
		copyFilter(DamageModifiersMap{
			combat.DamageTypeKinetic: true,
		}),
		copyFilter(DamageModifiersMap{
			combat.DamageTypeThermal: true,
		}),
		copyFilter(DamageModifiersMap{
			combat.DamageWeaponPrimary: true,
		}),
		copyFilter(DamageModifiersMap{
			combat.DamageWeaponPrimary:   false,
			combat.DamageWeaponSecondary: false,
			combat.DamageCollision:       false,
			combat.DamageCrit:            false,
			combat.DamageIgoreShield:     false,
		}),
		copyFilter(DamageModifiersMap{
			combat.DamageIgoreShield: true,
		}),
		copyFilter(DamageModifiersMap{
			combat.DamageCollision: true,
		}),
		copyFilter(DamageModifiersMap{
			combat.DamageModule: true,
		}),
	}
}

func (s *Service) getDamageFilters(ctx context.Context, lvl *splitter.Level) (res []*PlayerDamageFilterConfig) {
	ctx, span := s.tr.Start(ctx, "getDamageFilters")
	defer span.End()

	if s.cfg.Damage != "" {
		var filters []*PlayerDamageFilterConfig
		if err := json.Unmarshal([]byte(s.cfg.Damage), &filters); err != nil {
			s.lg.For(ctx).Errorw("json.Unmarshal", "err", err)
			return nil
		}

		return filters
	}

	for _, players := range lvl.Teams {
		for _, player := range players {
			var enemies []splitter.Player
			for _, enemyPlayers := range lvl.Teams {
				enemies = append(enemies, enemyPlayers...)
			}

			filters := []*PlayerDamageFilterConfig{
				{
					InitiatorName: player.Name,
				},
				{
					InitiatorName: player.Name,
					FriendlyFire:  true,
				},
			}

			for _, enemy := range enemies {
				filters = append(filters,
					&PlayerDamageFilterConfig{
						InitiatorName: player.Name,
						RecipientName: enemy.Name,
					},
					&PlayerDamageFilterConfig{
						InitiatorName: player.Name,
						RecipientName: enemy.Name,
						FriendlyFire:  true,
					})
			}

			for _, filter := range filters {
				res = append(res, makeDamageFilters(filter)...)
			}
		}
	}

	s.lg.For(ctx).Debugw("created filters for players damage", "count", len(res))
	return res
}

func (s *Service) getHealFilters(ctx context.Context, lvl *splitter.Level) (res []*PlayerHealFilterConfig) {
	ctx, span := s.tr.Start(ctx, "getHealFilters")
	defer span.End()

	if s.cfg.Heal != "" {
		var filters []*PlayerHealFilterConfig
		if err := json.Unmarshal([]byte(s.cfg.Heal), &filters); err != nil {
			s.lg.For(ctx).Errorw("json.Unmarshal", "err", err)
			return nil
		}

		return filters
	}

	for _, players := range lvl.Teams {
		for _, player := range players {
			var enemies []splitter.Player
			for _, enemyPlayers := range lvl.Teams {
				enemies = append(enemies, enemyPlayers...)
			}

			res = append(res, &PlayerHealFilterConfig{InitiatorName: player.Name})
			for _, enemy := range enemies {
				res = append(res, &PlayerHealFilterConfig{
					InitiatorName: player.Name,
					RecipientName: enemy.Name,
				})
			}
		}
	}

	s.lg.For(ctx).Debugw("created filters for players heal", "count", len(res))
	return res
}

func (s *Service) getKillsFilters(ctx context.Context, lvl *splitter.Level) (res []*PlayerKillsFilterConfig) {
	ctx, span := s.tr.Start(ctx, "getKillsFilters")
	defer span.End()

	if s.cfg.Kill != "" {
		var filters []*PlayerKillsFilterConfig
		if err := json.Unmarshal([]byte(s.cfg.Kill), &filters); err != nil {
			s.lg.For(ctx).Errorw("json.Unmarshal", "err", err)
			return nil
		}

		return filters
	}

	for _, players := range lvl.Teams {
		for _, player := range players {
			var enemies []splitter.Player
			for _, enemyPlayers := range lvl.Teams {
				enemies = append(enemies, enemyPlayers...)
			}

			res = append(res, []*PlayerKillsFilterConfig{
				{InitiatorName: player.Name},
				{InitiatorName: player.Name, FriendlyFire: true},
			}...)
			for _, enemy := range enemies {
				res = append(res, []*PlayerKillsFilterConfig{
					{
						InitiatorName: player.Name,
						RecipientName: enemy.Name,
					},
					{
						InitiatorName: player.Name,
						RecipientName: enemy.Name,
						FriendlyFire:  true,
					},
				}...)
			}
		}
	}

	s.lg.For(ctx).Debugw("created filters for players Kills", "count", len(res))
	return res
}

func (s *Service) saveToDir(ctx context.Context, levels []*splitter.Level) error {
	ctx, span := s.tr.Start(ctx, "saveToDir")
	defer span.End()

	for _, level := range levels {
		fileName := filepath.Join(s.cfg.OutputDir, level.StartLevelTime.Format("2006-01-02_15-04-05")+".json")

		data, err := json.Marshal(level)
		if err != nil {
			return fmt.Errorf("json.Marshal: %w", err)
		}

		const fileMode = 0600

		s.lg.For(ctx).Infow("write to", "file", fileName)
		if err := os.WriteFile(fileName, data, fileMode); err != nil {
			return fmt.Errorf("os.WriteFile: %w", err)
		}
	}

	return nil
}
