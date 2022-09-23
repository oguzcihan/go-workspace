package dtos

type Search struct {
	Column string `json:"column"`
	Query  string `json:"query"`
}
