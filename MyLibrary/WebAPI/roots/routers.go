package roots

import (
	"MyLibrary/WebAPI/handlers"
	"MyLibrary/WebAPI/middleware"
	"MyLibrary/infrastructure/auth"
	"github.com/gin-gonic/gin"
)

func AllRoutes() {
	var appServices, redisService = Load()
	r := gin.Default()
	r.Use(middleware.CORSMiddleware()) //For CORS

	token := auth.NewToken()
	user := handlers.NewUserHandler(appServices.User, redisService.Auth, token)

	r.POST("/register", user.Register)

	StartApp(r)
}
