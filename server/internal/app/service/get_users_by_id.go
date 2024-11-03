package service

import (
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/repository"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/utils"
)

type GetUsersByIdService struct {
	UsersRepository repository.UserRepository
}

type GetUsersByIdServiceResponseDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *GetUsersByIdService) Get(ids []string) ([]GetUsersByIdServiceResponseDTO, error) {
	users, err := s.UsersRepository.GetManyById(ids)

	if err != nil {
		return nil, err
	}

	usersDTO := []GetUsersByIdServiceResponseDTO{}

	for _, user := range users {
		userID, _ := utils.ExtractObjectID(user.ID)

		usersDTO = append(usersDTO, GetUsersByIdServiceResponseDTO{
			ID:   userID,
			Name: user.Name,
		})
	}

	return usersDTO, nil
}
