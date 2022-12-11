package roots

import (
	"MyLibrary/infrastructure/auth"
	"MyLibrary/infrastructure/persistence/repositories"
	"MyLibrary/infrastructure/zap_logger"
	"go.uber.org/zap"
)

func Load() (*repositories.Repositories, *auth.RedisService) {
	services, err := repositories.NewRepositories()
	if err != nil {
		panic(err)
	}

	_redisService, err := auth.NewRedisDB()
	if err != nil {
		zap_logger.Logger.Fatal("auth_error", zap.Error(err))
	}

	return services, _redisService
}
