package router

import (
	"cron_tab_c/master/global"
	"cron_tab_c/master/handler"
	"github.com/gin-gonic/gin"
)

func Task(CoreRouter *gin.Engine) {
	var taskGroup *gin.RouterGroup
	global.SLogger.Info("初始Task路由")
	taskGroup = CoreRouter.Group("job/")

	{
		taskGroup.GET("/list", handler.GetTaskList)
		taskGroup.DELETE("/del", handler.DelJob)
		taskGroup.POST("/save", handler.UpdateTask)
	}

	return
}
