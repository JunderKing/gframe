package router

import (
	"gframe/controller"
	"gframe/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)
	// 通用中间件
	router.Use(middleware.Recover())
	// 探活
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 路由组
	apiGroup := router.Group("/demo")
	{
		demoController := controller.Demo{}
		apiGroup.POST("/user/create", demoController.CreateUser)
	}

	return router
}
