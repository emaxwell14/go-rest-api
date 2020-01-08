package main

import (
	"log"
	"net/http"

	"github.com/emaxwell14/go-rest-api/model"
	"github.com/emaxwell14/go-rest-api/service"

	"github.com/emaxwell14/go-rest-api/api"
)

func main() {
	// Mock data on startup
	service.AddTask(model.Task{ID: 10})

	server := api.InitServer()
	serveAddr := ":8080"

	log.Println("REST server starting on " + serveAddr)
	err := http.ListenAndServe(serveAddr, server)
	log.Fatal(err)
}
