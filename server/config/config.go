package config

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

// Config содержит конфигурацию сервера.
type Config struct {
	RunAddress  string `env:"RUN_ADDRESS"`
	DatabaseURI string `env:"DATABASE_URI"`
	LoggerDebug bool   `env:"LOGGER_DEBUG"`
}

// GetConfig возвращает новый экземпляр конфигурации.
func GetConfig() Config {
	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		newPrint := fmt.Sprintf("unable to parse environment values for config : %s", err.Error())
		log.Error().Msg(newPrint)
	}

	flag.StringVar(&cfg.RunAddress, "a", cfg.RunAddress, "адрес GRPC сервера")
	flag.StringVar(&cfg.DatabaseURI, "d", cfg.DatabaseURI, "адрес БД")
	flag.BoolVar(&cfg.LoggerDebug, "l", cfg.LoggerDebug, "нужно ли использовать логгер в режиме отладки")
	flag.Parse()

	return cfg
}
