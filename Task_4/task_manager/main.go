package main

import (
	"context"
	"log"
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	userCollection := client.Database("taskmanager").Collection("users")
	taskCollection := client.Database("taskmanager").Collection("tasks")

	userService := data.NewUserService(userCollection)
	taskService := data.NewTaskService(taskCollection)

	controllers.SetServices(userService, taskService)

	r := router.SetUpRouter()

	r.Run(":8080")
}
