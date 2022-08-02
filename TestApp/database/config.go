package database

import "fmt"

//Db properties
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DbName)
	return connectionString

}
