package routes

import (
	"ScanForLogin/config"
	"github.com/gin-gonic/gin"
	"ScanForLogin/controller"
)

// InitRouter 初始化路由
func InitRouter()  {
	// 设置gin的启动模式
	gin.SetMode(config.ServerCfg.AppMode)
	r := gin.New()

	// 设置路由组
	router := r.Group("/")
	{
		router.GET("/", controller.)
	}
}