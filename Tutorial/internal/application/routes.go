package application

import "github.com/gin-gonic/gin"

func PersonRoute(r *gin.Engine) {
	application, err := NewApplication()
	if err != nil {
		return
	}

	r.POST("/people/", func(context *gin.Context) {
		application.PersonContext.Create(context)
	})
}
