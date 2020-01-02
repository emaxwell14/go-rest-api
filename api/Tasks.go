package api

import (
	"io"
	"net/http"
)

/*
Tasks All tasks endpoints
*/
func Tasks() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "All Tasks")
	})
}
