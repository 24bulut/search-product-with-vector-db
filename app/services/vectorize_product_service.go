package services

import (
	"context"
	"fmt"
	"log"

	"product_search_and_vectorize_service/app/Structures"
	"product_search_and_vectorize_service/app/openai"
	"product_search_and_vectorize_service/app/qdrant"
)

func VectorizeProductService(ctx context.Context, product Structures.Product) (*Structures.VectorizedProduct, error) {
	// Create OpenAI client
	openaiClient, err := openai.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create openai client: %w", err)
	}

	// Create Qdrant client
	qdrantClient, err := qdrant.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create qdrant client: %w", err)
	}
	defer qdrantClient.Close()

	// Create collection for user if not exists
	if err := qdrantClient.CreateCollection(ctx, product.UserID); err != nil {
		return nil, fmt.Errorf("failed to create collection for user %d: %w", product.UserID, err)
	}

	// Vectorize product
	vectorizedProduct, err := openaiClient.VectorizeProduct(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("failed to vectorize product: %w", err)
	}

	// Store in Qdrant
	if err := qdrantClient.UpsertProduct(ctx, product.UserID, *vectorizedProduct); err != nil {
		return nil, fmt.Errorf("failed to store product in qdrant: %w", err)
	}

	log.Printf("Successfully stored product '%s' (ID: %d) in Qdrant collection 'products_user_%d'", product.Name, product.ID, product.UserID)
	return vectorizedProduct, nil
}
