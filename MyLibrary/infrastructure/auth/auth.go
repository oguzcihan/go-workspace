package auth

type IAuthentication interface {
	Create(uint64, *TokenDetails)
	Fetch(string) (uint64, error)
	DeleteRefresh(string) error
	DeleteTokens(*AccessDetails) error
}

type AccessDetails struct {
	TokenUuid string
	UserId    uint64
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUuid    string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}
