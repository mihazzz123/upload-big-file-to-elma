package config

import (
	"fmt"
	"git.elewise.com/elma365/common/pkg/config"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/vporoshok/envcfg"
)

// Config of application
type Config struct {
	config.Config

	customConfig
}

type customConfig struct {
	ElmaAddr        string `yaml:"elma_addr"`
	DirId           string `yaml:"dir_id"`
	Token           string `yaml:"token"`
	DefaultPartSize int    `yaml:"default_part_size"`
}

// New создать новый конфиг
func New(name string) (Config, error) {
	ccfg, err := config.New(name)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{Config: ccfg}
	err = envcfg.Read(&cfg, envcfg.WithPrefix(config.Prefix), envcfg.WithDefault(nil))

	customCfg, err := newCustomConfig()
	if err != nil {
		return Config{}, err
	}
	cfg.customConfig = customCfg

	return cfg, err
}

// newCustomConfig ...
func newCustomConfig() (customConfig, error) {
	cfg := customConfig{}
	if err := cleanenv.ReadConfig("./config.yml", &cfg); err != nil {
		return customConfig{}, err
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return customConfig{}, err
	}

	return cfg, nil
}

func (c customConfig) GetElmaURL() string {
	return fmt.Sprintf("https://%s/pub/v1/disk/directory/%s/upload?hash=", c.ElmaAddr, c.DirId)
}
