package formatter_test

import (
	"context"
	"errors"
	"syscall"
	"testing"

	"github.com/Feresey/sclogparser/cmd/sclogparser/config"
	"github.com/Feresey/sclogparser/internal/logger"
	"github.com/Feresey/sclogparser/internal/mytrace"
	"github.com/Feresey/sclogparser/internal/splitter"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Suite struct {
	suite.Suite

	app *fxtest.App

	splitter *splitter.Splitter
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
			fx.Annotated{
				Name:   "service",
				Target: "sclogparser",
			},
		),
		fx.Provide(
			logger.NewFactory,
			mytrace.NewTraceProvider,
			splitter.NewSplitter,
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
