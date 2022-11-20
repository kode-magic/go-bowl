package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func GenerateJWTForEmail(email string, duration time.Duration) string {
	secretKey := os.Getenv("SECRET")

	var mySigningKey = []byte(secretKey)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		err = fmt.Errorf("something Went Wrong: %s", err.Error())
		return ""
	}

	return tokenString
}
