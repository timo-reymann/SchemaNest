package database

import (
	"context"
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

// Query the database and map the results using the provided rowMapper function.
// The rowMapper function gets the sql.Rows#Scan method as an argument, if it returns false, the loop will break, the same if it returns an error.
func (dbc *DBConnection) Query(ctx context.Context, rowMapper func(scan func(dest ...any) error) (bool, error), sql string, parameters ...any) error {
	q, err := dbc.sql.QueryContext(ctx, sql, parameters...)

	if err != nil {
		return errors.Join(errors.New("failed to run query"), err)
	}
	defer q.Close()

	for q.Next() {
		cont, err := rowMapper(q.Scan)
		if err != nil {
			return errors.Join(errors.New("failed to read results"), err)
		}

		if !cont {
			break
		}
	}

	return nil
}

func (dbc *DBConnection) Insert(sql string, parameters ...any) error {
	stmt, err := dbc.sql.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(parameters...)
	return err
}

func Connect(connStr string) (*DBConnection, error) {
	typ := databaseType(connStr)
	if typ == "" {
		return nil, errors.New("unsupported database type")
	}

	conn, err := sql.Open(typ, strings.ReplaceAll(connStr, "sqlite3://", "file:"))
	if err != nil {
		return nil, errors.Join(errors.New("failed to create database connection"), err)
	}

	return &DBConnection{
		ConnType:   typ,
		connString: connStr,
		sql:        conn,
	}, nil
}
