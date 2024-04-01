package config

import (
	"flag"
)

type Config struct {
	InputDir  string
	OutputDir string
	TextOut   string

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
	flag.StringVar(&c.Trace.ServiceName, "service", "lux-panel", "service name")
	flag.BoolVar(&c.Trace.Enabled, "trace", false, "enable tracing")
	flag.Parse()

	return &c
}
