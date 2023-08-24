package config

import (
	"github.com/zenorachi/dynamic-user-segmentation/pkg/database/postgres"
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
