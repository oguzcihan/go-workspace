package config

import (
	helper "AwesomeProject/internal/helpers"
	"go.uber.org/zap"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func GetDatabase() *gorm.DB {
	//her yerde kullanmak için ayrı alınmalı
	pgConnection := os.Getenv("PgConnection")
	connection, err := gorm.Open(pg.Open(pgConnection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		helper.Logger.Fatal("database_url_error", zap.Error(err))
		return nil
	}

	sqlDB, errConnection := connection.DB()
	errConnection = sqlDB.Ping()
	if errConnection != nil {
		helper.Logger.Fatal("connection_error", zap.Error(errConnection))
	}

	helper.Logger.Info("database_connection_success")

	return connection
}

func InitialMigration(x interface{}) {
	databaseConnection := GetDatabase()
	defer closeDatabase(databaseConnection)
	err := databaseConnection.AutoMigrate(x)
	helper.Logger.Info("database_migrate_success")
	if err != nil {
		helper.Logger.Fatal("migrate_error", zap.Error(err))
	}
}

func closeDatabase(connection *gorm.DB) {
	sqlDB, err := connection.DB()
	err = sqlDB.Close()
	if err != nil {
		helper.Logger.Fatal("closeDatabase_error", zap.Error(err))
	}
}
