// Package common provides ...
package common

import (
	"github.com/dgrijalva/jwt-go"
	"goblog/config"
	"goblog/model"
	"time"
)

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 生成用户认证token
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(config.EXPRIATION)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "wongxy",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JWT_KEY))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 校验token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(config.JWT_KEY), nil
	})

	return token, claims, err
}
