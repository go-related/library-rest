package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-related/library-rest/internal/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Handler struct {
	Service services.Service
	Engine  *gin.Engine
}

type Response struct {
	StatusCode int
	Err        error
}

func NewHandler(bookService services.Service, router *gin.Engine) *Handler {
	handler := &Handler{
		Service: bookService,
		Engine:  router,
	}
	SetupHealth(router)
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

	// if custom validation error update status and message
	var badRequest *services.ServiceError
	if errors.As(err, &badRequest) {
		status = http.StatusBadRequest
		message = err.Error()
	}

	c.AbortWithStatusJSON(status, Response{
		StatusCode: status,
		Err:        errors.New(message),
	})
}

func getParamUInt(c *gin.Context, paramName string) (uint, error) {
	id := c.Params.ByName(paramName)
	idValue, err := strconv.ParseUint(id, 10, 32)
	return uint(idValue), err
}
