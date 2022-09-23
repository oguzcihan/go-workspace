package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int
	UserName  string
	Firstname string
	Lastname  string
	Email     string
	Password  string
	IsActive  bool
	Role      string
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:false"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
