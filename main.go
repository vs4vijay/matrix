package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vs4vijay/matrix/pkg/handlers"
	server "github.com/vs4vijay/matrix/pkg/web"
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

	logger.Printf("Starting Matrix Server at %s\n", port)

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
