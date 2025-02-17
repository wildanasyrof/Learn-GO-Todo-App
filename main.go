package main

import (
	"to-do-api/config"
	"to-do-api/routes"
)

func main() {
	// Initialize database
	config.InitDB()

	// Setup routes
	r := routes.SetupRouter()

	// Run server
	r.Run(":8080")
}
