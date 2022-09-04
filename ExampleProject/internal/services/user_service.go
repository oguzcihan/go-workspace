package services

import (
	"ExampleProject/internal/dtos"
	. "ExampleProject/internal/models"
	. "ExampleProject/internal/repository"
	. "ExampleProject/internal/utils"
	"time"
)

func NewUserService(_repository UserRepository) UserService {
	return UserService{repository: _repository}
}

type UserService struct {
	repository UserRepository
}

var (
	ErrorUserAlreadyExists = NewError("user_already_exists", 400)
	ErrorUserNotFound      = NewError("user_not_found", 400)
	SuccessUserDelete      = NewError("success_user_delete", 200)
)

func (service UserService) Create(dto dtos.UserDto) (*dtos.Response, error) {

	createData := User{
		TcNo:      dto.TcNo,
		UserName:  dto.UserName,
		Firstname: dto.Firstname,
		Lastname:  dto.Lastname,
		Email:     dto.Email,
		Password:  dto.Password,
		CreatedAt: time.Now(),
		IsActive:  true,
	}
	err := service.checkUserName(dto.UserName, "")
	if err != nil {
		return nil, ErrorUserAlreadyExists
	}
	userCreate := service.repository.Create(&createData)
	var userData = userCreate.Result.(*User)

	return &dtos.Response{Success: true, Data: userData}, nil
}

func (service UserService) checkUserName(newUserName string, oldUserName string) error {
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
