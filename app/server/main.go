package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// Simple hello world endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello World!",
		})
	})

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "healthy",
		})
	})

	log.Println("Server starting on port 8080...")
	log.Println("Try: curl http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
