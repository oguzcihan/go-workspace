package Restfull

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var (
	getUrl  = "https://jsonplaceholder.typicode.com/todos/1"
	postUrl = "https://jsonplaceholder.typicode.com/todos"
)

func Demo() {
	response, err := http.Get(getUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() //işlem bittikten sonra kanalı kapat. gereksiz kanal açıklığnda performansı etkiler

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	//bodyString := string(bodyBytes)
	//json.NewDecoder(response.Body)
	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)

}

func Demo2() {
	todo := Todo{UserId: 1, Id: 2, Title: "Test", Completed: false}
	jsonTodo, _ := json.Marshal(todo)
	response, _ := http.Post(postUrl, "application/json", bytes.NewBuffer(jsonTodo))
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	//bodyString := string(bodyBytes)
	//json.NewDecoder(response.Body)
	var todoRes Todo
	json.Unmarshal(bodyBytes, &todoRes)
	fmt.Println(todo)
}
