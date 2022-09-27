package user

type UserList struct {
	TotalCount int64       `json:"totalCount"`
	Users      interface{} `json:"users"`
}
