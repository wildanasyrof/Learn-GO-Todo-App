package controllers

import (
	"net/http"
	"strconv"
	"to-do-api/config"
	"to-do-api/models"
	"to-do-api/utils"

	"github.com/gin-gonic/gin"
)

// Get tasks with pagination, filtering, and search
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	var total int64

	// Get query parameters
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		utils.Logger.Warn("Invalid page number")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		utils.Logger.Warn("Invalid limit value")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	status := c.Query("status") // Filter by status
	search := c.Query("search") // Search by title

	// Pagination logic
	offset := (page - 1) * limit
	query := config.DB.Model(&models.Task{})

	// Apply status filter if provided
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Apply search filter if provided
	if search != "" {
		query = query.Where("title ILIKE ?", "%"+search+"%")
	}

	// Count total tasks
	query.Count(&total)

	// Fetch tasks
	result := query.Offset(offset).Limit(limit).Find(&tasks)
	if result.Error != nil {
		utils.Logger.Error("Database error: ", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Log successful request
	utils.Logger.Infof("Fetched %d tasks", len(tasks))

	// Response with pagination metadata
	c.JSON(http.StatusOK, gin.H{
		"tasks":      tasks,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}

// Get task by ID
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	result := config.DB.First(&task, id)

	if result.Error != nil {
		utils.Logger.Warnf("Task with Id %s is not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	utils.Logger.Infof("Fetched task with Id %s", id)
	c.JSON(http.StatusOK, task)
}

// Create a new task
func CreateTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		utils.Logger.Error("Invalid request adding Task")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	config.DB.Create(&newTask)
	utils.Logger.Infof("Created new task with Id %d", newTask.Id)
	c.JSON(http.StatusCreated, newTask)
}

// Update task
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := config.DB.First(&task, id).Error; err != nil {
		utils.Logger.Warnf("Task with Id %s is not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		utils.Logger.Error("Invalid request updating Task")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	config.DB.Save(&task)
	utils.Logger.Infof("Updated task with Id %s", id)
	c.JSON(http.StatusOK, task)
}

// Delete task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := config.DB.First(&task, id).Error; err != nil {
		utils.Logger.Warnf("Task with Id %s is not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	config.DB.Delete(&task)
	utils.Logger.Infof("Deleted task with Id %s", id)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
