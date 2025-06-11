package config

import (
	"testing"
)

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid config with auth disabled",
			config: Config{
				EnableUploadAuthentication: false,
				APIKeys:                    nil,
			},
			wantErr: false,
		},
		{
			name: "valid config with auth enabled",
			config: Config{
				EnableUploadAuthentication: true,
				APIKeys: []ApiKey{
					{
						Identifier: "test-key",
						Key:        "secret123",
						Patterns:   []string{"*.json"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "auth enabled but no api keys",
			config: Config{
				EnableUploadAuthentication: true,
				APIKeys:                    []ApiKey{},
			},
			wantErr: true,
			errMsg:  "at least one API key must be configured when upload authentication is enabled",
		},
		{
			name: "auth enabled with invalid api key",
			config: Config{
				EnableUploadAuthentication: true,
				APIKeys: []ApiKey{
					{
						Identifier: "",
						Key:        "secret123",
						Patterns:   []string{"*.json"},
					},
				},
			},
			wantErr: true,
			errMsg:  "invalid API key configuration for key #0: api key identifier cannot be empty",
		},
		{
			name: "auth enabled with invalid api key",
			config: Config{
				EnableUploadAuthentication: true,
				APIKeys: []ApiKey{
					{
						Identifier: "",
						Key:        "secret123",
						Patterns:   []string{"*.json"},
					},
				},
			},
			wantErr: true,
			errMsg:  "invalid API key configuration for key #0: api key identifier cannot be empty",
		},
		{
			name: "auth enabled with invalid api key and identifier",
			config: Config{
				EnableUploadAuthentication: true,
				APIKeys: []ApiKey{
					{
						Identifier: "foo",
						Key:        "",
						Patterns:   []string{"*.json"},
					},
				},
			},
			wantErr: true,
			errMsg:  "invalid API key configuration for key foo: api key cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("Config.Validate() error message = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}

func TestConfig_mapKeys(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		want   map[string]*ApiKey
	}{
		{
			name: "empty api keys",
			config: Config{
				APIKeys:    []ApiKey{},
				keyMapping: make(map[string]*ApiKey),
			},
			want: map[string]*ApiKey{},
		},
		{
			name: "single api key",
			config: Config{
				APIKeys: []ApiKey{
					{
						Identifier: "test-key",
						Key:        "secret123",
						Patterns:   []string{"*.json"},
					},
				},
				keyMapping: make(map[string]*ApiKey),
			},
			want: map[string]*ApiKey{
				"secret123": {
					Identifier: "test-key",
					Key:        "secret123",
					Patterns:   []string{"*.json"},
				},
			},
		},
		{
			name: "multiple api keys",
			config: Config{
				APIKeys: []ApiKey{
					{
						Identifier: "test-key-1",
						Key:        "secret123",
						Patterns:   []string{"*.json"},
					},
					{
						Identifier: "test-key-2",
						Key:        "secret456",
						Patterns:   []string{"*.yaml"},
					},
				},
				keyMapping: make(map[string]*ApiKey),
			},
			want: map[string]*ApiKey{
				"secret123": {
					Identifier: "test-key-1",
					Key:        "secret123",
					Patterns:   []string{"*.json"},
				},
				"secret456": {
					Identifier: "test-key-2",
					Key:        "secret456",
					Patterns:   []string{"*.yaml"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.config.mapKeys()
			if len(tt.config.keyMapping) != len(tt.want) {
				t.Errorf("Config.mapKeys() resulted in keyMapping with length %v, want %v", len(tt.config.keyMapping), len(tt.want))
				return
			}
			for k, v := range tt.want {
				mapped, exists := tt.config.keyMapping[k]
				if !exists {
					t.Errorf("Config.mapKeys() key %v not found in mapping", k)
					continue
				}
				if mapped.Identifier != v.Identifier || mapped.Key != v.Key || len(mapped.Patterns) != len(v.Patterns) {
					t.Errorf("Config.mapKeys() for key %v = %v, want %v", k, mapped, v)
				}
			}
		})
	}
}

func TestConfig_LookupApiKey(t *testing.T) {
	tests := []struct {
		name       string
		config     *Config
		lookupKey  string
		wantApiKey *ApiKey
		wantFound  bool
	}{
		{
			name: "existing key",
			config: &Config{
				EnableUploadAuthentication: true,
				APIKeys: []ApiKey{
					{
						Identifier: "test-key",
						Key:        "abc123",
						Patterns:   []string{"*.json"},
					},
				},
			},
			lookupKey: "abc123",
			wantApiKey: &ApiKey{
				Identifier: "test-key",
				Key:        "abc123",
				Patterns:   []string{"*.json"},
			},
			wantFound: true,
		},
		{
			name: "non-existent key",
			config: &Config{
				EnableUploadAuthentication: true,
				APIKeys: []ApiKey{
					{
						Identifier: "test-key",
						Key:        "abc123",
						Patterns:   []string{"*.json"},
					},
				},
			},
			lookupKey:  "xyz789",
			wantApiKey: nil,
			wantFound:  false,
		},
		{
			name: "empty config",
			config: &Config{
				EnableUploadAuthentication: false,
				APIKeys:                    []ApiKey{},
			},
			lookupKey:  "any-key",
			wantApiKey: nil,
			wantFound:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize the key mapping before looking up
			tt.config.mapKeys()

			gotApiKey, gotFound := tt.config.LookupApiKey(tt.lookupKey)

			if gotFound != tt.wantFound {
				t.Errorf("LookupApiKey() found = %v, want %v", gotFound, tt.wantFound)
			}

			if !tt.wantFound {
				if gotApiKey != nil {
					t.Errorf("LookupApiKey() returned non-nil ApiKey when not found")
				}
				return
			}

			if gotApiKey == nil {
				t.Errorf("LookupApiKey() returned nil ApiKey when found")
				return
			}

			if gotApiKey.Identifier != tt.wantApiKey.Identifier {
				t.Errorf("LookupApiKey() Identifier = %v, want %v",
					gotApiKey.Identifier, tt.wantApiKey.Identifier)
			}

			if gotApiKey.Key != tt.wantApiKey.Key {
				t.Errorf("LookupApiKey() Key = %v, want %v",
					gotApiKey.Key, tt.wantApiKey.Key)
			}

			if len(gotApiKey.Patterns) != len(tt.wantApiKey.Patterns) {
				t.Errorf("LookupApiKey() Patterns length = %v, want %v",
					len(gotApiKey.Patterns), len(tt.wantApiKey.Patterns))
				return
			}

			for i, pattern := range tt.wantApiKey.Patterns {
				if gotApiKey.Patterns[i] != pattern {
					t.Errorf("LookupApiKey() Patterns[%d] = %v, want %v",
						i, gotApiKey.Patterns[i], pattern)
				}
			}
		})
	}
}
