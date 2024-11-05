package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware logs each request method and URL.
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next(w, r)
	}
}
