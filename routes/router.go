package routes

import (
	"ScanForLogin/config"
	"ScanForLogin/controller"
	"ScanForLogin/utils"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() {
	// 设置gin的启动模式
	gin.SetMode(config.ServerCfg.AppMode)
	r := gin.New()

	// 设置路由组
	router := r.Group("/")
	{
		//router.GET("/login", controller.Login)
		router.GET("/", controller.Index)
		router.GET("/pc-login", controller.PcLogin)
		router.GET("/qr-code/:uid", controller.QrCode)
		router.GET("/cellphone-confirm/:uid", controller.ConfirmScanStatus)
	}

	err := r.Run(config.ServerCfg.HttpPort)
	if err != nil {
		utils.Logger.Panic("服务启动失败")
	}
}
