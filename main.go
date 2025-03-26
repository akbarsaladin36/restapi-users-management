package main

import (
	"fmt"
	"restapi-users-management/database"
	"restapi-users-management/migrations"
	"restapi-users-management/routes"

	"github.com/joho/godotenv"
)

// @title Rest API - User Management
// @version 1.0
// @description The API backend for personal project
// @host localhost:8002
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token
func main() {
	loadEnv()
	database.ConnectDB()
	migrations.MigrateTables()
	routes.ConnectRoutes()
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Failed to load .env file")
	}
}
