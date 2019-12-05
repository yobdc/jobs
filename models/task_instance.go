package models

import (
	"log"
	"os/exec"
	"time"
)

// TaskInstance 任务执行的实例对象
type TaskInstance struct {
	task            *Task
	result          TaskResult
	cmd             *exec.Cmd
	childInstances  []*TaskInstance
	parentInstances map[*TaskInstance]bool
	env             map[string]string
	createTime      time.Time
	startTime       time.Time
	endTime         time.Time
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
	taskInstance.startTime = time.Now()
	log.Println(taskInstance.task.Name, "start")
	defer func() {
		taskInstance.endTime = time.Now()
	}()

	taskInstance.cmd = exec.Command("bash", "-c", taskInstance.task.Cmd)
	err := taskInstance.cmd.Run()
	if err != nil {
		out <- TaskFailed(err.Error())
		log.Println(taskInstance.task.Name, "error:", err.Error())
		return
	}

	log.Println(taskInstance.task.Name, "end with success")
	out <- TaskOK("success")
}

// Stop 杀死执行任务进程
func (taskInstance *TaskInstance) Stop() {
	if taskInstance.cmd != nil {
		taskInstance.cmd.Process.Kill()
		log.Println(taskInstance.task.Name, "try to stop")
	} else {
		log.Println(taskInstance.task.Name, "not started")
	}
}

// ListInstances 列出任务实例自身及其说有子任务的实例
func (taskInstance *TaskInstance) ListInstances() []*TaskInstance {
	instanceMap := make(map[*TaskInstance]bool)
	taskInstance.allInstances(instanceMap)
	resultInstances := make([]*TaskInstance, len(instanceMap))
	i := 0
	for k := range instanceMap {
		resultInstances[i] = k
		i++
	}
	return resultInstances
}

func (taskInstance *TaskInstance) allInstances(instanceMap map[*TaskInstance]bool) {
	instanceMap[taskInstance] = true
	for i := range taskInstance.childInstances {
		taskInstance.childInstances[i].allInstances(instanceMap)
	}
}
