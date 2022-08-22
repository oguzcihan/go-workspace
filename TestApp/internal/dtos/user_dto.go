package dtos

type UserDto struct {
	ID        int    `json:"id"`
	TcNo      string `json:"tcNo" validate:"required,gt=10,lt=12"`
	UserName  string `json:"userName" validate:"required,gt=3"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gt=4,lt=9"`
	IsActive  bool   `json:"isActive"`
}

/*
gt=4 en az 4 karakter
lt=9 en fazla 8 karakter
*/
