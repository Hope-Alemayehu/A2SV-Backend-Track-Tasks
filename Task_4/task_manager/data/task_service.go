package data

import (
	"context"
	"errors"
	"task_manager/task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(collection *mongo.Collection) *TaskService {
	return &TaskService{collection: collection}
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) GetTaskById(ctx context.Context, id string) (models.Task, error) {
	var task models.Task
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (s *TaskService) AddTask(ctx context.Context, task models.Task) (string, error) {
	task.ID = primitive.NewObjectID().Hex()
	_, err := s.collection.InsertOne(ctx, task)
	if err != nil {
		return "", err
	}
	return task.ID, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id string, updateFields bson.M) error {
	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": updateFields,
	}

	result, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	result, err := s.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
