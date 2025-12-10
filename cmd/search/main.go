package main

import (
	"context"
	"flag"
	"log"

	"product_search_and_vectorize_service/app/services"
)

func main() {
	// Define flags
	search := flag.String("search", "", "Search query for products")
	userID := flag.Uint64("user_id", 0, "User ID to search products for")
	limit := flag.Uint64("limit", 5, "Maximum number of results to return")
	flag.Parse()

	// Check required arguments
	if *search == "" {
		log.Fatal("Error: -search argument is required. Usage: go run main.go -user_id 1 -search \"laptop\"")
	}
	if *userID == 0 {
		log.Fatal("Error: -user_id argument is required. Usage: go run main.go -user_id 1 -search \"laptop\"")
	}

	log.Println("Starting search product command")
	log.Printf("Search query: %s (user_id: %d, limit: %d)", *search, *userID, *limit)

	products, err := services.SearchProductsService(context.Background(), *userID, *search, *limit)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Found %d products:", len(products))
	for i, product := range products {
		log.Printf("  %d. %s (ID: %d, Price: $%.2f, Category: %s)", i+1, product.Name, product.ID, product.Price, product.Category)
	}
	log.Println("Search product command completed successfully")
}
