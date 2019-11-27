package models

// TaskSchedule 任务调度设置
type TaskSchedule struct {
	task *Task
	cron string
}
