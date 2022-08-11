package dtos

type UserDto struct {
	TcNo      int64  `json:"tcNo" validate:"required"`
	UserName  string `json:"userName" validate:"required"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password"`
}
