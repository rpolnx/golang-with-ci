package adapter

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rpolnx.com.br/golang-with-ci/src/model/dto"
	"rpolnx.com.br/golang-with-ci/src/model/entities"
	"rpolnx.com.br/golang-with-ci/src/ports/out"
	"rpolnx.com.br/golang-with-ci/src/repository"
)

type userAdapter struct {
	userRepository repository.UserDBRepository
}

func (a userAdapter) FindAllUsers(pagination dto.PaginationDTO) ([]entities.User, error) {
	return a.userRepository.FindAllUsers(pagination)
}

func (a userAdapter) FindUserById(id string) (*entities.User, error) {
	mongoId, err := getMongoId(id)

	if err != nil {
		return nil, err
	}

	return a.userRepository.FindUserById(mongoId)
}

func (a userAdapter) CreateUser(e entities.User) (string, error) {
	result, err := a.userRepository.CreateUser(e)

	if err != nil {
		return "", err
	}

	id, _ := result.InsertedID.(primitive.ObjectID)

	if id.IsZero() {
		return "", errors.New("created user did not returned id")
	}

	hex := id.Hex()

	return hex, nil
}

func (a userAdapter) UpdateUser(id string, e entities.User) error {
	mongoId, err := getMongoId(id)

	if err != nil {
		return err
	}

	e.ID = mongoId
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

//func InitializeUserAdapter2(userRepository repository.UserDBRepository) out.UserPort {
//	return &userAdapter{
//		userRepository,
//	}
//}