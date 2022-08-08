package app

import (
	"TestApp/internal/config"
	"TestApp/internal/handler"
	. "TestApp/internal/model"
	"TestApp/internal/repository"
	"TestApp/internal/service"
	. "context"
)

//var database *gorm.DB

type ApplicationContext struct {
	User handler.UserHandler
}

func NewApplication(context Context) (*ApplicationContext, error) {

	database := config.DatabaseConnection(&User{}) //dönüş olarak database gönderir

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(*userRepository)
	userHandler := handler.NewUserHandler(*userService)
	//exit yapılabilir app boş dönebilir error msg

	return &ApplicationContext{User: userHandler}, nil
}
