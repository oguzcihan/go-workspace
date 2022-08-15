package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int    `json:"id"`
	TcNo      string `json:"tcNo"`
	UserName  string `json:"userName"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	//IsActive  bool           `json:"isActive"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

//updatedat ve deletedat hiç oluşmadan tarih atmamalı
