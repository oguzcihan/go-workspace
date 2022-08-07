package server

import (
	"TestApp/pkg/config"
	. "TestApp/pkg/models"
	"TestApp/pkg/routes"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

func RunServer() {
	dbConnection(&Book{})

	r := mux.NewRouter()
	//TODO:route başlamazsa ne olucak?
	fmt.Println("connecting to routes")
	routes.RegisterBookStoreRoutes(r)
	fmt.Println("connected")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9090", r))
}

func dbConnection(x interface{}) {
	config.ConnectDB()
	db = config.GetDB()
	err := db.AutoMigrate(x)
	if err != nil {
		return
	}
}
