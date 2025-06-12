package config

func Default() *Config {
	return &Config{
		APIKeys:                    []*ApiKey{},
		EnableUploadAuthentication: false,
		DBConnectionString:         DefaultDBConnectionString,
	}
}
