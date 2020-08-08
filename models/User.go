package models

// User Struct
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}

// Authentication Struct
type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
