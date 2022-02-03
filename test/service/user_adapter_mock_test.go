package service

import (
	"github.com/stretchr/testify/mock"
	"rpolnx.com.br/golang-with-ci/src/model/entities"
	"rpolnx.com.br/golang-with-ci/src/ports/out"
)

type userAdapterMock struct {
	mock.Mock
	out.UserPort
}

func (m *userAdapterMock) FindAllUsers() ([]entities.User, error) {
	args := m.Called()
	s, _ := args.Get(0).([]entities.User)
	return s, args.Error(1)
}

func (m *userAdapterMock) FindUserById(id string) (*entities.User, error) {
	args := m.Called(id)
	s, _ := args.Get(0).(*entities.User)
	return s, args.Error(1)
}

func (m *userAdapterMock) CreateUser(e entities.User) (string, error) {
	args := m.Called(e)
	s, _ := args.Get(0).(string)
	return s, args.Error(1)
}

func (m *userAdapterMock) UpdateUser(id string, e entities.User) error {
	args := m.Called(id, e)
	return args.Error(0)
}

func (m *userAdapterMock) DeleteUserById(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
