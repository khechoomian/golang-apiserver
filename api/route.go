package api

import (
	"github.com/gin-gonic/gin"
)

// Routes of group api v1
func Routes(router *gin.Engine) {
	router.Group("/api/v1")
}
