package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"product_search_and_vectorize_service/app/Structures"
	"product_search_and_vectorize_service/app/services"
)

func findProduct(products []Structures.Product, userID, productID uint64) (*Structures.Product, error) {
	for _, product := range products {
		if product.UserID == userID && product.ID == productID {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("product not found: user_id=%d, product_id=%d", userID, productID)
}

func vectorizeProduct(ctx context.Context, userID, productID uint64) error {
	products, err := services.GetProductsService(ctx)
	if err != nil {
		return err
	}
	log.Printf("Loaded %d total products from file", len(products))

	// Find the specific product
	product, err := findProduct(products, userID, productID)
	if err != nil {
		return err
	}

	vectorizedProduct, err := services.VectorizeProductService(ctx, *product)
	if err != nil {
		return err
	}
	log.Printf("Vectorized product: %s (ID: %d)", vectorizedProduct.Product.Name, vectorizedProduct.Product.ID)
	return nil
}

func main() {
	// Define flags
	userID := flag.Uint64("user_id", 0, "User ID of the product owner")
	productID := flag.Uint64("product_id", 0, "Product ID to vectorize")
	flag.Parse()

	if *userID == 0 {
		log.Fatal("Error: -user_id argument is required. Usage: go run main.go -user_id 1 -product_id 1")
	}
	if *productID == 0 {
		log.Fatal("Error: -product_id argument is required. Usage: go run main.go -user_id 1 -product_id 1")
	}

	log.Printf("Starting vectorize product command for user %d, product %d", *userID, *productID)
	err := vectorizeProduct(context.Background(), *userID, *productID)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("Vectorize product command completed successfully")
}
