package models

import (
	"ExampleProject/internal/dtos/user"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int
	TcNo      string
	UserName  string
	Firstname string
	Lastname  string
	Email     string
	Password  string
	IsActive  bool
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:false"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	userslist []user.UserList
}
