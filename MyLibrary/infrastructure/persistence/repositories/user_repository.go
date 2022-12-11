package repositories

import (
	"MyLibrary/domain/entities"
	"MyLibrary/domain/repositories"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB) UserRepository {
	return UserRepository{DB: gormDb}
}

var _ repositories.IUserRepository = &UserRepository{}

func (u UserRepository) Create(user *entities.User) (*entities.User, error) {
	err := u.DB.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
