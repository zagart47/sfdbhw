package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Cfg struct {
	Dsn string `yaml:"dsn"`
}

var Config Cfg

func init() {
	err := cleanenv.ReadConfig("./config/config.yaml", &Config)
	if err != nil {
		fmt.Println("Error reading Config")
	}
}
