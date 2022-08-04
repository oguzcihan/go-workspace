package model

type User struct {
	//json formatında ve adı id olacak
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname""`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
