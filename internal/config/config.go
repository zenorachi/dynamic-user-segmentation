package config

import (
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/database/postgres"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/logger"
)

type Config struct {
	HTTP   HTTPConfig
	Auth   AuthConfig
	GDrive GDriveConfig
	GIN    GINConfig
	DB     postgres.DBConfig
}

type (
	HTTPConfig struct {
		Host         string
		Port         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	AuthConfig struct {
		AccessTokenTTL  time.Duration
		RefreshTokenTTL time.Duration
		Salt            string
		Secret          string
	}

	GDriveConfig struct {
		Credentials string
	}

	GINConfig struct {
		Mode string
	}
)

var (
	config = new(Config)
	once   sync.Once
)

func New() *Config {
	once.Do(func() {
		if err := viper.Unmarshal(config); err != nil {
			logger.Fatal("viper config", err.Error())
		}

		if err := envconfig.Process("db", &config.DB); err != nil {
			logger.Fatal("db config", err.Error())
		}

		if err := envconfig.Process("hash", &config.Auth); err != nil {
			logger.Fatal("hash envs", err.Error())
		}

		if err := envconfig.Process("gdrive", &config.GDrive); err != nil {
			logger.Fatal("gdrive", err.Error())
		}

		if err := envconfig.Process("gin", &config.GIN); err != nil {
			logger.Fatal("gin config", err.Error())
		}
	})

	return config
}
