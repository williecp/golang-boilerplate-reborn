package router

import (
	// "example_app/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine, handler *controllers.UserController) {
	group := router.Group("v1/users")
	group.GET("", handler.GetUsers)
	group.GET(":id", handler.GetUserByID)
	group.PUT(":id", handler.UpdateUsersByID)
}
