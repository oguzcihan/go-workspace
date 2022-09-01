package models

type Person struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	City      string `json:"city"`
}

func (e *Person) TableName() string {
	return "person"
}
