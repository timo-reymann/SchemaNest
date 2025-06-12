package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

func ParseFromToml(tomlData string) (*Config, error) {
	config := &Config{}
	_, err := toml.Decode(tomlData, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func ParseFromFile(configFile string, parser func(data string) (*Config, error)) (*Config, error) {
	if configFile == "" {
		return Default(), nil
	}

	cfgContent, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err)
	}

	cfg, err := parser(string(cfgContent))
	if err != nil {
		return nil, fmt.Errorf("failed to parse toml: %s", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate config: %s", err)
	}

	return cfg, nil
}
