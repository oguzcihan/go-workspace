package application

import "github.com/gin-gonic/gin"

func PersonRoute(r *gin.Engine, ctx *gin.Context) {
	application, err := NewApplication(ctx)
	if err != nil {
		return
	}

	r.GET("/people/", func(context *gin.Context) {
		application.PersonContext.Create(ctx)
	})
}
