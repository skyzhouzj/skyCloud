package configs

import (
	"fmt"
	"testing"
)

func TestDotEnv(t *testing.T) {
	conf := config.Get()

	fmt.Println(conf)
}
