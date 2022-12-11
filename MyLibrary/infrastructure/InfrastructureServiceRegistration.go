package infrastructure

import (
	"MyLibrary/domain/entities"
	"MyLibrary/infrastructure/persistence/dotenv"
	"MyLibrary/infrastructure/persistence/migration"
	"MyLibrary/infrastructure/zap_logger"
)

func AddInfrastructureService() {
	zap_logger.InitializeLogger()
	dotenv.LoadDotEnvFile()
	migration.InitialMigration(entities.User{})
}
