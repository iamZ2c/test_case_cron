package initialize

import (
	"cron_tab_c/master/global"
	"fmt"
	etcd "go.etcd.io/etcd/client/v3"
	"time"
)

func LoadEtcdClient() {

	var (
		config etcd.Config
		client *etcd.Client
		err    error
	)

	config = etcd.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 1 * time.Second,
	}

	if client, err = etcd.New(config); err != nil {
		fmt.Println(err)
		goto ERR
	}

	global.EtcdClient = client
	global.KV = etcd.NewKV(client)
	global.Watcher = etcd.NewWatcher(client)
	return
ERR:
	panic(err)
}
