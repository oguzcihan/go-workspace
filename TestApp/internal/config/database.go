package config

import (
	pq "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var database *gorm.DB

func ConnectDatabase() {
	pgConnection := os.Getenv("PgConnection")
	DB, err := gorm.Open(pq.Open(pgConnection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database = DB
}

func DatabaseConnection(x interface{}) *gorm.DB {
	//database classında çalışmalı
	//migrate ayrı olmalı
	ConnectDatabase() //return database olarak dönmeli
	err := database.AutoMigrate(x)
	if err != nil {
		return nil
	}
	return database
}
