package dto

import "rpolnx.com.br/golang-with-ci/src/model/entities"

type UserDTO struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (d *UserDTO) ToEntity() entities.User {
	return entities.User{
		Name: d.Name, Age: d.Age,
	}
}

type UserDtoList []UserDTO

func UserDtoListFromEntity(entities []entities.User) UserDtoList {
	list := make(UserDtoList, 0)

	for _, value := range entities {
		list = append(list, UserDtoFromEntity(value))
	}

	return list
}

func UserDtoFromEntity(entity entities.User) UserDTO {
	return UserDTO{
		ID: entity.ID.Hex(), Name: entity.Name, Age: entity.Age,
	}
}
