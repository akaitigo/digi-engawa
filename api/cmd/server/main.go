package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/akaitigo/digi-engawa/api/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := handler.NewRouter()

	log.Printf("Starting server on :%s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
