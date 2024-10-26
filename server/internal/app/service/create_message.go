package service

import (
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/repository"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"
)

type CreateMessageServiceDTO struct {
	ChatID     string `json:"chat_id"`
	ReceiverID string `json:"receiver_id"`
	SenderID   string `json:"sender_id"`
	Content    string `json:"content"`
}

type CreateMessageService struct {
	MessageRepository repository.MessagesRepository
}

func (service *CreateMessageService) Create(dto *CreateMessageServiceDTO) error {
	message := &entity.Message{
		ChatID:     dto.ChatID,
		ReceiverID: dto.ReceiverID,
		SenderID:   dto.SenderID,
		Content:    dto.Content,
	}

	return service.MessageRepository.Create(message)
}
