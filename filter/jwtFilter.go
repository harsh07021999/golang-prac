package filter

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("JWT secret key")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": username,
		"exprity":  time.Now().Add(time.Minute * 1).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		log.Fatal("error signing token")
		return "", err
	}
	return tokenString, nil

}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	fmt.Print(token.Claims)

	if err != nil {
		log.Fatal("error verifying token", err)
		return err
	}
	if !(token.Valid) {
		return fmt.Errorf("invalid token")
	}
	return nil
}
