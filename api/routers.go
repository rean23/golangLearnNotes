package main

import (
	"github.com/gin-gonic/gin"
	RegisterRoutes "api/routes"
)

func setupRouter() *gin.Engine {

	//创建初始路由
	router := gin.New()//创建一个无中间件的路由
	// 如果要创建一个带日志和恢复的中间件路由  router := gin.Default()

	router.Use(gin.Logger())// 使用日志中间件

	router.Use(gin.Recovery())// 使用恢复中间件

	// 注册api路由
	RegisterRoutes.RegisterApiRoutes(router)

	//返回路由对象
	return router
}