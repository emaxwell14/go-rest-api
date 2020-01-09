package api

import (
	"encoding/json"
	"net/http"
	"strconv"

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
