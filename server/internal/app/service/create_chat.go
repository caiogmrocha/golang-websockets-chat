package service

import (
	"errors"

	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/repository"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"
)

type CreateChatServiceDTO struct {
	UsersIDs [2]string `json:"users_id"`
}

type CreateChatService struct {
	ChatRepository repository.ChatRepository
}

func (service *CreateChatService) Create(dto *CreateChatServiceDTO) error {
	chat, err := service.ChatRepository.GetByUsersIDs(dto.UsersIDs)

	if err != nil {
		return err
	}

	if chat != nil {
		return errors.New("CHAT_ALREADY_EXISTS")
	}

	chat = &entity.Chat{
		UsersIDs:    dto.UsersIDs,
		MessagesIDs: []string{},
	}

	err = service.ChatRepository.Create(chat)

	if err != nil {
		return err
	}

	return nil
}
