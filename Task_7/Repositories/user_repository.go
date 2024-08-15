package repositories

import (
	domain "Task_7/Domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) CreateUser(c context.Context, user *domain.User) error {
	if user.UserID == "" {
		user.UserID = primitive.NewObjectID().Hex()
	} else {
		// Validate and convert the user-provided ID
		if _, err := primitive.ObjectIDFromHex(user.UserID); err != nil {
			return errors.New("invalid ID format")
		}
	}
	collection := ur.database.Collection(ur.collection)
	_, err := collection.InsertOne(c, user)
	return err
}

func (ur *userRepository) GetUserByUsername(c context.Context, username string) (domain.User, error) {
	var user domain.User
	collection := ur.database.Collection(ur.collection)
	err := collection.FindOne(c, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (ur *userRepository) GetUserByID(c context.Context, userID string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return domain.User{}, err
	}

	err = collection.FindOne(c, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}

func (ur *userRepository) PromoteUser(c context.Context, userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"role": "admin"}}
	collection := ur.database.Collection(ur.collection)
	_, err = collection.UpdateOne(c, bson.M{"_id": objID}, update)
	return err
}

func (ur *userRepository) DeleteUser(ctx context.Context, userID string) error {
	collection := ur.database.Collection(ur.collection)
	filter := bson.M{"_id": userID}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}

func (ur *userRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	collection := ur.database.Collection(ur.collection)
	cursor, err := collection.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user domain.User
		if err = cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
