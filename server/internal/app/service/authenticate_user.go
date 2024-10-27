package service

import (
	"encoding/json"
	"errors"

	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/jwt"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/repository"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticateUserService struct {
	UserRepository repository.UserRepository
	JWTProvider    jwt.JWTProvider
}

type AuthenticateUserServiceParamsDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

type AuthenticateUserServiceTokenPayload struct {
	ID string `json:"id"`
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

	userID, _ := utils.ExtractObjectID(user.ID)

	marshalledPayload, _ := json.Marshal(AuthenticateUserServiceTokenPayload{ID: userID})

	token, err = service.JWTProvider.GenerateToken([]byte(marshalledPayload))

	if err != nil {
		return "", errors.New("ERROR_GENERATING_TOKEN")
	}

	return token, nil
}
