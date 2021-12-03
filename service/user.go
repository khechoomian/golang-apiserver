package service

import "golang-apiserver/model"

// userService
type userService struct {
	userRepository model.UserRepository
}

// NewUserService create new service
func NewUserService(r model.UserRepository) model.UserService {
	return &userService{
		userRepository: r,
	}
}

// Login
func (srv *userService) Login(user model.User) (model.User, bool) {
	u := srv.userRepository.GetByEmail(user.Email)
	if u.ID == 0 {
		return model.User{}, false
	}
	return u, user.Password == u.Password
}

// Register
func (srv *userService) Register(user model.User) (model.User, bool) {
	if user.Email == "" || user.Password == "" {
		return model.User{}, false
	}
	// check if user exists
	u := srv.userRepository.GetByEmail(user.Email)
	if u.ID != 0 {
		return model.User{}, false
	}
	data := srv.userRepository.Create(user)
	//if user not create
	if data.ID == 0 {
		return model.User{}, false
	}
	return data, true
}
