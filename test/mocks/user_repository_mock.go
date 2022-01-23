package mocks

import (
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rpolnx.com.br/golang-with-ci/src/model/entities"
	"rpolnx.com.br/golang-with-ci/src/repository"
)

type UserRepositoryMock struct {
	mock.Mock
	repository.UserDBRepository
}

func (m *UserRepositoryMock) FindAllUsers() ([]entities.User, error) {
	args := m.Called()
	s, _ := args.Get(0).([]entities.User)
	return s, args.Error(1)
}

func (m *UserRepositoryMock) FindUserById(id primitive.ObjectID) (*entities.User, error) {
	args := m.Called(id)
	s, _ := args.Get(0).(*entities.User)
	return s, args.Error(1)
}

func (m *UserRepositoryMock) CreateUser(entity entities.User) (*mongo.InsertOneResult, error) {
	args := m.Called(entity)
	s, _ := args.Get(0).(*mongo.InsertOneResult)
	return s, args.Error(1)
}

func (m *UserRepositoryMock) UpsertUser(id primitive.ObjectID, entity entities.User) (*mongo.UpdateResult, error) {
	args := m.Called(id, entity)
	s, _ := args.Get(0).(*mongo.UpdateResult)
	return s, args.Error(1)
}

func (m *UserRepositoryMock) DeleteUserById(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	args := m.Called(id)
	s, _ := args.Get(0).(*mongo.DeleteResult)
	return s, args.Error(1)
}
