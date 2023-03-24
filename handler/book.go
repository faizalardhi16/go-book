package handler

import (
	"fmt"
	"go-book/book"
	"go-book/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) CreateBookHandler(c *gin.Context) {
	var input book.CreateBookInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Input cannot be allowed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newBook, err := h.bookService.SaveBook(input)

	fmt.Println(newBook, "NEW BOOK")

	if err != nil {
		response := helper.APIResponse("Failed to ceate book", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	format := book.FormatResponseBook(newBook)

	response := helper.APIResponse("Success to create book", http.StatusOK, "success", format)

	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) GetAllBook(c *gin.Context) {
	allBook, err := h.bookService.GetAllBook()

	if err != nil {
		response := helper.APIResponse("Failed to load Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := book.FormatGetAllBook(allBook)

	response := helper.APIResponse("Success to load data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
