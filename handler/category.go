package handler

import (
	"go-book/category"
	"go-book/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService category.Service
}

func NewCategoryHandler(service category.Service) *categoryHandler {
	return &categoryHandler{service}
}

func (h *categoryHandler) CreateCategoryHandler(c *gin.Context) {
	var input category.CategoryInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create category", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isAvailable, err := h.categoryService.FindCategoryName(input.CategoryName)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create category", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if !isAvailable {
		response := helper.APIResponse("Failed to save category, category is registered", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCategory, err := h.categoryService.SaveCategory(input)

	if err != nil {
		response := helper.APIResponse("Failed to create category", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := category.FormatCategory(newCategory)

	response := helper.APIResponse("Success to Create category", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)
}
