package products

import (
	products "ejercicioTest/pkg/models"
	"encoding/json"
	"fmt"
	"os"
)

func LoadProducts(filename string) ([]products.Product, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var products []products.Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return products, nil
}
