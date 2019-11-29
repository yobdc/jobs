package models_test

import (
	"github.com/yobdc/jobs/models"
	"testing"
)

func TestStart(t *testing.T) {
	t1 := models.NewTask("task1", "i am task1", "")
	ti1 := t1.NewInstance()
	ti1.Start()
}
