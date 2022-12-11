package dotenv

import (
	zaplog "MyLibrary/infrastructure/zap_logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func LoadDotEnvFile() {
	if err := godotenv.Load(); err != nil {
		zaplog.Logger.Error("no_env_file", zap.Error(err))
	}
}
