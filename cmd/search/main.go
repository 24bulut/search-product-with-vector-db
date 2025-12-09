package main

import (
	"context"
	"flag"
	"log"
	"product_search_and_vectorize_service/app/services"
)

func main() {
	// Define search flag
	search := flag.String("search", "", "Search query for products")
	flag.Parse()

	// Check if search argument is provided
	if *search == "" {
		log.Fatal("Error: -search argument is required. Usage: go run main.go -search \"laptop\"")
	}

	log.Println("Starting search product command")
	log.Printf("Search query: %s", *search)
	products, err := services.SearchProductsService(context.Background(), *search)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("Products:", products)
	log.Println("Search product command completed successfully")
}
