package main

// import (
// 	"context"
// 	"errors"
// 	"flag"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"time"

// 	"github.com/Feresey/sclogparser/pkg/formatter"
// 	"github.com/Feresey/sclogparser/pkg/logger"
// 	"github.com/Feresey/sclogparser/pkg/parser"
// 	"github.com/tealeg/xlsx/v3"
// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// )

func main() {
	// var (
	// 	inputDir   string
	// 	outputFile string
	// )

	// flag.StringVar(&inputDir, "i", "", "input directory")
	// flag.StringVar(&outputFile, "o", "", "output file")
	// flag.Parse()

	// if err := run(inputDir, outputFile); err != nil {
	// 	panic(err)
	// }
}

// func run(inputDir string, outputFile string) (err error) {
// 	ctx := context.TODO()
// 	_, baseDir := filepath.Split(inputDir)

// 	log, close, err := logger.NewFactory(zap.NewDevelopmentConfig())
// 	if err != nil {
// 		return fmt.Errorf("create logger: %w", err)
// 	}
// 	defer close()

// 	sessionStartTime, err := time.Parse("2006.01.02 15.04.05.000", baseDir)
// 	if err != nil {
// 		return fmt.Errorf("parse base dir time: %w", err)
// 	}

// 	combatLog, err := os.Open(filepath.Join(inputDir, "combat.log"))
// 	if err != nil {
// 		return fmt.Errorf("os.Open(combat.log): %w", err)
// 	}
// 	defer combatLog.Close()

// 	gameLog, err := os.Open(filepath.Join(inputDir, "game.log"))
// 	if err != nil {
// 		return fmt.Errorf("os.Open(game.log): %w", err)
// 	}
// 	defer gameLog.Close()

// 	out := xlsx.NewFile()

// 	p := parser.New(log, combatLog, gameLog, sessionStartTime)

// 	lc := zap.NewDevelopmentConfig()
// 	lc.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
// 	lg, err := lc.Build()
// 	if err != nil {
// 		return fmt.Errorf("log.Build: %w", err)
// 	}
// 	defer lg.Sync()

// 	var levelsParsed int

// 	for {
// 		level, err := p.ParseLevel(ctx, time.Time{})
// 		if err != nil {
// 			return fmt.Errorf("parser.ParseLevel: %w", err)
// 		}
// 		if level.Combat.Finished == nil || level.Game.Finished == nil {
// 			lg.Info("level is not finished")
// 			break
// 		}

// 		format := formatter.New(level)

// 		if err := format.Format(out); err != nil {
// 			return fmt.Errorf("formatter.Format: %w", err)
// 		}
// 		levelsParsed++
// 	}

// 	if levelsParsed == 0 {
// 		lg.Warn("no levels parsed")
// 		return nil
// 	}

// 	outFile, err := os.Create(outputFile)
// 	if err != nil {
// 		return fmt.Errorf("os.Create(%s): %w", outputFile, err)
// 	}
// 	defer func() {
// 		if closeErr := outFile.Close(); closeErr != nil {
// 			err = errors.Join(err, closeErr)
// 		}
// 	}()

// 	if err := out.Write(outFile); err != nil {
// 		return fmt.Errorf("out.Write: %w", err)
// 	}

// 	return nil
// }
