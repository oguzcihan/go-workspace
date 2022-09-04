package repository

import (
	. "TestApp/internal/model"
	. "context"
	"gorm.io/gorm"
)

//type UserRepository interface {
//	Create(context Context, user *User) (*User, error)
//}
/*
	Repository her model için ayrı mı olmalı?
*/

func NewUserRepository(database *gorm.DB) UserRepository {
	//error olmalı
	return UserRepository{DB: database}
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
	saveUser := u.DB.Omit("created_at").Where("id", user.ID).Save(&user)
	if saveUser == nil {
		return nil, saveUser.Error
	}
	return user, nil
}

func (u UserRepository) Delete(ctx Context, id int) error {
	var user User
	deleteUser := u.DB.Where("id=?", id).Delete(&user)
	//silinirse deleteUser dolu geliyor
	if deleteUser.Error != nil {
		return deleteUser.Error
	}
	//error verilecek
	return nil
}

func (u UserRepository) Patch(ctx Context, user *User) (int, error) {
	var userModel User
	res := u.DB.Model(&userModel).Where("id = ?", user.ID).Updates(user)
	if res.Error != nil {
		return 0, res.Error
	}
	//success/error
	return user.ID, nil
}

func (u UserRepository) GetAllUser(ctx Context) ([]User, error) {
	var users []User
	result := u.DB.Where("is_active=true").Find(&users)

	if result != nil {
		return users, nil
	}

	return nil, result.Error
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
