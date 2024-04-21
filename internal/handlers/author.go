package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-related/library-rest/internal/models"
	"net/http"
)

type AuthorDto struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func convertAuthorToDTO(sc *models.Author) *AuthorDto {
	return &AuthorDto{
		Id:   sc.Model.ID,
		Name: sc.PublicName,
	}
}

func bindToAuthor(c *gin.Context) (*models.Author, error) {
	var input AuthorDto
	err := c.BindJSON(&input)
	if err != nil {
		return nil, err
	}
	authorData := models.Author{
		PublicName: input.Name,
	}
	return &authorData, nil
}

func (h *Handler) GetAuthors(c *gin.Context) {
	//type QueryParameter struct {
	//	Limit  string `form:"limit,default=5" binding:"numeric"`
	//	Offset string `form:"offset,default=0" binding:"numeric"`
	//}
	//TODO make uses of the pagination
	result, err := h.Service.GetAllAuthors(c.Request.Context())
	if err != nil {
		AbortWithMessage(c, http.StatusInternalServerError, err, "failed to load authors")
		return
	}
	var output []*AuthorDto
	for _, item := range result {
		output = append(output, convertAuthorToDTO(item))
	}
	c.IndentedJSON(http.StatusOK, output)
}

func (h *Handler) GetAuthor(c *gin.Context) {
	id, err := getParamUInt(c, "id")
	if err != nil {
		AbortWithMessage(c, http.StatusBadRequest, err, "error converting id to int")
		return
	}
	result, err := h.Service.GetAuthorById(c.Request.Context(), id)
	if err != nil {
		AbortWithMessage(c, http.StatusInternalServerError, err, "failed to load author")
		return
	}
	c.IndentedJSON(http.StatusOK, convertAuthorToDTO(result))
}

func (h *Handler) DeleteAuthor(c *gin.Context) {
	id, err := getParamUInt(c, "id")
	if err != nil {
		AbortWithMessage(c, http.StatusBadRequest, err, "error converting id to int")
		return
	}

	err = h.Service.DeleteAuthor(c.Request.Context(), id)
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
	input, err := bindToAuthor(c)
	if err != nil {
		AbortWithMessage(c, http.StatusBadRequest, err, "error binding to json")
		return
	}
	// execute
	data, err := h.Service.CreateAuthor(c.Request.Context(), input)
	if err != nil {
		AbortWithMessage(c, http.StatusInternalServerError, err, "failed to create author")
		return
	}
	c.IndentedJSON(http.StatusCreated, convertAuthorToDTO(data))
}

func (h *Handler) UpdateAuthor(c *gin.Context) {
	id, err := getParamUInt(c, "id")
	if err != nil {
		AbortWithMessage(c, http.StatusBadRequest, err, "error converting id to int")
		return
	}
	// prepare the input

	input, err := bindToAuthor(c)
	if err != nil {
		AbortWithMessage(c, http.StatusBadRequest, err, err.Error())
		return
	}
	input.ID = id

	// update
	err = h.Service.UpdateAuthor(c.Request.Context(), input)
	if err != nil {
		AbortWithMessage(c, http.StatusInternalServerError, err, "failed to update author")
		return
	}
	c.IndentedJSON(http.StatusOK, convertAuthorToDTO(input))
}
