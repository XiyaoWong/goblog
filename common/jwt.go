// Package common provides ...
package common

import (
	"goblog/config"
	"goblog/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims ...
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

// ReleaseToken 生成用户认证token
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(config.Expriation)
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "wongxy",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JWTKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 校验token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(config.JWTKey), nil
	})

	return token, claims, err
}
