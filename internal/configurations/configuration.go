package configurations

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Library struct {
	DbConnectionString string
	Port               string
	EnableMigrations   bool
	DbHost             string `mapstructure:"db-host"`
	DbName             string `mapstructure:"db-name"`
	DbPort             string `mapstructure:"db-port"`
	DbUser             string `mapstructure:"db-user"`
	DbPassword         string `mapstructure:"db-password"`
}

func NewLibrary() (*Library, error) {
	v := viper.New()
	v.SetConfigName("library")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file because we
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.Is(err, &configFileNotFoundError) {
			logrus.WithError(err).Warning("error loading config file")
		}
	}

	var config Library
	err := v.UnmarshalExact(&config)
	// since we expect the db config from env variables we will overwrite the library configuration
	host := os.Getenv("db-host")
	if err != nil && host != "" {
		config.DbConnectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, os.Getenv("db-user"), os.Getenv("db-password"), os.Getenv("db-name"), os.Getenv("db-port"))
		logrus.WithField("conn", config.DbConnectionString).Info("updating default connection string")
	}
	return &config, err
}
