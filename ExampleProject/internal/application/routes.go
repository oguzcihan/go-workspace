package application

import (
	"ExampleProject/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UserRoute(r *gin.Engine) {
	application, err := NewApplication()
	if err != nil {
		utils.Logger.Fatal("application_error", zap.Error(err))
		return
	}

	r.POST("/user/", application.UserHandler.Create)

	r.PUT("/user/:id", application.UserHandler.Save)

	r.GET("/user/", application.UserHandler.GetAllUser)

	r.DELETE("/user/:id", application.UserHandler.GetAllUser)
}
