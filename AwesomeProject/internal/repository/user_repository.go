package repository

import (
	. "AwesomeProject/internal/models"
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

func (u UserRepository) GetUsername(userName string) (*User, error) {
	//The current username is queried from the user table
	var user User
	errUser := u.DB.Where("user_name=?", userName).FirstOrInit(&user)
	if errUser == nil {
		return nil, errUser.Error
	}
	return &user, nil
}
