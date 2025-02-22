package repository

import (
	productsModel "ejercicioScaffolding/pkg/models"
	"encoding/json"
	"errors"
	"os"
)

type ProductRepository interface {
	FindByID(id int) (*productsModel.Product, error)
	FindAll() ([]productsModel.Product, error)
	FindByPriceGreaterThan(price int) ([]productsModel.Product, error)
	Save(product productsModel.Product) error
}

type MapProductRepository struct {
	filePath string
}

func NewMapProductRepository(filePath string) ProductRepository {
	return &MapProductRepository{filePath: filePath}
}

func (r *MapProductRepository) loadProducts() ([]productsModel.Product, error) {
	file, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var products []productsModel.Product
	err = json.Unmarshal(file, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *MapProductRepository) saveProducts(products []productsModel.Product) error {
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}

	err = os.WriteFile(r.filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r *MapProductRepository) FindByID(id int) (*productsModel.Product, error) {
	products, err := r.loadProducts()
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		if product.ID == id {
			return &product, nil
		}
	}

	return nil, errors.New("product not found")
}

func (r *MapProductRepository) FindAll() ([]productsModel.Product, error) {
	return r.loadProducts()
}

func (r *MapProductRepository) FindByPriceGreaterThan(price int) ([]productsModel.Product, error) {
	products, err := r.loadProducts()
	if err != nil {
		return nil, err
	}

	var filteredProducts []productsModel.Product
	for _, product := range products {
		if product.Price > float64(price) {
			filteredProducts = append(filteredProducts, product)
		}
	}

	return filteredProducts, nil
}

func (r *MapProductRepository) Save(product productsModel.Product) error {
	products, err := r.loadProducts()
	if err != nil {
		return err
	}

	for _, p := range products {
		if p.CodeValue == product.CodeValue {
			return errors.New("code value must be unique")
		}
	}

	maxID := 0
	for _, p := range products {
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	product.ID = maxID + 1

	products = append(products, product)

	return r.saveProducts(products)
}
