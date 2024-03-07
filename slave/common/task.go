package common

import (
	"golang.org/x/net/context"
	"time"
)

type TaskVar struct {
	Name     string    `json:"name"`
	CronExpr string    `json:"cron_expr"`
	Command  string    `json:"command"`
	NextTime time.Time `json:"next_time,omitempty"`
}

type TaskRequest struct {
	TaskVar TaskVar `json:"task_var"`
}

type RunningTask struct {
	Name      string
	StartTime time.Time
	Cancel    context.CancelFunc
}
