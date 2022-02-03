package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserDBRepository2 interface {
	UserRepository
}

type mongoUserRepository2 struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func InitializeUserDatabaseClient2() (*mongoUserRepository2, error) {
	repo, err := newMongoUserRepository2()

	if err != nil {
		return nil, err
	}

	return repo, nil
}

func newMongoUserRepository2() (*mongoUserRepository2, error) {
	repo := &mongoUserRepository2{
		timeout:  time.Duration(timeout) * time.Second,
		database: database,
	}

	client, err := newMongoClient(timeout, uri)

	if err != nil {
		return nil, err
	}

	repo.client = client

	return repo, nil
}
