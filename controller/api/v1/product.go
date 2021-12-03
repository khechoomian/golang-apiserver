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

// ProductHandler
type ProductHandler struct {
	productService model.ProductService
}

// NewProductHandler
func NewProductHandler(r *gin.RouterGroup) {
	userRepo := repository.NewProductRepository(db.DB.Con)
	productService := service.NewProductService(userRepo)
	productHandler := ProductHandler{
		productService: productService,
	}
	// Authentication Middleware
	user := r.Group("/product")
	{
		user.GET("", productHandler.ProductList)
	}
}

// Login
func (h *ProductHandler) ProductList(c *gin.Context) {
	data := h.productService.List()
	c.JSON(http.StatusBadRequest, resources.WebResponse{
		Code:    http.StatusBadRequest,
		Results: data,
	})
	return
}
