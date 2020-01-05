package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/emaxwell14/go-rest-api/service"
)

func getAllHandler(w http.ResponseWriter, _ *http.Request) {
	tasks := service.AllTasks()
	js, err := json.Marshal(tasks)
	if err != nil {
		SetErrorResponse(w, http.StatusInternalServerError, "")
		return
	}
	w.Write(js)
}

func getOneHandler(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/tasks/") // TODO this wont handle longer paths
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
}

/*
Tasks All tasks endpoints
*/
func Tasks(mux *http.ServeMux) {
	mux.HandleFunc("/tasks", getAllHandler)
	mux.HandleFunc("/tasks/", getOneHandler)
}
