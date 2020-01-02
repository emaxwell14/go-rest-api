package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("Request recived")
		io.WriteString(w, "Hello World")
	})
	http.ListenAndServe(":8080", nil)
}
