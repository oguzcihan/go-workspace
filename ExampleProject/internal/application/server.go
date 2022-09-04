package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.SetTrustedProxies([]string{" 127.0.0.1"})

	fmt.Println("------gin routes connecting-------")
	UserRoute(router)
	fmt.Println("-----connected-------")

	router.Run(":9090")
}
