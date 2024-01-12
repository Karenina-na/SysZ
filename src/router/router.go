package router

import (
	"SysZ/src/controller"
	"SysZ/src/controller/interception"
	"github.com/gin-gonic/gin"
)

// WebAPI
//
//	@Description: web api
//	@param r	*gin.Engine
func WebAPI(r *gin.Engine) {
	// hello
	txHello := r.Group("/test/hello")
	txHello.Use(interception.Interception())
	{
		// /v1/hello/hello		测试
		txHello.GET("/", controller.HelloController)
	}
}

// API
//
//	@Description: api
//	@param r	*gin.Engine
func API(r *gin.Engine) {
	// login
	txLogin := r.Group("/v1/login")
	txLogin.Use(interception.Interception())
	{
		// /v1/login/login		登录
		txLogin.POST("/", controller.RootController)
	}

	// operator
	txOperator := r.Group("/v1/operator")
	txOperator.Use(interception.Interception(), interception.RootInterception())
	{
		// /v1/operator/get		获取服务器状态
		txOperator.GET("/getAll", controller.GetStatusController)
	}
}
