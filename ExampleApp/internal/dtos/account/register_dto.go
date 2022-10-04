package account

type Register struct {
	UserName  string `json:"userName" validate:"required,gt=3"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Password  string `json:"password" validate:"required,gt=4,lt=9"`
	Role      string `json:"role"`
	//password repeat
	Email string `json:"email" validate:"required,email"`
}
