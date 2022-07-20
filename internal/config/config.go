package config

import (
	"git.elewise.com/elma365/common/pkg/config"

	"github.com/vporoshok/envcfg"
)

// Config of application
type Config struct {
	config.Config

	// TODO: добавьте свои параметры 
	// MyParameter string
}

// New создать новый конфиг
func New(name string) (Config, error) {
	ccfg, err := config.New(name)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{Config: ccfg}
	err = envcfg.Read(&cfg, envcfg.WithPrefix(config.Prefix), envcfg.WithDefault(nil))

	return cfg, err
}
