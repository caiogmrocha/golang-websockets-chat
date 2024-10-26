package infra_jwt

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct{}

const (
	JWT_TOKEN_EXPIRATION = time.Hour
)

func (jwtProvider *JWTProvider) GenerateToken(payload []byte) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": payload,
		"exp": time.Now().Add(JWT_TOKEN_EXPIRATION).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jwtProvider *JWTProvider) ValidateToken(token string) ([]byte, error) {
	t, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	sub, _ := t.Claims.(jwt.MapClaims)["sub"].(string)

	decodedSub, err := base64.StdEncoding.DecodeString(sub)

	if err != nil {
		return nil, err
	}

	return []byte(decodedSub), nil
}
