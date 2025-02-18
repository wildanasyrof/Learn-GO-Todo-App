package main

import (
	"to-do-api/config"
	"to-do-api/routes"
	"to-do-api/utils"
)

func main() {

	// Initialize Logger
	utils.InitLogger()
	utils.Logger.Info("Starting server...")

	// Initialize database
	config.InitDB()
	utils.Logger.Info("Database connected successfully")

	// Setup routes
	r := routes.SetupRouter()

	// Run server
	utils.Logger.Info("Server running on port 8080")
	r.Run(":8080")
}
