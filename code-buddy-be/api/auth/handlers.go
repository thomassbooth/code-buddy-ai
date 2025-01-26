package auth

import (
	"log"
	"net/http"
)

var Route = http.HandlerFunc(GetUsers)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Print("Executing finalHandler")
	w.Write([]byte("OK"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Handler logic
}
