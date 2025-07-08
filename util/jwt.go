package util

import (
	"time"

	config "github.com/OctavianoRyan25/e-ticketing/configs"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(role, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})

	t, err := token.SignedString(config.JWT_SECRET)

	if err != nil {
		return "", err
	}
	return t, nil
}
