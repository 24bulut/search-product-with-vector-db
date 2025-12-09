package services

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"

	"product_search_and_vectorize_service/app/Structures"
)

// getProjectRoot returns the project root directory
func getProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	// Go up from app/services/ to project root
	return filepath.Join(filepath.Dir(b), "..", "..")
}

func GetProductsService(ctx context.Context) ([]Structures.Product, error) {
	filePath := filepath.Join(getProjectRoot(), "example_products.json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var products []Structures.Product
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, err
	}
	return products, nil
}
