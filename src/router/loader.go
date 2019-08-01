package router

import (
	"github.com/gin-gonic/gin"
)

func LoadRouter(router *gin.Engine) {
	UserRouter(router)	
}