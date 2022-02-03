package service

import (
	"log"
	"rpolnx.com.br/golang-with-ci/src/model/dto"
	"rpolnx.com.br/golang-with-ci/src/model/entities"
	"rpolnx.com.br/golang-with-ci/src/ports/in"
	"rpolnx.com.br/golang-with-ci/src/ports/out"
)

type userService struct {
	userPort out.UserPort
}

func (s *userService) GetAllUsers(pagination dto.PaginationDTO) ([]entities.User, error) {
	users, err := s.userPort.FindAllUsers(pagination)

	if err != nil {
		log.Println("Error getting all users ", err)
	}

	return users, err
}

func (s *userService) GetOneUser(id string) (*entities.User, error) {
	user, err := s.userPort.FindUserById(id)

	if err != nil {
		log.Printf("Error getting user id %s with error %v\n", id, err)
	}

	return user, err
}

func (s *userService) PostUser(e entities.User) (string, error) {
	created, err := s.userPort.CreateUser(e)

	if err != nil {
		log.Printf("Error creating user with object %v got error %v \n", e, err)
	}

	return created, err
}

func (s *userService) PutUser(id string, e entities.User) error {
	err := s.userPort.UpdateUser(id, e)

	if err != nil {
		log.Printf("Error updating user id %s got error %v \n", id, err)
	}

	return err
}

func (s *userService) DeleteUser(id string) error {
	err := s.userPort.DeleteUserById(id)

	if err != nil {
		log.Printf("Error deleting user id %s got error %v \n", id, err)
	}

	return err
}

func InitializeUserService(userPort out.UserPort) in.UserUsecase {
	return &userService{
		userPort,
	}
}
