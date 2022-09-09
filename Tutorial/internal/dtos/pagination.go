package dtos

type Pagination struct {
	Limit        int         `json:"-"`
	Page         int         `json:"-"`
	Sort         string      `json:"-"`
	TotalRows    int64       `json:"-"`
	FirstPage    string      `json:"-"`
	PreviousPage string      `json:"-"`
	NextPage     string      `json:"-"`
	LastPage     string      `json:"-"`
	FromRow      int64       `json:"-"`
	ToRow        int64       `json:"-"`
	Searches     []Search    `json:"search"`
	Rows         interface{} `json:"rows"`
}
