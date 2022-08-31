package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	fmt.Println("Gin routes connected")
	PersonRoute(r)
}
