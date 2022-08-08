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

func (u UserRepository) Create(context Context, user *User) (*User, error) {
	//TODO: db ye kayıt yaparken exception kontrolü olmalı
	resUser := u.DB.Create(&user)
	if resUser == nil {
		return nil, resUser.Error
	}
	return user, nil
}
