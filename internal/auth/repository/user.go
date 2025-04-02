package repository

import (
	"context"
	"time"

	models "github.com/edorguez/payment-reminder/internal/auth/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var collectionName = "users"

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id uint) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint) error
}

type userRepository struct {
	client *mongo.Client
  dbName string
}

func NewUserRepository(client *mongo.Client, dbName string) UserRepository {
  return &userRepository{
    client: client,
    dbName: dbName,
  }
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
  collection := r.client.Database(r.dbName).Collection(collectionName)

	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

  return nil
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
  collection := r.client.Database(r.dbName).Collection(collectionName)

	var result models.User
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	collection := r.client.Database(r.dbName).Collection(collectionName)

	var result models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
  collection := r.client.Database(r.dbName).Collection(collectionName)

	opts := options.UpdateOne().SetUpsert(true)
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": bson.M{
		"name":        user.Email,
		"modifiedAt":  time.Now(),
	},
	}

	_, err := collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
  collection := r.client.Database(r.dbName).Collection(collectionName)

	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}
