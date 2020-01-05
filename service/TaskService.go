package service

import (
	"github.com/emaxwell14/go-rest-api/model"
)

var tempTask model.Task = model.Task{ID: 10}
var tempTasks = []model.Task{tempTask}

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
