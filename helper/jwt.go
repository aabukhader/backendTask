package helper

import (
	"os"
	"time"

	"github.com/aabukhader/backEnd/models"
	jwt "github.com/dgrijalva/jwt-go"

	// loaded for SECRET_KEY
	_ "github.com/joho/godotenv"
)

// GetToken provide access token
func GetToken(username, password string) (*models.Token, error) {
	secretKey := os.Getenv("SECRET_KEY")
	signingKey := []byte(secretKey)
	// create the access token
	token := jwt.New(jwt.SigningMethodHS256)
	// set the access token Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	// sgin the access token with the secret key
	accesstoken, err := token.SignedString(signingKey)
	if err != nil {
		return nil, err
	}
	// create the refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	// set the refresh token Claims
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// sgin the refresh token with the secret key
	refreshtoken, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}
	tokens := &models.Token{}
	tokens.AccessToken = accesstoken
	tokens.RefreshToken = refreshtoken
	return tokens, nil
}
