package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `form:"name" json:"name" binding:"required"`
}

// ProductService
type ProductService interface {
	List() []Product
}

// ProductRepository
type ProductRepository interface {
	All() []Product
}
