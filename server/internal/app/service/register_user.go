package service

import (
	"errors"

	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/repository"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"
)

type RegisterUserService struct {
	UserRepository repository.UserRepository
}

type RegisterUserServiceDTO struct {
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

func (service *RegisterUserService) Create(dto *RegisterUserServiceDTO) error {
	user, err := service.UserRepository.GetByEmail(dto.Email)

	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("USER ALREADY EXISTS")
	}

	user = &entity.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

	return service.UserRepository.Create(user)
}
