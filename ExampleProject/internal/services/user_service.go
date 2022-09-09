package services

import (
	"ExampleProject/internal/dtos"
	. "ExampleProject/internal/models"
	. "ExampleProject/internal/repository"
	. "ExampleProject/internal/utils"
	"fmt"
	"net/http"
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
	//code olmamalı
	return &dtos.Response{Success: true, Code: 201, Data: userData}, nil
}

func (service UserService) Save(dto dtos.UserDto, userId int) (*dtos.Response, error) {
	getUser, _ := service.getUserById(userId)
	if getUser.ID == userId {
		//nil değilse username db'de mevcut
		err := service.checkUserName(dto.UserName, getUser.UserName)
		if err == nil {
			updateData := User{
				ID:        userId,
				TcNo:      dto.TcNo,
				UserName:  dto.UserName,
				Firstname: dto.Firstname,
				Lastname:  dto.Lastname,
				Email:     dto.Email,
				Password:  dto.Password,
				UpdatedAt: time.Now(),
				IsActive:  dto.IsActive,
			}

			userSave := service.repository.Save(&updateData)
			var userData = userSave.Result.(*User)
			return &dtos.Response{Success: true, Code: http.StatusOK, Data: userData}, nil
		}

	}
	return nil, ErrorUserAlreadyExists
}

func (service UserService) Delete(id int) (*dtos.Response, error) {
	getUser, _ := service.getUserById(id)
	if getUser.ID != id {
		return nil, ErrorUserNotFound
	}

	result := service.repository.Delete(id)
	if result != nil {
		return nil, NewError(result.Error.Error(), http.StatusBadRequest)
	}
	//data nil gidiyor nasıl olmalı?
	return &dtos.Response{Success: true, Code: http.StatusOK}, nil

}

func (service UserService) GetAllUser(userList *dtos.UserList) (*dtos.UserList, error) {
	getResult := service.repository.GetAll(userList)
	if getResult.Error != nil {
		return nil, NewError(getResult.Error.Error(), http.StatusBadRequest)
	}

	var data = getResult.Result.(*dtos.UserList)
	searchQuery := ""

	for _, search := range userList.Searches {
		searchQuery += fmt.Sprintf("&%s=%s", search.Column, search.Query)
	}
	usersCount := len(data.Users)
	//search silinecek
	return &dtos.UserList{Users: data.Users,
		Searches: data.Searches, TotalCount: usersCount}, nil
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
func (service UserService) getUserById(userId int) (*User, error) {
	resGetUserById, err := service.repository.GetUserId(userId)
	if err != nil {
		return nil, err
	}
	return resGetUserById, nil
}
