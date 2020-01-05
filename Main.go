package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/emaxwell14/go-rest-api/model"
	"github.com/emaxwell14/go-rest-api/service"

	"github.com/emaxwell14/go-rest-api/api"
)

// Cant get this working
// func loggerMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		timestamp := time.Now().Format("2006-01-02 15:04:05")
// 		fmt.Printf("%s: %s request recieved on endpoint %s\n", timestamp, r.Method, r.URL.Path)
// 		next.ServeHTTP(w, r)
// 	})
// }

func loggerFunc(_ http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s: %s request recieved on endpoint %s\n", timestamp, r.Method, r.URL.Path)
}

func main() {
	service.AddTask(model.Task{ID: 10})

	mux := http.NewServeMux()
	api.Tasks(mux)

	mux.HandleFunc("/", loggerFunc) // Only happens on exact path
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
