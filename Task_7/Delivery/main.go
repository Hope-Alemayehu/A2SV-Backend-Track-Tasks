package main

import (
	"context"
	"log"
	"time"

	routers "Task_7/Delivery/routers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("taskdb")

	// Pass the timeout duration and the database to the router
	routers.SetupRouter(10*time.Second, db)
}
