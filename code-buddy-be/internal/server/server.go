package server

import (
	"fmt"
	"net/http"

	"github.com/thomassbooth/code-buddy-be/api/auth"
	"github.com/thomassbooth/code-buddy-be/config"
)

type Server struct {
	Router   *http.ServeMux
	server   *http.Server
	shutdown chan struct{}
}

func NewServer() *Server {
	mux := http.NewServeMux()
	mux.Handle("/", UserAuthenticator(auth.Route))
	return &Server{
		Router:   mux,
		shutdown: make(chan struct{}),
	}
}

func (s *Server) Start() error {
	port := config.AppConfig.Server.Port
	addr := ":" + port

	s.server = &http.Server{
		Addr:    addr,
		Handler: s.Router,
	}
	fmt.Printf("Server listening on port %s...\n", port)
	err := s.server.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Close() error {
	fmt.Println("Server stopped gracefully")
	return nil
}
