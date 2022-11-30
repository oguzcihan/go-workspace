package repositories

import (
	"MyLibrary/domain/entities"
	"gorm.io/gorm"
)

func NewUserRepository(gormDb *gorm.DB) UserRepository {
	return UserRepository{DB: gormDb}
}

type UserRepository struct {
	DB *gorm.DB
}

func (u UserRepository) Create(user *entities.User) (*entities.User, error) {
	err := u.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
