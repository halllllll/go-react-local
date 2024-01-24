package config

import "github.com/caarlos0/env/v10"


type Config struct{
	Port int `env:"PORT" env_default:"3056"`
	Dir string `env:"DATA_DIR" env_default:"data"`
}

func New()(*Config, error){
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil{
		return nil, err
	}
	return cfg, nil
}