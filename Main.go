package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/emaxwell14/go-rest-api/api"
)

func loggerMiddleware(_ http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s: %s request recieved on endpoint %s\n", timestamp, r.Method, r.URL.Path)
}

func main() {
	mux := http.NewServeMux()
	api.Tasks(mux)
	// Only happens on exact path
	mux.HandleFunc("/", loggerMiddleware)
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
