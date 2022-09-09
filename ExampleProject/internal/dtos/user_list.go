package dtos

import "ExampleProject/internal/models"

type UserList struct {
	Limit      int           `json:"-"`
	Page       int           `json:"-"`
	Sort       string        `json:"-"`
	TotalCount int           `json:"totalCount"`
	Searches   []Search      `json:"search"`
	Users      []models.User `json:"userList"`
}
