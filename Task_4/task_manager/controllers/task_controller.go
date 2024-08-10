package controllers

import (
	"context"
	"net/http"
	"task_manager/task_manager/data"
	"task_manager/task_manager/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type TaskController struct {
	service *data.TaskService
}

func NewTaskController(service *data.TaskService) *TaskController {
	return &TaskController{service: service}
}

func (tc *TaskController) GetTasks(ctx *gin.Context) {
	tasks, err := tc.service.GetAllTasks(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := tc.service.GetTaskById(context.Background(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (tc *TaskController) PostTask(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := tc.service.AddTask(context.Background(), newTask)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (tc *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateData models.Task
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a map to hold the fields to be updated
	updateFields := bson.M{
		"title":       updateData.Title,
		"description": updateData.Description,
		"due_date":    updateData.DueDate,
		"status":      updateData.Status,
	}

	// Perform the update operation
	err := tc.service.UpdateTask(context.Background(), id, updateFields)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func (tc *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	err := tc.service.DeleteTask(context.Background(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
