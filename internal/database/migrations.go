package database

import (
	"errors"
	"fmt"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/database/postgres"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	sourceMigrationsURL = "file://scripts/migrations"
)

func DoMigrations(cfg *postgres.DBConfig) error {
	database := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode)

	m, err := migrate.New(sourceMigrationsURL, database)
	if err != nil {
		return err
	}

	err = m.Up()
	defer func() { _, _ = m.Close() }()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
