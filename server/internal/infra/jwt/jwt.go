package infra_jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct {}

func (jwtProvider *JWTProvider) GenerateToken(payload []byte) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": payload,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

  tokenString, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))

  if err != nil {
    return "", err
  }

	return tokenString, nil
}

func (jwtProvider *JWTProvider) ValidateToken(token string) ([]byte, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("JWT_SECRET"), nil
	})

	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return []byte(t.Claims.(jwt.MapClaims)["sub"].(string)), nil
}
