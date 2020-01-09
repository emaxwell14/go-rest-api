package api

import "net/http"

// SetErrorResponse uses default status if msg string is empty
func SetErrorResponse(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	if msg == "" {
		msg = http.StatusText(status)
	}
	w.Write([]byte(msg))
}
