package openai

import (
	"context"
	"fmt"
	"os"

	"product_search_and_vectorize_service/app/Structures"

	"product_search_and_vectorize_service/app/utils"

	"github.com/sashabaranov/go-openai"
)

// Client wraps the OpenAI client for embedding operations
type Client struct {
	client *openai.Client
	model  openai.EmbeddingModel
}

// NewClient creates a new OpenAI client
// Uses OPENAI_API_KEY environment variable by default
func NewClient() (*Client, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	return &Client{
		client: openai.NewClient(apiKey),
		model:  openai.SmallEmbedding3, // text-embedding-3-small (1536 dimensions)
	}, nil
}

// SetModel allows changing the embedding model
// Available models:
// - openai.SmallEmbedding3 (text-embedding-3-small) - 1536 dimensions
// - openai.LargeEmbedding3 (text-embedding-3-large) - 3072 dimensions
// - openai.AdaEmbeddingV2 (text-embedding-ada-002) - 1536 dimensions
func (c *Client) SetModel(model openai.EmbeddingModel) {
	c.model = model
}

// GetEmbedding creates an embedding for a single text
func (c *Client) GetEmbedding(ctx context.Context, text string) ([]float32, error) {
	resp, err := c.client.CreateEmbeddings(ctx, openai.EmbeddingRequest{
		Input: []string{text},
		Model: c.model,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create embedding: %w", err)
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no embedding returned")
	}

	if err := utils.LogEmbedding(c.model, text, resp); err != nil {
		return nil, fmt.Errorf("failed to log embedding: %w", err)
	}

	return resp.Data[0].Embedding, nil
}

// VectorizeProduct creates an embedding for a single product
// The embedding is created from the product's name and description
func (c *Client) VectorizeProduct(ctx context.Context, product Structures.Product) (*Structures.VectorizedProduct, error) {
	// Combine name and description for better semantic representation
	text := fmt.Sprintf("%s: %s (Category: %s)", product.Name, product.Description, product.Category)

	embedding, err := c.GetEmbedding(ctx, text)
	if err != nil {
		return nil, fmt.Errorf("failed to vectorize product %s: %w", product.Name, err)
	}

	return &Structures.VectorizedProduct{
		Product:   product,
		Embedding: embedding,
	}, nil
}
