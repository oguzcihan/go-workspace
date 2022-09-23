package config

import (
	"ExampleProject/internal/utils"
	"go.uber.org/zap"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func connectDatabase() *gorm.DB {
	//her yerde kullanmak için ayrı alınmalı
	pgConnection := os.Getenv("PgConnection")
	DB, err := gorm.Open(pg.Open(pgConnection), &gorm.Config{})
	utils.Logger.Info("database_connection_success")
	if err != nil {
		utils.Logger.Fatal("connectDatabase()_error", zap.Error(err))
		return nil
	}
	//database = DB
	return DB
}

func DatabaseConnection(x interface{}) *gorm.DB {
	database := connectDatabase()
	//defer CloseDatabase(database)
	err := database.AutoMigrate(x)
	utils.Logger.Info("database_migrate_success")
	if err != nil {
		utils.Logger.Fatal("DatabaseConnection()_error", zap.Error(err))
		return nil
	}
	return database
}
