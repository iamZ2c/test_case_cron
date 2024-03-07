package initialize

import (
	"context"
	"cron_tab_c/slave/common"
	"cron_tab_c/slave/global"
	"cron_tab_c/slave/utils"
	"encoding/json"
	"go.etcd.io/etcd/api/v3/mvccpb"
	etcd "go.etcd.io/etcd/client/v3"
)

// ExistTask 初始化所有任务到内存
func ExistTask() (err error) {
	var (
		getResponse  *etcd.GetResponse
		kv           *mvccpb.KeyValue
		CurrentTasks map[string]*common.TaskVar
	)
	global.SLogger.Info("加载etcd存在的任务")
	if getResponse, err = global.KV.Get(context.Background(), common.TaskSaveDir, etcd.WithPrefix()); err != nil {
		return err
	}
	CurrentTasks = make(map[string]*common.TaskVar)
	for _, kv = range getResponse.Kvs {
		Task := common.TaskVar{}
		if err = json.Unmarshal(kv.Value, &Task); err != nil {
			panic(err)
		}
		CurrentTasks[utils.HandleEventName(string(kv.Key))] = &Task
	}
	global.CurrentExistTask = CurrentTasks
	return
}

func RunningTask() {
	// 初始化正在执行的任务map
	global.SLogger.Info("初始化执行表")
	RunningMap := make(map[string]*common.RunningTask)
	global.RunningTaskMap = RunningMap
}
