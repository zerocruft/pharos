package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func New(configFile string) *Config {
	var cfg Config
	_, err := toml.DecodeFile(configFile, &cfg)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &cfg
}

type Config struct {
	Sources  []Source `toml:"source"`
	Port     string   `toml:"port"`
	BuildDir string   `toml:"build-dir"`
}

type Source struct {
	Name   string `toml:"name"`
	GitURL string `toml:"url"`
}
