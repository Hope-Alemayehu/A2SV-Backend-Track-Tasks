package delivery

import (
	domain "Task_7/Domain"
	infrastructure "Task_7/Infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controller struct {
	UserUsecase domain.UserUsecase
	TaskUsecase domain.TaskUsecase
}

// UserController methods (unchanged)
func (uc *Controller) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user.Password = hash

	if err := uc.UserUsecase.CreateUser(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *Controller) LoginUser(c *gin.Context) {
	var loginInfo domain.LoginRequest

	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserUsecase.GetUserByUsername(c, loginInfo.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := infrastructure.CheckPassword(user.Password, loginInfo.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	tokenString, err := infrastructure.GenerateToken(user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (uc *Controller) PromoteUser(c *gin.Context) {
	userID := c.Param("id")

	err := uc.UserUsecase.PromoteUser(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Promoted"})
}

func (uc *Controller) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	if err := uc.UserUsecase.DeleteUser(c, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (uc *Controller) GetAllUsers(c *gin.Context) {
	users, err := uc.UserUsecase.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc *Controller) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	user, err := uc.UserUsecase.GetUserByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// TaskController methods
func (tc *Controller) CreateTask(c *gin.Context) {
	var task domain.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.TaskID = primitive.NewObjectID().Hex()

	if err := tc.TaskUsecase.CreateTask(c, &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *Controller) GetAllTasks(c *gin.Context) {
	tasks, err := tc.TaskUsecase.GetAllTasks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *Controller) GetTaskByID(c *gin.Context) {
	taskID := c.Param("id")

	task, err := tc.TaskUsecase.GetTaskByID(c, taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *Controller) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	var task domain.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.TaskUsecase.UpdateTask(c, taskID, task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (tc *Controller) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")

	if err := tc.TaskUsecase.DeleteTask(c, taskID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
