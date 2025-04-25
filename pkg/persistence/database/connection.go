package database

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

const Sqlite = "sqlite3"
const Postgres = "postgres"

func databaseType(connStr string) string {
	parts := strings.SplitN(connStr, ":", 2)
	if len(parts) < 2 {
		return ""
	}

	prefix := parts[0]
	if prefix == Sqlite || prefix == Postgres {
		return prefix
	}

	return ""
}

type DBConnection struct {
	sql        *sql.DB
	ConnType   string
	connString string
}

func (dbc *DBConnection) MigrateUp() error {
	m, err := NewMigrationRunner(dbc)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && errors.Is(migrate.ErrNoChange, err) {
		return nil
	}
	return err
}

func Connect(connStr string) (*DBConnection, error) {
	typ := databaseType(connStr)
	if typ == "" {
		return nil, errors.New("unsupported database type")
	}

	conn, err := sql.Open(typ, connStr)
	if err != nil {
		return nil, errors.Join(errors.New("failed to create database connection"), err)
	}

	return &DBConnection{
		ConnType:   typ,
		connString: connStr,
		sql:        conn,
	}, nil
}
