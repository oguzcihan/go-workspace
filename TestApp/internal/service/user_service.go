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
//iş akışı gerektiren şeyler service de olmalı

var (
	ErrorUserAlreadyExists = NewError("user_already_exists", 404)
	ErrorUserNotFound      = NewError("user_not_found", 404)
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
	createData := User{
		TcNo:      userDto.TcNo,
		UserName:  userDto.UserName,
		Firstname: userDto.Firstname,
		Lastname:  userDto.Lastname,
		Email:     userDto.Email,
		Password:  userDto.Password,
		CreatedAt: time.Now(),
		IsActive:  true,
	}
	err := service.CheckUserName(userDto.UserName)
	if err != nil {
		return nil, ErrorUserAlreadyExists
	}
	userCreate, _ := service.repository.Create(context, &createData)
	return userCreate, nil
}

func (service UserService) Update(ctx Context, userDto dtos.UserDto, userId int) (*User, error) {
	//user var mı getusername

	//userName := service.CheckUserName(userDto.UserName)
	//Id := service.GetUserById(userId)

	updateData := User{
		ID:        userId,
		TcNo:      userDto.TcNo,
		UserName:  userDto.UserName,
		Firstname: userDto.Firstname,
		Lastname:  userDto.Lastname,
		Email:     userDto.Email,
		Password:  userDto.Password,
		UpdatedAt: time.Now(),
		IsActive:  userDto.IsActive,
	}

	return service.repository.Update(ctx, &updateData)

	//user := service.repository.

}

func (service UserService) CheckUserName(userName string) error {
	//resUsername, err := service.GetUsername(userName)
	resUsername, err := service.repository.GetUsername(userName)
	//send to service layer for username
	if err != nil {
		return nil
	} else if resUsername.UserName == userName {
		//Throw an error if the response from the DB and the response from the request are equal
		return ErrorUserAlreadyExists
	}
	return nil
}

func (service UserService) GetUserById(userId int) (*User, error) {
	resUserId, err := service.repository.GetUserId(userId)
	if err != nil {
		return nil, err
	}
	return resUserId, nil
}
