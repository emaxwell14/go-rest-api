package service

import (
	"testing"

	"github.com/emaxwell14/go-rest-api/model"
)

func TestAllTasks_empty(t *testing.T) {
	AddTask(model.Task{ID: 10})
	tasks := AllTasks()
	if len(tasks) != 1 {
		t.Errorf("Get all tasks returned incorrectly, got: %d, wanted: %d.", len(tasks), 1)
	}
}
