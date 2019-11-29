package models

// TaskResultStatus 定义任务执行结果状态码
type TaskResultStatus int

// TaskResultStatus任务状态常量定义
const (
	None TaskResultStatus = iota
	OK
	Failed
	Cancelled
)

// TaskResult 任务执行结果
type TaskResult struct {
	code TaskResultStatus
	msg  string
}

// IsOK 判断任务结果是否成功
func (result *TaskResult) IsOK() bool {
	return result.code == OK
}

// IsFailed 判断任务结果是否失败
func (result *TaskResult) IsFailed() bool {
	return result.code == Failed
}

// IsCancelled 判断任务结果是否成功
func (result *TaskResult) IsCancelled() bool {
	return result.code == Cancelled
}

// TaskOK 定义成功的任务执行结果
func TaskOK(msg string) TaskResult {
	return TaskResult{
		code: OK,
		msg:  msg,
	}
}

// TaskFailed 定义失败的任务执行结果
func TaskFailed(msg string) TaskResult {
	return TaskResult{
		code: Failed,
		msg:  msg,
	}
}

// TaskCancelled 定义取消的任务执行结果
func TaskCancelled(msg string) TaskResult {
	return TaskResult{
		code: Cancelled,
		msg:  msg,
	}
}
