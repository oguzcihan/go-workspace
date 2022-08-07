package app

import (
	"TestApp/internal/config"
	"TestApp/internal/handler"
	. "TestApp/internal/model"
	"TestApp/internal/repository"
	"TestApp/internal/service"
	. "context"
	"gorm.io/gorm"
)

var db *gorm.DB

type ApplicationContext struct {
	User handler.UserHandler
}

func NewApplication(context Context) (*ApplicationContext, error) {
	dbConnection(&User{})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	return &ApplicationContext{User: userHandler}, nil
}

func dbConnection(x interface{}) {
	config.ConnectDB()
	db = config.GetDB()
	err := db.AutoMigrate(x)
	if err != nil {
		return
	}
}
