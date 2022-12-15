package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"govue.demo/go_web_0/model"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTima := time.Now().Add(7 * 24 * time.Hour)
	Claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTima.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "govue.demo",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return " ", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
