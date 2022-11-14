package migration

import (
	"MyLibrary/infrastructure/zap_logger"
	"MyLibrary/persistence/postgres"
)

func InitialMigration(entity interface{}) {
	databaseConnection := postgres.ConnectDatabase()
	err := databaseConnection.AutoMigrate(entity)
	zap_logger.Logger.Info("database_migrate_success")
	if err != nil {
		zap_logger.Logger.Fatal("migration_error")
	}
}
