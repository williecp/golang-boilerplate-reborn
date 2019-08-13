package controller

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"example_app/router"
)

var router *gin.Engine

func LoadRouterMock() (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	resp := httptest.NewRecorder()
	context, router := gin.CreateTestContext(resp)
	// routers.UserRouter(router,&UserController{
	// 	UserService: UserServiceMock{},
	// })
	return context, router, resp
}