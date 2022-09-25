package account

type Login struct {
	Username string `json:"username" validate:"required,gt=3"`
	Password string `json:"password" validate:"required,gt=4,lt=9"`
}
