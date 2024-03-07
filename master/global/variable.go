package global

import (
	"cron_tab_c/master/common"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis"
	etcd "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

var (
	// EtcdClient etcd操作
	EtcdClient *etcd.Client
	KV         etcd.KV
	Watcher    etcd.Watcher

	// SLogger 日志
	SLogger *zap.SugaredLogger

	// TotalConfig nacos配置
	TotalConfig   common.TotalConfig
	LimiterConfig common.RateLimitConfig

	// Rs redis操作
	Rs        *redsync.Redsync
	RedisPool redis.Pool
)
