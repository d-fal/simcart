package middleware

import (
	"net/http"
)

// MiddlewareExample example middleware for extra http calls
func MiddlewareExample(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
