package models

// User Struct
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Token     *Token `json:"token"`
}

// Token Struct
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Authentication Struct
type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
