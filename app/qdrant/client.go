package qdrant

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"product_search_and_vectorize_service/app/Structures"

	"github.com/qdrant/go-client/qdrant"
)

const (
	DefaultVectorSize = 1536 // text-embedding-3-small dimension
)

// GetVectorSize returns the vector size from env or default
func GetVectorSize() uint64 {
	sizeStr := os.Getenv("OPENAI_EMBEDDING_DIMENSION")
	if sizeStr == "" {
		return DefaultVectorSize
	}
	size, err := strconv.ParseUint(sizeStr, 10, 64)
	if err != nil {
		return DefaultVectorSize
	}
	return size
}

// Client wraps the Qdrant client
type Client struct {
	client *qdrant.Client
}

// NewClient creates a new Qdrant client
func NewClient() (*Client, error) {
	host := os.Getenv("QDRANT_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("QDRANT_PORT")
	if port == "" {
		port = "6334"
	}

	client, err := qdrant.NewClient(&qdrant.Config{
		Host: host,
		Port: 6334,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create qdrant client: %w", err)
	}

	return &Client{client: client}, nil
}

// GetCollectionName returns the collection name for a specific user
func GetCollectionName(userID uint64) string {
	return fmt.Sprintf("products_user_%d", userID)
}

// CreateCollection creates the products collection for a user if it doesn't exist
func (c *Client) CreateCollection(ctx context.Context, userID uint64) error {
	collectionName := GetCollectionName(userID)

	exists, err := c.client.CollectionExists(ctx, collectionName)
	if err != nil {
		return fmt.Errorf("failed to check collection: %w", err)
	}

	if exists {
		return nil
	}

	err = c.client.CreateCollection(ctx, &qdrant.CreateCollection{
		CollectionName: collectionName,
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     GetVectorSize(),
			Distance: qdrant.Distance_Cosine,
		}),
	})
	if err != nil {
		return fmt.Errorf("failed to create collection: %w", err)
	}

	return nil
}

// UpsertProduct adds or updates a single vectorized product in Qdrant
func (c *Client) UpsertProduct(ctx context.Context, userID uint64, product Structures.VectorizedProduct) error {
	collectionName := GetCollectionName(userID)

	points := []*qdrant.PointStruct{
		{
			Id:      qdrant.NewIDNum(product.Product.ID),
			Vectors: qdrant.NewVectors(product.Embedding...),
			Payload: qdrant.NewValueMap(map[string]any{
				"id":          product.Product.ID,
				"user_id":     product.Product.UserID,
				"name":        product.Product.Name,
				"description": product.Product.Description,
				"price":       product.Product.Price,
				"category":    product.Product.Category,
			}),
		},
	}

	_, err := c.client.Upsert(ctx, &qdrant.UpsertPoints{
		CollectionName: collectionName,
		Points:         points,
	})
	if err != nil {
		return fmt.Errorf("failed to upsert product: %w", err)
	}

	return nil
}

// UpsertProducts adds or updates multiple vectorized products in Qdrant
func (c *Client) UpsertProducts(ctx context.Context, userID uint64, products []Structures.VectorizedProduct) error {
	collectionName := GetCollectionName(userID)

	points := make([]*qdrant.PointStruct, len(products))

	for i, product := range products {
		points[i] = &qdrant.PointStruct{
			Id:      qdrant.NewIDNum(product.Product.ID),
			Vectors: qdrant.NewVectors(product.Embedding...),
			Payload: qdrant.NewValueMap(map[string]any{
				"id":          product.Product.ID,
				"user_id":     product.Product.UserID,
				"name":        product.Product.Name,
				"description": product.Product.Description,
				"price":       product.Product.Price,
				"category":    product.Product.Category,
			}),
		}
	}

	_, err := c.client.Upsert(ctx, &qdrant.UpsertPoints{
		CollectionName: collectionName,
		Points:         points,
	})
	if err != nil {
		return fmt.Errorf("failed to upsert products: %w", err)
	}

	return nil
}

// SearchSimilar searches for similar products based on a query vector
func (c *Client) SearchSimilar(ctx context.Context, userID uint64, queryVector []float32, limit uint64) ([]Structures.Product, error) {
	collectionName := GetCollectionName(userID)

	results, err := c.client.Query(ctx, &qdrant.QueryPoints{
		CollectionName: collectionName,
		Query:          qdrant.NewQuery(queryVector...),
		Limit:          qdrant.PtrOf(limit),
		WithPayload:    qdrant.NewWithPayload(true),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to search: %w", err)
	}

	products := make([]Structures.Product, len(results))
	for i, result := range results {
		payload := result.Payload
		products[i] = Structures.Product{
			ID:          uint64(payload["id"].GetIntegerValue()),
			UserID:      uint64(payload["user_id"].GetIntegerValue()),
			Name:        payload["name"].GetStringValue(),
			Description: payload["description"].GetStringValue(),
			Price:       payload["price"].GetDoubleValue(),
			Category:    payload["category"].GetStringValue(),
		}
	}

	return products, nil
}

// Close closes the Qdrant client connection
func (c *Client) Close() error {
	return c.client.Close()
}
