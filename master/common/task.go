package common

import "time"

type TaskVar struct {
	Name     string    `json:"name"`
	CronExpr string    `json:"cron_expr"`
	Command  string    `json:"command"`
	NextTime time.Time `json:"next_time"`
}

type TaskRequest struct {
	TaskVar TaskVar `json:"task_var"`
}
