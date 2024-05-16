package config

import (
	"flag"
)

const (
	ShowDamage = "damage"
	ShowHeal   = "heal"
	ShowKill   = "kill"
)

type Config struct {
	InputDir  string
	OutputDir string
	TextOut   string
	Show      string
	Damage    string
	Heal      string
	Kill      string

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
	flag.StringVar(&c.Damage, "damage", "", "custom filter")
	flag.StringVar(&c.Heal, "heal", "", "custom filter")
	flag.StringVar(&c.Kill, "kill", "", "custom filter")
	flag.StringVar(&c.Trace.ServiceName, "service", "lux-panel", "service name")
	flag.BoolVar(&c.Trace.Enabled, "trace", false, "enable tracing")
	flag.Parse()

	return &c
}
