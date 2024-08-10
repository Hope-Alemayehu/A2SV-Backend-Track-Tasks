package router

import (
	"task_manager/task_manager/controllers"
	"task_manager/task_manager/data"
	"task_manager/task_manager/database"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	// Creating TaskService and TaskController instance
	taskService := data.NewTaskService(database.Collection)
	taskController := controllers.NewTaskController(taskService)

	r.GET("/tasks", taskController.GetTasks)
	r.GET("/tasks/:id", taskController.GetTaskById)
	r.POST("/tasks", taskController.PostTask)
	r.PUT("/tasks/:id", taskController.UpdateTask)
	r.DELETE("/tasks/:id", taskController.DeleteTask)

	return r
}
