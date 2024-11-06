package service

import (
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/repository"
	"github.com/caiogmrocha/golang-websockets-chat/server/pkg/utils"
)

type GetUserByIdService struct {
	UsersRepository repository.UserRepository
}

type GetUserByIdServiceResponseDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *GetUserByIdService) Get(id string) (*GetUserByIdServiceResponseDTO, error) {
	user, err := s.UsersRepository.GetByID(id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	userID, _ := utils.ExtractObjectID(user.ID)

	userDTO := GetUserByIdServiceResponseDTO{
		ID:   userID,
		Name: user.Name,
	}

	return &userDTO, nil
}
