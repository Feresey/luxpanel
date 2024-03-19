package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"go.opentelemetry.io/otel/trace"

	"github.com/Feresey/luxpanel/cmd/luxpanel/config"
	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/Feresey/luxpanel/internal/splitter"
)

type Service struct {
	cfg      *config.Config
	lg       logger.Factory
	tr       trace.Tracer
	splitter *splitter.Splitter
	parser   *parser.Parser
}

func NewService(
	cfg *config.Config,
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

	levels, err := s.parseLevels(ctx)
	if err != nil {
		return fmt.Errorf("parseLevels: %w", err)
	}

	if s.cfg.OutputDir != "" {
		if err := s.saveToDir(ctx, levels); err != nil {
			return fmt.Errorf("saveToDir: %w", err)
		}
	}

	for _, level := range levels {
		s.showLevelStatistics(ctx, level)
		fmt.Println("========")
	}

	return nil
}

func (s *Service) showLevelStatistics(ctx context.Context, lvl *splitter.Level) {
	ctx, span := s.tr.Start(ctx, "showLevelStatistics")
	defer span.End()

	filters := s.getPlayersFilters(ctx, lvl)
	for _, filter := range filters {
		res := ProcessArray(lvl.CombatLog.Damage, FilterPlayerDamage(filter), Sum)
		if res > 0 {
			fmt.Printf("%s: %f\n", filter.String(), res)
		}
	}
}

func makeDamageFilters(filter *PlayerDamageFilterConfig) []*PlayerDamageFilterConfig {
	copyFilter := func(modifiers DamageModifiersMap) *PlayerDamageFilterConfig {
		return &PlayerDamageFilterConfig{
			InitiatorName:   filter.InitiatorName,
			RecipientName:   filter.RecipientName,
			RecipientObject: filter.RecipientObject,
			DamageType:      filter.DamageType,
			DamageModifiers: modifiers,
		}
	}
	return []*PlayerDamageFilterConfig{
		copyFilter(DamageModifiersMap{}),
		copyFilter(DamageModifiersMap{
			parser.DamageCrit: true,
		}),
		copyFilter(DamageModifiersMap{
			parser.DamageExplosion: true,
		}),
		copyFilter(DamageModifiersMap{
			parser.DamageWeaponPrimary:   false,
			parser.DamageWeaponSecondary: false,
		}),
		copyFilter(DamageModifiersMap{
			parser.DamageIgoreShield: true,
		}),
		copyFilter(DamageModifiersMap{
			parser.DamageCollision: true,
		}),
		copyFilter(DamageModifiersMap{
			parser.DamageTypeEMP: true,
		}),
		copyFilter(DamageModifiersMap{
			parser.DamageTypeKinetic: true,
		}),
		copyFilter(DamageModifiersMap{
			parser.DamageTypeThermal: true,
		}),
		copyFilter(DamageModifiersMap{
			parser.DamageTypeTrueDamage: true,
		}),
	}
}

func (s *Service) getPlayersFilters(ctx context.Context, lvl *splitter.Level) (res []*PlayerDamageFilterConfig) {
	ctx, span := s.tr.Start(ctx, "getPlayersFilters")
	defer span.End()

	for _, players := range lvl.Teams {
		for _, player := range players {
			var enemies []splitter.Player
			for otherTeamID, enemyPlayers := range lvl.Teams {
				if otherTeamID != player.TeamID {
					enemies = append(enemies, enemyPlayers...)
				}
			}

			res = append(res, makeDamageFilters(&PlayerDamageFilterConfig{
				InitiatorName: player.Name,
			})...)
			res = append(res, makeDamageFilters(&PlayerDamageFilterConfig{
				InitiatorName: player.Name,
				DamageType:    DamageTypeShield,
			})...)
			res = append(res, makeDamageFilters(&PlayerDamageFilterConfig{
				InitiatorName: player.Name,
				DamageType:    DamageTypeHull,
			})...)
			for _, enemy := range enemies {
				res = append(res, makeDamageFilters(&PlayerDamageFilterConfig{
					InitiatorName: player.Name,
					RecipientName: enemy.Name,
				})...)
				res = append(res, makeDamageFilters(&PlayerDamageFilterConfig{
					InitiatorName: player.Name,
					RecipientName: enemy.Name,
					DamageType:    DamageTypeShield,
				})...)
				res = append(res, makeDamageFilters(&PlayerDamageFilterConfig{
					InitiatorName: player.Name,
					RecipientName: enemy.Name,
					DamageType:    DamageTypeHull,
				})...)
			}
		}
	}

	s.lg.For(ctx).Debugw("created filters for players damage", "count", len(res))
	return res
}

func (s *Service) parseLevels(ctx context.Context) ([]*splitter.Level, error) {
	ctx, span := s.tr.Start(ctx, "parseLevels")
	defer span.End()

	combatLog, err := os.Open(filepath.Join(s.cfg.InputDir, "combat.log"))
	if err != nil {
		return nil, fmt.Errorf("os.Open(combat.log): %w", err)
	}
	defer combatLog.Close()

	gameLog, err := os.Open(filepath.Join(s.cfg.InputDir, "game.log"))
	if err != nil {
		return nil, fmt.Errorf("os.Open(game.log): %w", err)
	}
	defer gameLog.Close()

	combatLogLines, err := s.parser.ParseCombatLog(ctx, combatLog)
	if err != nil {
		return nil, fmt.Errorf("parser.ParseCombatLog: %w", err)
	}
	gameLogLines, err := s.parser.ParseGameLog(ctx, gameLog)
	if err != nil {
		return nil, fmt.Errorf("parser.ParseGameLog: %w", err)
	}

	levels, err := s.splitter.SplitLevels(ctx, gameLogLines, combatLogLines)
	if err != nil {
		return nil, fmt.Errorf("splitter.SplitLevels: %w", err)
	}

	return levels, nil
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
