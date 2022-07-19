package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}
	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}
	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	config := &Config{}

	if err := cleanenv.ReadConfig("./config/config.yml", config); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	if err := cleanenv.ReadEnv(config); err != nil {
		return nil, err
	}

	return config, nil
}
