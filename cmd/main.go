package main

import (
	"github.com/go-related/library-rest/internal"
	"github.com/go-related/library-rest/internal/configurations"
	"github.com/sirupsen/logrus"
)

func main() {
	configs, err := configurations.NewLibrary()
	if err != nil {
		logrus.WithError(err).Error("failed to load configurations.")
	}
	_, err = internal.NewServer(configs)
	if err != nil {
		logrus.WithError(err).Error("failed to setup server.")
	}
}
