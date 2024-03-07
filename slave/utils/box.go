package utils

import (
	"cron_tab_c/slave/common"
	"cron_tab_c/slave/global"
	"encoding/json"
	etcd "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

func EventToTask(event *etcd.Event) (task common.TaskVar, err error) {
	if err = json.Unmarshal(event.Kv.Value, &task); err != nil {
		return task, err
	}
	return task, nil
}

func CountRecentTime() (minTime time.Duration) {
	now := time.Now()
	minTime = 10 * time.Second
	for _, task := range global.CurrentExistTask {
		k := task.NextTime.Sub(now)
		if k < minTime {
			minTime = task.NextTime.Sub(now)
		}
	}
	return minTime
}

func HandleEventName(dirName string) string {
	return strings.Split(dirName, "/")[3]
}

func HandleEventGetKey() {

}
