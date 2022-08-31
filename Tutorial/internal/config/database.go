package config

import (
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func loadDatabase() *gorm.DB {
	pgConnection := "user=postgres password=postgres host=localhost port=5432 dbname=testData sslmode=disable"
	DB, _ := gorm.Open(pg.Open(pgConnection), &gorm.Config{})
	return DB
}
func DatabaseConnection(x interface{}) *gorm.DB {
	database := loadDatabase()
	err := database.AutoMigrate(x)
	if err != nil {
		return nil
	}
	return database
}
