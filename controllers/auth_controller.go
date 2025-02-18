package controllers

import (
	"net/http"

	"to-do-api/config"
	"to-do-api/models"
	"to-do-api/utils"

	"github.com/gin-gonic/gin"
)

// Register a new user
func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Logger.Warn("Invalid registration input")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := user.HashPassword(); err != nil {
		utils.Logger.Error("Failed to hash password")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	config.DB.Create(&user)
	utils.Logger.Infof("User registered: %s", user.Username)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login and get JWT token
func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Logger.Warn("Invalid login input")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		utils.Logger.Warn("Invalid login attempt for user: ", input.Username)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if !user.CheckPassword(input.Password) {
		utils.Logger.Warn("Invalid password attempt for user: ", input.Username)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		utils.Logger.Error("Failed to generate token")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	utils.Logger.Infof("User logged in: %s", user.Username)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
