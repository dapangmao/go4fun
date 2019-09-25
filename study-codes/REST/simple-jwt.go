// https://stackoverflow.com/questions/36236109/go-and-jwt-simple-authentication

package main

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)

const (
	secret = "WOW,MuchShibe,ToDogge"
)

func main() {
	createdToken, err := Encode()
	if err != nil {
		fmt.Println("Creating token failed")
	}
	Decode(createdToken)
}

func Encode() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["foo"] = "bar"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func Decode(myToken string) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err == nil && token.Valid {
		fmt.Println("Your token is valid. ")
		fmt.Println(token.Claims)
	} else {
		fmt.Println("This token is NOT valid.")
	}
}
