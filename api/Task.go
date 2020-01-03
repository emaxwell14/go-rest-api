package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/emaxwell14/go-rest-api/model"
)

var tempTask model.Task = model.Task{ID: 10}
var tempTasks = []model.Task{tempTask}

func getAllHandler(w http.ResponseWriter, _ *http.Request) {
	js, err := json.Marshal(tempTasks)
	if err != nil {
		log.Print("Error parsing json")
		w.Write([]byte("Error retreiving tasks."))
		w.WriteHeader(400)
	}
	w.Write(js)
}

func getOneHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if id == "" {
		w.Write([]byte("Id not found in query."))
		w.WriteHeader(400)
	}
	for _, v := range tempTasks {
		if idToInt, err := strconv.Atoi(id); err == nil && idToInt == v.ID {
			js, err := json.Marshal(v)
			if err != nil {
				log.Print("Error parsing json")
				w.Write([]byte("Error retreiving tasks."))
				w.WriteHeader(400)
			}
			w.Write(js)
			return
		}
	}
	io.WriteString(w, "Could not find task.")
}

/*
Tasks All tasks endpoints
*/
func Tasks(mux *http.ServeMux) {
	mux.HandleFunc("/tasks", getAllHandler)
	mux.HandleFunc("/tasks/", getOneHandler)
}
