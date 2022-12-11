package entities

import (
	"MyLibrary/infrastructure/security"
	"gorm.io/gorm"
	"html"
	"strings"
	"time"
)

type User struct {
	ID        uint64
	UserName  string
	Firstname string
	Lastname  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:false;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (user *User) Prepare() {
	user.Firstname = html.EscapeString(strings.TrimSpace(user.Firstname))
	user.Lastname = html.EscapeString(strings.TrimSpace(user.Lastname))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

type PublicUser struct {
	UserName  string `json:"userName"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (user *User) BeforeSave() error {
	hashPassword, err := security.GenerateHashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPassword
	return nil
}

type Users []User

func (users Users) PublicUsers() []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.PublicUser()
	}
	return result
}

func (user *User) PublicUser() interface{} {
	return &PublicUser{
		UserName:  user.UserName,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	}
}
