package services

import (
	"context"
	"encoding/json"

	"product_search_and_vectorize_service/app/Structures"
	"product_search_and_vectorize_service/app/utils"
)

func GetProductsService(ctx context.Context) ([]Structures.Product, error) {
	data, err := utils.ReadFile("example_products.json")
	if err != nil {
		return nil, err
	}
	var products []Structures.Product
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, err
	}
	return products, nil
}
