package database

import (
	"testing"
	"time"
)

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupPostgresContainer(t *testing.T) (string, func()) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(30 * time.Second),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Failed to start PostgreSQL container: %v", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get container host: %v", err)
	}

	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatalf("Failed to get container port: %v", err)
	}

	connStr := "postgres://test:test@" + host + ":" + port.Port() + "/testdb?sslmode=disable"

	cleanup := func() {
		container.Terminate(ctx)
	}

	return connStr, cleanup
}

func TestDBConnection_Insert(t *testing.T) {
	// Setup PostgreSQL container
	pgConnStr, cleanup := setupPostgresContainer(t)
	defer cleanup()

	tests := []struct {
		name     string
		connStr  string
		setup    func(*DBConnection) error
		expected bool
	}{
		{
			name:    "SQLite Insert",
			connStr: "sqlite3://" + t.TempDir() + "/insert.db",
			setup: func(dbc *DBConnection) error {
				_, err := dbc.sql.Exec("CREATE TABLE test_table (id INTEGER PRIMARY KEY, name TEXT)")
				return err
			},
			expected: true,
		},
		{
			name:    "Postgres Insert",
			connStr: pgConnStr,
			setup: func(dbc *DBConnection) error {
				_, err := dbc.sql.Exec("DROP TABLE IF EXISTS test_table; CREATE TABLE test_table (id SERIAL PRIMARY KEY, name TEXT)")
				return err
			},
			expected: true,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbc, err := Connect(tt.connStr)
			if err != nil {
				t.Fatalf("Failed to connect: %v", err)
			}
			defer dbc.sql.Close()

			if err := tt.setup(dbc); err != nil {
				t.Fatalf("Setup failed: %v", err)
			}

			err = dbc.Insert("INSERT INTO test_table (name) VALUES ($1)", "test")
			if (err == nil) != tt.expected {
				t.Errorf("Insert() error = %v, expected success = %v", err, tt.expected)
			}
		})
	}
}

func TestDBConnection_Query(t *testing.T) {
	tests := []struct {
		name     string
		connStr  string
		setup    func(*DBConnection) error
		expected int
	}{
		{
			name:    "SQLite Query",
			connStr: "sqlite3://" + t.TempDir() + "/query.db",
			setup: func(dbc *DBConnection) error {
				_, err := dbc.sql.Exec("CREATE TABLE test_table (id INTEGER PRIMARY KEY, name TEXT)")
				if err != nil {
					return err
				}
				_, err = dbc.sql.Exec("INSERT INTO test_table (name) VALUES ('test1'), ('test2')")
				return err
			},
			expected: 2,
		},
		{
			name:    "Postgres Query",
			connStr: "",
			setup: func(dbc *DBConnection) error {
				_, err := dbc.sql.Exec("DROP TABLE IF EXISTS test_table; CREATE TABLE test_table (id SERIAL PRIMARY KEY, name TEXT)")
				if err != nil {
					return err
				}
				_, err = dbc.sql.Exec("INSERT INTO test_table (name) VALUES ('test1'), ('test2')")
				return err
			},
			expected: 2,
		},
	}

	// Setup PostgreSQL container
	connStr, cleanup := setupPostgresContainer(t)
	defer cleanup()
	tests[1].connStr = connStr

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbc, err := Connect(tt.connStr)
			if err != nil {
				t.Fatalf("Failed to connect: %v", err)
			}
			defer dbc.sql.Close()

			if err := tt.setup(dbc); err != nil {
				t.Fatalf("Setup failed: %v", err)
			}

			count := 0
			err = dbc.Query(context.Background(), func(scan func(dest ...any) error) (bool, error) {
				var name string
				if err := scan(&name); err != nil {
					return false, err
				}
				count++
				return true, nil
			}, "SELECT name FROM test_table")
			if err != nil {
				t.Errorf("Query() error = %v", err)
			}
			if count != tt.expected {
				t.Errorf("Query() count = %d, expected = %d", count, tt.expected)
			}
		})
	}
}

func TestDBConnection_MigrateUp(t *testing.T) {
	// Setup PostgreSQL container
	pgConnstr, cleanup := setupPostgresContainer(t)
	defer cleanup()

	tests := []struct {
		name    string
		connStr string
	}{
		{"SQLite MigrateUp", "sqlite3://" + t.TempDir() + "/migrate.db"},
		{"Postgres MigrateUp", pgConnstr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbc, err := Connect(tt.connStr)
			if err != nil {
				t.Fatalf("Failed to connect: %v", err)
			}
			defer dbc.sql.Close()

			err = dbc.MigrateUp()
			if err != nil {
				t.Errorf("MigrateUp() error = %v", err)
			}
		})
	}
}
