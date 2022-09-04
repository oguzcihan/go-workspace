package models

import "time"

type Person struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	City      string `json:"city"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Persons []Person
