package initialize

import (
	"cron_tab_c/master/global"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

func RedisClient() {
	global.SLogger.Info("初始化redis客户端")
	client := goredislib.NewClient(
		&goredislib.Options{
			Addr: "127.0.0.1:6397",
		},
	)
	pool := goredis.NewPool(client)
	global.Rs = redsync.New(pool)
	global.RedisPool = pool
}
