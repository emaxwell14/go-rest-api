package main

import (
	"io"
	"log"
	"net/http"

	"github.com/emaxwell14/go-rest-api/api"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		log.Print("Hello World request received.")
		io.WriteString(w, "Hello World")
	})
	api.Tasks()
	http.ListenAndServe(":8080", nil)
}
