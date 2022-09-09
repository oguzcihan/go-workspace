package repository

import (
	"ExampleProject/internal/dtos"
	. "ExampleProject/internal/models"
	"fmt"
	"gorm.io/gorm"
)

func NewUserRepository(database *gorm.DB) UserRepository {
	//error olmalı
	return UserRepository{DB: database}
}

type UserRepository struct {
	DB *gorm.DB
}

func (u UserRepository) Create(user *User) (*User, error) {
	//service in anlayacağı şeyi göndermeli
	createUser := u.DB.Create(user).Error
	if createUser != nil {
		return nil, createUser
	}
	return user, nil
}

// Save update the user
func (u UserRepository) Save(user *User) dtos.RepositoryResult {
	//update kullanarak hangi kolonların güncellenmek istediği verilmeli
	saveUser := u.DB.Omit("created_at").Where("id", user.ID).Save(&user).Error
	if saveUser != nil {
		return dtos.RepositoryResult{Error: saveUser}
	}
	return dtos.RepositoryResult{Result: user}
}

func (u UserRepository) Delete(id int) *dtos.RepositoryResult {
	var user User
	deleteErr := u.DB.Where("id", id).Delete(&user).Error
	if deleteErr != nil {
		return &dtos.RepositoryResult{Error: deleteErr}
	}
	return nil
}

func (u UserRepository) GetAll(userList *dtos.UserList) dtos.RepositoryResult {
	var users []User

	//get data with limit,order
	//offset(page)
	findUser := u.DB.Limit(userList.Limit).Order(userList.Sort)
	searches := userList.Searches
	if searches != nil {
		for _, value := range searches {
			column := value.Column
			query := value.Query

			buildQuery := fmt.Sprintf("%s=?", column)
			findUser = findUser.Where(buildQuery, query)
		}
	}

	findUser = findUser.Find(&users)
	userList.Users = users

	return dtos.RepositoryResult{Result: userList}
}

func (u UserRepository) GetUsername(userName string) (*User, error) {
	//The current username is queried from the user table
	var user User
	errUser := u.DB.Where("user_name=?", userName).FirstOrInit(&user)
	if errUser == nil {
		return nil, errUser.Error
	}
	return &user, nil
}

func (u UserRepository) GetUserId(userId int) (*User, error) {
	var user User
	errUser := u.DB.Where("id=?", userId).First(&user)
	if errUser == nil {
		return nil, errUser.Error
	}
	return &user, nil
}
