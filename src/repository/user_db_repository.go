package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const uri = "mongodb://root:password@localhost:27017/?authSource=admin"
const timeout = 30
const database = "golang-poc"

type UserDBRepository interface {
	UserRepository
}

type mongoUserRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func InitializeUserDatabaseClient() (UserDBRepository, error) {
	repo, err := newMongoUserRepository()

	if err != nil {
		return nil, err
	}

	return repo, nil
}

func newMongoUserRepository() (UserDBRepository, error) {
	repo := &mongoUserRepository{
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
