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

// Save update the user
func (u UserRepository) Save(ctx Context, user *User) (*User, error) {
	//update kullanarak hangi kolonların güncellenmek istediği verilmeli
	saveUser := u.DB.Omit("created_at").Save(&user)
	if saveUser == nil {
		return nil, saveUser.Error
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

func (u UserRepository) GetUserId(userId int) (*User, error) {
	var user User
	errUserId := u.DB.Where("id=?", userId).First(&user)
	if errUserId == nil {
		return nil, errUserId.Error
	}
	return &user, nil
}
