package initialize

import (
	"cron_tab_c/master/global"
	"cron_tab_c/master/handler"
	"cron_tab_c/master/router"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	var (
		CoreRouter *gin.Engine
	)
	global.SLogger.Info("初始化主路由")
	CoreRouter = gin.Default()

	CoreRouter.LoadHTMLGlob("./master/templates/*")

	CoreRouter.GET("index/", handler.Index)

	{
		router.Task(CoreRouter)
	}

	return CoreRouter

}
