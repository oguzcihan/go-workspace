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
	//birden fazla model olduğu durumu için generic yapmak gerekiyor
	database := config.DatabaseConnection(&User{})

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(*userRepository)
	userHandler := handler.NewUserHandler(*userService)
	//exit yapılabilir app boş dönebilir error msg

	//birden fazla model olduğunda applicationcontext nasıl yapılmalı
	return &ApplicationContext{User: userHandler}, nil
}
