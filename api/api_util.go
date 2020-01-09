package api

import (
	"fmt"
	"net/http"
	"time"
)

// SetErrorResponse uses default status if msg string is empty
func SetErrorResponse(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	if msg == "" {
		msg = http.StatusText(status)
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s: error response sent: \"%s\"\n", timestamp, msg)
	w.Write([]byte(msg))
}
