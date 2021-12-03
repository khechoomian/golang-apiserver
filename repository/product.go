package repository

import (
	"golang-apiserver/model"
	"log"

	"gorm.io/gorm"
)

// ProductRepository Product implement struct
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository create new repository
func NewProductRepository(DB *gorm.DB) model.ProductRepository {
	return &productRepository{DB}
}

// Create
func (r *productRepository) All() []model.Product {
	var product []model.Product
	result := r.db.Find(&product)
	if result.Error != nil {
		log.Println(result.Error)
		return nil
	}
	return product
}
