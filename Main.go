package main

import (
	"log"
	"net/http"

	"github.com/emaxwell14/go-rest-api/api"
)

func main() {
	mux := http.NewServeMux()
	api.Tasks(mux)

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
