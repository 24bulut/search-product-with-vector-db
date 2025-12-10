package Structures

// VectorizedProduct contains the original product with its embedding and usage info
type VectorizedProduct struct {
	Product        Product        `json:"product"`
	Embedding      []float32      `json:"embedding"`
	Model          string         `json:"model"`
	Usage          EmbeddingUsage `json:"usage"`
	StoredInQdrant bool           `json:"stored_in_qdrant"`
}
