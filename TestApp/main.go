package main

import (
	"TestApp/controllers"
	"TestApp/database"
	"TestApp/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	initialDatabase()
	log.Println("HTTP server 9090 started")

	router := mux.NewRouter().StrictSlash(true) //localhost:9090/create, true ise erişim sağlanır
	initializeHandlers(router)
	log.Fatal(http.ListenAndServe(":9090", router))
}

func initializeHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST") //Url yolunu belirler
}

func initialDatabase() {
	config :=
		database.Config{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			DbName:   "testAppDb",
		}
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	database.Migrate(&entity.Person{})
}
