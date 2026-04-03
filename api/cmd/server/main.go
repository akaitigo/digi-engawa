package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/akaitigo/digi-engawa/api/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = "./data"
	}

	h, err := handler.NewRouter(dataDir)
	if err != nil {
		log.Fatalf("Failed to create router: %v", err)
	}

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           h,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	log.Printf("Starting server on :%s", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
