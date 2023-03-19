package handler

import (
	"go-book/helper"
	"go-book/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewHandlerUser(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUserHandler(c *gin.Context) {

	var input user.RegisterInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register User Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.CreateUser(input)

	if err != nil {
		response := helper.APIResponse("Register User Failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatRegisterUser(newUser)

	response := helper.APIResponse("Success to Register User", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailibility(c *gin.Context) {
	var email user.CheckEmailInput

	err := c.ShouldBindJSON(&email)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register User Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	isRegister, err := h.userService.GetUserByEmail(email)

	if err != nil {
		response := helper.APIResponse("Email has been registered", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	messageData := gin.H{
		"is_available": isRegister,
	}

	c.JSON(http.StatusOK, messageData)
}
