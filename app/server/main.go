package main

import (
	"encoding/json"
	"log"
	"net/http"

	"product_search_and_vectorize_service/app/controllers"
)

func main() {
	// Initialize controllers
	productController := controllers.NewProductController()

	// Simple hello world endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Product Search and Vectorize Service",
		})
	})

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "healthy",
		})
	})

	// Product endpoints
	http.HandleFunc("/api/products/vectorize", productController.VectorizeProduct)

	log.Println("Server starting on port 8091...")
	log.Println("Endpoints:")
	log.Println("  GET  /           - Service info")
	log.Println("  GET  /health     - Health check")
	log.Println("  POST /api/products/vectorize - Vectorize a product")
	log.Fatal(http.ListenAndServe(":8091", nil))
}
