package core

import (
	"AwesomeProject/internal/config"
	"AwesomeProject/internal/handlers"
	helper "AwesomeProject/internal/helpers"
	"AwesomeProject/internal/models"
	"AwesomeProject/internal/repository"
	"AwesomeProject/internal/services"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

var (
	port   = ":9090"
	router = mux.NewRouter()
)

func StartServer() {
	helper.InitializeLogger()
	config.InitialMigration(models.User{})

	helper.Logger.Info("---connecting routes---")
	UserAndAccountRoutes()
	helper.Logger.Info("---connected---")

	errRunServe := http.ListenAndServe(port, router)
	if errRunServe != nil {
		helper.Logger.Fatal("router_run_error:", zap.Error(errRunServe))
	}
}

func UserAndAccountRoutes() {
	//UserRepository
	userRepository := repository.NewUserRepository(config.GetDatabase())
	//UserService and UserHandler
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	UserRoute(userHandler, router)

	//AccountService and AccountHandler
	accountService := services.NewAccountService(userRepository)
	accountHandler := handlers.NewAccountHandler(accountService)
	AccountRoute(accountHandler, router)

}
