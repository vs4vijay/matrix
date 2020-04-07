package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"matrix/pkg/handlers"
	server "matrix/pkg/web"
)

var (
	logger *log.Logger
	port   string = os.Getenv("SERVER_PORT")
)

func init() {
	logger = log.New(os.Stdout, "matrix: ", log.LstdFlags|log.Lshortfile)
}

func main() {
	logger.Println("Initializing Matrix")

	mux := http.NewServeMux()

	h := handlers.NewHandler(logger)
	h.SetupRoutes(mux)

	if port == "" {
		port = "9090"
	}

	address := ":" + port
	srv := server.New(mux, address)

	go func () {
		logger.Printf("Starting Matrix Server at %s\n", port)
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatalf("Server failed to start: %v", err)
			os.Exit(1)
		}
	}()

	// Trap sigterm or interrupt to gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received.
	sig := <-sigChan
	logger.Printf("Got signal: %v, shutting down the server\n", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
}
