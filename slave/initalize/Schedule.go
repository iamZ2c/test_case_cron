package initialize

import (
	"cron_tab_c/slave/executor"
	"cron_tab_c/slave/global"
	"cron_tab_c/slave/utils"
	"github.com/gorhill/cronexpr"
	"time"
)

func ScheduleLoop() {
	var (
		restTime  time.Duration
		restTimer *time.Timer
	)
	global.SLogger.Info("开始加载调度器")
	global.ListenChan = make(chan int, 1)
	for {

		for _, task := range global.CurrentExistTask {
			now := time.Now()
			if task.NextTime.Before(now) || task.NextTime.Equal(now) {
				expr, _ := cronexpr.Parse(task.CronExpr)
				task.NextTime = expr.Next(now)
				//fmt.Println(fmt.Sprintf("%v,%v", task.Name, task.NextTime))
				// TODO 发送消息给 任务执行器，执行任务操作，编写调度器
				go func() {
					executor.ExecuteTask(task.Name)
				}()
			}
		}
		// 防止过快遍历对服务器产生过大压力
		restTime = utils.CountRecentTime()
		restTimer = time.NewTimer(restTime)
		// 有变动和执行都需要重新计算最近时间
		select {
		case <-restTimer.C:

			global.SLogger.Infow(
				"调度器心跳&当前任务有:",
				"task", global.CurrentExistTask,
			)
		case <-global.ListenChan:
			global.SLogger.Info("调度器收到消息")
		}
	}
}
