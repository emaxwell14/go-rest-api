package main

import (
	"log"

	"github.com/emaxwell14/go-rest-api/model"
	"github.com/emaxwell14/go-rest-api/service"

	"github.com/emaxwell14/go-rest-api/api"
)

func main() {
	// Mock data on startup
	service.AddTask(model.Task{ID: 10})

	server := api.InitServer()

	log.Println("REST server starting on", server.Addr)
	log.Fatal(server.ListenAndServe())
}
