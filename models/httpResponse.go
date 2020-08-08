package models

// UserStatusResSuccss Struct
type UserStatusResSuccss struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   User   `json:"data"`
}

// UserStatusResFail struct
type UserStatusResFail struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   [1]string `json:"data"`
}

// StatusRes struct
type StatusRes struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}
