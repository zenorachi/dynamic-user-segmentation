package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/zenorachi/dynamic-user-segmentation/internal/app"
	"github.com/zenorachi/dynamic-user-segmentation/internal/config"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/logger"
)

const (
	envFile   = ".env"
	directory = "configs"
	ymlFile   = "main"
)

func main() {
	if err := godotenv.Load(envFile); err != nil {
		logger.Fatal("config", ".env initialization failed")
	}

	viper.AddConfigPath(directory)
	viper.SetConfigName(ymlFile)

	app.Run(config.New())
}
