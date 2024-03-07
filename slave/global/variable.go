package global

import (
	"cron_tab_c/slave/common"
	"github.com/go-redsync/redsync/v4"
	etcd "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

var (
	EtcdClient       *etcd.Client
	KV               etcd.KV
	Watcher          etcd.Watcher
	CurrentExistTask map[string]*common.TaskVar
	ListenChan       chan int
	RunningTaskMap   map[string]*common.RunningTask
	Rs               *redsync.Redsync
	SLogger          *zap.SugaredLogger
)
