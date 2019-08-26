package routes

import (
	"github.com/gin-gonic/gin"
	"api/controllers"
)

func RegisterApiRoutes(router *gin.Engine) {

	//定义路由组
	apiRouter := router.Group("api")

	// 获取用户信息
	apiRouter.GET("/user/:username", controllers.UserInfo)

	// 新增用户
	apiRouter.GET("/createuser", controllers.CreateUser)
}
