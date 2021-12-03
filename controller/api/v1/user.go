package apiv1

import (
	"golang-apiserver/db"
	"golang-apiserver/model"
	"golang-apiserver/repository"
	"golang-apiserver/resources"
	"golang-apiserver/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler
type UserHandler struct {
	userService model.UserService
}

// NewUserHandler
func NewUserHandler(r *gin.RouterGroup) {
	userRepo := repository.NewUserRepository(db.DB.Con)
	userService := service.NewUserService(userRepo)
	userHandler := UserHandler{
		userService: userService,
	}
	// Authentication Middleware
	user := r.Group("/user")
	{
		user.POST("/login", userHandler.Login)
		user.POST("/register", userHandler.Register)
	}
}

// Login
func (h *UserHandler) Login(c *gin.Context) {
	//TODO check email format and password length
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, resources.WebResponse{
			Code: http.StatusBadRequest,
		})
		return
	}
	if _, ok := h.userService.Login(user); ok {
		return
	} else {
		c.JSON(http.StatusBadRequest, resources.WebResponse{
			Code: http.StatusBadRequest,
		})
		return
	}
}

// Register
func (h *UserHandler) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, resources.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "User not registered",
		})
		return
	}
	data, ok := h.userService.Register(user)
	if !ok {
		c.JSON(http.StatusBadRequest, resources.WebResponse{
			Code:    http.StatusCreated,
			Message: "User not registered",
		})
		return
	}
	c.JSON(http.StatusCreated, resources.WebResponse{
		Code:    http.StatusCreated,
		Message: "User Registered",
		Results: data,
	})
}
