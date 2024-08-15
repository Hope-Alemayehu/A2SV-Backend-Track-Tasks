package domain

import "context"

type Task struct {
	TaskID      string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Status      string `json:"completed" bson:"completed"`
}

type User struct {
	UserID   string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

type LoginRequest struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type UserRepository interface {
	CreateUser(c context.Context, user *User) error
	GetUserByUsername(c context.Context, username string) (User, error)
	GetUserByID(c context.Context, userID string) (User, error)
	PromoteUser(c context.Context, userID string) error
	DeleteUser(c context.Context, userID string) error
	GetAllUsers(c context.Context) ([]User, error)
}

type UserUsecase interface {
	CreateUser(c context.Context, user *User) error
	GetUserByUsername(c context.Context, username string) (User, error)
	GetUserByID(c context.Context, userID string) (User, error)
	PromoteUser(c context.Context, userID string) error
	DeleteUser(c context.Context, userID string) error
	GetAllUsers(c context.Context) ([]User, error)
}

type TaskRepository interface {
	CreateTask(c context.Context, task *Task) error
	UpdateTask(c context.Context, taskID string, task Task) error // Change UpdateTask signature
	DeleteTask(c context.Context, taskID string) error
	GetTaskByID(c context.Context, taskID string) (Task, error)
	GetAllTasks(c context.Context) ([]Task, error) // Return a slice of Task
}

type TaskUsecase interface {
	CreateTask(c context.Context, task *Task) error
	UpdateTask(c context.Context, taskID string, task Task) error // Change UpdateTask signature
	DeleteTask(c context.Context, taskID string) error
	GetTaskByID(c context.Context, taskID string) (Task, error)
	GetAllTasks(c context.Context) ([]Task, error) // Return a slice of Task
}
