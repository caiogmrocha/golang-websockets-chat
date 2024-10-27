package service

import (
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/repository"
)

type GetAllChatMessagesService struct {
	MessagesRepository repository.MessagesRepository
}

type GetAllChatMessagesServiceResponseDTO struct {
	Content string `json:"content"`
	Owner   string `json:"owner"`
}

func (s *GetAllChatMessagesService) Get(senderID, receiverID string) ([]GetAllChatMessagesServiceResponseDTO, error) {
	messages, err := s.MessagesRepository.GetBySenderIdAndReceiverId(senderID, receiverID)

	if err != nil {
		return nil, err
	}

	var messagesDTO []GetAllChatMessagesServiceResponseDTO

	var owner string

	for _, message := range messages {
		if message.SenderID == senderID {
			owner = "sender"
		} else {
			owner = "receiver"
		}

		messagesDTO = append(messagesDTO, GetAllChatMessagesServiceResponseDTO{
			Content: message.Content,
			Owner:   owner,
		})
	}

	return messagesDTO, nil
}
