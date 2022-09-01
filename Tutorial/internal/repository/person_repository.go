package repository

import (
	. "Tutorial/internal/models"
	"gorm.io/gorm"
)

func NewPersonRepository(database *gorm.DB) *PersonRepository {
	//error olmalı
	return &PersonRepository{DB: database}
}

type PersonRepository struct {
	DB *gorm.DB
}

func (u PersonRepository) Create(person *Person) *Person {
	u.DB.Create(&person)
	return person
}
