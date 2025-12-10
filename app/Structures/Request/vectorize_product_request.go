package Request

// VectorizeProductRequest represents the request body for vectorizing a product
type VectorizeProductRequest struct {
	UserID      uint64   `json:"user_id"`
	ProductID   uint64   `json:"product_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Category    *string  `json:"category"` // nullable
	Price       *float64 `json:"price"`    // nullable
}
