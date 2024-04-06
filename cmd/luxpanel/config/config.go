package config

import (
	"flag"
)

const (
	ShowTypeDamage = "damage"
	ShowTypeHeal   = "heal"
	ShowTypeKill   = "kill"
)

type Config struct {
	InputDir  string
	OutputDir string
	TextOut   string
	Show      string
	Player    string

	Trace TraceConfig
}

type TraceConfig struct {
	ServiceName string
	Enabled     bool
}

func GetConfig() *Config {
	var c Config

	flag.StringVar(&c.InputDir, "i", "", "input directory")
	flag.StringVar(&c.OutputDir, "o", "", "output directory")
	flag.StringVar(&c.TextOut, "txt", "", "output to text file")
	flag.StringVar(&c.Show, "show", "", "show type of logs (kill,heal,damage)")
	flag.StringVar(&c.Player, "player", "", "show logs by player")
	flag.StringVar(&c.Trace.ServiceName, "service", "lux-panel", "service name")
	flag.BoolVar(&c.Trace.Enabled, "trace", false, "enable tracing")
	flag.Parse()

	return &c
}
