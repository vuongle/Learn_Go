package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "blog_golang_api"

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    issuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", err
	}

	return result, nil
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(
		cookie,
		jwt.RegisteredClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return "", err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	return claims.Issuer, nil
}
