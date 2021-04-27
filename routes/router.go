package routes

import (
	"ScanForLogin/config"
	"ScanForLogin/controller"
	"ScanForLogin/utils"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

// InitRouter 初始化路由
func InitRouter() {
	// 设置gin的启动模式
	gin.SetMode(config.ServerCfg.AppMode)
	r := gin.New()
	// 加载模板
	r.HTMLRender = loadTemplates("./static")

	// 设置路由组
	router := r.Group("/")
	{
		//router.GET("/login", controller.Login)
		router.GET("/", controller.Index)
		router.GET("/hello", controller.Hello)
		router.GET("/pc-login", controller.PcLogin)
		router.GET("/qr-code/:uid", controller.QrCode)
		router.GET("/cellphone-confirm/:uid", controller.ConfirmScanStatus)
	}

	err := r.Run(config.ServerCfg.HttpPort)
	if err != nil {
		utils.Logger.Panic("服务启动失败")
	}
}

// 加载模板
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	// 返回与模式匹配的所有文件名(字符串列表)
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		utils.Logger.Panic("layouts模板加载失败")
		return nil
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		utils.Logger.Panic("includes模板加载失败")
		return nil
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		// filepath.Base(path)返回路径中的最后一个元素
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
