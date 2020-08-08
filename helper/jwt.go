package helper

import (
	"os"
	jwt "github.com/dgrijalva/jwt-go"
	// loaded for SECRET_KEY
	_ "github.com/joho/godotenv"
)

// GetToken provide access token
func GetToken(username, password string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	signingKey := []byte(secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

// VerifyToken checks if the access token is vaild
func VerifyToken(tokenString string) (jwt.Claims, error) {
	secretKey := os.Getenv("SECRET_KEY")
	signingKey := []byte(secretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
