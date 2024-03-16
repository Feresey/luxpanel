package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"go.opentelemetry.io/otel/trace"

	"github.com/Feresey/sclogparser/cmd/sclogparser/config"
	"github.com/Feresey/sclogparser/internal/formatter"
	"github.com/Feresey/sclogparser/internal/logger"
	"github.com/Feresey/sclogparser/internal/parser"
)

type Service struct {
	cfg       *config.Config
	lg        logger.Factory
	tr        trace.Tracer
	formatter *formatter.Formatter
	parser    *parser.Parser
}

func NewService(
	cfg *config.Config,
	lg logger.Factory,
	tr trace.Tracer,
	formatter *formatter.Formatter,
	parser *parser.Parser,
) *Service {
	return &Service{
		cfg:       cfg,
		lg:        lg,
		tr:        tr,
		formatter: formatter,
		parser:    parser,
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

	return nil
}

func (s *Service) parseLevels(ctx context.Context) ([]*formatter.Level, error) {
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

	levels, err := s.formatter.SplitLevels(ctx, gameLogLines, combatLogLines)
	if err != nil {
		return nil, fmt.Errorf("formatter.SplitLevels: %w", err)
	}

	return levels, nil
}

func (s *Service) saveToDir(ctx context.Context, levels []*formatter.Level) error {
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
