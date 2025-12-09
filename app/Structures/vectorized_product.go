package Structures

// VectorizedProduct contains the original product with its embedding
type VectorizedProduct struct {
	Product   Product   `json:"product"`
	Embedding []float32 `json:"embedding"`
}
