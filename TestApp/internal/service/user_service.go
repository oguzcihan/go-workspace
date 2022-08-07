package service

import (
	. "TestApp/internal/model"
	. "TestApp/internal/repository"
	. "context"
)

type UserService interface {
	Create(context Context, user *User) (*User, error)
}

func NewUserService(_repository UserRepository) UserService {
	return &userService{repository: _repository}
}

type userService struct {
	repository UserRepository
}

func (service *userService) Create(context Context, user *User) (*User, error) {
	//TODO:id li user kayıt edildi bildirimi verilebilir
	return service.repository.Create(context, user)
}
