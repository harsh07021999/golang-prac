package filter

import (
	"fmt"
	"log"
	"time"
	"urlShrtGo/util"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("JWT secret key")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 10).Unix(),
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
	fmt.Print(token.Claims.GetExpirationTime())

	if err != nil {
		util.Error.Println("error verifying token", err)
		return err
	}
	if !(token.Valid) {
		util.Error.Println("invalid token")
	}
	return nil
}
