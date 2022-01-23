package in

import "rpolnx.com.br/golang-with-ci/src/model/entities"

type UserUsecase interface {
	GetAllUsers() ([]entities.User, error)

	GetOneUser(id string) (*entities.User, error)

	PostUser(e entities.User) (string, error)

	PutUser(id string, e entities.User) error

	DeleteUser(id string) error
}
