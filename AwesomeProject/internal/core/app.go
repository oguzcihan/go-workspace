package core

import (
	"AwesomeProject/internal/config"
	"AwesomeProject/internal/handlers"
	. "AwesomeProject/internal/models"
	"AwesomeProject/internal/repository"
	"AwesomeProject/internal/services"
)

type AppContext struct {
	UserHandler    handlers.IUserHandler
	AccountHandler handlers.IAccountHandler
}

func NewApplication() (AppContext, error) {

	config.InitialMigration(User{})

	//UserRepository
	userRepository := repository.NewUserRepository(config.GetDatabase())
	//UserService and UserHandler
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	//AccountService and AccountHandler
	accountService := services.NewAccountService(userRepository)
	accountHandler := handlers.NewAccountHandler(accountService)

	return AppContext{
		UserHandler:    userHandler,
		AccountHandler: accountHandler,
	}, nil
}
