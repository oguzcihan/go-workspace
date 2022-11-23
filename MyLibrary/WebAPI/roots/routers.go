package roots

import (
	zaplog "MyLibrary/infrastructure/zap_logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

func AllRoutes() {
	r := gin.Default()

	appPort := os.Getenv("API_PORT")
	if appPort == "" {
		appPort = "8888"
	}
	zaplog.Logger.Fatal("port_error", zap.Error(r.Run(":"+appPort)))
}
