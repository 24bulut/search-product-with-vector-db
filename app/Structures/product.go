package Structures

// Product represents a product to be vectorized
type Product struct {
	ID          uint64  `json:"id"`
	UserID      uint64  `json:"user_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}
