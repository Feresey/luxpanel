package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed templates/combat_regexps.ymltpl
var combatRaw string

//go:embed templates/game_regexps.ymltpl
var gameRaw string

func main() {
	if err := generateByTemplate("combat", combatRaw); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if err := generateByTemplate("game", gameRaw); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func generateByTemplate(tplName string, tplRaw string) error {
	g, err := NewGenerator(&Config{
		OutputDir:     "../parser/" + tplName,
		OutputPackage: tplName,
	})
	if err != nil {
		return fmt.Errorf("create generator: %w", err)
	}

	if err := g.Generate(tplRaw); err != nil {
		return fmt.Errorf("generate %s: %w", tplName, err)
	}

	return nil
}
