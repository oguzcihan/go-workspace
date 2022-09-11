package repository

import (
	. "ExampleProject/internal/dtos"
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
func (u UserRepository) Save(user *User) (*User, error) {
	//update kullanarak hangi kolonların güncellenmek istediği verilmeli
	saveUser := u.DB.Omit("created_at").Where("id", user.ID).Save(&user).Error
	if saveUser != nil {
		return nil, saveUser
	}
	return user, nil
}

func (u UserRepository) Delete(id int) error {
	var user User
	deleteErr := u.DB.Where("id", id).Delete(&user).Error
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}

func (u UserRepository) GetAll(pagination *Pagination) (*Pagination, error) {
	var users []User
	var totalRows int64 = 0
	//get data with limit,order
	offset := (pagination.Page * pagination.Limit) - pagination.Limit

	order := fmt.Sprintf("%s %s", pagination.OrderBy, pagination.Sort)
	findUser := u.DB.Limit(pagination.Limit).Offset(offset).Order(order)
	filters := pagination.Filters
	if filters != nil {
		for _, value := range filters {
			column := value.Column
			query := value.Query

			buildQuery := fmt.Sprintf("%s=?", column)
			findUser = findUser.Where(buildQuery, query)
		}
	}
	errCount := u.DB.Model(&User{}).Count(&totalRows).Error
	if errCount != nil {
		return nil, errCount
	}
	pagination.TotalRows = totalRows

	findUser = findUser.Find(&users)
	pagination.Rows = users
	pagination.TotalCount = findUser.RowsAffected

	return pagination, nil
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
