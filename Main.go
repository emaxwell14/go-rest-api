package main

import (
	"log"

	"github.com/emaxwell14/go-rest-api/api"
)

func main() {
	server := api.InitServer()

	log.Println("REST server starting on:", server.Addr)
	log.Fatal(server.ListenAndServe())
}
