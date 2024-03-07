package main

import (
	"cron_tab_c/master/common"
	"cron_tab_c/master/global"
	"cron_tab_c/master/initialize"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		server *gin.Engine
		err    error
		csc    common.CronServerConfig
	)
	if err = initialize.NacosClient(); err != nil {
		goto ERR
	}

	initialize.Logger()

	initialize.LoadEtcdClient()

	server = initialize.Router()

	initialize.RateLimitClient()

	initialize.RedisClient()

	csc = global.TotalConfig.CronServerConfig

	if err = server.Run(csc.Host + ":" + csc.IP); err != nil {
		goto ERR
	}
ERR:
	fmt.Println(err)
	panic(err)
}
