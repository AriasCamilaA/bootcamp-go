package domain

import products "ejercicioScaffolding/pkg/models"

type productsService interface {
	GetProductByIDHandler(id int) (products.Product, error)
	GetAllProductsHandler() ([]products.Product, error)
	SearchProductsHandler() ([]products.Product, error)
	AddProductHandler() (products.Product, error)
}
