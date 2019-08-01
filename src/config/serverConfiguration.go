package config

import (
	"github.com/gin-gonic/gin"
	"example_app/util/middleware"
)

// default server router configuration
func SetupRouter() *gin.Engine {
	defaultMiddleware := middleware.DefaultMiddleware{}
	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())
	return router
}