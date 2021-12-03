/*
Copyright 2021
*/

// apiserver is the main api server
package main

import (
	"golang-apiserver/api"
	"golang-apiserver/db"
	"log"
	"os"

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
	DB := db.NewDB()
	if v := os.Getenv("API_PORT"); v != "" {
		api.APIPort = v
	}
	if v := os.Getenv("DATABASE_USER"); v != "" {
		DB.User = v
	}
	if v := os.Getenv("DATABASE_HOST"); v != "" {
		DB.Host = v
	}
	if v := os.Getenv("DATABASE_PASSWORD"); v != "" {
		DB.Password = v
	}
	if v := os.Getenv("DATABASE_NAME"); v != "" {
		DB.DBName = v
	}
	if v := os.Getenv("DATABASE_PORT"); v != "" {
		DB.Port = v
	}
	if v := os.Getenv("DATABASE_TIMEZONE"); v != "" {
		DB.TimeZone = v
	}
	DB.Connect()

	//TODO this function is not good place
	db.MigrateModels()
}
