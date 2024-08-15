package repositories

import (
	domain "Task_7/Domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	database   *mongo.Database // Pointer to mongo.Database
	collection string
}

func NewTaskRepository(db *mongo.Database, collection string) domain.TaskRepository {
	return &TaskRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *TaskRepository) CreateTask(c context.Context, task *domain.Task) error {
	collection := tr.database.Collection(tr.collection)
	_, err := collection.InsertOne(c, task)
	return err
}

func (tr *TaskRepository) UpdateTask(c context.Context, taskID string, task domain.Task) error {
	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}

	collection := tr.database.Collection(tr.collection)
	update := bson.M{}

	// Only add the fields to update if they are not empty
	if task.Title != "" {
		update["title"] = task.Title
	}

	if task.Description != "" {
		update["description"] = task.Description
	}

	if task.Status != "" {
		update["status"] = task.Status
	}

	if len(update) == 0 {
		return errors.New("no field to update")
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": objID}, bson.M{"$set": update})
	return err
}

func (tr *TaskRepository) GetAllTasks(c context.Context) ([]domain.Task, error) { // Return slice of domain.Task
	collection := tr.database.Collection(tr.collection)
	cursor, err := collection.Find(c, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(c)
	var tasks []domain.Task
	for cursor.Next(c) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (tr *TaskRepository) GetTaskByID(c context.Context, taskID string) (domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return domain.Task{}, err
	}

	collection := tr.database.Collection(tr.collection)
	var task domain.Task
	err = collection.FindOne(c, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (tr *TaskRepository) DeleteTask(c context.Context, taskID string) error {
	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}
	collection := tr.database.Collection(tr.collection)
	_, err = collection.DeleteOne(c, bson.M{"_id": objID})
	return err
}
