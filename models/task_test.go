package models_test

import (
	"github.com/yobdc/jobs/models"
	"testing"
)

func TestNewtask(t *testing.T) {
	var t1 *models.Task
	t1 = models.NewTask("", "")
	if t1 != nil {
		t.Error("t1 = models.NewTask(\"\", \"\") => fail")
	}
}

func TestAddChild(t *testing.T) {
	var t1, t2, t3 *models.Task

	if _, ok := t1.AddChild(t2); ok == nil {
		t.Error("t1, t2 is empty, t1.AddChild(t2) should return error => fail")
	}

	t1 = models.NewTask("task1", "i am task1")
	if _, ok := t1.AddChild(t2); ok == nil {
		t.Error("t2 is empty, t1.AddChild(t2) should return error => fail")
	}

	t2 = models.NewTask("task2", "i am task2")
	if _, ok := t1.AddChild(t2); ok != nil {
		t.Error("t1.AddChild(t2) should not return error => fail")
	}

	t3 = models.NewTask("task3", "i am task3")
	t1.AddChild(t2)
	t2.AddChild(t3)
	if _, ok := t3.AddChild(t1); ok == nil {
		t.Error("t1.AddChild(t2), t2.AddChild(t3), t3.AddChild(t1) should return error => fail")
	}
}
