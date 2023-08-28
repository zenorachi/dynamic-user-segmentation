package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zenorachi/dynamic-user-segmentation/internal/service/storage"

	_ "github.com/zenorachi/dynamic-user-segmentation/docs"
	"github.com/zenorachi/dynamic-user-segmentation/internal/config"
	"github.com/zenorachi/dynamic-user-segmentation/internal/database"
	"github.com/zenorachi/dynamic-user-segmentation/internal/repository"
	"github.com/zenorachi/dynamic-user-segmentation/internal/server"
	"github.com/zenorachi/dynamic-user-segmentation/internal/service"
	"github.com/zenorachi/dynamic-user-segmentation/internal/transport"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/auth"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/database/postgres"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/hash"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/logger"
)

const (
	shutdownTimeout = 5 * time.Second
)

// @title           			Dynamic User Segmentation Service
// @version         			1.0
// @description     			This is a service for segmenting users with the ability to automatically add and remove users from segments.

// @contact.name   				Maksim Sonkin
// @contact.email  				msonkin33@gmail.com

// @host      					localhost:8080
// @BasePath  					/

// @securityDefinitions.apikey  Bearer
// @in 						    header
// @name 					    Authorization
// @description					Type "Bearer" followed by a space and JWT token.

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
	logger.Info("database", "postgres connected")

	/* INIT TOKEN MANAGER */
	tokenManager := auth.NewManager(cfg.Auth.Secret)

	/* INIT SERVICES & DEPS */
	services := service.New(service.Deps{
		Repos:           repository.New(db),
		Hasher:          hash.NewSHA1Hasher(cfg.Auth.Salt),
		TokenManager:    tokenManager,
		AccessTokenTTL:  cfg.Auth.AccessTokenTTL,
		RefreshTokenTTL: cfg.Auth.RefreshTokenTTL,
		Storage:         storage.NewProvider(&cfg.GDrive),
	})

	/* INIT HTTP HANDLER */
	handler := transport.NewHandler(services, tokenManager)

	/* INIT HTTP SERVER */
	srv := server.New(cfg, handler.InitRoutes())
	go func() {
		if err = srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("server", err)
		}
	}()
	logger.Info("server", "server started")

	/* GRACEFUL SHUTDOWN */
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	/* WAITING FOR SYSCALL */
	<-quit

	/* SHUTTING DOWN */
	logger.Info("server", "shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		logger.Fatal("server", err)
	}
	logger.Info("server", "server stopped")
}
