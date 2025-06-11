package config

import "testing"

func TestApiKey_Validate(t *testing.T) {
	tests := []struct {
		name    string
		apiKey  ApiKey
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid api key",
			apiKey: ApiKey{
				Identifier: "test-key",
				Key:        "secret123",
				Patterns:   []string{"my-schema"},
			},
			wantErr: false,
		},
		{
			name: "empty identifier",
			apiKey: ApiKey{
				Identifier: "",
				Key:        "secret123",
				Patterns:   []string{"@foo/*"},
			},
			wantErr: true,
			errMsg:  "api key identifier cannot be empty",
		},
		{
			name: "empty key",
			apiKey: ApiKey{
				Identifier: "test-key",
				Key:        "",
				Patterns:   []string{"foo/*"},
			},
			wantErr: true,
			errMsg:  "api key cannot be empty",
		},
		{
			name: "no patterns",
			apiKey: ApiKey{
				Identifier: "test-key",
				Key:        "secret123",
				Patterns:   []string{},
			},
			wantErr: true,
			errMsg:  "at least one pattern must be specified for api key test-key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.apiKey.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiKey.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("ApiKey.Validate() error message = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}
