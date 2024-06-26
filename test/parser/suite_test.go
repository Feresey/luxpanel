package parser_test

import (
	"context"
	"errors"
	"syscall"
	"testing"

	"github.com/Feresey/luxpanel/cmd/luxpanel/config"
	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/mytrace"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Suite struct {
	suite.Suite

	app *fxtest.App

	parser *parser.Parser
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

func (s *Suite) SetupSuite() {
	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	s.app = fxtest.New(
		s.T(),
		fx.NopLogger,
		fx.Supply(
			config.GetConfig(),
			logConfig,
		),
		fx.Provide(
			logger.NewFactory,
			mytrace.NewTraceProvider,
			parser.NewParser,
		),
		fx.Populate(
			&s.parser,
		),
	)

	s.app.RequireStart()
}

func (s *Suite) TearDownSuite() {
	if stopErr := s.app.Stop(context.Background()); stopErr != nil && !errors.Is(stopErr, syscall.ENOTTY) {
		// s.NoError(stopErr)
	}
}
