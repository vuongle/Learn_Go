package security

import (
	"github-trending-api/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "my_go_api_server_skey"

func GenerateToken(user models.User) (string, error) {
	claims := &models.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", err
	}

	return result, nil
}
