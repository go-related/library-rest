package configurations

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Library struct {
	DbConnectionString string
	Port               string
	EnableMigrations   bool
}

func NewLibrary() (*Library, error) {
	v := viper.New()
	v.SetConfigName("library")
	v.SetConfigType("yml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file because we
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.Is(err, &configFileNotFoundError) {
			logrus.WithError(err).Warning("error loading config file")
		}
	}
	v.AutomaticEnv()
	var config Library
	err := v.UnmarshalExact(&config)
	return &config, err
}
