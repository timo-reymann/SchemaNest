package config

import "github.com/BurntSushi/toml"

func ParseFromToml(tomlData string) (*Config, error) {
	config := &Config{}
	_, err := toml.Decode(tomlData, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
