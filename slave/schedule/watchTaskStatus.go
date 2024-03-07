package schedule

import (
	"context"
	"cron_tab_c/slave/common"
	"cron_tab_c/slave/executor"
	"cron_tab_c/slave/global"
	"cron_tab_c/slave/utils"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	etcd "go.etcd.io/etcd/client/v3"
	"strings"
)

func ListenEtcdStatusAndHandleSchedule() {
	var (
		watchChan     etcd.WatchChan
		err           error
		watchResponse etcd.WatchResponse
		event         *etcd.Event
	)
	global.SLogger.Info("初始化监听器")
	if watchChan = global.Watcher.Watch(context.Background(), common.TaskDir, etcd.WithPrefix()); err != nil {
		// TODO 处理异常
	}
	for watchResponse = range watchChan {
		for _, event = range watchResponse.Events {
			switch event.Type {
			// 修改或者新增
			case mvccpb.PUT:
				if strings.Split(string(event.Kv.Key), "/")[2] == "kill" {
					// 有就杀死任务，没有就不管。
					executor.KillTask(utils.HandleEventName(string(event.Kv.Key)))
				}
				if strings.Split(string(event.Kv.Key), "/")[2] == "save" {
					UpdateTaskMap(event, common.TaskEventSave)
				}
			case mvccpb.DELETE:
				UpdateTaskMap(event, common.TaskEventDelete)
			}
			fmt.Println("监听到了")
			global.ListenChan <- 1
		}
	}
	fmt.Println("监听满载")
}

func UpdateTaskMap(event *etcd.Event, taskType int) {
	var (
		task common.TaskVar
		err  error
	)
	switch taskType {
	case common.TaskEventSave:
		task, err = utils.EventToTask(event)
		if err != nil {
			panic(err)
		}
		// 修改全局 map。
		global.CurrentExistTask[task.Name] = &task
	case common.TaskEventDelete:
		taskName := utils.HandleEventName(string(event.Kv.Key))
		if _, exist := global.CurrentExistTask[taskName]; exist {
			delete(global.CurrentExistTask, taskName)
		}
	}
	k := global.CurrentExistTask
	fmt.Println(k)

}
