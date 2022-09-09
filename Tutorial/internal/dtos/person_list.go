package dtos

import "Tutorial/internal/models"

type PersonList struct {
	TotalCount int             `json:"totalCount"`
	Searches   []Search        `json:"search"`
	Persons    []models.Person `json:"personList"`
}
