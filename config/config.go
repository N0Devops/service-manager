package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Http    Http `json:"http" yaml:"http"`
	Account map[string]string
}

func Load() Config {
	b, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	var config Config
	if err := yaml.Unmarshal(b, &config); err != nil {
		panic(err)
	}
	return config
}
