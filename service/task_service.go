package service

import (
	"github.com/emaxwell14/go-rest-api/model"
)

var tempTasks = []model.Task{}

// AllTasks gets every task that has been created
func AllTasks() []model.Task {
	return tempTasks
}

// OneTask gets a task by id. If it doesnt exist returns empty object and boolean false
func OneTask(id int) (model.Task, bool) {
	for _, v := range tempTasks {
		if id == v.ID {
			return v, true
		}
	}
	return model.Task{}, false
}

// CreateTask adds a new task and returns task and ok
func CreateTask(task model.Task) (model.Task, bool) {
	task.ID = len(tempTasks) // TODO: maybe use uuid?
	tempTasks = append(tempTasks, task)
	return task, true // return ok as other impl may be a db and could have an issue
}
