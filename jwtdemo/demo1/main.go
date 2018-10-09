package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var key = []byte("abcdefasdf")


func main() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 2),
		"uid": 11111,
	})
	tokenStr, err := token.SignedString(key)
	fmt.Print(tokenStr, err)
}


