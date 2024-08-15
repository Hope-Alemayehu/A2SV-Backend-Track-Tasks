package routers

import (
	"time"

	controllers "Task_7/Delivery/controllers"
	infrastructure "Task_7/Infrastructure"
	"Task_7/repositories"
	"Task_7/usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(timeout time.Duration, db *mongo.Database) *gin.Engine {
	r := gin.Default()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db, "users")
	taskRepo := repositories.NewTaskRepository(db, "tasks")

	// Initialize usecases
	userUsecase := usecases.NewUserUsecase(userRepo, timeout)
	taskUsecase := usecases.NewTaskUsecase(taskRepo, timeout)

	// Initialize controllers
	controller := &controllers.Controller{
		UserUsecase: userUsecase,
		TaskUsecase: taskUsecase,
	}

	// Public routes
	publicRouter := r.Group("/v1")
	{
		publicRouter.POST("/register", controller.CreateUser)
		publicRouter.POST("/login", controller.LoginUser)
	}

	// Protected routes
	protectedRouter := r.Group("/v2")
	protectedRouter.Use(infrastructure.AuthMiddleware())
	{
		protectedRouter.GET("/users", controller.GetAllUsers)
		protectedRouter.GET("/users/:id", controller.GetUserByID)
		protectedRouter.DELETE("/users/:id", infrastructure.AdminMiddleware(), controller.DeleteUser)
		protectedRouter.POST("/promote/:id", infrastructure.AdminMiddleware(), controller.PromoteUser)

		protectedRouter.POST("/tasks", controller.CreateTask)
		protectedRouter.GET("/tasks", controller.GetAllTasks)
		protectedRouter.GET("/tasks/:id", controller.GetTaskByID)
		protectedRouter.PUT("/tasks/:id", controller.UpdateTask)
		protectedRouter.DELETE("/tasks/:id", controller.DeleteTask)
	}

	r.Run(":8080")
	return r
}
