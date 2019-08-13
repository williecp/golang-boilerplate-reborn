package router

import (
	"github.com/gin-gonic/gin"
	"example_app/controller"
	"example_app/service"
)

func LoadRouter(router *gin.Engine) {
	UserRouter(router,&controllers.UserController{
		UserService: services.UserServiceHandler(),
	})
}