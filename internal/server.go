package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-related/library-rest/internal/configurations"
	"github.com/go-related/library-rest/internal/handlers"
	"github.com/go-related/library-rest/internal/persistance"
	"github.com/sirupsen/logrus"
)

type Server struct {
	handler *handlers.Handler
}

func NewServer(config *configurations.Library) (*Server, error) {
	router := gin.Default()
	booksDb, err := persistance.NewBooks(config.DbConnectionString)
	if err != nil {
		logrus.WithError(err).Fatal("couldn't connect to db")
	}
	handler := handlers.NewHandler(booksDb, router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	err = router.Run(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		logrus.WithError(err).Errorf("Setting up service failed.")
		return nil, err
	}
	logrus.Infof("Application is running on port:%s", config.Port)
	return &Server{handler: handler}, nil
}
