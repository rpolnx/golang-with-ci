package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rpolnx.com.br/golang-with-ci/src/adapter"
	"rpolnx.com.br/golang-with-ci/src/model/entities"
	"testing"
)

//FindAllUsers
func Test_ShouldGetUnexpectedErrorWhenGettingUsers(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	adapterMock.On("FindAllUsers").Return(nil, errors.New("general error users"))

	//Execution Phase
	users, err := userService.FindAllUsers()

	//Assert Phase
	expectedErrorMsg := "general error users"

	assert.Nil(t, users)
	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}

//FindAllUsers
func Test_ShouldGetAListOfUsersWhenGettingUsers(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	data := []entities.User{
		{ID: primitive.NewObjectID(), Name: "Vegetta", Age: 10},
		{ID: primitive.NewObjectID(), Name: "Goku", Age: 15},
	}
	adapterMock.On("FindAllUsers").Return(data, nil)

	//Execution Phase
	users, err := userService.FindAllUsers()

	//Assert Phase
	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, users, data)
}

//FindUserById
func Test_ShouldGetInvalidIdErrorWhenGettingUserWithAnInvalidId(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	id := "error-id"
	adapterMock.On("FindUserById", id).Return(nil, nil)

	//Execution Phase
	users, err := userService.FindUserById(id)

	//Assert Phase
	expectedErrorMsg := "the provided hex string is not a valid ObjectID"

	assert.Nil(t, users)
	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}

//FindUserById
func Test_ShouldGetUnexpectedErrorWhenGettingOneUser(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	id := primitive.NewObjectID()
	adapterMock.On("FindUserById", id).Return(nil, errors.New("general error one user"))

	//Execution Phase
	users, err := userService.FindUserById(id.Hex())

	//Assert Phase
	expectedErrorMsg := "general error one user"

	assert.Nil(t, users)
	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}

//FindUserById
func Test_ShouldReceiveAnUserWhenGettingOneUserById(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	data := &entities.User{ID: primitive.NewObjectID(), Name: "Vegetta", Age: 10}
	id := primitive.NewObjectID()
	adapterMock.On("FindUserById", id).Return(data, nil)

	//Execution Phase
	users, err := userService.FindUserById(id.Hex())

	//Assert Phase
	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, users, data)
}

//CreateUser
func Test_ShouldGetIdErrorWhenCreatingUserAndDidntReceiveAnID(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	data := entities.User{ID: primitive.NewObjectID(), Name: "Vegetta", Age: 10}
	expectedSendData := data
	adapterMock.On("CreateUser", expectedSendData).Return(new(mongo.InsertOneResult), nil)

	//Execution Phase
	createdId, err := userService.CreateUser(data)

	//Assert Phase
	expectedErrorMsg := "created user did not returned id"

	assert.Equal(t, "", createdId)
	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}

//CreateUser
func Test_ShouldGetUnexpectedErrorWhenCreatingUser(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	data := entities.User{ID: primitive.NewObjectID(), Name: "Vegetta", Age: 10}
	expectedSendData := data
	adapterMock.On("CreateUser", expectedSendData).Return(nil, errors.New("general error create user"))

	//Execution Phase
	createdId, err := userService.CreateUser(data)

	//Assert Phase
	expectedErrorMsg := "general error create user"

	assert.Equal(t, "", createdId)
	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}

//CreateUser
func Test_ShouldCreateAnUserWhenCreatingUser(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	data := entities.User{Name: "Vegetta", Age: 10}
	id := primitive.NewObjectID()
	adapterMock.On("CreateUser", data).Return(&mongo.InsertOneResult{InsertedID: id}, nil)

	//Execution Phase
	result, err := userService.CreateUser(data)

	//Assert Phase
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id.Hex(), result)
}

//UpdateUser
func Test_ShouldGetInvalidIdErrorErrorWhenUpdatingUserWithAnInvalidId(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	id := "fail-id"
	data := entities.User{ID: primitive.NewObjectID(), Name: "Vegetta", Age: 10}
	adapterMock.On("UpdateUser", id, data).Return(nil)

	//Execution Phase
	err := userService.UpdateUser(id, data)

	//Assert Phase
	expectedErrorMsg := "the provided hex string is not a valid ObjectID"

	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}

//UpdateUser
func Test_ShouldGetUnexpectedErrorWhenUpdatingAnUser(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	id := primitive.NewObjectID()
	data := entities.User{ID: id, Name: "Vegetta", Age: 10}
	adapterMock.On("UpsertUser", id, data).Return(nil, errors.New("general error update user"))

	//Execution Phase
	err := userService.UpdateUser(id.Hex(), data)

	//Assert Phase
	expectedErrorMsg := "general error update user"

	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}

//UpdateUser
func Test_ShouldUpdateAnUserWhenReceiveValidIdEvenIfDoesntExists(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	id := primitive.NewObjectID()
	data := entities.User{ID: id, Name: "Vegetta", Age: 10}
	updateResult := &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 1, UpsertedID: id}
	adapterMock.On("UpsertUser", id, data).Return(updateResult, nil)

	//Execution Phase
	err := userService.UpdateUser(id.Hex(), data)

	//Assert Phase
	assert.Nil(t, err)
}

//UpdateUser
func Test_ShouldUpdateAnUserWhenReceiveValidId(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	id := primitive.NewObjectID()
	data := entities.User{ID: id, Name: "Vegetta", Age: 10}
	updateResult := &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1, UpsertedCount: 0, UpsertedID: id}
	adapterMock.On("UpsertUser", id, data).Return(updateResult, nil)

	//Execution Phase
	err := userService.UpdateUser(id.Hex(), data)

	//Assert Phase
	assert.Nil(t, err)
}

//DeleteUserById
func Test_ShouldGetInvalidIdErrorErrorWhenDeletingAnUserWithAnInvalidId(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	id := "fail-id"
	adapterMock.On("DeleteUserById", id).Return(errors.New("general error update user"))

	//Execution Phase
	err := userService.DeleteUserById(id)

	//Assert Phase
	expectedErrorMsg := "the provided hex string is not a valid ObjectID"

	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}

//DeleteUserById
func Test_ShouldGetUnexpectedErrorWhenDeletingAnUser(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	id := primitive.NewObjectID()
	adapterMock.On("DeleteUserById", id).Return(nil, errors.New("general error update user"))

	//Execution Phase
	err := userService.DeleteUserById(id.Hex())

	//Assert Phase
	expectedErrorMsg := "general error update user"

	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}

//DeleteUserById
func Test_ShouldSucceedDeletingAnExistingUserAnUserWhenReceiveValidID(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	id := primitive.NewObjectID()
	mongoDeletedResult := &mongo.DeleteResult{DeletedCount: 1}
	adapterMock.On("DeleteUserById", id).Return(mongoDeletedResult, nil)

	//Execution Phase
	err := userService.DeleteUserById(id.Hex())

	//Assert Phase
	assert.Nil(t, err)
}

//DeleteUserById
func Test_ShouldSucceedDeletingANonExistingUserAnUserWhenReceiveAValidIdThatIsNotInTheDB(t *testing.T) {
	// Initializations Phase
	adapterMock := new(userRepositoryMock)
	userService := adapter.InitializeUserAdapter(adapterMock)

	//Mock Setups Phase
	id := primitive.NewObjectID()
	mongoDeletedResult := &mongo.DeleteResult{DeletedCount: 0}
	adapterMock.On("DeleteUserById", id).Return(mongoDeletedResult, nil)

	//Execution Phase
	err := userService.DeleteUserById(id.Hex())

	//Assert Phase
	assert.Nil(t, err)
}
