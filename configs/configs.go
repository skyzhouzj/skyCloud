package configs

import (
	cfg "github.com/skyzhouzj/skyCloud/pkg/config"
)

var config = new(Config)

type Config struct {
	SkyCloud   cfg.SkyCloud   `mapstructure:"SkyCloud"`
	DataSource cfg.DataSource `mapstructure:"DataSource"`
	Redis      cfg.Redis      `mapstructure:"Redis"`
}

func init() {
	cfg.Init(config, "../", "dev")
}

func Get() Config {
	return *config
}
