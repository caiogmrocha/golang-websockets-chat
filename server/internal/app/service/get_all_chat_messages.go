package service

import (
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/repository"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"
)

type GetAllChatMessagesService struct {
	messagesRepository repository.MessagesRepository
}

func (s *GetAllChatMessagesService) Get(senderID, receiverID string) ([]entity.Message, error) {
	return s.messagesRepository.GetBySenderIdAndReceiverId(senderID, receiverID)
}
