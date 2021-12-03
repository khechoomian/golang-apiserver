/*
Copyright 2021
*/

// apiserver is the main api server
package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
}

func init() {
	//TODO this section must be change into another file and bether construct
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}
