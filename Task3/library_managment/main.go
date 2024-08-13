package main

import (
	"library_managment/controllers"
	"library_managment/services"
)

func main() {
	service := services.NewLibrary()
	controller := controllers.NewLibraryController(service)
	controllers.RunLibraryManagementSystem(controller)
}
