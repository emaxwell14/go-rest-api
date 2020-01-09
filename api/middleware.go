package api

import (
	"fmt"
	"net/http"
	"time"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("%s: %s request recieved on endpoint: %s\n", timestamp, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
