package service

import (
	"errors"
	"log"
	"rpolnx.com.br/golang-with-ci/src/model/entities"
	"rpolnx.com.br/golang-with-ci/src/ports/in"
	"rpolnx.com.br/golang-with-ci/src/ports/out"
)

type userService struct {
	userPort out.UserPort
}

func (s *userService) GetAllUsers() ([]entities.User, error) {
	users, err := s.userPort.FindAllUsers()

	if err != nil {
		log.Println("Error getting all users ", err)
		return nil, err
	}

	return users, nil
}

func (s *userService) GetOneUser(id string) (*entities.User, error) {
	user, err := s.userPort.FindUserById(id)

	if err != nil {
		log.Printf("Error getting user id %s with error %v\n", id, err)
		return nil, err
	}

	return user, nil
}

func (s *userService) PostUser(e entities.User) (*string, error) {
	created, err := s.userPort.CreateUser(e)

	if err != nil {
		log.Printf("Error creating user with object %v got error %v \n", e, err)
		return nil, err
	}

	if created == nil {
		return nil, errors.New("created user did not returned id")
	}

	return created, nil
}

func (s *userService) PutUser(id string, e entities.User) error {
	err := s.userPort.UpdateUser(id, e)

	if err != nil {
		log.Printf("Error updating user id %s got error %v \n", id, err)
		return err
	}

	return nil
}

func (s *userService) DeleteUser(id string) error {
	err := s.userPort.DeleteUserById(id)

	if err != nil {
		log.Printf("Error deleting user id %s got error %v \n", id, err)
		return err
	}

	return nil
}

func InitializeUserService(userPort out.UserPort) in.UserUsecase {
	return &userService{
		userPort,
	}
}
