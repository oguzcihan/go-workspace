package dtos

type Pagination struct {
	Limit        int         `json:"limit"`
	Page         int         `json:"page"`
	Sort         string      `json:"sort"`
	TotalRows    int64       `json:"-"`
	TotalCount   int64       `json:"totalCount"`
	FirstPage    string      `json:"-"`
	PreviousPage string      `json:"-"`
	NextPage     string      `json:"-"`
	LastPage     string      `json:"-"`
	FromRow      int         `json:"-"`
	ToRow        int         `json:"-"`
	Start        int         `json:"start"`
	OrderBy      string      `json:"orderBy"`
	Rows         interface{} `json:"rows"`
	Filters      []Filter    `json:"filters"`
}
