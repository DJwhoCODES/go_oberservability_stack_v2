package repository

import (
	"context"
	"fmt"
	"instagram/user-service/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return fmt.Errorf("failed to convert inserted id")
	}

	user.ID = oid.Hex()

	return nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.collection.FindOne(ctx, map[string]string{"email": email}).Decode(&user)
	return &user, err
}
