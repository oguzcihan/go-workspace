package controllers

import (
	"TestApp/database"
	"TestApp/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Yeni kullanıcı oluşturur
func CreatePerson(writer http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body) //data gelir
	var person *entity.Person                      //entity class
	json.Unmarshal(requestBody, &person)

	database.Connector.Create(person)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated) //HTTP code gelecektir.
	json.NewEncoder(writer).Encode(person) //NewEncoder?
}
