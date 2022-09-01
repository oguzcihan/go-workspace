package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	fmt.Println("Gin routes connected")
	PersonRoute(r)
	fmt.Println("Connected")

	r.Run(":9090")
}
