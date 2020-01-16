package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/emaxwell14/go-rest-api/model"
	"github.com/emaxwell14/go-rest-api/service"
	"github.com/gorilla/mux"
)

// TasksGetAll returns all of the tasks
// TODO: could be improved with pagination etc.
func TasksGetAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		tasks := service.AllTasks()
		js, err := json.Marshal(tasks)
		if err != nil {
			SetErrorResponse(w, http.StatusInternalServerError, "")
			return
		}
		w.Write(js)
	})
}

// TasksGetOne returns a single task based on the id in the path
func TasksGetOne() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idString := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idString)
		if err != nil {
			SetErrorResponse(w, http.StatusBadRequest, "Could not find id in request")
			return
		}

		task, ok := service.OneTask(id)
		if !ok {
			SetErrorResponse(w, http.StatusNotFound, "Task with id not found")
			return
		}

		js, jsonErr := json.Marshal(task)
		if jsonErr != nil {
			SetErrorResponse(w, http.StatusInternalServerError, "")
			return
		}
		w.Write(js)
	})
}

// TasksCreate adds a new task and returns the task in the response
func TasksCreate() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		task := model.Task{}
		err := decoder.Decode(&task)
		if err != nil {
			// TODO: how to validate the task fields
			SetErrorResponse(w, http.StatusBadRequest, "Validation error on task")
			fmt.Println(err)
			return
		}

		task, ok := service.CreateTask(task)
		if !ok {
			SetErrorResponse(w, http.StatusInternalServerError, "Unable to create task")
			return
		}

		js, err := json.Marshal(task)
		if err != nil {
			SetErrorResponse(w, http.StatusInternalServerError, "")
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(js)
	})
}

// TasksUpdate updates an existing task
// Returns 200 and the resource if successful
// Returns 204 if the resource does not exist
func TasksUpdate() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idString := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idString)
		if err != nil {
			SetErrorResponse(w, http.StatusBadRequest, "Could not find id in request")
			return
		}

		decoder := json.NewDecoder(r.Body)
		task := model.Task{}
		err = decoder.Decode(&task)
		if err != nil {
			// TODO: how to validate the task fields
			SetErrorResponse(w, http.StatusBadRequest, "Validation error on task")
			fmt.Println(err)
			return
		}

		task.ID = id // ensure that the id from the url is set
		ok := service.UpdateTask(task)
		if !ok {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		js, err := json.Marshal(task)
		if err != nil {
			SetErrorResponse(w, http.StatusInternalServerError, "")
			return
		}
		w.Write(js)
	})
}

// TasksDelete removes a task
// Returns 204 if successful
// Returns 404 if resource does not exist
func TasksDelete() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idString := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idString)
		if err != nil {
			SetErrorResponse(w, http.StatusBadRequest, "Could not find id in request")
			return
		}

		ok := service.DeleteTask(id)
		if !ok {
			SetErrorResponse(w, http.StatusNotFound, "Unable to find task to delete")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}
