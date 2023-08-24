package postgres

import (
	"database/sql"
	"fmt"
)

const (
	driverName = "postgres"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	SSLMode  string
	Password string
}

func NewDB(cfg *DBConfig) (*sql.DB, error) {
	source := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.SSLMode, cfg.Password)

	db, err := sql.Open(driverName, source)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
