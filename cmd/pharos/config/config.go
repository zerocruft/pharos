package config

type Config struct {
	Source string `toml:"source"`
	Port   string `toml:"port"`
}
