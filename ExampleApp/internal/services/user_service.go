package services

import (
	"ExampleApp/internal/dtos"
	"ExampleApp/internal/dtos/user"
	. "ExampleApp/internal/helpers"
	. "ExampleApp/internal/models"
	. "ExampleApp/internal/repository"
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

//func (service UserService) Create(dto user.UserDto) (*User, error) {
//
//	createData := User{
//		UserName:  dto.UserName,
//		Firstname: dto.Firstname,
//		Lastname:  dto.Lastname,
//		Email:     dto.Email,
//		Password:  dto.Password,
//		CreatedAt: time.Now(),
//		IsActive:  true,
//	}
//	checkUserErr := service.checkUserName(dto.UserName, "")
//	if checkUserErr != nil {
//		return nil, ErrorUserAlreadyExists
//	}
//	createUser, err := service.repository.Create(&createData)
//	if err != nil {
//		return nil, err
//	}
//
//	return createUser, nil
//}

func (service UserService) Save(userDto user.UserDto, userId int) (*User, error) {
	getUser, _ := service.getUserById(userId)
	if getUser.ID == userId {
		//nil değilse username db'de mevcut
		err := service.checkUserName(userDto.UserName, getUser.UserName)
		if err == nil {
			updateData := User{
				ID:        userId,
				UserName:  userDto.UserName,
				Firstname: userDto.Firstname,
				Lastname:  userDto.Lastname,
				Email:     userDto.Email,
				Password:  userDto.Password,
				UpdatedAt: time.Now(),
				IsActive:  userDto.IsActive,
			}

			userSave, saveErr := service.repository.Save(&updateData)
			if saveErr != nil {
				return nil, saveErr
			}
			return userSave, nil
		}

	}
	return nil, ErrorUserAlreadyExists
}

func (service UserService) Delete(id int) error {
	getUser, _ := service.getUserById(id)
	if getUser.ID != id {
		return ErrorUserNotFound
	}

	result := service.repository.Delete(id)
	if result != nil {
		return result
	}
	//data nil gidiyor nasıl olmalı?
	return SuccessUserDelete

}

func (service UserService) GetAllUser(pagination *dtos.Pagination) (*user.UserList, error) {
	getResult, err := service.repository.GetAll(pagination)
	if err != nil {
		return nil, err
	}
	//var totalPages, fromRow, toRow int64 = 0, 0, 0, 0
	//var totalPages int64 = 0
	//usersCount := len(data.Users)
	//calculate total pages
	//totalPages = int64(math.Ceil(float64(getResult.TotalRows)/float64(pagination.Limit))) - 1

	if getResult.TotalCount != 0 {
		users := user.UserList{
			Users:      getResult.Rows,
			TotalCount: getResult.TotalCount,
		}
		return &users, nil
	}
	return nil, ErrorUserNotFound
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
