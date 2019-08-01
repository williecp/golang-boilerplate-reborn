package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	// userService services.V1UserService
	// errorHelper helpers.ErrorHelper
}

func (handler *UserController) TestFunction(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}