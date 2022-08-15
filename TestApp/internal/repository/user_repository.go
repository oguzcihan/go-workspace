package repository

import (
	. "TestApp/internal/model"
	. "context"
	"gorm.io/gorm"
)

//type UserRepository interface {
//	Create(context Context, user *User) (*User, error)
//}

func NewUserRepository(database *gorm.DB) *UserRepository {
	//error olmalı
	return &UserRepository{DB: database}
}

type UserRepository struct {
	DB *gorm.DB
}

// Create the user
func (u UserRepository) Create(context Context, user *User) (*User, error) {
	//TODO: db ye kayıt yaparken exception kontrolü olmalı
	//Create a user
	createUser := u.DB.Create(&user)
	if createUser == nil {
		return nil, createUser.Error
	}
	return user, nil
}

func (u UserRepository) GetUsername(userName string) (*User, error) {
	var user User
	//The current username is queried from the user table
	err := u.DB.Where("user_name=?", userName).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update the user
func (u UserRepository) Update(ctx Context, user *User) (*User, error) {
	saveUser := u.DB.Save(&user)
	if saveUser != nil {
		return nil, saveUser.Error
	}
	return user, nil
}
