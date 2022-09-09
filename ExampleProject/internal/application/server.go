package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func StartServer() {
	router := gin.Default()
	errRouter := router.SetTrustedProxies([]string{" 127.0.0.1"})
	if errRouter != nil {
		log.Fatal("Log_SetTrustedProxies:", errRouter)
		return
	}

	fmt.Println("------gin routes connecting-------")
	UserRoute(router)
	fmt.Println("-----connected-------")

	errRun := router.Run(":9090")
	if errRun != nil {
		log.Fatal("Log_router_run:", errRun)
		return
	}

	//token entegre
	//auth
	//register
	//login ve login ile erişilen endpoint
	//jwt
	//zap kullanılacak
}
