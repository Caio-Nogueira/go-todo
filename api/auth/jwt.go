package auth

import (
	"fmt"
	"os"
	"github.com/golang-jwt/jwt"
)


var SecretKey = []byte(os.Getenv("SECRET_KEY"))

var Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func GenerateJWT(ID string) (string, error) {
	Claims.ID = ID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	fmt.Println(Claims)
	return token.SignedString(SecretKey)
}