package service

import (
	productsRepo "ejercicioTest/internal/infraestructure/repository"
	productsModel "ejercicioTest/pkg/models"
)

type ProductService interface {
	GetProductByID(id int) (*productsModel.Product, error)
	GetAllProducts() ([]productsModel.Product, error)
	SearchProducts(priceGt int) ([]productsModel.Product, error)
	AddProduct(product productsModel.Product) error
	UpdateProduct(id int, product productsModel.Product) error
	ReplaceProduct(id int, product productsModel.Product) error
	DeleteProduct(id int) error
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

func (s *DefaultProductService) UpdateProduct(id int, product productsModel.Product) error {
	return s.repo.Update(id, product)
}

func (s *DefaultProductService) ReplaceProduct(id int, product productsModel.Product) error {
	return s.repo.Replace(id, product)
}

func (s *DefaultProductService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}
