package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	RunAddress  string `env:"RUN_ADDRESS"`
	DatabaseURI string `env:"DATABASE_URI"`
}

func GetConfig() Config {
	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&cfg.RunAddress, "a", cfg.RunAddress, "HTTP server start address")
	flag.StringVar(&cfg.DatabaseURI, "d", cfg.DatabaseURI, "the base address of the resulting shortened URL")
	flag.Parse()

	return cfg
}
