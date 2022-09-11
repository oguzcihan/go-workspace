package user

type UserResponseBody struct {
	ID        int    `json:"id"`
	UserName  string `json:"userName" validate:"required,gt=3"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email" validate:"required,email"`
}
