package config

import (
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func connectDatabase() *gorm.DB {
	//her yerde kullanmak için ayrı alınmalı
	pgConnection := os.Getenv("PgConnection")
	DB, err := gorm.Open(pg.Open(pgConnection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//database = DB
	return DB
}

func DatabaseConnection(x interface{}) *gorm.DB {
	database := connectDatabase()
	err := database.AutoMigrate(x)
	if err != nil {
		return nil
	}
	return database
}
