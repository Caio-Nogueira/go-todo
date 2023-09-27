package auth

import (
	"fmt"
	"os"
	"github.com/golang-jwt/jwt"
)


var SecretKey = []byte(os.Getenv("SECRET_KEY"))

var Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	Claims.Username = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	fmt.Println(Claims)
	return token.SignedString(SecretKey)

}