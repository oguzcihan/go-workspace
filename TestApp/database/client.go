package database

import (
	"TestApp/entity"
	"log"

	"github.com/jinzhu/gorm"
)

//CRUD operasyonları için değişken oluşturuldu
var Connector *gorm.DB

//Connect Postgres
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	log.Println("Connection was successful")
	return nil
}

func Migrate(table *entity.Person) {

	Connector.AutoMigrate(&table)
	log.Println("Table Migrated Success")
}
