package account

type Register struct {
	UserName  string `json:"userName" validate:"required,gt=3"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Password  string `json:"password"`
	//password repeat
	Email string `json:"email" validate:"required,email"`
}
