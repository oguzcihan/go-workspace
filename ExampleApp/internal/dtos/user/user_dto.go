package user

type UserDto struct {
	ID        int    `json:"id"`
	UserName  string `json:"userName" validate:"required,gt=3,lt=10"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gt=4,lt=9"`
	IsActive  bool   `json:"isActive"`
	Role      string `json:"role"`
}
