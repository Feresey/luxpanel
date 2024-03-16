package config

import "flag"

type Config struct {
	InputDir  string
	OutputDir string
}

func GetConfig() *Config {
	var c Config

	flag.StringVar(&c.InputDir, "i", "", "input directory")
	flag.StringVar(&c.OutputDir, "o", "", "output directory")
	flag.Parse()

	return &c
}
