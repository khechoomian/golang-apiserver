package repository

import (
	"golang-apiserver/model"
	"log"

	"gorm.io/gorm"
)

// userRepository user implement struct
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository create new repository
func NewUserRepository(DB *gorm.DB) model.UserRepository {
	return &userRepository{DB}
}

// Create
func (r *userRepository) Create(u model.User) model.User {
	result := r.db.Create(&u)
	if result.Error != nil {
		log.Println(result.Error)
		return model.User{}
	}
	return u
}

// GetByEmail
func (r *userRepository) GetByEmail(userEmail string) model.User {
	user := model.User{}
	result := r.db.Where("email = ?", userEmail).First(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return model.User{}
	}
	return user
}
