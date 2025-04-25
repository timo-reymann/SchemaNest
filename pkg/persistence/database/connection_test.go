package database

import "testing"

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
