package main

import (
	"github.com/gin-gonic/gin"
	RegisterRoutes "api/routes"
)

func setupRouter() *gin.Engine {

	//创建初始路由
	router := gin.New()

	// 注册api路由
	RegisterRoutes.RegisterApiRoutes(router)

	//返回路由对象
	return router
}