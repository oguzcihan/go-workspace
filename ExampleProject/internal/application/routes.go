package application

import (
	"github.com/gin-gonic/gin"
	"log"
)

func UserRoute(r *gin.Engine) {
	application, err := NewApplication()
	if err != nil {
		log.Fatal("Log_NewApplication():", err)
		return
	}

	r.POST("/user/", application.UserHandler.Create)

	r.PUT("/user/:id", application.UserHandler.Save)

	r.GET("/user/", application.UserHandler.GetAllUser)

	r.DELETE("/user/:id", application.UserHandler.GetAllUser)
}
