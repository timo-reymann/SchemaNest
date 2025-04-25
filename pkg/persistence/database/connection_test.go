package database

import (
	"testing"
)

func TestDatabaseType(t *testing.T) {
	tests := []struct {
		name     string
		connStr  string
		expected string
	}{
		{"Valid SQLite", "sqlite3:example.db", Sqlite},
		{"Valid Postgres", "postgres://user:pass@localhost/db", Postgres},
		{"Invalid Type", "mysql://user:pass@localhost/db", ""},
		{"Empty Connection String", "", ""},
		{"Malformed Connection String", "invalid", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := databaseType(tt.connStr)
			if result != tt.expected {
				t.Errorf("databaseType(%q) = %q; want %q", tt.connStr, result, tt.expected)
			}
		})
	}
}

func TestConnect(t *testing.T) {
	tests := []struct {
		name    string
		connStr string
		wantErr bool
	}{
		{"Valid SQLite", "sqlite3:example.db", false},
		{"Valid Postgres", "postgres://user:pass@localhost/db", false},
		{"Invalid Type", "mysql://user:pass@localhost/db", true},
		{"Empty Connection String", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Connect(tt.connStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect(%q) error = %v, wantErr %v", tt.connStr, err, tt.wantErr)
			}
		})
	}
}
