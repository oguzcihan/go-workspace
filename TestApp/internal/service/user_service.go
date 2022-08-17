package service

import (
	"TestApp/internal/dtos"
	. "TestApp/internal/model"
	. "TestApp/internal/repository"
	. "TestApp/internal/utils"
	. "context"
	"time"
)

//	type UserService interface {
//		Create(context Context, user *User) (*User, error)
//	}
var (
	ErrorUserAlreadyExists = BadRequest("user_already_exists")
)

func NewUserService(_repository UserRepository) *UserService {
	return &UserService{repository: _repository}
}

type UserService struct {
	repository UserRepository
}

func (service UserService) Create(context Context, userDto dtos.UserDto) (*User, *CustomError) {
	//TODO:id li user kayıt edildi bildirimi verilebilir
	//username unique check
	//iş akışı gerektiren şeyler service de olmalı
	createData := User{
		TcNo:      userDto.TcNo,
		UserName:  userDto.UserName,
		Firstname: userDto.Firstname,
		Lastname:  userDto.Lastname,
		Email:     userDto.Email,
		Password:  userDto.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}
	err := service.CheckUserName(userDto.UserName)
	if err != nil {
		return nil, ErrorUserAlreadyExists
	}
	userCreate, _ := service.repository.Create(context, &createData)
	return userCreate, nil
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
		UpdatedAt: time.Now(),
	}
	return service.repository.Update(ctx, &updateData)
}

func (service *UserService) CheckUserName(userName string) *CustomError {
	//service katmanında olmalı
	//send to service layer for username
	resUsername, err := service.GetUsername(userName)
	if err != nil {
		return nil
	} else if resUsername.UserName == userName {
		//Throw an error if the response from the DB and the response from the request are equal
		return ErrorUserAlreadyExists
	}
	return nil
}
