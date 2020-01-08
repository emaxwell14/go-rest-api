package api

import "net/http"

func InitServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/tasks", loggerMiddleware(TasksGetAll()))
	mux.Handle("/tasks/", loggerMiddleware(TasksGetOne()))
	return mux
}
