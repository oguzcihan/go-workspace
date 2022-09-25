package account

type Token struct {
	Role        string `json:"role"`
	UserName    string `json:"userName"`
	TokenString string `json:"token"`
}
