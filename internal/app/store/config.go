package store

type Config struct {
	dbURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
