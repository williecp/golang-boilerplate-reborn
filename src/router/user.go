package router

import (
	"example_app/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	handler := &controllers.UserController{

	}
	group := router.Group("v1/users")
	// group.Use(defaultMiddleware.AuthenticationMiddleware()){
		group.GET("", handler.TestFunction)
		// group.POST(":id", handler.UpdateById)
	// }
}
