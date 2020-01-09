package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// InitServer creates a server and adds routes
func InitServer() *http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", TasksGetAll()).Methods("GET")
	r.HandleFunc("/tasks/{id}", TasksGetOne()).Methods("GET")
	r.HandleFunc("/tasks", TasksCreate()).Methods("POST")

	r.Use(loggerMiddleware)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv
}
