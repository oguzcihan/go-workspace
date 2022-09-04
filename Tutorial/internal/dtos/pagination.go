package dtos

type Pagination struct {
	Limit        int         `json:"limit"`
	Page         int         `json:"page"`
	Sort         string      `json:"sort"`
	TotalRows    int64       `json:"total_rows"`
	FirstPage    string      `json:"first_page"`
	PreviousPage string      `json:"previous_page"`
	NextPage     string      `json:"next_page"`
	LastPage     string      `json:"last_page"`
	FromRow      int64       `json:"from_row"`
	ToRow        int64       `json:"to_row"`
	Rows         interface{} `json:"rows"`
	Searches     []Search    `json:"searches"`
}
