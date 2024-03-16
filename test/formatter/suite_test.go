package formatter_test

import (
	"testing"

	"github.com/Feresey/sclogparser/internal/formatter"
	"github.com/Feresey/sclogparser/internal/logger"
	"github.com/stretchr/testify/suite"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
	"go.uber.org/zap"
)

type Suite struct {
	suite.Suite

	lf logger.Factory
	tp trace.TracerProvider

	formatter *formatter.Formatter
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

func (s *Suite) SetupSuite() {
	r := s.Require()

	lc := zap.NewDevelopmentConfig()
	lf, close, err := logger.NewFactory(lc)
	r.NoError(err)
	defer close()

	s.lf = lf
	s.tp = noop.NewTracerProvider()

	s.formatter = formatter.NewFormatter(lf, s.tp.Tracer("test"))
}
