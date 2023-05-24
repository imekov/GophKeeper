package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

// Config содержит данные адреса сервера и имени файла для сохранения.
type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	Filename      string `env:"FILENAME"`
}

// GetConfig является конструктором конфига
func GetConfig() Config {
	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&cfg.ServerAddress, "s", cfg.ServerAddress, "адрес GRPC сервера")
	flag.StringVar(&cfg.Filename, "f", cfg.Filename, "имя файла для сохранения данных")
	flag.Parse()

	return cfg
}
