package service

import (
	. "TestApp/internal/model"
	. "TestApp/internal/repository"
	. "context"
)

//type UserService interface {
//	Create(context Context, user *User) (*User, error)
//}

func NewUserService(_repository UserRepository) *UserService {
	return &UserService{repository: _repository}
}

type UserService struct {
	repository UserRepository
}

func (service UserService) Create(context Context, user *User) (*User, error) {
	//TODO:id li user kayıt edildi bildirimi verilebilir
	return service.repository.Create(context, user)
}
