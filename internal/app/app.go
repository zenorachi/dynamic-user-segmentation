package app

import (
	"github.com/zenorachi/dynamic-user-segmentation/internal/config"
	"github.com/zenorachi/dynamic-user-segmentation/internal/database"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/database/postgres"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/logger"
)

func Run(cfg *config.Config) {
	/* DO MIGRATIONS */
	err := database.DoMigrations(&cfg.DB)
	if err != nil {
		logger.Fatal("migrations", "migrations failed")
	}
	logger.Info("migrations", "migrations done")

	/* INIT POSTGRES-DB */
	db, err := postgres.NewDB(&cfg.DB)
	defer func() { _ = db.Close() }()
	if err != nil {
		logger.Fatal("database-connection", err)
	}
	logger.Info("database", "postgres started")
}
