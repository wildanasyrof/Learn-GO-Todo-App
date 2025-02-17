package config

import (
	"fmt"
	"log"
	"to-do-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=wildan123 dbname=todo_app port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate models
	DB.AutoMigrate(&models.Task{})
	DB.AutoMigrate(&models.User{})

	fmt.Println("Database connected!")
}
