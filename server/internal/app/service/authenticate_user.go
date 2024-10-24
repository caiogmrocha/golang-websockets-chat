package services

import (
	"encoding/json"
	"errors"

	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/jwt"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticateUserService struct {
  UserRepository repositories.UserRepository
  JWTProvider jwt.JWTProvider
}

type AuthenticateUserServiceParamsDTO struct {
  Email string `json:"email" validate:"required,email"`
  Password string `json:"password" validate:"required,min=6,max=255"`
}

type AuthenticateUserServiceTokenPayload struct {
  ID int `json:"id"`
}

func (service *AuthenticateUserService) Authenticate(dto *AuthenticateUserServiceParamsDTO) (token string, err error) {
  user, err := service.UserRepository.GetByEmail(dto.Email)

  if err != nil {
    return "", errors.New("ERROR_FETCHING_USER")
  }

  if user == nil {
    return "", errors.New("USER_NOT_FOUND")
  }

  err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))

  if err != nil {
    return "", err
  }

  marshalledPayload, _ := json.Marshal(AuthenticateUserServiceTokenPayload{ID: user.ID})

  token, err = service.JWTProvider.GenerateToken([]byte(marshalledPayload))

  if err != nil {
    return "", errors.New("ERROR_GENERATING_TOKEN")
  }

  return token, nil
}
