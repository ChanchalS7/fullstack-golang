package auth

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(user_id uint32) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.GetEnv("API_SECRET")))
}
