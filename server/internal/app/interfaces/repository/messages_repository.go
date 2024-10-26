package repository

import "github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"

type MessagesRepository interface {
	Create(message *entity.Message) error
}
