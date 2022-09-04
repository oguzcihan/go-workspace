package repository

import (
	"ExampleProject/internal/dtos"
	"ExampleProject/internal/models"
	"gorm.io/gorm"
)

func NewUserRepository(database *gorm.DB) UserRepository {
	//error olmalı
	return UserRepository{DB: database}
}

type UserRepository struct {
	DB *gorm.DB
}

func (u UserRepository) Create(user *models.User) dtos.RepositoryResult {

	createUser := u.DB.Create(user).Error
	if createUser != nil {
		return dtos.RepositoryResult{Error: createUser}
	}
	return dtos.RepositoryResult{Result: user}
}

func (u UserRepository) GetUsername(userName string) (*models.User, error) {
	//The current username is queried from the user table
	var user models.User
	errUser := u.DB.Where("user_name=?", userName).FirstOrInit(&user)
	if errUser == nil {
		return nil, errUser.Error
	}
	return &user, nil
}
