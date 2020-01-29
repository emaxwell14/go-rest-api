package main

import (
	"log"

	"github.com/emaxwell14/go-rest-api/api"
	"github.com/emaxwell14/go-rest-api/model"
	"github.com/emaxwell14/go-rest-api/service"
)

func main() {
	server := api.InitServer()
	// TODO remove after testing
	service.CreateTask(model.Task{})
	service.CreateTask(model.Task{})
	log.Println("REST server starting on:", server.Addr)
	log.Fatal(server.ListenAndServe())
}
