package routes

import (
	"to-do-api/controllers"
	"to-do-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Authentication routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/tasks", controllers.GetTasks)
		protected.GET("/tasks/:id", controllers.GetTaskByID)
		protected.POST("/tasks", controllers.CreateTask)
		protected.PUT("/tasks/:id", controllers.UpdateTask)
		protected.DELETE("/tasks/:id", controllers.DeleteTask)
	}

	return r
}
