package api

import "net/http"

// InitServer creates a server and adds routes
func InitServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/tasks", loggerMiddleware(TasksGetAll()))
	mux.Handle("/tasks/", loggerMiddleware(TasksGetOne()))
	return mux
}
