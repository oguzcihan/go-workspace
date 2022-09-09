package config

import (
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func connectDatabase() *gorm.DB {
	//her yerde kullanmak için ayrı alınmalı
	pgConnection := os.Getenv("PgConnection")
	DB, err := gorm.Open(pg.Open(pgConnection), &gorm.Config{})
	if err != nil {
		log.Fatal("connectDatabase():", err)
		return nil
	}
	//database = DB
	return DB
}

func DatabaseConnection(x interface{}) *gorm.DB {
	database := connectDatabase()
	err := database.AutoMigrate(x)
	if err != nil {
		log.Fatal("Log_DatabaseConnection:", err)
		return nil
	}
	return database
}
