package user

type SuccessMessage struct {
	UserId string `json:"userId"`
	Status int    `json:"code"`
	Err    error  `json:"-"`
}

func (c SuccessMessage) Error() string {
	return c.Err.Error()
}

func NewSuccessMessage(_userId string, code int) error {
	return SuccessMessage{
		UserId: _userId,
		Status: code,
	}
}
