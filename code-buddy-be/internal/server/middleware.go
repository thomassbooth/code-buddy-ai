package server

import (
	"fmt"
	"net/http"
)

func UserAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic
		fmt.Println("Authenticating user")
		next.ServeHTTP(w, r)
	})
}
