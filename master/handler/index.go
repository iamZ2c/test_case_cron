package handler

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	var (
		e *base.SentinelEntry
		b *base.BlockError
	)
	if e, b = sentinel.Entry("[TASK_INDEX_HTML]", sentinel.WithTrafficType(base.Inbound)); b != nil {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"msg":   "success",
				"tasks": "流量限制",
			},
		)
		return
	}
	ctx.HTML(http.StatusOK, "index.html", nil)
	e.Exit()
	return
}
