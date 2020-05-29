package apiserver

import "go-http-rest-api/internal/app/store"

type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `toml:"log_level"`
	Store       *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddress: "*8008",
		LogLevel:    "debug",
		Store:       store.NewConfig(),
	}
}
