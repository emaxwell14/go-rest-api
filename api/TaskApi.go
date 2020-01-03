package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/emaxwell14/service"
)

func getAllHandler(w http.ResponseWriter, _ *http.Request) {
	tasks := service.AllTasks()
	js, err := json.Marshal(tasks)
	if err != nil {
		w.Write([]byte("Error retreiving tasks."))
		w.WriteHeader(400)
	}
	w.Write(js)
}

func getOneHandler(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/tasks/") // TODO this wont handle longer paths
	id, err := strconv.Atoi(idString)
	if err != nil {
		w.Write([]byte("Id not found in query"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task, ok := service.OneTask(id)
	if !ok {
		w.Write([]byte("Task not found"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	js, jsonErr := json.Marshal(task)
	if jsonErr != nil {
		w.Write([]byte("Error retrieving tasks."))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(js)
	return
}

/*
Tasks All tasks endpoints
*/
func Tasks(mux *http.ServeMux) {
	mux.HandleFunc("/tasks", getAllHandler)
	mux.HandleFunc("/tasks/", getOneHandler)
}
