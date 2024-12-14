package parser_test

import (
	"context"
	"errors"
	"syscall"
	"testing"

	"github.com/Feresey/luxpanel/config"
	"github.com/Feresey/luxpanel/internal/logger"
	"github.com/Feresey/luxpanel/internal/mytrace"
	"github.com/Feresey/luxpanel/internal/parser"
	"github.com/Feresey/luxpanel/internal/splitter"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Suite struct {
	suite.Suite

	app *fxtest.App

	cfg      config.Config
	splitter *splitter.Splitter
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

func (s *Suite) SetupSuite() {
	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	s.cfg = config.Config{}

	s.app = fxtest.New(
		s.T(),
		fx.NopLogger,
		fx.Supply(
			s.cfg,
			logConfig,
		),
		fx.Provide(
			logger.NewFactory,
			mytrace.NewTraceProvider,
			splitter.NewSplitter,
			parser.NewParser,
		),
		fx.Populate(
			&s.splitter,
		),
	)

	s.app.RequireStart()
}

func (s *Suite) TearDownSuite() {
	if stopErr := s.app.Stop(context.Background()); stopErr != nil && !errors.Is(stopErr, syscall.ENOTTY) {
		// s.NoError(stopErr)
	}
}
