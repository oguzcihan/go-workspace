package database

import (
	"TestApp/entity"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

//CRUD operasyonları için değişken oluşturuldu
var Connector *gorm.DB

//Connect Postgres
func Connect(connectionString string) error {
	var err error
	godotenv.Load(".env") //load yapılamsı gerekir
	Connector, err = gorm.Open(os.Getenv("SQL_DIALECT"), os.Getenv("CONNECTION_STRING"))
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
