package models

import (
	"errors"
	"github.com/satori/go.uuid"
	"time"
)

// Task 用户任务
type Task struct {
	ID          uuid.UUID
	Name        string
	Desc        string
	Cmd         string
	childTasks  []*Task // 子任务
	parentTasks []*Task // 父任务
}

// NewTask 创建任务并初始化
func NewTask(name, desc, cmd string) *Task {
	if name == "" {
		return nil
	}
	task := new(Task)
	task.ID = uuid.NewV4()
	task.Name = name
	task.Desc = desc
	task.Cmd = cmd
	task.childTasks = make([]*Task, 0)
	task.parentTasks = make([]*Task, 0)
	return task
}

// NewInstance 创建任务实例
func (task *Task) NewInstance() *TaskInstance {
	allInstances := make(map[*Task]*TaskInstance)
	return task.makeInstance(allInstances)
}
func (task *Task) makeInstance(allInstances map[*Task]*TaskInstance) *TaskInstance {
	if _, ok := allInstances[task]; ok {
		return allInstances[task]
	}
	instance := &TaskInstance{}
	instance.task = task
	instance.childInstances = make([]*TaskInstance, len(task.childTasks))
	instance.parentInstances = make(map[*TaskInstance]bool)
	instance.createTime = time.Now()
	allInstances[task] = instance
	for i := range task.childTasks {
		instance.childInstances[i] = task.childTasks[i].makeInstance(allInstances)
		instance.childInstances[i].parentInstances[instance] = false
	}
	return instance
}

// AddChild 添加子任务
func (task *Task) AddChild(child *Task) (*Task, error) {
	if task == nil {
		return task, errors.New("task is nil, cannot add child")
	}
	if child == nil {
		return task, errors.New("child task is nil, cannot be added")
	}
	for _, item := range task.childTasks {
		if item.ID == child.ID {
			return task, errors.New("child task already exists")
		}
	}
	if checkTaskCircle(task, child) {
		return task, errors.New("child task is nil, task circle exixts")
	}
	task.childTasks = append(task.childTasks, child)
	child.parentTasks = append(child.parentTasks, task)
	return task, nil
}

// checkTaskCircle 在parent.AddChild(child)之前检查是否会导致任务循环依赖
func checkTaskCircle(parent, child *Task) bool {
	if parent == nil || child == nil {
		return false
	}
	if parent.ID == child.ID {
		return true
	}
	for _, item := range child.childTasks {
		if checkTaskCircle(parent, item) {
			return true
		}
	}
	return false
}
