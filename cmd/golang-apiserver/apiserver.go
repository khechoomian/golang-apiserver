/*
Copyright 2021
*/

// apiserver is the main api server
package main

import (
	"golang-apiserver/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	api.Run()
}

func init() {
	//TODO this section must be change into another file and bether construct
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}
