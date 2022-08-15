package service

import (
	"TestApp/internal/dtos"
	. "TestApp/internal/model"
	. "TestApp/internal/repository"
	. "context"
	"errors"
)

//	type UserService interface {
//		Create(context Context, user *User) (*User, error)
//	}
var (
	ErrorUserAlreadyExist = errors.New("user_already_exists")
)

func NewUserService(_repository UserRepository) *UserService {
	return &UserService{repository: _repository}
}

type UserService struct {
	repository UserRepository
}

func (service UserService) Create(context Context, userDto dtos.UserDto) (*User, error) {
	//TODO:id li user kayıt edildi bildirimi verilebilir
	//username unique check
	//service katmanında func ile birbirini çağırmalı
	//iş akışı gerektiren şeyler service de olmalı
	err := service.CheckUserName(userDto.UserName)
	if errors.Is(ErrorUserAlreadyExist, err) {
		return nil, err.Error()
	}

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

func (service UserService) GetUsername(userName string) (*User, error) {
	//Send to incoming username repository layer
	resUsername, err := service.repository.GetUsername(userName)
	if err != nil {
		return nil, err
	}
	return resUsername, nil
}

func (service UserService) Update(ctx Context, userDto dtos.UserDto) (*User, error) {
	//user var mı getusername
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

func (service *UserService) CheckUserName(userName string) error {
	//service katmanında olmalı
	//send to service layer for username
	resUsername, err := service.GetUsername(userName)
	if err != nil {
		return err
	} else if resUsername.UserName == userName {
		//Throw an error if the response from the DB and the response from the request are equal
		return ErrorUserAlreadyExist
	}
	return nil
}
