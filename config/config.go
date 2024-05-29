package config

import (
	"flag"

	"go.uber.org/fx"
)

type Config struct {
	fx.Out

	Service *ServiceConfig
	Trace   *TraceConfig
}

type TraceConfig struct {
	ServiceName string
	Enabled     bool
	Addr        string
}

func GetConfig() Config {
	c := Config{
		Service: &ServiceConfig{},
		Trace:   &TraceConfig{},
	}

	flag.StringVar(&c.Service.InputDir, "i", "", "input directory")
	flag.StringVar(&c.Service.OutputDir, "o", "", "output directory")
	flag.StringVar(&c.Service.TextOut, "txt", "", "output to text file")
	flag.StringVar(&c.Service.Show, "show", "", "show type of logs (kill,heal,damage)")
	flag.StringVar(&c.Service.Damage, "damage", "", "custom filter")
	flag.StringVar(&c.Service.Heal, "heal", "", "custom filter")
	flag.StringVar(&c.Service.Kill, "kill", "", "custom filter")
	flag.StringVar(&c.Trace.ServiceName, "service", "lux-panel", "service name")
	flag.BoolVar(&c.Trace.Enabled, "trace", false, "enable tracing")
	flag.StringVar(&c.Trace.Addr, "trace_addr", "", "jaeger addr")
	flag.Parse()

	return c
}
