package dtos

type Filter struct {
	Column string `json:"column"`
	Query  string `json:"query"`
}
