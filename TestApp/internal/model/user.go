package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int            `json:"id"`
	TcNo      string         `json:"tcNo"`
	UserName  string         `json:"userName"`
	Firstname string         `json:"firstName"`
	Lastname  string         `json:"lastName"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	IsActive  bool           `json:"isActive"`
	Role      string         `json:"role"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:false"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type Authentication struct {
	UserName string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	UserName    string `json:"email"`
	TokenString string `json:"token"`
}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}
