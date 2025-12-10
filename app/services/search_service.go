package services

import (
	"context"
	"fmt"
	"log"

	"product_search_and_vectorize_service/app/Structures"
	"product_search_and_vectorize_service/app/openai"
	"product_search_and_vectorize_service/app/qdrant"
)

func SearchProductsService(ctx context.Context, userID uint64, query string, limit uint64) ([]Structures.Product, error) {
	log.Printf("Searching for products with query: %s (user_id: %d)", query, userID)

	// Create OpenAI client to vectorize query
	openaiClient, err := openai.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create openai client: %w", err)
	}

	// Vectorize the search query
	queryResult, err := openaiClient.VectorizePlainText(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to vectorize query: %w", err)
	}
	log.Printf("Query vectorized, dimension: %d, tokens used: %d", len(queryResult.Embedding), queryResult.Usage.TotalTokens)

	// Create Qdrant client to search
	qdrantClient, err := qdrant.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create qdrant client: %w", err)
	}
	defer qdrantClient.Close()

	// Search similar products in user's collection
	products, err := qdrantClient.SearchSimilar(ctx, userID, queryResult.Embedding, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to search products: %w", err)
	}

	return products, nil
}
