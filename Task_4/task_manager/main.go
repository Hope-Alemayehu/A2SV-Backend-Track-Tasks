package main

import (
	"fmt"
	"task_manager/router"
)

func main() {
	fmt.Println(("Task Manager API is starting..."))
	r := router.SetUpRouter()

	//we can adjust this as needed
	port := ":8080"

	if err := r.Run(port); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

}
