package adapter

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rpolnx.com.br/golang-with-ci/src/model/entities"
	"rpolnx.com.br/golang-with-ci/src/ports/out"
	"rpolnx.com.br/golang-with-ci/src/repository"
)

type userAdapter struct {
	userRepository repository.UserDBRepository
}

func (a userAdapter) FindAllUsers() ([]entities.User, error) {
	return a.userRepository.FindAllUsers()
}

func (a userAdapter) FindUserById(id string) (*entities.User, error) {
	mongoId, err := getMongoId(id)

	if err != nil {
		return nil, err
	}

	return a.userRepository.FindUserById(mongoId)
}

func (a userAdapter) CreateUser(e entities.User) (*string, error) {
	result, err := a.userRepository.CreateUser(e)

	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	hex := id.Hex()

	return &hex, nil
}

func (a userAdapter) UpdateUser(id string, e entities.User) error {
	mongoId, err := getMongoId(id)

	if err != nil {
		return err
	}

	_, err = a.userRepository.UpsertUser(mongoId, e)

	return err
}

func (a userAdapter) DeleteUserById(id string) error {
	mongoId, err := getMongoId(id)

	if err != nil {
		return err
	}

	_, err = a.userRepository.DeleteUserById(mongoId)

	return err
}

func getMongoId(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

func InitializeUserAdapter(userRepository repository.UserDBRepository) out.UserPort {
	return &userAdapter{
		userRepository,
	}
}