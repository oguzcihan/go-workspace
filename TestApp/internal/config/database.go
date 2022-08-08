package config

import (
	"github.com/joho/godotenv"
	pq "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var database *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	DB, err := gorm.Open(pq.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database = DB
}

func DatabaseConnection(x interface{}) *gorm.DB {
	//database classında çalışmalı
	ConnectDatabase()
	err := database.AutoMigrate(x)
	if err != nil {
		return nil
	}
	return database
}
