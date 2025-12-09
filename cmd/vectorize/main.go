package main

import (
	"context"
	"log"

	"product_search_and_vectorize_service/app/services"
)

func vectorizeProducts(ctx context.Context) error {
	products, err := services.GetProductsService(ctx)
	log.Println("Products:", products)
	if err != nil {
		return err
	}
	vectorizedProducts, err := services.VectorizeProductsService(ctx, products)
	log.Println("Vectorized products:", vectorizedProducts)
	if err != nil {
		return err
	}
	log.Println("Vectorized products:", vectorizedProducts)
	return nil
}

func main() {
	log.Println("Starting vectorize product command")
	err := vectorizeProducts(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("Vectorize products command completed successfully")
}
