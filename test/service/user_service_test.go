package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rpolnx.com.br/golang-with-ci/src/model/entities"
	"rpolnx.com.br/golang-with-ci/src/service"
	"rpolnx.com.br/golang-with-ci/test/mocks"
	"testing"
)

//GetAllUsers
func Test_ShouldGetUnexpectedErrorWhenGettingUsers(t *testing.T) {
	// Initializations Phase
	adapterMock := new(mocks.UserAdapterMock)
	userService := service.InitializeUserService(adapterMock)

	//Mock Setups Phase
	adapterMock.On("FindAllUsers").Return(nil, errors.New("general error users"))

	//Execution Phase
	users, err := userService.GetAllUsers()

	//Assert Phase
	expectedErrorMsg := "general error users"

	assert.Nil(t, users)
	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}

//GetAllUsers
func Test_ShouldGetAListOfUsersWhenGettingUsers(t *testing.T) {
	// Initializations Phase
	adapterMock := new(mocks.UserAdapterMock)
	userService := service.InitializeUserService(adapterMock)

	//Mock Setups Phase
	data := []entities.User{
		{ID: primitive.NewObjectID(), Name: "Vegetta", Age: 10},
		{ID: primitive.NewObjectID(), Name: "Goku", Age: 15},
	}
	adapterMock.On("FindAllUsers").Return(data, nil)

	//Execution Phase
	users, err := userService.GetAllUsers()

	//Assert Phase
	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, users, data)
}

//GetOneUser
func Test_ShouldGetUnexpectedErrorWhenGettingOneUser(t *testing.T) {
	// Initializations Phase
	adapterMock := new(mocks.UserAdapterMock)
	userService := service.InitializeUserService(adapterMock)

	//Mock Setups Phase
	id := "error-id"
	adapterMock.On("FindUserById", id).Return(nil, errors.New("general error one user"))

	//Execution Phase
	users, err := userService.GetOneUser(id)

	//Assert Phase
	expectedErrorMsg := "general error one user"

	assert.Nil(t, users)
	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}

//GetOneUser
func Test_ShouldReceiveAnUserWhenGettingOneUserById(t *testing.T) {
	// Initializations Phase
	adapterMock := new(mocks.UserAdapterMock)
	userService := service.InitializeUserService(adapterMock)

	//Mock Setups Phase
	data := &entities.User{ID: primitive.NewObjectID(), Name: "Vegetta", Age: 10}
	id := "test-id"
	adapterMock.On("FindUserById", id).Return(data, nil)

	//Execution Phase
	users, err := userService.GetOneUser(id)

	//Assert Phase
	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, users, data)
}

//PostUser
func Test_ShouldGetUnexpectedErrorWhenCreatingUser(t *testing.T) {
	// Initializations Phase
	adapterMock := new(mocks.UserAdapterMock)
	userService := service.InitializeUserService(adapterMock)

	//Mock Setups Phase
	data := entities.User{Name: "Vegetta", Age: 10}
	adapterMock.On("CreateUser", data).Return("", errors.New("general error create user"))

	//Execution Phase
	users, err := userService.PostUser(data)

	//Assert Phase
	expectedErrorMsg := "general error create user"

	assert.Equal(t, "", users)
	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}

//PostUser
func Test_ShouldCreateAnUserWhenCreatingUser(t *testing.T) {
	// Initializations Phase
	adapterMock := new(mocks.UserAdapterMock)
	userService := service.InitializeUserService(adapterMock)

	//Mock Setups Phase
	data := entities.User{Name: "Vegetta", Age: 10}
	hex := primitive.NewObjectID().Hex()
	adapterMock.On("CreateUser", data).Return(hex, nil)

	//Execution Phase
	result, err := userService.PostUser(data)

	//Assert Phase
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, hex, result)
}

//PutUser
func Test_ShouldGetUnexpectedErrorWhenUpdatingAnUser(t *testing.T) {
	// Initializations Phase
	adapterMock := new(mocks.UserAdapterMock)
	userService := service.InitializeUserService(adapterMock)

	//Mock Setups Phase
	id := "fail-id"
	data := entities.User{ID: primitive.NewObjectID(), Name: "Vegetta", Age: 10}
	adapterMock.On("UpdateUser", id, data).Return(errors.New("general error update user"))

	//Execution Phase
	err := userService.PutUser(id, data)

	//Assert Phase
	expectedErrorMsg := "general error update user"

	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}

//PutUser
func Test_ShouldUpdateAnUserWhenReceiveValidId(t *testing.T) {
	// Initializations Phase
	adapterMock := new(mocks.UserAdapterMock)
	userService := service.InitializeUserService(adapterMock)

	//Mock Setups Phase
	id := "test-id"
	data := entities.User{ID: primitive.NewObjectID(), Name: "Vegetta", Age: 10}
	adapterMock.On("UpdateUser", id, data).Return(nil)

	//Execution Phase
	err := userService.PutUser(id, data)

	//Assert Phase
	assert.Nil(t, err)
}

//DeleteUser
func Test_ShouldGetUnexpectedErrorWhenDeletingAnUser(t *testing.T) {
	// Initializations Phase
	adapterMock := new(mocks.UserAdapterMock)
	userService := service.InitializeUserService(adapterMock)

	//Mock Setups Phase
	id := "fail-id"
	adapterMock.On("DeleteUserById", id).Return(errors.New("general error update user"))

	//Execution Phase
	err := userService.DeleteUser(id)

	//Assert Phase
	expectedErrorMsg := "general error update user"

	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}

//DeleteUser
func Test_ShouldDeleteAnUserWhenReceiveValidID(t *testing.T) {
	// Initializations Phase
	adapterMock := new(mocks.UserAdapterMock)
	userService := service.InitializeUserService(adapterMock)

	//Mock Setups Phase
	id := "test-id"
	adapterMock.On("DeleteUserById", id).Return(nil)

	//Execution Phase
	err := userService.DeleteUser(id)

	//Assert Phase
	assert.Nil(t, err)
}
