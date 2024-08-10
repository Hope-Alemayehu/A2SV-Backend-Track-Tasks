package main

import (
	"fmt"
	"log"
	"task_manager/task_manager/database"
	"task_manager/task_manager/router"
)

func main() {
	// Connect to the database
	database.ConnectDB()
	defer database.DisconnectDB()

	// Set up the router
	r := router.SetUpRouter()

	fmt.Println("Task manager API running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
