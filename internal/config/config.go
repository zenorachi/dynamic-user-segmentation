package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/database/postgres"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/logger"
	"sync"
	"time"
)

type Config struct {
	HTTP HTTPConfig
	Auth AuthConfig
	DB   postgres.DBConfig
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
)

var (
	config = new(Config)
	once   sync.Once
)

func New() *Config {
	once.Do(func() {
		if err := viper.ReadInConfig(); err != nil {
			logger.Fatal("config", "viper initialization failed")
		}

		if err := viper.Unmarshal(config); err != nil {
			logger.Fatal("config", "viper initialization failed")
		}

		if err := envconfig.Process("db", &config.DB); err != nil {
			logger.Fatal("config", "dbConfig initialization failed")
		}

		if err := envconfig.Process("hash", &config.Auth); err != nil {
			logger.Fatal("config", "hashConfig initialization failed")
		}
	})

	return config
}
