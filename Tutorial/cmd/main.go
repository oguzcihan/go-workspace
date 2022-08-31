package main

import (
	"Tutorial/internal/application"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	application.RunServer()

	//r := gin.Default()
	//r.POST("/people", CreatePerson)
}

//func CreatePerson(c *gin.Context) {
//
//	var person models.Person
//	c.BindJSON(&person)
//
//	db.Create(&person)
//	c.JSON(200, person)
//}
