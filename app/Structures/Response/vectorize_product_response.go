package Response

import "product_search_and_vectorize_service/app/Structures"

// VectorizeProductResponse represents the response for vectorizing a product
type VectorizeProductResponse struct {
	Success bool                          `json:"success"`
	Message string                        `json:"message"`
	Data    *Structures.VectorizedProduct `json:"data,omitempty"`
}
