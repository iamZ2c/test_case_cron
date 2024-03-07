package main

import (
	initialize "cron_tab_c/slave/initalize"
	"cron_tab_c/slave/schedule"
)

func main() {
	var (
		err error
	)
	initialize.Logger()

	initialize.LoadEtcdClient()

	if err = initialize.ExistTask(); err != nil {
		panic(err)
	}

	initialize.RunningTask()

	initialize.RedisClient()

	go func() {
		initialize.ScheduleLoop()
	}()
	schedule.ListenEtcdStatusAndHandleSchedule()

}
