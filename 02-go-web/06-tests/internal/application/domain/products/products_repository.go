package domain

import products "ejercicioTest/pkg/models"

type productsRepository interface {
	GetProductByIDHandler(id int) (products.Product, error)
	GetAllProductsHandler() ([]products.Product, error)
	SearchProductsHandler() ([]products.Product, error)
	AddProductHandler() (products.Product, error)
}
