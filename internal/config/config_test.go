package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/database/postgres"
	"reflect"
	"testing"
	"time"
)

const (
	envFile   = "./fixtures/.env"
	directory = "./fixtures"
	ymlFile   = "main"
)

func TestConfig(t *testing.T) {
	type args struct {
		directoryPath string
		ymlFile       string
	}

	initConfigs := func() error {
		if err := godotenv.Load(envFile); err != nil {
			return err
		}

		viper.AddConfigPath(directory)
		viper.SetConfigName(ymlFile)
		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		return nil
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "test config",
			args: args{
				directoryPath: directory,
				ymlFile:       ymlFile,
			},
			want: &Config{
				HTTP: HTTPConfig{
					Host:         "0.0.0.0",
					Port:         "8080",
					ReadTimeout:  time.Second * 10,
					WriteTimeout: time.Second * 10,
				},
				Auth: AuthConfig{
					AccessTokenTTL:  time.Minute * 1,
					RefreshTokenTTL: time.Hour * 1,
					Salt:            "salt",
					Secret:          "secret",
				},
				GDrive: GDriveConfig{Credentials: "credentials"},
				GIN:    GINConfig{Mode: "release"},
				DB: postgres.DBConfig{
					Host:     "host",
					Port:     "port",
					User:     "user",
					Name:     "name",
					SSLMode:  "disable",
					Password: "password",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := initConfigs()

			if (err != nil) != tt.wantErr {
				t.Errorf("initConfigs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got := New()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() got = %v, want %v", got, tt.want)
			}
		})
	}
}
