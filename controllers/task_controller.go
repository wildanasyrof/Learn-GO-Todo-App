package controllers

import (
	"net/http"
	"to-do-api/config"
	"to-do-api/models"

	"github.com/gin-gonic/gin"
)

// Get all tasks
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	config.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// Get task by ID
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	result := config.DB.First(&task, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Create a new task
func CreateTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	config.DB.Create(&newTask)
	c.JSON(http.StatusCreated, newTask)
}

// Update task
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	config.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// Delete task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	config.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
