package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	s, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return s, nil
}
