package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	NatsURL string `yaml:"natsUrl"`
}

func New(cfgFile string) (*Config, error) {
	data, err := os.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, nil
	}

	return &cfg, nil
}
