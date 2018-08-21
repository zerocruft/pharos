package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
)

func New(configFile string, logger *zap.Logger) *Config {
	var cfg Config
	_, err := toml.DecodeFile(configFile, &cfg)
	if err != nil {
		logger.Error("failure to decode config file", zap.Error(err))
		return nil
	}
	if cfg.Workspace == "" {
		cwd, err := os.Getwd()
		if err != nil {
			logger.Error("os error, obtaining current working directory", zap.Error(err))
			return nil
		}
		cfg.Workspace = cwd
	}
	cfg.PathSep = string(os.PathSeparator)
	return &cfg
}

type Config struct {
	Sources   []Source `toml:"source"`
	Port      string   `toml:"port"`
	Workspace string   `toml:"workspace"`
	PathSep   string
}

type Source struct {
	Name       string `toml:"name"`
	GitURL     string `toml:"url"`
	RemoteHead string `toml:"head"`
}
