package config

import (
	"github.com/caarlos0/env/v10"
)

type Config struct {
	Port    int    `env:"PORT" envDefault:"3056"`
	Dir     string `env:"DATA_DIRNAME" envDefault:"data"`
	Env     string `env:"ENV" envDefault:"dev"`
	Address string `env:"ADDRESS,expand" envDefault:"localhost:${PORT}"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}