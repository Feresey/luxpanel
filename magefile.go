//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	// mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = GenerateGramma

var grammaDirs = struct {
	grammaDirs []string
	prefix     string
	suffig     string
}{
	grammaDirs: []string{"game", "combat"},
	prefix:     "internal/parser",
	suffig:     "ext",
}

func GenerateGramma() error {
	for _, dir := range grammaDirs.grammaDirs {
		inPath := filepath.Join(grammaDirs.prefix, dir, grammaDirs.suffig)
		outPath := filepath.Join(grammaDirs.prefix, dir)

		ragelConfig := ragelConfig{
			Input:  filepath.Join(inPath, "parser.rl"),
			Output: filepath.Join(outPath, "parser.gen.go"),
			Flags:  []string{"-Z", "-G2"},
		}
		if err := runRagel(ragelConfig); err != nil {
			return fmt.Errorf("runRagel: %w", err)
		}

		yaccConfig := yaccConfig{
			Input:   filepath.Join(inPath, "lexer.y"),
			Output:  filepath.Join(outPath, "lexer.gen.go"),
			Verbose: filepath.Join(inPath, "y.output"),
			Prefix:  "Yacc",
			NoLines: true,
		}
		if err := runYacc(yaccConfig); err != nil {
			return fmt.Errorf("runYacc: %w", err)
		}
	}

	return nil
}

func ShowGramma() error {
	for _, dir := range grammaDirs.grammaDirs {
		inPath := filepath.Join(grammaDirs.prefix, dir, grammaDirs.suffig)
		outPath := inPath

		ragelConfig := ragelConfig{
			Input:  filepath.Join(inPath, "parser.rl"),
			Output: filepath.Join(outPath, "parser.dot"),
			Flags:  []string{"-Vp"},
		}
		if err := runRagel(ragelConfig); err != nil {
			return fmt.Errorf("runRagel: %w", err)
		}

		dotConfig := dotConfig{
			Input:  filepath.Join(outPath, "parser.dot"),
			Output: filepath.Join(outPath, "parser.svg"),
			Flags:  []string{"-T", "svg"},
		}
		if err := runDot(dotConfig); err != nil {
			return fmt.Errorf("runDot: %w", err)
		}
	}

	return nil
}

func GOJS() error {
	if err := sh.RunWith(map[string]string{"GOOS": "js", "GOARCH": "wasm"}, "go", "build", "-o", "mysite/code/gojs.wasm", "./cmd/gojs"); err != nil {
		return err
	}
	_ = sh.Run("mkdir", "-p", "src/dist")
	return sh.Run("cp", "mysite/code/gojs.wasm", "src/dist/gojs.wasm")
}

func Start() error {
	mg.Deps(GOJS)
	return sh.RunV("yarn", "start")
}

// Site: webpack (src/js → docs/) + .nojekyll + docs/gojs.wasm для GitHub Pages.
func Site() error {
	if _, err := os.Stat("src"); err != nil {
		return fmt.Errorf("ожидается каталог src/: %w", err)
	}
	if err := os.RemoveAll("docs"); err != nil {
		return err
	}
	// Как `SITE_OUT=docs npm run build` — npm есть везде с Node; CI после `yarn install` тоже ок.
	if err := sh.RunWith(map[string]string{"SITE_OUT": "docs"}, "npm", "run", "build"); err != nil {
		return fmt.Errorf("webpack → docs: %w", err)
	}
	if err := os.WriteFile(filepath.Join("docs", ".nojekyll"), nil, 0o644); err != nil {
		return err
	}
	return sh.RunWith(map[string]string{"GOOS": "js", "GOARCH": "wasm"}, "go", "build", "-o", "docs/gojs.wasm", "./cmd/gojs")
}

type ragelConfig struct {
	Input, Output string
	Flags         []string
}

func runRagel(rc ragelConfig) error {
	var args []string
	args = append(args, rc.Flags...)
	args = append(args, "-o", rc.Output, rc.Input)
	return sh.Run("ragel", args...)
}

type yaccConfig struct {
	Input, Output string
	Verbose       string
	NoLines       bool
	Prefix        string
}

func runYacc(yc yaccConfig) error {
	var args []string

	args = append(args, "run", "./internal/parser/yacc")

	if yc.Verbose != "" {
		args = append(args, "-v", yc.Verbose)
	}
	if yc.NoLines {
		args = append(args, "-l")
	}
	args = append(args, "-p", yc.Prefix)
	args = append(args, "-o", yc.Output, yc.Input)
	return sh.Run("go", args...)
}

type dotConfig struct {
	Input, Output string
	Flags         []string
}

func runDot(dc dotConfig) error {
	var args []string
	args = append(args, dc.Flags...)
	args = append(args, "-o", dc.Output, dc.Input)
	return sh.Run("dot", args...)
}
