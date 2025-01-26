package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/thomassbooth/code-buddy-be/internal/server"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	server := server.NewServer()
	go func() {
		if err := server.Start(); err != nil {
			// this closes our application on Fatal error
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	sig := <-signals
	log.Printf("Received signal: %v. Shutting down...", sig)
	server.Close()

}
