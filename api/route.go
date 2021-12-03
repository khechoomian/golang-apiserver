package api

import (
	apiv1 "golang-apiserver/controller/api/v1"

	"github.com/gin-gonic/gin"
)

// Routes of group api v1
func Routes(router *gin.Engine) {
	apiV1 := router.Group("/api/v1")
	apiv1.NewUserHandler(apiV1)
	apiv1.NewProductHandler(apiV1)
}
