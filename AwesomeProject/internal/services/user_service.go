package services

import (
	"AwesomeProject/internal/dtos/userDtos"
	. "AwesomeProject/internal/helpers"
	"AwesomeProject/internal/models"
	. "AwesomeProject/internal/repository"
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

func (service UserService) Create(dto userDtos.UserDto) (*models.User, error) {

	createData := models.User{
		UserName:  dto.UserName,
		Firstname: dto.Firstname,
		Lastname:  dto.Lastname,
		Email:     dto.Email,
		Password:  dto.Password,
		CreatedAt: time.Now(),
		IsActive:  true,
	}
	checkUserErr := service.checkUserName(dto.UserName, "")
	if checkUserErr != nil {
		return nil, ErrorUserAlreadyExists
	}
	createUser, err := service.repository.Create(&createData)
	if err != nil {
		return nil, err
	}

	return createUser, nil
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
