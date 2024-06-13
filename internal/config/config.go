package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	Host         string        `env:"HOST" env-required:"true"`
	Port         int           `env:"PORT" env-required:"true"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" env-required:"true"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" env-required:"true"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT" env-required:"true"`

	DBConnectionString string `env:"DB_CONNECTION_STRING" env-required:"true"`
}

func New(configPath string) (*Config, error) {
	cfg := Config{}
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
