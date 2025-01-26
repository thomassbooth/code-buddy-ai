package server

import (
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func (s *Server) RegisterRoute(route Route) {
	// Register routes from all files in the given directory
}
