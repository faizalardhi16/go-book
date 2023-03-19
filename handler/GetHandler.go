package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Output(c gin.Context) {
	c.JSON(http.StatusOK, "HELLO")
}
