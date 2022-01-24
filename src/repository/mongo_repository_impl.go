package repository

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient(timeout int, uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	getenv := os.Getenv("MONGO_URI")

	if getenv != "" {
		uri = getenv
	}

	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, err
	}

	return client, nil
}
