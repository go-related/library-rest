package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-related/library-rest/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) GetAuthors(c *gin.Context) {
	type QueryParameter struct {
		Limit  string `form:"limit,default=5" binding:"numeric"`
		Offset string `form:"offset,default=0" binding:"numeric"`
	}
	//TODO make uses of the pagination
	result, err := h.BookDb.GetAllAuthors(c.Request.Context())
	if err != nil {
		AbortWithMessage(c, http.StatusInternalServerError, err, "failed to load authors")
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func (h *Handler) GetAuthor(c *gin.Context) {
	id := c.Params.ByName("id")

	idValue, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		AbortWithMessage(c, http.StatusBadRequest, err, "error converting id to int")
		return
	}
	result, err := h.BookDb.GetAuthorById(c.Request.Context(), uint(idValue))
	if err != nil {
		AbortWithMessage(c, http.StatusInternalServerError, err, "failed to load author")
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func (h *Handler) DeleteAuthor(c *gin.Context) {
	id := c.Params.ByName("id")
	idValue, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		AbortWithMessage(c, http.StatusBadRequest, err, "error converting id to int")
		return
	}

	err = h.BookDb.DeleteAuthor(c.Request.Context(), uint(idValue))
	if err != nil {
		AbortWithMessage(c, http.StatusInternalServerError, err, "failed to delete author")
		return
	}
	c.IndentedJSON(http.StatusNoContent, gin.H{
		"message": "Resource deleted successfully",
	})
}

func (h *Handler) CreateAuthor(c *gin.Context) {
	// prepare input
	type Author struct {
		Name string `json:"name"`
	}
	var input Author
	err := c.BindJSON(&input)
	if err != nil {
		AbortWithMessage(c, http.StatusBadRequest, err, "error binding to json")
		return
	}
	authorData := models.Author{
		PublicName: input.Name,
	}
	// execute
	data, err := h.BookDb.CreateAuthor(c.Request.Context(), authorData)
	if err != nil {
		AbortWithMessage(c, http.StatusInternalServerError, err, "failed to create author")
		return
	}
	c.IndentedJSON(http.StatusCreated, data)
}

func (h *Handler) UpdateAuthor(c *gin.Context) {
	id := c.Params.ByName("id")
	idValue, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		AbortWithMessage(c, http.StatusBadRequest, err, "error converting id to int")
		return
	}
	// prepare the input
	type Author struct {
		Name string `json:"name"`
	}
	var input Author
	err = c.BindJSON(&input)
	if err != nil {
		AbortWithMessage(c, http.StatusBadRequest, err, err.Error())
		return
	}
	authorData := models.Author{
		PublicName: input.Name,
	}
	authorData.ID = uint(idValue)

	// update
	err = h.BookDb.UpdateAuthor(c.Request.Context(), authorData)
	if err != nil {
		AbortWithMessage(c, http.StatusInternalServerError, err, "failed to update author")
		return
	}
	c.IndentedJSON(http.StatusOK, authorData)
}
