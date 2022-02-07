package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// https://tools.ietf.org/html/rfc7519

var jwtSecret = []byte("password")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	expireAt := time.Now().Add(time.Second * 180).Unix()

	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    "wuji",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func main() {
	token, err := GenerateToken("wujimaster")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)

	rawClaim, err := ParseToken(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rawClaim.Username)
}