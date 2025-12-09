package services

import (
	"context"
	"log"
	"product_search_and_vectorize_service/app/Structures"
	"product_search_and_vectorize_service/app/openai"
)

func SearchProductsService(ctx context.Context, query string) ([]Structures.Product, error) {
	log.Println("Searching for products with query:", query)
	client, err := openai.NewClient()
	if err != nil {
		return nil, err
	}
	vectorizedQuery, err := client.VectorizePlainText(ctx, query)
	if err != nil {
		return nil, err
	}
	log.Println("Vectorized query:", vectorizedQuery)

	return nil, nil
}
