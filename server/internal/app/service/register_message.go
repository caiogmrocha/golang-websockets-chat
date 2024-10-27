package service

import (
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/repository"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"
)

type RegisterMessageServiceDTO struct {
	ReceiverID string `json:"receiver_id"`
	SenderID   string `json:"sender_id"`
	Content    string `json:"content"`
}

type RegisterMessageService struct {
	MessagesRepository repository.MessagesRepository
	ChatsRepository    repository.ChatsRepository
}

func (service *RegisterMessageService) Create(dto *RegisterMessageServiceDTO) error {
	chat, err := service.ChatsRepository.GetByUsersIDs([2]string{dto.SenderID, dto.ReceiverID})

	if err != nil {
		return err
	}

	if chat == nil {
		chat = &entity.Chat{
			UsersIDs:    [2]string{dto.SenderID, dto.ReceiverID},
			MessagesIDs: []string{},
		}

		err = service.ChatsRepository.Create(chat)

		if err != nil {
			return err
		}
	}

	message := &entity.Message{
		ChatID:     chat.ID.String(),
		ReceiverID: dto.ReceiverID,
		SenderID:   dto.SenderID,
		Content:    dto.Content,
	}

	return service.MessagesRepository.Create(message)
}
