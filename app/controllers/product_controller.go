package controllers

import (
	"encoding/json"
	"net/http"

	"product_search_and_vectorize_service/app/Structures"
	"product_search_and_vectorize_service/app/Structures/Request"
	"product_search_and_vectorize_service/app/Structures/Response"
	"product_search_and_vectorize_service/app/services"
	"product_search_and_vectorize_service/app/utils"
)

// ProductController handles product-related HTTP requests
type ProductController struct{}

// NewProductController creates a new ProductController
func NewProductController() *ProductController {
	return &ProductController{}
}

// VectorizeProduct handles POST request to vectorize a product
func (c *ProductController) VectorizeProduct(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		utils.SendError(w, http.StatusMethodNotAllowed, "Method not allowed. Use POST.")
		return
	}

	// Parse request body
	var req Request.VectorizeProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid JSON body: "+err.Error())
		return
	}

	// Validate required fields
	if req.UserID == 0 {
		utils.SendError(w, http.StatusBadRequest, "user_id is required")
		return
	}
	if req.ProductID == 0 {
		utils.SendError(w, http.StatusBadRequest, "product_id is required")
		return
	}
	if req.Name == "" {
		utils.SendError(w, http.StatusBadRequest, "name is required")
		return
	}
	if req.Description == "" {
		utils.SendError(w, http.StatusBadRequest, "description is required")
		return
	}

	// Build product struct
	product := Structures.Product{
		ID:          req.ProductID,
		UserID:      req.UserID,
		Name:        req.Name,
		Description: req.Description,
	}

	// Set optional fields if provided
	if req.Category != nil {
		product.Category = *req.Category
	}
	if req.Price != nil {
		product.Price = *req.Price
	}

	// Vectorize product
	vectorizedProduct, err := services.VectorizeProductService(r.Context(), product)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to vectorize product: "+err.Error())
		return
	}

	// Send success response
	utils.SendJSON(w, http.StatusOK, Response.VectorizeProductResponse{
		Success: true,
		Message: "Product vectorized successfully",
		Data:    vectorizedProduct,
	})
}
