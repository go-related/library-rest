package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-related/library-rest/internal/configurations"
	"github.com/go-related/library-rest/internal/handlers"
	"github.com/go-related/library-rest/internal/persistance"
	"github.com/go-related/library-rest/internal/services"
	"github.com/sirupsen/logrus"
)

type Server struct {
	handler *handlers.Handler
}

func NewServer(config *configurations.Library) (*Server, error) {
	router := gin.Default()
	router.Use(handlers.CORSMiddleware())

	booksDb, err := persistance.NewBooks(config.DbConnectionString)
	if err != nil {
		logrus.WithError(err).Fatal("couldn't connect to db")
	}
	booksService, err := services.NewService(booksDb)
	if err != nil {
		logrus.WithError(err).Fatal("couldn't setup server")
	}
	handler := handlers.NewHandler(booksService, router)

	err = router.Run(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		logrus.WithError(err).Errorf("Setting up service failed.")
		return nil, err
	}
	logrus.Infof("Application is running on port:%s", config.Port)
	return &Server{handler: handler}, nil
}
