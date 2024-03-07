package executor

import (
	"context"
	"cron_tab_c/slave/common"
	"cron_tab_c/slave/global"
	"fmt"
	"os/exec"
	"time"
)

//func ExecuteListener() {
//	for {
//		select {
//		case c:
//
//		}
//	}
//}

func ExecuteTask(taskName string) {
	//TODO 先拿分布式锁，拿到执行，没拿到直接退出，
	var (
		err  error
		isOk bool
	)
	mutex := global.Rs.NewMutex(taskName)
	if err = mutex.Lock(); err != nil {
		global.SLogger.Errorw(
			"获取分布式锁失败",
			"task_name", taskName,
		)
		return
	}
	//并且记录该goroutine的context，记录日志信息
	var (
		ctx     context.Context
		cancelF context.CancelFunc
	)
	ctx, cancelF = context.WithCancel(context.Background())

	cmd := exec.CommandContext(ctx, "python3")
	// TODO 表示该任务有被执行过，存一个正在执行的map，

	global.RunningTaskMap[taskName] = &common.RunningTask{
		Name:      taskName,
		StartTime: time.Now(),
		Cancel:    cancelF,
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	// TODO 写执行日志进入非关数据库
	fmt.Println(string(output))
	if isOk, err = mutex.Unlock(); err != nil || !isOk {
		global.SLogger.Errorw(
			"解锁失败",
			"task_name", taskName,
		)
		return
	}
	return
}

func KillTask(taskName string) error {
	var (
		task *common.RunningTask
		ok   bool
		//delResponse *etcd.DeleteResponse
		err error
	)
	if task, ok = global.RunningTaskMap[taskName]; !ok {
		return nil
	}
	task.Cancel()
	delete(global.CurrentExistTask, taskName)
	if _, err = global.KV.Delete(context.Background(), common.TaskKillDir+taskName); err != nil {
		return err
	}
	return nil
}
