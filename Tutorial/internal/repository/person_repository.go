package repository

import "gorm.io/gorm"

func NewPersonRepository(database *gorm.DB) *PersonRepository {
	//error olmalı
	return &PersonRepository{DB: database}
}

type PersonRepository struct {
	DB *gorm.DB
}
