package service

import (
	productsRepo "ejercicioScaffolding/internal/infraestructure/repository"
	productsModel "ejercicioScaffolding/pkg/models"
)

type ProductService interface {
	GetProductByID(id int) (*productsModel.Product, error)
	GetAllProducts() ([]productsModel.Product, error)
	SearchProducts(priceGt int) ([]productsModel.Product, error)
	AddProduct(product productsModel.Product) error
}

type DefaultProductService struct {
	repo productsRepo.ProductRepository
}

func NewProductService(repo productsRepo.ProductRepository) ProductService {
	return &DefaultProductService{repo: repo}
}

func (s *DefaultProductService) GetProductByID(id int) (*productsModel.Product, error) {
	return s.repo.FindByID(id)
}

func (s *DefaultProductService) GetAllProducts() ([]productsModel.Product, error) {
	return s.repo.FindAll()
}

func (s *DefaultProductService) SearchProducts(priceGt int) ([]productsModel.Product, error) {
	return s.repo.FindByPriceGreaterThan(priceGt)
}

func (s *DefaultProductService) AddProduct(product productsModel.Product) error {
	return s.repo.Save(product)
}
