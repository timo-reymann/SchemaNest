package config

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestParseFromToml(t *testing.T) {
	tests := []struct {
		name     string
		tomlData string
		want     *Config
		wantErr  bool
	}{
		{
			name: "valid config with authentication enabled and API keys",
			tomlData: `
enable_upload_authentication = true

[[api_keys]]
identifier = "frontend"
key = "abc123"
patterns = ["*"]

[[api_keys]]
identifier = "backend"
key = "xyz789"
patterns = ["@deepl/*"]`,
			want: &Config{
				EnableUploadAuthentication: true,
				APIKeys: []*ApiKey{
					{
						Identifier: "frontend",
						Key:        "abc123",
						Patterns:   []string{"*"},
					},
					{
						Identifier: "backend",
						Key:        "xyz789",
						Patterns:   []string{"@deepl/*"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "valid config with authentication disabled",
			tomlData: `
enable_upload_authentication = false
api_keys = []`,
			want: &Config{
				EnableUploadAuthentication: false,
				APIKeys:                    []*ApiKey{},
			},
			wantErr: false,
		},
		{
			name: "invalid toml syntax",
			tomlData: `
enable_upload_authentication = true
[[api_keys]
invalid-toml-here`,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFromToml(tt.tomlData)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFromToml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if got.EnableUploadAuthentication != tt.want.EnableUploadAuthentication {
				t.Errorf("ParseFromToml() EnableUploadAuthentication = %v, want %v",
					got.EnableUploadAuthentication, tt.want.EnableUploadAuthentication)
			}

			if len(got.APIKeys) != len(tt.want.APIKeys) {
				t.Errorf("ParseFromToml() APIKeys length = %v, want %v",
					len(got.APIKeys), len(tt.want.APIKeys))
				return
			}

			for i, wantKey := range tt.want.APIKeys {
				gotKey := got.APIKeys[i]
				if gotKey.Identifier != wantKey.Identifier {
					t.Errorf("ParseFromToml() APIKeys[%d].Identifier = %v, want %v",
						i, gotKey.Identifier, wantKey.Identifier)
				}
				if gotKey.Key != wantKey.Key {
					t.Errorf("ParseFromToml() APIKeys[%d].Key = %v, want %v",
						i, gotKey.Key, wantKey.Key)
				}
				if len(gotKey.Patterns) != len(wantKey.Patterns) {
					t.Errorf("ParseFromToml() APIKeys[%d].Patterns length = %v, want %v",
						i, len(gotKey.Patterns), len(wantKey.Patterns))
					continue
				}
				for j, pattern := range wantKey.Patterns {
					if gotKey.Patterns[j] != pattern {
						t.Errorf("ParseFromToml() APIKeys[%d].Patterns[%d] = %v, want %v",
							i, j, gotKey.Patterns[j], pattern)
					}
				}
			}
		})
	}
}

func TestParseFromFile(t *testing.T) {
	testParser := func(data string) (*Config, error) {
		if data == "invalid" {
			return nil, errors.New("parse error")
		}
		return &Config{DBConnectionString: data}, nil
	}

	tmpDir, err := os.MkdirTemp("", "config-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name        string
		configFile  string
		fileContent string
		wantErr     bool
		setup       func(string) error
	}{
		{
			name:       "empty config file path returns default config",
			configFile: "",
			wantErr:    false,
		},
		{
			name:       "non-existent file returns error",
			configFile: filepath.Join(tmpDir, "nonexistent.toml"),
			wantErr:    true,
		},
		{
			name:        "invalid content returns error",
			configFile:  filepath.Join(tmpDir, "invalid.toml"),
			fileContent: "invalid",
			wantErr:     true,
			setup: func(path string) error {
				return os.WriteFile(path, []byte("invalid"), 0644)
			},
		},
		{
			name:        "valid content returns config",
			configFile:  filepath.Join(tmpDir, "valid.toml"),
			fileContent: "valid content",
			wantErr:     false,
			setup: func(path string) error {
				return os.WriteFile(path, []byte("valid content"), 0644)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup if needed
			if tt.setup != nil {
				if err := tt.setup(tt.configFile); err != nil {
					t.Fatalf("setup failed: %v", err)
				}
			}

			cfg, err := ParseFromFile(tt.configFile, testParser)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if cfg == nil {
					t.Error("ParseFromFile() returned nil config when no error was expected")
				}

				if tt.configFile == "" && cfg == nil {
					t.Error("ParseFromFile() with empty path did not return default config")
				}

				if tt.configFile != "" && cfg.DBConnectionString != tt.fileContent {
					t.Errorf("ParseFromFile() got config with DBConnectionString = %v, want %v",
						cfg.DBConnectionString, tt.fileContent)
				}
			}
		})
	}
}
