package models

// TaskResult 任务执行结果
type TaskResult struct {
	code TaskResultStatus
	msg  string
	task *Task
}

// TaskResultStatus 定义任务执行结果状态码
type TaskResultStatus int

// TaskResultStatus任务状态常量定义
const (
	OK TaskResultStatus = iota + 1
	Failed
	Cancelled
)
