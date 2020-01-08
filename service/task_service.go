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

// AddTask will add a new task
// TODO: Decide later what happens for id. Add/edit etc
func AddTask(task model.Task) (model.Task, bool) {
	tempTasks = append(tempTasks, task)
	return task, true // return ok as other impl may be a db and could have an issue
}
