package data

import (
	"context"
	"task_manager/models"

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

func (s *TaskService) GetTasks() ([]models.Task, error) {
	cur, err := s.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var tasks []models.Task
	for cur.Next(context.Background()) {
		var task models.Task
		if err := cur.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *TaskService) GetTaskById(id string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var task models.Task
	err = s.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) CreateTask(task *models.Task) error {
	task.ID = primitive.NewObjectID()
	_, err := s.collection.InsertOne(context.Background(), task)
	return err
}

func (s *TaskService) UpdateTask(id string, task *models.Task) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = s.collection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": task})
	return err
}

func (s *TaskService) DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = s.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	return err
}
