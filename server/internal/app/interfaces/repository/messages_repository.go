package repository

import "github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"

type MessagesRepository interface {
	GetByChatID(chatID string) ([]entity.Message, error)
	GetBySenderIdAndReceiverId(senderID, receiverID string) ([]entity.Message, error)
	Create(message *entity.Message) error
}
