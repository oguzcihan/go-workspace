package repository

import (
	. "TestApp/internal/model"
	"context"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

// Constructor
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

type userRepository struct {
	DB *gorm.DB
}

// Create database veri kayıt işlemi yapar
func (u *userRepository) Create(ctx context.Context, user *User) (int64, error) {
	res := u.DB.Create(&user)
	//TODO: context kullanımı araştırılacak
	return res.RowsAffected, nil
}

// Update database de olan veriyi güncelleme işlemi yapar
func (u *userRepository) Update(ctx context.Context, user *User) (int64, error) {
	res := u.DB.Save(&user)
	//error olmalı
	return res.RowsAffected, nil
}

// Delete database de olan veriyi silme işlemi yapar
func (u *userRepository) Delete(ctx context.Context, id string) (int64, error) {
	var user User
	res := u.DB.Where("id=?", id).Delete(&user)
	return res.RowsAffected, nil
}
