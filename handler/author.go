package handler

import (
	"fmt"
	"go-book/author"
	"go-book/constant"
	"go-book/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authorHandler struct {
	authorService author.Service
}

func NewAuthorHandler(authorService author.Service) *authorHandler {
	return &authorHandler{authorService}
}

func (h *authorHandler) CreateAuthorHandler(c *gin.Context) {
	var input author.AuthorInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create author", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	checkAuthor, err := h.authorService.CheckAuthor(input.AuthorName)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create author", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if !checkAuthor {
		response := helper.APIResponse("Cannot save two identical author name", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newAuthor, err := h.authorService.SaveAuthor(input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create author", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := author.FormatterAuthor(newAuthor)

	response := helper.APIResponse("Success to save an author", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *authorHandler) DeleteAuthorHandler(c *gin.Context) {
	var input author.DeleteAuthorInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to delete author", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	responseAuthor, err := h.authorService.DestroyAuthor(input)

	fmt.Println(responseAuthor, "RESPONSE AUTHOR")

	if err != nil {
		response := helper.APIResponse("Failed to delete author", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if responseAuthor == constant.ResponseEnum.NotFound {
		response := helper.APIResponse("Author not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Success to delete an author", http.StatusOK, "success", nil)

	c.JSON(http.StatusOK, response)

}
