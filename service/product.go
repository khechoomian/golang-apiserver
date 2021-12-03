package service

import "golang-apiserver/model"

// ProductService
type productService struct {
	productRepository model.ProductRepository
}

// NewProductService create new service
func NewProductService(r model.ProductRepository) model.ProductService {
	return &productService{
		productRepository: r,
	}
}

// List
func (srv *productService) List() []model.Product {
	data := srv.productRepository.All()
	return data
}
