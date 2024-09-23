package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Env                  string `env:"APP_ENV" envDefault:"dev"`
	Port                 int    `env:"PORT" envDefault:"80"`
	DBHost               string `env:"APP_DB_HOST" envDefault:"127.0.0.1"`
	DBPort               int    `env:"APP_DB_PORT" envDefault:"13306"`
	DBUser               string `env:"APP_DB_USER" envDefault:"mysql"`
	DBPassword           string `env:"APP_DB_PASSWORD" envDefault:"password"`
	DBName               string `env:"APP_DB_NAME" envDefault:"example"`
	AllowNativePasswords bool   `env:"APP_DB_ALLOW_NATIVE_PASSWORDS" envDefault:"true"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
