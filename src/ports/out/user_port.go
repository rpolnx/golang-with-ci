package out

import (
	"rpolnx.com.br/golang-with-ci/src/model/dto"
	"rpolnx.com.br/golang-with-ci/src/model/entities"
)

type UserPort interface {
	FindAllUsers(dto.PaginationDTO) ([]entities.User, error)

	FindUserById(id string) (*entities.User, error)

	CreateUser(e entities.User) (string, error)

	UpdateUser(id string, e entities.User) error

	DeleteUserById(id string) error
}
