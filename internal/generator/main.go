package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	g, err := NewGenerator(&Config{
		OutputDir:     "../parser/combat",
		OutputPackage: "combat",
	})
	if err != nil {
		return fmt.Errorf("create generator: %w", err)
	}

	if err := g.Generate(combatRaw); err != nil {
		return fmt.Errorf("generate combat: %w", err)
	}

	return nil
}
