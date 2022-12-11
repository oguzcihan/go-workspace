package migration

import (
	"MyLibrary/infrastructure/persistence/postgres"
	"MyLibrary/infrastructure/zap_logger"
)

func InitialMigration(entity interface{}) {
	databaseConnection := postgres.ConnectDatabase()
	err := databaseConnection.AutoMigrate(entity)
	zap_logger.Logger.Info("database_migrate_success")
	if err != nil {
		zap_logger.Logger.Fatal("migration_error")
	}
}
