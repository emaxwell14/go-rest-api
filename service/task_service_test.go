package service

import (
	"testing"
	"time"

	"github.com/emaxwell14/go-rest-api/model"
)

func TestAllTasks(t *testing.T) {
	CreateTask(model.Task{ID: 10})
	tasks := AllTasks()
	if len(tasks) != 1 {
		t.Errorf("Get all tasks returned incorrectly, got: %d, wanted: %d.", len(tasks), 1)
	}
}

func TestOneTask_empty(t *testing.T) {
	tempTasks = []model.Task{}
	task, ok := OneTask(0)
	if ok {
		t.Errorf("One task should return false when the task does not exist, got task with id: %d.", task.ID)
	}
}
func TestOneTask_taskDoesntExist(t *testing.T) {
	CreateTask(model.Task{ID: 10})
	task, ok := OneTask(111)
	if ok {
		t.Errorf("One task should return false when the task does not exist, got task with id: %d.", task.ID)
	}
}
func TestOneTask_taskExists(t *testing.T) {
	CreateTask(model.Task{})
	task, ok := OneTask(1)
	if !ok {
		t.Errorf("One task should return task with id: %d.", task.ID)
	}
}

func TestCreateTask_create(t *testing.T) {
	tempTasks = []model.Task{}
	task := model.Task{
		Name:        "TODO task",
		IsDone:      true,
		CreatedDate: time.Now().AddDate(0, -1, 0),
	}

	createdTask, ok := CreateTask(task)

	if !ok || task != createdTask {
		t.Errorf("Create task should create task correctly.\nExpected: %+v\nActual:   %+v", task, createdTask)
	}

}
