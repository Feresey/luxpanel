package config

import "flag"

type Config struct {
	InputDir  string
	OutputDir string

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
	flag.StringVar(&c.Trace.ServiceName, "service", "sclogparser", "service name")
	flag.BoolVar(&c.Trace.Enabled, "trace_enabled", true, "enable tracing")
	flag.Parse()

	return &c
}
