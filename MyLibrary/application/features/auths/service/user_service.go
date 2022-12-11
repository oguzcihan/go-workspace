package service

import (
	"MyLibrary/domain/entities"
	"MyLibrary/domain/repositories"
)

type UserService struct {
	userService repositories.IUserRepository
}

var _ IUserService = &UserService{}

type IUserService interface {
	Create(*entities.User) (*entities.User, error)
}

func (us UserService) Create(user *entities.User) (*entities.User, error) {
	//write here business logic code

	userData, dbErr := us.userService.Create(user)
	if dbErr != nil {
		return nil, dbErr
	}

	return userData, nil
}
