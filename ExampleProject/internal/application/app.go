package application

import (
	"ExampleProject/internal/config"
	"ExampleProject/internal/handlers"
	"ExampleProject/internal/models"
	"ExampleProject/internal/repository"
	"ExampleProject/internal/services"
)

type AppContext struct {
	UserContext handlers.IUserHandler
}

func NewApplication() (AppContext, error) {

	database := config.DatabaseConnection(&models.User{})

	userRepository := repository.NewUserRepository(database)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	return AppContext{UserContext: userHandler}, nil
}
