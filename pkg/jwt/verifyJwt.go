package jwt

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyJwt(jwtString string) (string, error) {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return "", fmt.Errorf("invalid token")
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	email := fmt.Sprint(claims["email"])
	if email == "" {
		return "", fmt.Errorf("invalid token payload")
	}
	return email, nil
}
