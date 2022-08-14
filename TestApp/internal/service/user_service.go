package service

import (
	"TestApp/internal/dtos"
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

func (service UserService) Create(context Context, userDto dtos.UserDto) (*User, error) {
	//TODO:id li user kayıt edildi bildirimi verilebilir
	createData := User{
		TcNo:      userDto.TcNo,
		UserName:  userDto.UserName,
		Firstname: userDto.Firstname,
		Lastname:  userDto.Lastname,
		Email:     userDto.Email,
		Password:  userDto.Password,
	}
	return service.repository.Create(context, &createData)
}

func (service UserService) GetUsername(ctx Context, userName string) (*User, error) {
	//Send to incoming username repository layer
	resUsername, err := service.repository.GetUsername(ctx, userName)
	return resUsername, err
}

func (service UserService) Update(ctx Context, userDto dtos.UserDto) (*User, error) {
	updateData := User{
		ID:        userDto.ID,
		TcNo:      userDto.TcNo,
		UserName:  userDto.UserName,
		Firstname: userDto.Firstname,
		Lastname:  userDto.Lastname,
		Email:     userDto.Email,
		Password:  userDto.Password,
	}
	return service.repository.Update(ctx, &updateData)
}
