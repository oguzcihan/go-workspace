package config

import (
	pq "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectDatabase() *gorm.DB {
	pgConnection := os.Getenv("PgConnection")
	DB, err := gorm.Open(pq.Open(pgConnection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//database = DB
	return DB
}

func DatabaseConnection(x interface{}) *gorm.DB {
	database := ConnectDatabase()
	err := database.AutoMigrate(x)
	if err != nil {
		return nil
	}
	return database
}
