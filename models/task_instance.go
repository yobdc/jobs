package models

import (
	"log"
	"strings"
	"time"
)

// TaskInstance 任务执行的实例对象
type TaskInstance struct {
	task            *Task
	result          TaskResult
	childInstances  []*TaskInstance
	parentInstances map[*TaskInstance]bool
	env             map[string]string
}

// IsReady 检查TaskInstance所有的父任务是否已完成
func (taskInstance *TaskInstance) IsReady() bool {
	result := true
	for parentItem := range taskInstance.parentInstances {
		result = result && taskInstance.parentInstances[parentItem]
		if !result {
			return false
		}
	}
	return result
}

// Start 启动任务实例执行
func (taskInstance *TaskInstance) Start() {
	out := make(chan TaskResult)
	defer close(out)

	go taskInstance.Exec(out)
	execResult := <-out

	for _, childInstance := range taskInstance.childInstances {
		if execResult.IsOK() {
			childInstance.parentInstances[taskInstance] = true
		}
		if childInstance.IsReady() {
			go childInstance.Start()
		}
	}
}

// Exec 任务执行
func (taskInstance *TaskInstance) Exec(out chan<- TaskResult) {
	log.Println(taskInstance.task.Name, "start")
	log.Println(taskInstance.task)
	if strings.Contains(taskInstance.task.Name, "2") {
		time.Sleep(time.Second * 3)
	}
	if strings.Contains(taskInstance.task.Name, "3") {
		time.Sleep(time.Second * 1)
	}
	log.Println(taskInstance.task.Name, "end")
	out <- TaskOK("success")
}
