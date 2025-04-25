package database

import (
	"github.com/golang-migrate/migrate/v4"
	"testing"
)

func TestMigrations(t *testing.T) {
	// Setup PostgreSQL container
	pgConnStr, cleanup := setupPostgresContainer(t)
	defer cleanup()

	// Define test matrix
	tests := []struct {
		name    string
		connStr string
	}{
		{"SQLite Migration", "sqlite3://" + t.TempDir() + "/db.db"},
		{"Postgres Migration", pgConnStr},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Connect to the database
			dbc, err := Connect(tt.connStr)
			if err != nil {
				t.Fatalf("Failed to connect to database: %v", err)
			}
			defer dbc.sql.Close()

			// Create a migration runner
			m, err := NewMigrationRunner(dbc)
			if err != nil {
				t.Fatalf("Failed to create migration runner: %v", err)
			}

			// Run migrations up
			if err := m.Up(); err != nil && err != migrate.ErrNoChange {
				t.Errorf("Up() failed: %v", err)
			}

			// Run migrations down
			if err := m.Down(); err != nil && err != migrate.ErrNoChange {
				t.Errorf("Down() failed: %v", err)
			}
		})
	}
}
