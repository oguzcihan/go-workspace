package config

import (
	"github.com/joho/godotenv"
	pq "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func ConnectDB() {
	godotenv.Load(".env")
	DB, err := gorm.Open(pq.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = DB
}

func GetDB() *gorm.DB {
	return db
}
