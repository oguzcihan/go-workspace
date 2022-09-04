package application

import "github.com/gin-gonic/gin"

func UserRoute(r *gin.Engine) {
	application, err := NewApplication()
	if err != nil {
		return
	}

	r.POST("/user/", func(context *gin.Context) {
		application.UserContext.Create(context)
	})

	r.GET("/user/", func(context *gin.Context) {
		application.UserContext.GetAllUser(context)
	})
}
