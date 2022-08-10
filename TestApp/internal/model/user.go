package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int            `json:"id"`
	TcNo      int64          `json:"tcNo" validate:"required"`
	UserName  string         `json:"userName" validate:"required"`
	Firstname string         `json:"firstName"`
	Lastname  string         `json:"lastName"`
	Email     string         `json:"email" validate:"required,email"`
	Password  string         `json:"password" validate:"passwd"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
