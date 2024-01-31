package config

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v10"
	_ "github.com/joho/godotenv/autoload"
)

type appEnv string
type logType string

var (
	EnvProd appEnv = "prod"
	EnvDev  appEnv = "dev"

	AppLog logType = "app"
	GinLog logType = "gin"

	applogName string = "app.log"
	ginlogName string = "gin.log"
)

type Config struct {
	Port    int    `env:"PORT" envDefault:"3056"`
	Dir     string `env:"DATA_DIRNAME" envDefault:"data"`
	Env     appEnv `env:"ENV" envDefault:"dev"`
	Address string `env:"ADDRESS,expand" envDefault:"localhost:${PORT}"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) CheckEnv() (string, error) {
	var datapath string
	if cfg.Env == EnvProd {
		exe, err := os.Executable()
		if err != nil {
			return "", err
		}
		datapath = filepath.Join(filepath.Dir(exe), cfg.Dir)
	} else if cfg.Env == EnvDev {
		datapath = filepath.Join(".", cfg.Dir)
	} else {
		return "", fmt.Errorf("unexpected env mode")
	}
	return datapath, nil
}

func (cfg *Config) CreateAppLog(datapath string) (map[string]*slog.Logger, func(), error) {
	applog, err := os.Create(filepath.Join(datapath, applogName))

	if err != nil {
		return nil, func() {}, err
	}

	ginlog, err := os.Create(filepath.Join(datapath, ginlogName))
	if err != nil {
		return nil, func() {}, err
	}
	defer ginlog.Close()
	appLogger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stderr, applog), nil))
	ginLogger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stderr, ginlog), nil))

	logs := make(map[string]*slog.Logger)
	logs[string(AppLog)] = appLogger
	logs[string(GinLog)] = ginLogger

	return logs, func() { _, _ = ginlog.Close(), applog.Close() }, nil
}