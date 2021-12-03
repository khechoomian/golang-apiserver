package model

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// UserService service of user model
type UserService interface {
	Login(User) (User, bool)
	Register(User) (User, bool)
}

// UserRepository repository of user model
type UserRepository interface {
	Create(User) User
	GetByEmail(userEmail string) User
}
