package services

import (
	"context"
	"product_search_and_vectorize_service/app/Structures"
	"product_search_and_vectorize_service/app/openai"
)

func VectorizeProductsService(ctx context.Context, products []Structures.Product) ([]Structures.VectorizedProduct, error) {
	client, err := openai.NewClient()
	if err != nil {
		return nil, err
	}

	var vectorizedProducts []Structures.VectorizedProduct
	for _, product := range products {
		vectorizedProduct, err := client.VectorizeProduct(ctx, product)
		if err != nil {
			return nil, err
		}
		vectorizedProducts = append(vectorizedProducts, *vectorizedProduct)
	}
	return vectorizedProducts, nil
}
