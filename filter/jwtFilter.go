package filter

import (
	"errors"
	"fmt"
	"log"
	"time"
	"urlShrtGo/util"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("JWT secret key")

func CreateToken(id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"uid": id,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
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

func GetTokenId(tokenString string) (float64, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		util.Error.Println("error getting token", err)
		return -1, errors.New(err.Error())
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		util.Error.Println("Failed to parse claims")
		return -1, errors.New("Failed to parse claims")
	}
	util.Debug.Println(claims)

	if uidRaw, ok := claims["uid"]; ok {
		if uid, ok := uidRaw.(float64); ok {
			return uid, nil
		} else {
			util.Error.Println("Failed to assert uid as float64")
			return -1, errors.New("Failed to assert uid as float64")
		}
	} else {
		util.Error.Println("Key 'uid' does not exist in the claims map")
		return -1, errors.New("Key 'uid' does not exist")
	}

}
