package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/emaxwell14/go-rest-api/model"
)

var tempTask model.Task = model.Task{ID: 10}
var tempTasks = []model.Task{tempTask}

/*
Tasks All tasks endpoints
*/
func Tasks() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, _ *http.Request) {
		js, err := json.Marshal(tempTasks)
		if err != nil {
			log.Panic("Error parsing json")
		}
		w.Write(js)
	})
}
