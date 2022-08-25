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
	SuccessUserDelete      = NewError("success_user_delete", 200)
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
	err := service.CheckUserName(userDto.UserName, "")
	if err != nil {
		return nil, ErrorUserAlreadyExists
	}
	userCreate, _ := service.repository.Create(context, &createData)
	return userCreate, nil
}

func (service UserService) Save(ctx Context, userDto dtos.UserDto, userId int) (*User, error) {

	getUser, _ := service.GetUserById(userId)
	if getUser.ID == userId {
		//nil değilse username db'de mevcut
		err := service.CheckUserName(userDto.UserName, getUser.UserName)
		if err == nil {
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
			return service.repository.Save(ctx, &updateData)
		}

	}
	return nil, ErrorUserAlreadyExists
}

func (service UserService) Delete(ctx Context, id int) error {
	getUser, _ := service.GetUserById(id)
	if getUser.ID != id {
		return ErrorUserNotFound
	}

	result := service.repository.Delete(ctx, id)
	if result == nil {
		return SuccessUserDelete
	}

	return nil
}

func (service UserService) CheckUserName(newUserName string, oldUserName string) error {

	if newUserName == oldUserName {
		return nil
	}

	resUsername, _ := service.repository.GetUsername(newUserName)
	//send to service layer for username
	if resUsername.UserName != "" {
		//Throw an error if the response from the DB and the response from the request are equal
		return ErrorUserAlreadyExists
	}
	return nil
}

func (service UserService) GetUserById(userId int) (*User, error) {
	resGetUserById, err := service.repository.GetUserId(userId)
	if err != nil {
		return nil, err
	}
	return resGetUserById, nil
}
