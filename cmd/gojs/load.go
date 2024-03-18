//go:build js

package main

import (
	"fmt"
	"io"
	"time"

	"github.com/Feresey/luxpanel/internal/parser"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func run(combatLog, gameLog io.Reader, sessionStartTime time.Time) (err error) {
	p := parser.New(combatLog, gameLog, sessionStartTime)

	lc := zap.NewDevelopmentConfig()
	lc.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	lg, err := lc.Build()
	if err != nil {
		return fmt.Errorf("log.Build: %w", err)
	}
	defer lg.Sync()

	var levelsParsed int

	for {
		level, err := p.ParseLevel()
		if err != nil {
			return fmt.Errorf("parser.ParseLevel: %w", err)
		}
		if level.Combat.Finished == nil || level.Game.Finished == nil {
			lg.Info("level is not finished")
			break
		}

		levelsParsed++
	}

	if levelsParsed == 0 {
		lg.Warn("no levels parsed")
	}

	return nil
}
