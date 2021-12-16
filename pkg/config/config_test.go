package config

import (
	"fmt"
	"github.com/skyzhouzj/skyCloud/pkg/env"
	"testing"
)

type ConfigT struct {
	Test       string     `mapstructure:"test"`
	SkyCloud   SkyCloud   `mapstructure:"SkyCloud"`
	DataSource DataSource `mapstructure:"DataSource"`
	Redis      Redis      `mapstructure:"Redis"`
}

func TestDotEnv(t *testing.T) {
	var configt = new(ConfigT)
	env.Init("dev")
	Init(configt, "../../", env.Active().Value())
	fmt.Println(configt)
}
