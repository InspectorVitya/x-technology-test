package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	PortHTTP string `env:"PORT_HTTP" env-default:"8080"`
	DbURL    string `env:"DB_URL"`
}

func New() (Config, error) {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, err
}
