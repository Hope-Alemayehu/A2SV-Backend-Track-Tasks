package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

// creating an in memory address to store the tasks
var taskService = data.NewTaskService()

// get all tasks
func GetTasks(ctx *gin.Context) {
	tasks := taskService.GetAllTasks()
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})

}

func GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := taskService.GetTaskById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error:": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)

}

func PostTask(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	taskService.AddTask(newTask)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task Created"})

}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var UpdatedTask models.Task

	if err := ctx.ShouldBindJSON(&UpdatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//calling task service to update the task with the given id
	err := taskService.UpdateTask(id, UpdatedTask)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	//there is no error, reponds with 200
	ctx.JSON(http.StatusOK, gin.H{"message": "Task Updated"})

}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	err := taskService.DeleteTask(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
