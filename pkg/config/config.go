package config

import (
	"fmt"
)

const DefaultDBConnectionString = "sqlite3://schema_nest.sqlite"

type Config struct {
	EnableUploadAuthentication bool     `toml:"enable_upload_authentication"`
	APIKeys                    []ApiKey `toml:"api_keys"`
	DBConnectionString         string   `toml:"database_dsn"`
	keyMapping                 map[string]*ApiKey
}

func (c *Config) mapKeys() {
	c.keyMapping = make(map[string]*ApiKey, len(c.APIKeys))
	for _, ak := range c.APIKeys {
		c.keyMapping[ak.Key] = &ak
	}
}

func (c *Config) Validate() error {
	if c.DBConnectionString == "" {
		c.DBConnectionString = DefaultDBConnectionString
	}

	if !c.EnableUploadAuthentication {
		return nil
	}

	if len(c.APIKeys) == 0 {
		return fmt.Errorf("at least one API key must be configured when upload authentication is enabled")
	}
	for idx, key := range c.APIKeys {
		if err := key.Validate(); err != nil {
			identifier := key.Identifier
			if identifier == "" {
				identifier = fmt.Sprintf("#%d", idx)
			}

			return fmt.Errorf("invalid API key configuration for key %s: %w", identifier, err)
		}
	}
	c.mapKeys()

	return nil
}

func (c *Config) LookupApiKey(key string) (*ApiKey, bool) {
	apiKey, ok := c.keyMapping[key]
	return apiKey, ok
}
