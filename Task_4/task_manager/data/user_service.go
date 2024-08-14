package data

import (
	"context"
	"errors"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(collection *mongo.Collection) *UserService {
	return &UserService{collection: collection}
}

func (s *UserService) CreateUser(user *models.User) error {
	user.ID = primitive.NewObjectID()
	_, err := s.collection.InsertOne(context.Background(), user)
	return err
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserService) PromoteUser(userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	update := bson.M{"$set": bson.M{"role": "admin"}}
	_, err = s.collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	return err
}
