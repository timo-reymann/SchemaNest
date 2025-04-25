package database

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/timo-reymann/SchemaNest/pkg/internal/migrations"
)

func NewMigrationRunner(databaseConnection *DBConnection) (*migrate.Migrate, error) {
	names, err := migrations.AssetDir(databaseConnection.ConnType)
	if err != nil {
		return nil, err
	}

	resource := bindata.Resource(
		names,
		func(name string) ([]byte, error) {
			return migrations.Asset(databaseConnection.ConnType + "/" + name)
		},
	)
	binData, err := bindata.WithInstance(resource)
	if err != nil {
		return nil, errors.Join(errors.New("failed to load embedded migrations"), err)
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", binData, databaseConnection.connString)
	if err != nil {
		return nil, errors.Join(errors.New("failed to create migration runner"), err)
	}
	return m, nil
}
