package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var key = []byte("abcdefasdf")


func main() {
	cToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 2),
		"uid": 11111,
	})
	tokenStr, err := cToken.SignedString(key)
	fmt.Print(tokenStr, err)

	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("not authorization")
		}
		return key, nil
	})
	if !token.Valid {
		fmt.Println("false")
	} else {
		mapData := token.Claims.(jwt.MapClaims)
		uid := mapData["uid"]
		exp := mapData["exp"]
		fmt.Println(uid, exp)
	}
	fmt.Println(token.Signature)

}


