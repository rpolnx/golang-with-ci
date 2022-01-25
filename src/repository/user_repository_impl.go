package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rpolnx.com.br/golang-with-ci/src/model/entities"
	"time"
)

const userCollectionName = "users"

type UserRepository interface {
	FindAllUsers() ([]entities.User, error)

	FindUserById(id primitive.ObjectID) (*entities.User, error)

	CreateUser(entity entities.User) (*mongo.InsertOneResult, error)

	UpsertUser(id primitive.ObjectID, entity entities.User) (*mongo.UpdateResult, error)

	DeleteUserById(id primitive.ObjectID) (*mongo.DeleteResult, error)
}

func (r *mongoUserRepository) FindAllUsers() ([]entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection(userCollectionName)

	users := make([]entities.User, 0)

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *mongoUserRepository) FindUserById(id primitive.ObjectID) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection(userCollectionName)

	currentUser := new(entities.User)

	filter := bson.M{"_id": id}

	err := collection.FindOne(ctx, filter).Decode(currentUser)

	return currentUser, err
}

func (r *mongoUserRepository) CreateUser(entity entities.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection(userCollectionName)

	entity.ID = primitive.NewObjectID()

	return collection.InsertOne(ctx, entity)
}

func (r *mongoUserRepository) UpsertUser(id primitive.ObjectID, entity entities.User) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection(userCollectionName)

	filter := bson.M{"_id": id}

	upsertEntityArg := bson.M{
		"$set": entity,
		"$setOnInsert": bson.M{
			"created_at": time.Now(),
			"updated_at": time.Now(),
		},
	}

	upsert := true
	var upsertOptions = &options.UpdateOptions{
		Upsert: &upsert,
	}

	return collection.UpdateOne(ctx, filter, upsertEntityArg, upsertOptions)
}

func (r *mongoUserRepository) DeleteUserById(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection(userCollectionName)

	filter := bson.M{"_id": id}

	return collection.DeleteOne(ctx, filter)
}
