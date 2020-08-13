package models

// UserStatusResSuccss Struct
type UserStatusResSuccss struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   User   `json:"data"`
}

// UserStatusResFail struct
type UserStatusResFail struct {
	Status int       `json:"status"`
	Msg    string    `json:"msg"`
	Data   [1]string `json:"data"`
}

// StatusRes struct
type StatusRes struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

// PostItemResSuccss Struct
type PostItemResSuccss struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   []*PostItem `json:"data"`
}

// Pagination struct
type Pagination struct {
	ItemPerPage  int     `json:"item_per_page"`
	Page         int     `json:"page"`
	TotalCount   int     `json:"total_count"`
	NextPageURL  string  `json:"next_page_num"`
	PrevPageURL  string  `json:"prev_page_num"`
	TotalPageNum float64 `json:"total_page_num"`
}

// PageStatusResSuccss struct
type PageStatusResSuccss struct {
	Status int        `json:"status"`
	Msg    string     `json:"msg"`
	Data   FliterData `json:"data"`
}

// FliterData struct
type FliterData struct {
	Meta  Pagination `json:"meta"`
	Items []PostItem `json:"items"`
}
