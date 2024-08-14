package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	// Public endpoints
	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.LoginUser)

	// Protected endpoints
	protected := r.Group("/tasks")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("", controllers.GetTasks)
		protected.GET("/:id", controllers.GetTaskById)
		protected.POST("", controllers.PostTask)
		protected.PUT("/:id", controllers.UpdateTask)
		protected.DELETE("/:id", controllers.DeleteTask)
		protected.POST("/promote/:id", controllers.PromoteUser)
	}

	return r
}
