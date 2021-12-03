package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	APIPort = "1234"
	IP      = "0.0.0.0"
)

//Run the function for run api  erver
func Run() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	Routes(router)
	err := router.Run(IP + ":" + APIPort)
	if err != nil {
		log.Fatal(err)
	}
}
