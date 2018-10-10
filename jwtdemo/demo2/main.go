package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"errors"
	"time"
)

func main() {
	cJwt := NewCJwt()
	cClaims := CustomClaims{ID: 111,
	StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 2).Unix(),
	}}

	tokenStr, err := cJwt.CreateToken(cClaims)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("token String: ", tokenStr)
	revClaims, err := cJwt.ParseToken(tokenStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", revClaims)
}

var (
	TokenExpired = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed = errors.New("That's not even a token")
	TokenInvalid = errors.New("Couldn't handle this token:")
)
const(
	keyStr = "aaabcccccsaser;iujqwer"
)

type CJWT struct {
	SigningKey []byte
}

func NewCJwt() *CJWT {
	return &CJWT{SigningKey: []byte(keyStr)}
}

func (c *CJWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(c.SigningKey)
}

func (c *CJWT) ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return c.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if (ve.Errors & jwt.ValidationErrorExpired) != 0 {
				// token is expired
				return nil, TokenExpired
			} else if (ve.Errors & jwt.ValidationErrorMalformed != 0) {
				return nil, TokenMalformed
			} else if (ve.Errors & jwt.ValidationErrorNotValidYet != 0) {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid

			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims, nil
	}
	return nil, TokenInvalid
}


type CustomClaims struct {
	jwt.StandardClaims
	ID int `json:"id"`
}