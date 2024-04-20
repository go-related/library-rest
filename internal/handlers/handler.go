package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-related/library-rest/internal/persistance"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	BookDb persistance.BooksDB
	Engine *gin.Engine
}

type Response struct {
	StatusCode int
	Err        error
}

func NewHandler(bookDb persistance.BooksDB, router *gin.Engine) *Handler {
	handler := &Handler{
		BookDb: bookDb,
		Engine: router,
	}
	v1 := router.Group("/v1/api")

	// register authors
	v1.GET("/authors", handler.GetAuthors)
	v1.GET("/authors/:id", handler.GetAuthor)
	v1.PUT("/authors/:id", handler.UpdateAuthor)
	v1.POST("/authors", handler.CreateAuthor)
	v1.DELETE("/authors/:id", handler.DeleteAuthor)

	return handler
}

func AbortWithMessage(c *gin.Context, status int, err error, message string) {
	logrus.WithError(err).Error(message)
	errorData := Response{
		StatusCode: status,
		Err:        errors.New(message),
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, errorData)
}
