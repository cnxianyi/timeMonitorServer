package router

import (
	"github.com/gin-gonic/gin"
	"os"
	"timeMonitorServer/controller"
)

func Init() *gin.Engine {

	// 设置模式
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "debug"
	}
	gin.SetMode(ginMode)
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("upload", controller.Upload)

	router.GET("/", controller.All)

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
	return router
}
